package lib

// Spanish only
var inPunctuationMarksEs = []rune{'¿', '¡'}
var outPunctuationMarksEs = []rune{'?', '!', '.', ',', ';', ':'}
var articles = []string{"el", "la", "los", "las", "un", "uno", "una", "unos", "unas", "lo"}

// Chain fields types
const fieldTypeLetter = "LETTER"
const fieldTypePunct = "PUNCT"
const fieldTypeSpace = "SPACE"

// Field character repeat max consecutively
// If the value is 0 the character will not be repeated consecutively
const maxRepeatLetter = 1
const maxRepeatPunct = 0
const maxRepeatSpace = 0

// Recommend
const maxDistanceInDic = 2
const maxDistanceInContext = 3
const maxResults = 3
