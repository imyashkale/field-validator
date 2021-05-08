package main

import (
	"log"
	"net/http"
	"strings"
	"fmt"
	"github.com/imyashkale/field-validator/validator"
)

func main() {

	_, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		log.Fatal(err)
	}

	// performs the checks on data and return what is failed
	f , err := validator.DataValidator(strings.NewReader(
		`[
			{"name":"Yash","bio":"I like coading"}
		]`),
	)
	if err != nil {
		log.Fatal(err)
	}


	fmt.Println("what failed :", f)

}
