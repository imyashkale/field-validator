package main

import (
	"fmt"
	"github.com/imyashkale/field-validator/validator"
	"log"
	"net/http"
	"strings"
)

func main() {

	_, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		log.Fatal(err)
	}

	// performs the checks on data and return what is failed
	f, err := validator.DataValidator(strings.NewReader(
		`[
			{"age":11}
		]`),
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("what failed :", f)

}
