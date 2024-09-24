package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)



type Response struct {
	Data Project `json:"data"`
}

type Project struct {
	GID      string   `json:"gid"`
	Members  []User   `json:"members"`
	Owner    User     `json:"owner"`
	Followers []User   `json:"followers"`
}

type User struct {
	GID           string `json:"gid"`
	Name          string `json:"name"`
	ResourceType  string `json:"resource_type"`
}


func main() {

	var timer int
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Fetching time in minutes(5 or 30): ")
	input, err := reader.ReadString('\n') // Read until newline
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	
	input = strings.TrimSpace(input)
	switch input {
	case "5":
		timer = 5
	case "30":
		timer = 30
	default:
		fmt.Println("Wrong input")
	}

	fmt.Println(strconv.Itoa(timer))
	StartFetching(timer)
}

func StartFetching(timer int) {
	ticker := time.NewTicker(time.Duration(timer) * time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			fmt.Println("\nFetching data every", timer, "minutes...")
			fetchData()
		}
	}
}

func fetchData(){

	url1:= "https://app.asana.com/api/1.0/projects/1208387562722983"
	url2:= "https://app.asana.com/api/1.0/projects/1208387562723121"
	url3:= "https://app.asana.com/api/1.0/projects/1208387562723128"

	urlSlice := []string{url1, url2, url3}
	// Make a slice of project/users to hold each json respectively

	projectSlice := make([]string, 3) // hold just 3 projects
	usersSlice := make([]string, 12) // 3 projects with 4 users each


	for _, url := range urlSlice {
		projectInfo := getProjectByGid(url)
		projectSlice = append(projectSlice, projectInfo)
	}
	for _, projectInfo := range projectSlice {
		fmt.Println(projectInfo)
	}
	// Testing projects Json info
	for _, projectInfo := range projectSlice {
		userPerProject := displayProjectUsers(projectInfo)	
			usersSlice = append(usersSlice, userPerProject...)
	}

	for _,user := range usersSlice {
		fmt.Println("_______________________________________")
		fmt.Println(user)
		fmt.Println("_______________________________________")

	}
}


func getProjectByGid(url string) string{
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("authorization", "Bearer 2/1208387562521060/1208388116305736:6371b7bc3e073c024697e3094d1ddcb0")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	return string(body)
}

func displayProjectUsers(projectUrl string) []string {
	var response Response
	var userInfoSlice = make([]string, 4)
	if err := json.Unmarshal([]byte(projectUrl), &response); err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return nil
	}
	for _, follower := range response.Data.Members {
		userInfoSlice = append(userInfoSlice,getUsersByGid(follower.GID))
	}

	return userInfoSlice
}

func getUsersByGid(userGid string) string {
	
	url := "https://app.asana.com/api/1.0/users/"
	parts := []string{url, userGid}
	fullUrl := strings.Join(parts, "")
	
	req, _ := http.NewRequest("GET", fullUrl, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("authorization", "Bearer 2/1208387562521060/1208388116305736:6371b7bc3e073c024697e3094d1ddcb0")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	return string(body)
}