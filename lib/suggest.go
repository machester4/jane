package lib

import (
	"fmt"
	"github.com/machester4/jane/sanatizer"
)

/*Suggest Get words sugestions from text*/
func Suggest(text string, dict string, rules string) map[string]string {
	// dictionary := NewDictionary(dict)
	// Sanatize text
	sanatizer.Sanatize(&text)

	// Build chain

	fmt.Println("value after sanatize", &text)
	return nil
}
