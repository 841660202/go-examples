package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	goodJSON := `{"example": 1}`
	badJSON := `{"example":2:]}}`

	fmt.Println("goodJSON: ")
	fmt.Println(json.Valid([]byte(goodJSON)))

	fmt.Println("badJSON: ")
	fmt.Println(json.Valid([]byte(badJSON)))
}
