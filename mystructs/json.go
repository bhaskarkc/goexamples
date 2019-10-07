package mystructs

import "fmt"
import "encoding/json"

func MarshallExample() string {
	p1 := Person{
		First: "Bhaskar",
		Last:  "K C",
		Age:   31,
	}

	p2 := Person{
		First: "Anita",
		Last:  "Jayana",
		Age:   31,
	}

	// p1.Speak()
	// p2.Speak()

	people := []Person{p1, p2}

	fmt.Printf("Slice of People struct: %v\n", people)

	bs, err := json.Marshal(people)

	if err != nil {
		err.Error()
	}

	// convert bytestring to JSON string.
	return string(bs)
}

func UnmarshallExample(m func() string) {
	// JSON string to be converted.
	jsonString := m()

	// GO struct that json is to be mapped to.
	people := []Person{}

	err := json.Unmarshal([]byte(jsonString), &people)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(people)
}
