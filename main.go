package main

import (
	"fmt"
	"github.com/imyashkale/field-validator/validator"
	"log"
	"strings"
)

func main() {

	// performs the checks on data and return what is failed
	f, err := validator.DataValidator(strings.NewReader(
		`[
			{"name": "yash" , "age":9}
		]`),
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("what failed :", f)

}
