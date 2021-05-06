package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/imyashkale/field-validator/validator"
)

func main() {

	_, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		log.Fatal(err)
	}

	mp, err := validator.DataValidator(strings.NewReader(`[{"title":"yash","body":"I like sabrina carpenter"}]`))
	if err != nil {
		log.Fatal(err)
	}

	// WORKING ON THIS
	f := map[int]map[string][]string{}
	for recordIndex, v := range mp {
		f[recordIndex] = map[string][]string{}
		for field, validationResults := range v {
			var validators []string
			for k, checks := range validationResults {
				if checks {
					validators = append(validators, k)
				}
			}
			f[recordIndex][field] = validators
		}
	}

	fmt.Println("what failed :", f)

}
