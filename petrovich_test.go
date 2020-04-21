package petrovich

import (
	"fmt"
	"testing"

	"github.com/yalosev/petrovich/rules"
)

func TestPerson_String(t *testing.T) {
	testcases := map[Person]string{
		{FirstName: "Иван", MiddleName: "Иванович", LastName: "Иванов"}:        "Иванов Иван Иванович",
		{FirstName: "Антонина", MiddleName: "павловна", LastName: "Рябушкина"}: "Рябушкина Антонина Павловна",
		{FirstName: "Екатерина", LastName: "Морозова"}:                         "Морозова Екатерина",
		{FirstName: "Василий", LastName: "Петров"}:                             "Петров Василий",
	}

	for p, result := range testcases {
		if p.String() != result {
			fmt.Printf("Expected: %s, got: %s", result, p.String())
			t.Fail()
		}
	}
}

func TestPerson_Short(t *testing.T) {
	testcases := map[Person]string{
		{FirstName: "Иван", MiddleName: "Иванович", LastName: "Иванов"}:        "Иванов И.И.",
		{FirstName: "Антонина", MiddleName: "Павловна", LastName: "Рябушкина"}: "Рябушкина А.П.",
		{FirstName: "Екатерина", LastName: "Морозова"}:                         "Морозова Е.",
		{FirstName: "Василий", LastName: "Петров"}:                             "Петров В.",
	}

	for p, result := range testcases {
		if p.Short() != result {
			fmt.Printf("Expected: %s, got: %s", result, p.Short())
			t.Fail()
		}
	}
}

func TestDetectGender(t *testing.T) {
	testcases := map[Person]rules.Gender{
		{FirstName: "Иван", MiddleName: "Иванович", LastName: "Иванов"}:        rules.Male,
		{FirstName: "Антонина", MiddleName: "Павловна", LastName: "Рябушкина"}: rules.Female,
		{FirstName: "Екатерина", LastName: "Морозова"}:                         rules.Female,
		{FirstName: "Катя", LastName: "Морозова"}:                              rules.Female,
		{FirstName: "Василий", LastName: "Петров"}:                             rules.Male,
		{FirstName: "Ахрыгынын", LastName: "Педальный"}:                        rules.Androgynous,
	}

	for p, g := range testcases {
		res := p.detectGender()
		if res != g {
			fmt.Printf("Expected: %s, got: %s", g, res)
			t.Fail()
		}
	}
}
