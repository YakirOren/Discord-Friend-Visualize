package main

import (
	"Discord_Friend_Graph/types"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	friendsFile, err := os.Open("../friends.json")
	if err != nil {
		log.Fatal(err)
	}

	AllFriends := []types.Friend{}
	if err := json.NewDecoder(friendsFile).Decode(&AllFriends); err != nil {
		log.Fatal(err)
	}

	for _, friend := range AllFriends {
		log.Println(friend.Id)
		result, err := GetMutalFriends(friend.Id)
		if err != nil {
			log.Fatal(err)
		}

		if err := CreateFriendFile(friend.Id, result.Body); err != nil {
			log.Fatal(err)
		}

		time.Sleep(10 * time.Second)
	}
}

func GetMutalFriends(id string) (*http.Response, error) {
	url := fmt.Sprintf("https://discord.com/api/v9/users/%s/relationships", id)
	return SendDiscordGETRequest(url)
}

func SendDiscordGETRequest(url string) (*http.Response, error) {
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Set("Authorization", os.Getenv("DISCORD_TOKEN"))

	result, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return result, nil
}

func CreateFriendFile(fileName string, result io.ReadCloser) error {
	out, err := os.Create(fmt.Sprintf("Friends_Mutal_friends/%s.json", fileName))
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer out.Close()

	io.Copy(out, result)
	defer result.Close()

	return nil
}
