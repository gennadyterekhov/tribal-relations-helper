package main

import (
	"encoding/json"
	"fmt"

	"github.com/gennadyterekhov/tribal-relations-core/entity"
)

func main() {
	fmt.Println("{")

	for key, relation := range entity.RelationNameToRelationMap {
		fmt.Printf("\"%v\":", key)

		jsonResource, err := json.Marshal(relation)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(jsonResource))
		fmt.Println(",")

	}
	fmt.Println("}")

}
