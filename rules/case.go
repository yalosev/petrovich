package rules

//go:generate enumer -json -transform=snake -type=Case -output=case_string.go case.go

// Case grammatical case (падежи)
type Case int

const (
	// Nominative именительный
	Nominative Case = -1 + iota
	// Genitive родительный
	Genitive
	// Dative дательный
	Dative
	// Accusative винительный
	Accusative
	// Instrumental творительный
	Instrumental
	// Prepositional предложный
	Prepositional
)
