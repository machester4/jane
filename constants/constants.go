package constants

// Spanish only
var InPunctuationMarksEs = []rune{'¿', '¡'}
var OutPunctuationMarksEs = []rune{'?', '!', '.', ',', ';', ':'}
var Articles = []string{"el", "la", "los", "las", "un", "uno", "una", "unos", "unas", "lo"}

// Chain fields types
const FieldTypeLetter = "LETTER"
const FieldTypePunct = "PUNCT"
const FieldTypeSpace = "SPACE"

// Field character repeat max consecutively
// If the value is 0 the character will not be repeated consecutively
const MaxRepeatLetter = 1
const MaxRepeatPunct = 0
const MaxRepeatSpace = 0

// Recommend
const MaxDistance = 3
const MaxResults = 3
