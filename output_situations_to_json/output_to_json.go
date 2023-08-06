package main

import (
	"encoding/json"
	"fmt"

	"github.com/gennadyterekhov/tribal-relations-core/entity"
)

func main() {
	fmt.Println("{")

	for key, situation := range entity.SituationNameToSituationMap {
		fmt.Printf("\"%v\":", key)

		jsonResource, err := json.Marshal(situation)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(jsonResource))
		fmt.Println(",")

	}
	fmt.Println("}")

}
