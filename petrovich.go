package petrovich

import (
	"bytes"
	"strings"
	"unicode/utf8"

	"github.com/yalosev/petrovich/rules"
)

//go:generate git submodule update --recursive --remote
//go:generate go run cmd/generator.go
//go:generate go fmt ./...

// Gender specifies Person`s gender - Male, Female, Androgynous
type Gender = rules.Gender

// Case specifies grammatical case (падежи)
type Case = rules.Case

const (
	// Androgynous gender
	Androgynous = rules.Androgynous
	// Male gender
	Male = rules.Male
	// Female gender
	Female = rules.Female

	// cases

	// Nominative именительный
	Nominative = rules.Nominative
	// Genitive родительный
	Genitive = rules.Genitive
	// Dative дательный
	Dative = rules.Dative
	// Accusative винительный
	Accusative = rules.Accusative
	// Instrumental творительный
	Instrumental = rules.Instrumental
	// Prepositional предложный
	Prepositional = rules.Prepositional
)

// Person contains firstname, lastname, middlename and gender
type Person struct {
	Gender     Gender
	FirstName  string
	MiddleName string
	LastName   string
}

// NewPerson returns Person. Grammary cases based on next sequence: gender -> middlename -> firstname
func NewPerson(lastname, firstname, middlename string, gender ...Gender) Person {
	p := Person{
		FirstName:  firstname,
		MiddleName: middlename,
		LastName:   lastname,
	}
	if len(gender) > 0 {
		p.Gender = gender[0]
	}

	return p
}

// String returns full person`s name
func (p Person) String() string {
	var buf bytes.Buffer

	if len(p.LastName) > 0 {
		buf.WriteString(strings.Title(p.LastName))
	}

	if len(p.FirstName) > 0 {
		buf.WriteString(" ")
		buf.WriteString(strings.Title(p.FirstName))

		if len(p.MiddleName) > 0 {
			buf.WriteString(" ")
			buf.WriteString(strings.Title(p.MiddleName))
		}
	}

	return strings.TrimSpace(buf.String())
}

// Short returns person`s lastname and firstname/middlename cutted to the first letter
func (p Person) Short() string {
	var buf bytes.Buffer

	if len(p.LastName) > 0 {
		buf.WriteString(strings.Title(p.LastName))
	}

	if len(p.FirstName) > 0 {
		buf.WriteString(" ")
		buf.WriteString(string([]rune(strings.Title(p.FirstName))[:1]))
		buf.WriteString(".")

		if len(p.MiddleName) > 0 {
			buf.WriteString(string([]rune(strings.Title(p.MiddleName))[:1]))
			buf.WriteString(".")
		}
	}

	return strings.TrimSpace(buf.String())
}

// Transform person credentials to a given grammatical case
func Transform(person Person, cas rules.Case) Person {
	if person.Gender == 0 && (len(person.MiddleName) > 0 || len(person.FirstName) > 0) {
		person.Gender = person.detectGender()
	}

	newPerson := Person{
		Gender:     person.Gender,
		LastName:   inflect(person.LastName, person.Gender, cas, rules.AllRules.Lastname),
		MiddleName: inflect(person.MiddleName, person.Gender, cas, rules.AllRules.Middlename),
		FirstName:  inflect(person.FirstName, person.Gender, cas, rules.AllRules.Firstname),
	}

	return newPerson
}

func (p Person) detectGender() Gender {
	gender := p.detectGenderByMiddlename()
	if gender != Androgynous {
		return gender
	}

	return p.detectGenderByName()
}

func (p Person) detectGenderByMiddlename() Gender {
	if len(p.MiddleName) == 0 {
		return Androgynous
	}

	ending := strings.ToLower(p.MiddleName)

	if strings.HasSuffix(ending, "ич") {
		return Male
	} else if strings.HasSuffix(ending, "на") {
		return Female
	}

	return Androgynous
}

func (p Person) detectGenderByName() Gender {
	val, ok := rules.AllNames[strings.ToLower(p.FirstName)]
	if !ok {
		return Androgynous
	}

	return val
}

func inflect(name string, g Gender, c rules.Case, rg rules.RulesGroup) string {
	if len(name) == 0 {
		return ""
	}
	if c == rules.Nominative {
		return name
	}

	ns := strings.Split(name, "-")
	r := make([]string, len(ns))

	for i, n := range ns {
		if e := checkException(n, g, c, rg); e != "" {
			r[i] = e
		} else {
			r[i] = find(n, g, c, rg)
		}
	}

	return strings.Join(r, "-")
}

func find(name string, g Gender, c rules.Case, rg rules.RulesGroup) string {
	for _, r := range rg.Suffixes {
		if r.Gender == Androgynous || r.Gender == g {
			for _, t := range r.Test {
				if strings.HasSuffix(name, t) && len(r.Mods) > int(c) {
					return applyRule(name, r.Mods[c])
				}
			}
		}
	}

	return name
}

func checkException(name string, g Gender, c rules.Case, rg rules.RulesGroup) string {
	ln := strings.ToLower(name)

	for _, r := range rg.Exceptions {
		if r.Gender == Androgynous || r.Gender == g {
			for _, t := range r.Test {
				if t == ln && len(r.Mods) > int(c) {
					return applyRule(name, r.Mods[c])
				}
			}
		}
	}

	return ""
}

func applyRule(name, mod string) string {
	if mod == "." {
		return name
	}

	if i := strings.LastIndex(mod, "-"); i >= 0 {
		l := utf8.RuneCountInString(name) - i - 1
		return string([]rune(name)[:l]) + mod[i+1:]
	}

	return name + mod
}
