package rules

//go:generate enumer -json -transform=snake -type=Gender -output=gender_string.go gender.go

// Gender specify Person`s gender
type Gender int8

const (
	Androgynous Gender = iota
	Male
	Female
)
