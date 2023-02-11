package main

type Actor struct {
	ID        string `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Gender    string `json:"gender"`
	Age       int    `json:"age"`
	Movies    string `json:"movies"`
}
