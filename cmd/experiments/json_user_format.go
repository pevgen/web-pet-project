package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

func main() {
	//person := Person{Name: "aaa", Age: json.Number(strconv.Itoa(1))}
	person := Person{Name: "aaa", Age: 1}

	b, err := json.Marshal(person)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))
	fmt.Println()

	personAsString := "{\"name\":\"aaa\",\"age\":\"1\"}"

	var p2 Person
	err = json.Unmarshal([]byte(personAsString), &p2)
	if err != nil {
		panic(err)
	}

	fmt.Println(p2)
}

type AgeInt int

type Person struct {
	Name string `json:"name"`
	//Age  json.Number `json:"age"`
	Age AgeInt `json:"age"`
}

func (a *AgeInt) UnmarshalJSON(b []byte) error {
	//convert the bytes into an interface
	//this will help us check the type of our value
	var item interface{}
	if err := json.Unmarshal(b, &item); err != nil {
		return err
	}

	switch v := item.(type) {
	case int:
		*a = AgeInt(v)
	case string:
		i, err := strconv.Atoi(v)
		if err != nil {

			return err

		}
		*a = AgeInt(i)
	}
	return nil
}
