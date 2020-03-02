package lib

import (
	"io/ioutil"
	"os"
	"strings"

	errormanager "github.com/machester4/jane/utils"
)

type Dictionary struct {
	Lang  string
	Words []string
}

func NewDictionary(lang string) (dic *Dictionary) {
	dic = &Dictionary{
		Lang: lang,
	}

	path, err := os.Getwd()
	errormanager.Check(err)

	words, err := ioutil.ReadFile(path + "/dict/" + lang + ".dic")
	errormanager.Check(err)

	dic.Words = strings.Split(string(words), "\n")

	return dic
}
