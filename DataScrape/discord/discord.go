package discord

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func SaveMutualFriends(friendId string) {
	result, err := GetMutualFriends(friendId)
	if err != nil {
		log.Fatal(err)
	}

	if err := CreateFriendFile(friendId, result.Body); err != nil {
		log.Fatal(err)
	}
}

func GET(url string) (*http.Response, error) {
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Set("Authorization", os.Getenv("DISCORD_TOKEN"))

	result, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return result, nil
}

func SaveToAFile(path string, result io.ReadCloser) error {
	out, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer out.Close()

	io.Copy(out, result)
	defer result.Close()

	return nil
}

func CreateFriendFile(friendID string, result io.ReadCloser) error {
	return SaveToAFile(fmt.Sprintf("Friends_Mutal_friends/%s.json", friendID), result)
}

func GetMutualFriends(id string) (*http.Response, error) {
	return GET(fmt.Sprintf("https://discord.com/api/v9/users/%s/relationships", id))
}
