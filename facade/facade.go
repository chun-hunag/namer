package facade

import (
	"example/name-generator/provider"
)

type GenerateItemProcess func(item string, index int)

type QueryItemProcess func(item string)

var (
	GenerateFirstNameNum int
	GenerateGender       string
)

func init() {
	reset()
}

func reset() {
	GenerateFirstNameNum = 2
	GenerateGender = ""
}

func Generate(path string, count int, process GenerateItemProcess) error {
	res, _ := provider.ParseFile(path)

	generator := provider.Create()
	generator.Resource = res

	for i := 0; i < count; i++ {
		name, err := generator.Name(GenerateGender, GenerateFirstNameNum)

		if err != nil {

		}

		process(name, i)
	}

	return nil
}

func Query(str string, process QueryItemProcess) error {
	dict, err := provider.Query(str)

	if err != nil {
		return err
	}

	for _, heteronym := range dict.Heteronyms {
		for _, definition := range heteronym.Definitions {
			process(definition)
		}
	}

	return nil
}
