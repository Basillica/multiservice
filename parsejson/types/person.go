package types

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Person struct {
	Firstname   string `json:"firstname"`
	Lastname    string `json:"lastname"`
	Birthday    string `json:"birthday"`
	Address     string `json:"address"`
	PhoneNumber int64  `json:"phone_number"`
}

// String returns a string representation of the Person struct
func (p *Person) String() string {
	return fmt.Sprintf("%#v", &p)
}

// FetchUsers reads the json files within the json directory and parses them into the Person struct
//
// It returns a slice of the Person struct and an error if there was an error reading the files or parsing the contents
func (p *Person) FetchUsers() ([]Person, error) {
	files, err := os.ReadDir("./json")
	var persons []Person
	if err != nil {
		return persons, err
	}

	for _, file := range files {
		if !file.IsDir() {
			if person, err := p.ParseJSON("./json/" + file.Name()); err == nil {
				persons = append(persons, person)
			}
		}
	}
	return persons, nil
}

// ParseJSON reads a JSON file and parses its content
//
// It marshals the content into a Person struct if successfully read.
// It returns an error if there was an error, otherwise returns a Person struct
func (p *Person) ParseJSON(filepath string) (Person, error) {
	jsonFile, err := os.Open(filepath)
	if err != nil {
		return Person{}, nil
	}
	defer jsonFile.Close()

	var person Person
	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return person, err
	}
	json.Unmarshal(byteValue, &person)
	return person, nil
}
