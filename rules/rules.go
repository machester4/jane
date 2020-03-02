package rules

import (
	"errors"
	"strings"
	"unicode"

	"github.com/machester4/jane/constants"
)

var (
	errorInvalidLang = errors.New("invalid lang")
)

type Facade struct {
	lang     string
	inPunct  []rune
	outPunct []rune
}

func (f *Facade) PunctSanitize(text *string) *string {
	for _, p := range f.inPunct {
		*text = strings.ReplaceAll(*text, string(p)+" ", string(p))
	}

	for _, p := range f.outPunct {
		*text = strings.ReplaceAll(*text, " "+string(p), string(p))
	}
	return text
}

func New(lang string) Facade {
	facade := Facade{
		lang: lang,
	}

	switch lang {
	case "es":
		facade.inPunct = constants.InPunctuationMarksEs
		facade.outPunct = constants.OutPunctuationMarksEs
	default:
		panic(errorInvalidLang)
	}

	return facade
}

func PreventPunctCollision(prevCharacter *rune, currentCharacter rune) rune {
	if unicode.IsPunct(*prevCharacter) && unicode.IsPunct(currentCharacter) {
		return rune(0)
	}
	*prevCharacter = currentCharacter
	return currentCharacter
}
