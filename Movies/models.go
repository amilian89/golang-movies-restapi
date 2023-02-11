package main

type Movie struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Year   string `json:"year"`
	Genre  string `json:"genre"`
	Actor  string `json:"actor"`
	Rating int    `json:"rating"`
}
