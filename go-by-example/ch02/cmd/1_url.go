package main

import (
	"fmt"
	"log"

	"github.com/cynicdog/gopherdojo/go-by-example/ch02/a_url"
)

func main() {
	url, err := a_url.Parse("https://github.com/cynicdog")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Scheme	:", url.Scheme)
	fmt.Println("Host	:", url.Host)
	fmt.Println("Path	:", url.Path)

	fmt.Println(url)
}
