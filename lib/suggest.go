package lib

import (
	"fmt"
)

/*Suggest Get words sugestions from text*/
func Suggest(text string, dict string, rules string) map[string]string {
	// dictionary := NewDictionary(dict)
	// Sanatize text
	Sanatize(&text)

	// Build chain

	fmt.Println("value after sanatize", &text)
	return nil
}
