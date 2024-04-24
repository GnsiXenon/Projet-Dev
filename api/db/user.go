package db

// User table struct
type User struct {
	Id       int    `json:"id-user"`
	Name     string `json:"name"`
	Mail     string `json:"mail,omitempty"`
	Password string `json:"password,omitempty"`
	Hash     string `json:"hash,omitempty"`
	Score    int    `json:"score"`
}
