package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Name    string   `json:"title"`
	Authors []string `json:"authors"'`
	Age     int      `json:"age"`
}

func main() {
	movie := Movie{"喜剧之王", []string{"周星驰", "张柏芝"}, 2000}
	marshal, err := json.Marshal(movie)
	if err == nil {
		fmt.Printf("json= %s\n", marshal)
	}
	newMovie := Movie{}
	err = json.Unmarshal(marshal, &newMovie)
	if err == nil {
		fmt.Printf("%v\n", newMovie)
	} else {
		fmt.Println("解析异常")
	}
}
