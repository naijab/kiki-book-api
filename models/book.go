package models

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Cover  string `json:"cover"`
	Price  int    `json:"price"`
}
