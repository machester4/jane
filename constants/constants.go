package constants

// ES
var InPunctuationMarksEs = []rune{'¿', '¡'}
var OutPunctuationMarksEs = []rune{'?', '!', '.', ',', ';', ':'}

// Chain block types
const BLOCK_TYPE_LETTER = "LETTER"
const BLOCK_TYPE_PUNCT = "PUNCT"
const BLOCK_TYPE_SPACE = "SPACE"

// Chain block repeat max
const MAX_REPEAT_LETER = 2
const MAX_REPEAT_PUNCT = 1
const MAX_REPEAT_SPACE = 1