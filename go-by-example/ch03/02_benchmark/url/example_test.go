package url_test

import (
	"fmt"
	"log"

	"github.com/cynicdog/gopherdojo/go-by-example/ch03/02_benchmark/url"
)

func ExampleParse() {
	uri, err := url.Parse("https://github.com/cynicdog")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(uri)
	// Output:
	// https://github.com/cynicdog
}
