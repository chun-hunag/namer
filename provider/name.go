package provider

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type Generator struct {
	rand     *rand.Rand
	Resource NamesResource
}

func (generator *Generator) Name(gender string, fisrtNameNum int) (name string, err error) {
	switch gender {
	case "":
		name = generator.LastName() + generator.firstName(fisrtNameNum)
	case "male":
		name = generator.LastName() + generator.firstNameMale(fisrtNameNum)
	case "female":
		name = generator.LastName() + generator.firstNameFemale(fisrtNameNum)
	default:
		err = errors.New(fmt.Sprintf("gender '%s' is invalid", gender))
	}

	return
}

func (generator *Generator) FirstNameMale() string {
	return generator.firstNameMale(0)
}

func (generator *Generator) FirstNameFemale() string {
	return generator.firstNameFemale(0)
}

func (generator *Generator) NameSingle() string {
	return generator.LastName() + generator.FirstNameSingle()
}

func (generator *Generator) NameDouble() string {
	return generator.LastName() + generator.FirstNameDouble()
}

func (generator *Generator) LastName() string {
	return generator.pickCharacter(generator.Resource.LastNames, 1)
}

func (generator *Generator) firstName(count int) string {
	if count == 0 {
		count = generator.rand.Intn(2) + 1
	}
	
	merge := append(generator.Resource.CharacterMale, generator.Resource.CharacterFemale...)
	return generator.pickCharacter(merge, count)
}

func (generator *Generator) FirstNameSingle() string {
	return generator.firstName(1)
}

func (generate *Generator) FirstNameDouble() string {
	return generate.firstName(2)
}

func (generator *Generator) firstNameMale(count int) string {
	if count == 0 {
		count = generator.rand.Intn(2) + 1
	}

	return generator.pickCharacter(generator.Resource.CharacterMale, count)
}

func (generator *Generator) firstNameFemale(count int) string {
	if count == 0 {
		count = generator.rand.Intn(2) + 1
	}

	return generator.pickCharacter(generator.Resource.CharacterFemale, count)
}

func (generator *Generator) pickCharacter(collection []string, count int) string {
	length := len(collection)
	chars := ""
	for i := 0; i < count; i++ {
		randomIndex := generator.rand.Intn(length)
		chars = chars + collection[randomIndex]
	}

	return chars
}

func Create() Generator {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	return Generator{
		rand: r,
	}
}
