package main

import (
	"Discord_Friend_Graph/DataScrape/discord"
	"Discord_Friend_Graph/types"
	"encoding/json"
	"log"
	"os"
	"time"
)

const RequestsDelay = 10

const RawDataPath = "raw/friends.json"

func main() {
	SaveRawFriendsData()

	friendsFile, err := os.Open(RawDataPath)
	if err != nil {
		log.Fatal(err)
	}

	var AllFriends []types.Friend
	if err := json.NewDecoder(friendsFile).Decode(&AllFriends); err != nil {
		log.Fatal(err)
	}

	for _, friend := range AllFriends {
		log.Println(friend.Id)

		discord.SaveMutualFriends(friend.Id)

		time.Sleep(RequestsDelay * time.Second)
	}
}

// SaveRawFriendsData gets a json from discord that describes the user's friends and saves it to raw/friends.json
func SaveRawFriendsData() {
	res, err := discord.GET("https://discord.com/api/v9/users/@me/relationships")
	if err != nil {
		log.Fatal(err)
	}

	if err := discord.SaveToAFile(RawDataPath, res.Body); err != nil {
		return
	}
}
