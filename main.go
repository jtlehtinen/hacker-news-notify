package main

import "fmt"

func main() {
	top, _ := fetchTop()
	//fmt.Println(top)

	id := top[0]
	story := fetchStory(id)
	fmt.Println(story)

	//fetchNews()
	//fetchBest()
}
