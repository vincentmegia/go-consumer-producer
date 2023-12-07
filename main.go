package main

import (
	"consumer-producer/models"
	"consumer-producer/repository"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

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
	for index := 1; index != 5000; index++ {
		user := models.User{Id: strconv.Itoa(index), Firstname: "John", Lastname: "Doe", Msisdn: "12345678"}
		log.Printf("Adding of users done.: %s", user)
		repository.AddUser(&user)
	}
	log.Println("producing data completed.")
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
					log.Println("file has been cleaned up.")
				}
			case error := <-watcher.Errors:
				fmt.Println("ERROR:", error)
			}
		}
	}()
	return watcher
}

func main() {
	repository.CreateTable()
	// fileWatchDone := make(chan bool)
	watcher := createWatcher()
	/* feedPath, error := os.Getwd()
	if error != nil {
		log.Println("Failed to get current working directory", error)
	} */
	error := watcher.Add("./feed")
	if error != nil {
		log.Println("Failed to add watch file users.txt", error)
	}

	defer watcher.Close()
	var command string
	fmt.Println("Enter a command:")
	for {
		_, error := fmt.Scanln(&command)
		if error != nil {
			log.Println("Reading user input error: ", error)
			continue
		}

		switch command {
		case "exit":
			os.Exit(0)
		default:
			log.Println("Enter a valid command")
		}
	}
	/* channel := make()
	produce(&waitGroup)
	log.Print("Waiting for go routines to finish")
	waitGroup.Wait()
	rows := repository.GetAllUsers()
	log.Printf("Success fetched users: %s", rows) */
}
