package gsconfig

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type bufType map[interface{}]interface{}
type configBufType map[string]string

func yamlToConfig(path string) (map[string]string, error) {
	buf := make(bufType)
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(data, &buf)
	if err != nil {
		return nil, err
	}

	configBuf := make(configBufType)
	bufAnalysis(buf, configBuf)

	return configBuf, nil

}

func bufAnalysis(buf bufType, configBuf configBufType) (configBufType, error) {
	for key, value := range buf {
		switch value.(type) {
		case bufType:
			err := valueAnalysis(key.(string), value.(bufType), configBuf)
			if err != nil {
				return nil, err
			}
		case string:
			//fmt.Println(key.(string)+":", value.(string))
			configBuf[key.(string)] = value.(string)
		default:
			return nil, fmt.Errorf("value type should be string or map[interface{}]interface{}")
		}
	}

	return configBuf, nil
}

func valueAnalysis(prefix string, value bufType, configBuf configBufType) error {
	for k, v := range value {
		switch v.(type) {
		case bufType:
			valueAnalysis(prefix+"."+k.(string), v.(bufType), configBuf)
		case string:
			configBuf[prefix+"."+k.(string)] = v.(string)
		default:
			return fmt.Errorf("value type should be string or map[interface{}]interface{}")
		}
	}
	return nil
}
