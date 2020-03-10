package helpers

import (
	"io/ioutil"
	"os"
	"strings"
)

func CheckError(err error)  {
	if err != nil {
		panic(err)
	}
}


func GetDictionary(lang string) []string {
	path, err := os.Getwd()
	CheckError(err)

	words, err := ioutil.ReadFile(path + "/dict/" + lang + ".dic")
	CheckError(err)

	return strings.Split(string(words), "\n")
}