# Discord-Friend-Visualize

* using neo4j and data extracted from the discord API

<img src="assets/graph.svg" width="60%" alt="final-graph">

### SetUp

All the scraped data is saved to json files and saved to neo4j.

0. Add env variable with your personal Discord token

`export DISCORD_TOKEN=<your_token>`

2. First get all the friends.

`go run ./DataScrape/getAllFriends.go`

2. For each friend get the mutual friends

`go run ./DataScrape/getMutualFriends.go`

3. Use neoj Aura or run locally with Docker

`bash ./ne4jdb.sh`

4. Insert it all into neo4j and visualize

`go run main.go`
    
