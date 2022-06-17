package main

import (
	"Discord_Friend_Graph/DataScrape/discord"
	"Discord_Friend_Graph/types"
	"encoding/json"
	"log"
)

func GetOwnUserData() {
	res, err := discord.GET("https://discord.com/api/v9/users/@me")
	if err != nil {
		log.Fatal(err)
	}

	user := types.Self{}
	if err := json.NewDecoder(res.Body).Decode(&user); err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	discord.SaveMutualFriends(user.Id)
}

func main() {
	GetOwnUserData()
}
