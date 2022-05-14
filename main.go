package main

import (
	"Discord_Friend_Graph/types"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

func GetAllFriends(session neo4j.Session) {
	Data := GetFriendDataFromJsonFile("./friends.json")

	log.Println(len(Data))
	for _, friend := range Data {
		value, err := CreateFriend(session, friend)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(value)
	}
}

func GetFriendDataFromJsonFile(fileName string) []types.Friend {
	friendsFile, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	var Data []types.Friend
	if err := json.NewDecoder(friendsFile).Decode(&Data); err != nil {
		log.Fatal(err)
	}
	return Data
}

func main() {
	driver, session, err := Connect("bolt://localhost:7687", "neo4j", "test")
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()
	defer driver.Close()

	GetAllFriends(session)

	files, err := ioutil.ReadDir("./DataScrape/Friends_Mutal_friends")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		Data := GetFriendDataFromJsonFile("./DataScrape/Friends_Mutal_friends/" + file.Name())

		id := strings.TrimSuffix(file.Name(), ".json")
		log.Println(id)
		for _, mutal_friend := range Data {
			log.Println("	", mutal_friend.Id)
			_, err := CreateConnection(session, id, mutal_friend)
			if err != nil {
				log.Fatal(err)
			}
		}

	}
}

func CreateConnection(session neo4j.Session, masterID string, slave types.Friend) (string, error) {
	_, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(
			"Match (f1: Friend), (f2: Friend)\nwhere f1.id=$id1 and f2.id=$id2\ncreate (f1)-[p:Friend]->(f2)\nreturn f1,f2",
			map[string]interface{}{"id1": masterID, "id2": slave.Id})
		if err != nil {
			return nil, err
		}

		return nil, result.Err()
	})
	if err != nil {
		return "", err
	}

	return "", nil
}

func CreateFriend(session neo4j.Session, friend types.Friend) (string, error) {
	greeting, err := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		result, err := transaction.Run(
			"CREATE (a:Friend{name: $name, id: $id}) RETURN a.id",
			map[string]interface{}{"name": friend.User.Username, "id": friend.Id})
		if err != nil {
			return nil, err
		}

		if result.Next() {
			return result.Record().Values[0], nil
		}

		return nil, result.Err()
	})
	if err != nil {
		return "", err
	}

	return greeting.(string), nil
}

func Connect(uri string, username string, password string) (neo4j.Driver, neo4j.Session, error) {
	driver, err := neo4j.NewDriver(uri, neo4j.BasicAuth(username, password, ""))
	if err != nil {
		return nil, nil, err
	}
	session := driver.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})

	return driver, session, nil
}
