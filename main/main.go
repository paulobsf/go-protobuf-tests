package main

import (
	"../schema"
	"code.google.com/p/goprotobuf/proto"
	"fmt"
	"log"
	"os"
)

func main() {
	firstMeal := &schema.Meal{
		proto.String("Piri Piri Chicken"),
		[]string{
			"Chicken",
			"Piri Piri Sauce",
		},
		proto.Float32(8.5),
		schema.Stars_FOUR.Enum(),
		nil,
	}

	data, err := proto.Marshal(firstMeal)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	os.Stdout.Write(data)
	fmt.Println()

	fmt.Println(firstMeal)

	secondMeal := &schema.Meal{}
	err = proto.Unmarshal(data, secondMeal)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}

	fmt.Println(secondMeal)
}
