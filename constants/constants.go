package constants

// ES
var InPunctuationMarksEs = []rune{'¿', '¡'}
var OutPunctuationMarksEs = []rune{'?', '!', '.', ',', ';', ':'}

// Chain block types
const BlockTypeLetter = "LETTER"
const BlockTypePunct = "PUNCT"
const BlockTypeSpace = "SPACE"

// Chain block repeat max consecutively
// If the value is 0 the character will not be repeated consecutively
const MaxRepeatLeter = 1
const MaxRepeatPunct = 0
const MaxRepeatSpace = 0
