package validator

import (
	"encoding/json"
	"github.com/imyashkale/field-validator/yaml"
	"io"
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
	var records []map[string]string

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
			for _, chks := range checks {
				switch chks {
				case "exist":
					// if its empty then check failed on this
					// this will be included in the data
					if currentValue == "" {
						mp[idx][currentField] = append(mp[idx][currentField], chks)
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
	return mp, nil
}
