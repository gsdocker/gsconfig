package gsconfig

import (
	"encoding/json"
	"io/ioutil"
)

// LoadJSON load config from json file
func LoadJSON(file string) error {
	var pairs map[string]string

	data, err := ioutil.ReadFile(file)

	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &pairs)

	if err != nil {
		return err
	}

	Save(pairs)

	return nil
}
