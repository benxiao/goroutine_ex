package main

import (
	"time"
	"math/rand"
	"fmt"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	results := Google("golang")
	elapsed := time.Since(start)
	fmt.Println(results)
	fmt.Println(elapsed)

}

var (
	Web = fakeSearch("web")
	Image = fakeSearch("image")
	Video = fakeSearch("video")
)

type Result struct {
	str string
}

type Search func(query string) Result

func fakeSearch(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return Result{fmt.Sprintf("%s result for %q\n", kind, query)}
	}
}

func Google(query string) []Result {
	results := make([]Result, 0, 3)
	c := make(chan Result)
	go func(){c<-Web(query)}()
	go func(){c<-Image(query)}()
	go func(){c<-Video(query)}()
	for i:=0; i!=3; i++{
		results=append(results, <-c)

	}
	return results
}