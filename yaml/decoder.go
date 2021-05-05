package yaml

import (
	"fmt"
	"github.com/imyashkale/field-validator/models"
	"gopkg.in/yaml.v2"
	"os"
)

//ConfigFile Reader Reads config file convert to go's native type's
func ConfigFileReader(filePath string) (models.PostConfig, error) {

	config := &models.PostConfig{}
	// opening config file
	file, err := os.Open(filePath)
	if err != nil {
		return *config, fmt.Errorf("unable to decode %v file ", err)
	}
	//deffering file close
	defer file.Close()

	d := yaml.NewDecoder(file)

	if err := d.Decode(&config); err != nil {
		return *config, err
	}

	return *config, nil
}
