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
	friends_file, err := os.Open("../friends.json")
	if err != nil {
		log.Fatal(err)
	}

	Data := []types.Friend{}
	if err := json.NewDecoder(friends_file).Decode(&Data); err != nil {
		log.Fatal(err)
	}

	for _, friend := range Data {
		log.Println(friend.Id)
		result, err := GetMutalFriends(friend)
		if err != nil {
			log.Fatal(err)
		}

		if err := CreateFriendFile(friend.Id, result.Body); err != nil {
			log.Fatal(err)
		}

		time.Sleep(10 * time.Second)
	}
}

func GetMutalFriends(friend types.Friend) (*http.Response, error) {
	url := fmt.Sprintf("https://discord.com/api/v9/users/%s/relationships", friend.Id)
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Set("Authorization", os.Getenv("DISCORD_TOKEN"))

	result, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer result.Body.Close()

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
	return nil
}
