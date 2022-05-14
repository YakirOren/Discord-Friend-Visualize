package types

type Friend struct {
	Id   string `json:"id"`
	User struct {
		Username string `json:"username"`
	} `json:"user"`
}
