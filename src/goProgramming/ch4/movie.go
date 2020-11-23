//json编码和解码

package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

func main() {
	var movies = []Movie{
		{Title: "卧虎藏龙", Year: 1990, Color: false, Actors: []string{"章子怡", "梁朝伟"}},
		{Title: "战狼", Year: 2019, Color: true, Actors: []string{"吴京", "张翰"}},
		{Title: "八佰", Year: 2020, Color: true, Actors: []string{"黄晓明", "欧豪", "管虎"}},
	}
	data, err := json.MarshalIndent(movies, "", "  ")
	if err != nil {
		log.Fatalf("json marshaling failed:%s\n", err)
	}
	fmt.Printf("%s\n", data)

	var titles []struct {
		Title string
		Year  int  `json:"released"`
		Color bool `json:"color,omitempty"`
	}
	if err := json.Unmarshal(data, &titles); err != nil {
		log.Fatalf("json unmarshaling failed:%s\n", err)
	}
	fmt.Printf("%v\n", titles)
}
