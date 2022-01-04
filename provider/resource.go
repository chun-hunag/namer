package provider

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type NamesResource struct {
	LastNames       []string `yaml:"lastNames"`
	CharacterMale   []string `yaml:"characterMale"`
	CharacterFemale []string `yaml:"characterFemale"`
}

func ParseFile(file string) (res NamesResource, err error) {
	readBytes, err := ioutil.ReadFile(file)
	if err != nil {
		return res, err
	}

	if err := yaml.Unmarshal(readBytes, &res); err != nil {
		return res, err
	}

	return res, nil
}
