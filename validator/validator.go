package validator

import (
	"encoding/json"
	"github.com/imyashkale/field-validator/models"
	"github.com/imyashkale/field-validator/yaml"
	"io"
	"log"
	"reflect"
)

type validatorResult map[int]map[string]map[string]bool

func DataValidator(data io.Reader) (validatorResult, error) {
	mp := validatorResult{}

	// Decoding yaml configuration to the go's native type
	ymlConfig, err := yaml.ConfigFileReader("./config.yaml")
	if err != nil {
		return mp, err
	}
	// Decoding data which on check will be happing
	var posts []models.Post
	// Decoding string to json
	err = json.NewDecoder(data).Decode(&posts)
	if err != nil {
		return mp, err
	}

	for idx, item := range posts {
		v := reflect.ValueOf(item)

		mp[idx] = map[string]map[string]bool{}
		for index := 0; index < v.NumField(); index++ {
			currentField := v.Type().Field(index).Name
			currentValue := v.Field(index).Interface()
			r := reflect.ValueOf(ymlConfig)
			chks := reflect.Indirect(r).FieldByName(currentField).Interface().([]string)
			// if no check found
			// check for next field
			if len(chks) == 0 {
				log.Printf("No checks found for record[%d] Field [%v]", idx, currentField)
				continue
			}
			mp[idx][currentField] = map[string]bool{}
			for _, chks := range chks {

				switch chks {
				case "exist":
					if currentValue.(string) != "" {
						mp[idx][currentField][chks] = true
						break
					}
					mp[idx][currentField][chks] = false
				default:
					mp[idx][currentField][chks] = false
				}
			}
			// field name
			// fmt.Printf("Record[%d] : [%d] : Field : %v :	 Value  : %v : Checks : %v 	\n", idx, index, currentField, currentValue, chks)

		}
	}

	return mp, nil
}
