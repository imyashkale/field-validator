package validator

import (
	"encoding/json"
	"fmt"
	"github.com/imyashkale/field-validator/yaml"
	"io"
	"log"
	"strconv"
)

//DataValidator This performs checks on data.
func DataValidator(input io.Reader) (map[int]map[string][]string, error) {

	// storing the validators result
	mp := map[int]map[string][]string{}

	// Decoding yaml configuration to the go's native type
	config, err := yaml.ConfigFileReader("./config.yaml")
	if err != nil {
		return mp, err
	}
	// Decoding data which on check will be happing
	var records []map[string]interface{}

	// Decoding string to json
	err = json.NewDecoder(input).Decode(&records)
	if err != nil {
		return mp, err
	}
	// performing checks on each record in records.
	// in js context : record is element of array
	// this is the records : [{} ,{} ,{}]
	for idx, record := range records {
		mp[idx] = map[string][]string{}
		for currentField, currentValue := range record {
			// looking for the checks to perform on this field
			// in cofig. because it holds all the what to check on what
			// config got this checks information from the config.yaml
			checks := config[currentField]
			mp[idx][currentField] = []string{}
			for _, check := range checks {
				for checkKey, checkValue := range check {
					switch checkKey {
					// for exist check
					case "exist":
						// if its empty then check failed on this
						// this will be included in the data
						switch t := currentValue.(type) {
						case string:
							if checkValue == "yes" && t == "" {
								mp[idx][currentField] = append(mp[idx][currentField], checkKey)
							}
							if checkValue == "no" && t != "" {
								mp[idx][currentField] = append(mp[idx][currentField], checkKey)
							}
						case float64:
							if checkValue == "yes" && t == 0 {
								mp[idx][currentField] = append(mp[idx][currentField], checkKey)
							}
							if checkValue == "no" && t != 0 {
								mp[idx][currentField] = append(mp[idx][currentField], checkKey)
							}
						}
					case "min":
						currentValueLen, err := strconv.Atoi(checkValue)
						if err != nil {
							log.Println(err)
						}
						switch t := currentValue.(type) {
						case string:
							// here t is string
							// t converted to the int
							if len(t) < currentValueLen {
								mp[idx][currentField] = append(mp[idx][currentField], checkKey)
							}
						case float64:
							if t < float64(currentValueLen) {
								mp[idx][currentField] = append(mp[idx][currentField], checkKey)
							}
						}
					case "max":
						currentValueLen, err := strconv.Atoi(checkValue)
						if err != nil {
							log.Println(err)
						}
						switch t := currentValue.(type) {
						case string:
							// here t is string
							// t converted to the int
							if len(t) > currentValueLen {
								mp[idx][currentField] = append(mp[idx][currentField], checkKey)
							}
						case float64:
							if t > float64(currentValueLen) {
								mp[idx][currentField] = append(mp[idx][currentField], checkKey)
							}
						}

					}
				}
			}
			// if no checks failed on this field
			// then why include
			// only return to caller whats failed
			if len(mp[idx][currentField]) == 0 {
				delete(mp[idx], currentField)
			}
		}
		// if no checks failed on any field of record
		// then why include
		// only return to caller whats failed
		if len(mp[idx]) == 0 {
			delete(mp, idx)
		}
	}
	fmt.Println(mp)
	return mp, nil
}
