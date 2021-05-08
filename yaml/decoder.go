package yaml

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

//ConfigFile Reader Reads config file convert to go's native type's
func ConfigFileReader(filePath string) (map[string][]map[string]string, error) {
	// file path is the configuration file which is written in yaml
	var config map[string][]map[string]string
	// opening config file
	file, err := os.Open(filePath)
	if err != nil {
		return config, fmt.Errorf("unable to parse config file %v file ", err)
	}

	defer file.Close()

	if yaml.NewDecoder(file).Decode(&config); err != nil {
		return config, fmt.Errorf("unable to decode file %v file ", err)
	}
	return config, nil
}
