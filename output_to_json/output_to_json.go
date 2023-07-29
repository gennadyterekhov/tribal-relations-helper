package main

import (
	"encoding/json"
	"fmt"

	"github.com/gennadyterekhov/tribal-relations-core/entity"
)

func main() {
	fmt.Println("{")

	for key, resource := range entity.ResourceNameToResourceMap {
		fmt.Println("\"", key, "\":")

		jsonResource, err := json.Marshal(resource)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(jsonResource))
		fmt.Println(",")

	}
	fmt.Println("}")

}
