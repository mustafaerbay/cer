package modals

type User struct {
	ID int64 		`json:"id"`
	Name string 	`json:"name"`
	Username string `json:"username"`
	State string 	`json:"state"`
}