package petrovich_test

import (
	"fmt"

	"github.com/yalosev/petrovich"
)

func ExampleTransform() {
	people := []petrovich.Person{
		{
			Gender:     petrovich.Male,
			FirstName:  "Сергей",
			MiddleName: "Иванович",
			LastName:   "Прокопенко",
		},
		{
			FirstName:  "Андрей",
			MiddleName: "Павлович",
			LastName:   "Николаев",
		},
		{
			Gender:     petrovich.Female,
			FirstName:  "Екатерина",
			MiddleName: "Алексеевна",
			LastName:   "Морозова",
		},
		{
			FirstName:  "Антонина",
			MiddleName: "Сергеевна",
			LastName:   "Павлова",
		},
	}

	for _, p := range people {
		fmt.Println(petrovich.Transform(p, petrovich.Dative))
	}
	// Output:
	// Прокопенко Сергею Ивановичу
	// Николаеву Андрею Павловичу
	// Морозовой Екатерине Алексеевне
	// Павловой Антонине Сергеевне
}

func ExampleTransformShort() {
	people := []petrovich.Person{
		{
			Gender:     petrovich.Male,
			FirstName:  "Сергей",
			MiddleName: "Иванович",
			LastName:   "Прокопенко",
		},
		{
			FirstName:  "Андрей",
			MiddleName: "Павлович",
			LastName:   "Николаев",
		},
		{
			Gender:     petrovich.Female,
			FirstName:  "Екатерина",
			MiddleName: "Алексеевна",
			LastName:   "Морозова",
		},
		{
			FirstName:  "Антонина",
			MiddleName: "Сергеевна",
			LastName:   "Павлова",
		},
	}

	for _, p := range people {
		newp := petrovich.Transform(p, petrovich.Instrumental)
		fmt.Println(newp.Short())
	}
	// Output:
	// Прокопенко С.И.
	// Николаевым А.П.
	// Морозовой Е.А.
	// Павловой А.С.
}
