package types

type User struct {
	Fullname  string `json:"Fullname"`
	Username  string `json:"Username"`
	Password  string `json:"Password"`
	PasswordC string `json:"PasswordC"`
}

type Book struct {
	Quantity int    `json:"quantity"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	ID       int    `json:"id"`
}
