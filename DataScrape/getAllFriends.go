package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	friends_file, err := os.Create("../friends.json")
	if err != nil {
		log.Fatal(err)
	}
	defer friends_file.Close()

	res, err := GetAllRelationShips()
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	io.Copy(friends_file, res.Body)
}

func GetAllRelationShips() (*http.Response, error) {
	return SendDiscordGETRequest("https://discord.com/api/v9/users/@me/relationships")
}
