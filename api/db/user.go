package db

// User table struct
type User struct {
	Id       int    `json:"id-user"`
	Name     string `json:"name"`
	Mail     string `json:"mail"`
	Password string `json:"password"`
	Hash     string `json:"hash"`
	Score    int    `json:"score"`
}
