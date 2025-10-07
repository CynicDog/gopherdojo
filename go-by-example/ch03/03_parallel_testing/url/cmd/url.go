package main

import (
	"fmt"
	"log"

	"github.com/cynicdog/gopherdojo/go-by-example/ch03/03_parallel_testing/url"
)

func main() {
	uri, err := url.Parse("https://github.com/inancgumus")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Scheme:", uri.Scheme)
	fmt.Println("Host  :", uri.Host)
	fmt.Println("Path  :", uri.Path)
	fmt.Println(uri)
}
