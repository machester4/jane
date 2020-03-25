package lib

// articles available in the Spanish language
var articles = []string{"el", "la", "los", "las", "un", "uno", "una", "unos", "unas", "lo"}

// fieldTypeLetter - category of a character
const fieldTypeLetter = "LETTER"

// fieldTypePunct - category of a character
const fieldTypePunct = "PUNCT"

// fieldTypeSpace - category of a character
const fieldTypeSpace = "SPACE"

// If the value is 0 the character will not be repeated consecutively

// maxRepeatLetter - maximum amount that a letter can be repeated consecutively
const maxRepeatLetter = 1

// maxRepeatPunct - maximum amount that a punctuation mark can be repeated consecutively
const maxRepeatPunct = 0

// maxRepeatSpace - maximum amount that a space can be repeated consecutively
const maxRepeatSpace = 0

// maxDistanceInDic - maximum distance tolerated in a language dictionary
const maxDistanceInDic = 1

// maxDistanceInContext - maximum distance tolerated in a context dictionary
const maxDistanceInContext = 3

// maxResults - maximum recommendations per field
const maxResults = 3
