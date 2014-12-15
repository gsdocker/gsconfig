package gsconfig

import (
	"fmt"
	"testing"
)

func TestTo(t *testing.T) {
	config, err := yamlToConfig("test.yaml")
	if err != nil {
		panic(err)
	}
	for k, v := range config {
		fmt.Println(k, v)
	}
}
