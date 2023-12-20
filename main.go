package main

import (
	"consumer-producer/models"
	"consumer-producer/repository"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/fsnotify/fsnotify"
)

/* Exercise:
1) Listen to a file that will contain instruction how many records will be pumped to database
2) the main thread will listen from user input to exit
3) producer will read values from database, change firstname and lastname and push to
channel (should replace blocking queue)
consuer will read values and insert to database.
*/
/*
Producer of data
*/
func generateData() {
	for index := 1; index <= 5000; index++ {
		user := models.User{Id: strconv.Itoa(index), Firstname: "John", Lastname: "Doe", Msisdn: "12345678"}
		fmt.Printf("Adding of users done.: %s, index: %s", user, strconv.Itoa(index))
		fmt.Println()
		repository.AddUser(&user)
	}
	fmt.Println("producing data completed.")
}

func createWatcher() *fsnotify.Watcher {
	watcher, error := fsnotify.NewWatcher()
	if error != nil {
		log.Printf("Failed to create new watcher: %s", error)
	}

	go func() {
		for {
			select {
			case event := <-watcher.Events:
				fmt.Printf("EVENT! %#v\n", event.Op.String())
				if event.Op.String() == "CREATE" && strings.Contains(event.Name, "users.txt") {
					generateData()
					os.Remove("./feed/users.txt")
					fmt.Println("file has been cleaned up.")
				} else if event.Op.String() == "CREATE" && strings.Contains(event.Name, "process.txt") {
					go batchProcess()
					os.Remove("./feed/process.txt")
					fmt.Println("file has been cleaned up. process.txt")
				} else {
					fmt.Println("File watcher default")
				}
			case error := <-watcher.Errors:
				fmt.Println("ERROR:", error)
			}
		}
	}()
	return watcher
}

/*
Producer function will retrieve data from database users,
modify data and pass to channel for consumers to retrieve
*/
func produce() {
	users := repository.GetAllUsers()
	if len(users) == 0 {
		fmt.Println("No records to process")
		return
	}
	log.Println("Processing number of users: ", len(users))
	var batchUsers []models.User
	for _, user := range users {
		timeString := time.Now().String()
		user.Firstname = user.Firstname + timeString
		user.Lastname = user.Lastname + timeString
		batchUsers = append(batchUsers, user)
		if len(batchUsers) == 20 {
			fmt.Println("Pushing to queue batchuser: ", batchUsers)
			queue <- batchUsers
			fmt.Println("Clearing queue")
			batchUsers = []models.User{}
		}
	}
}

/*
retrieve from channel and update in database
*/
func consume() {
	for {
		users := <-queue
		fmt.Println("received users from queue: ", users)
		for _, user := range users {
			fmt.Println("receive from queue: ", user)
		}
	}
}

func batchProcess() {
	go produce()
	go consume()
}

var (
	queue        chan []models.User
	commandQueue chan string
)

func main() {
	log.Println("============== START ================")
	queue = make(chan []models.User)
	repository.CreateTable()
	watcher := createWatcher()
	error := watcher.Add("./feed")
	if error != nil {
		log.Println("Failed to add watch file users.txt", error)
	}
	select {}
}
