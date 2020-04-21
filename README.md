Go port of petrovich library(https://github.com/petrovich) which inflects Russian names to a given grammatical case

## Installation
```
go get github.com/yalosev/petrovich
```

## Usage

```go
package main

import (
    "fmt"
    "github.com/yalosev/petrovich"
)

func main() {
	ivan := petrovich.Person{
		FirstName: "Иван",
		MiddleName: "Иванович",
		LastName: "Иванов",
		Gender: petrovich.Male,
	}
	result := petrovich.Transform(ivan, petrovich.Genitive)
	
	fmt.Println(result) // Output: Иванова Ивана Ивановича
	fmt.Println(result.Short()) // Output: Иванова И. И.

    // OR
    petr := petrovich.NewPerson("Басов", "Петр", "Андреевич") 

    fmt.Println(petrovich.Transform(petr, petrovich.Prepositional)) // Output: Басове Петре Андреевиче
}
```

---
Possible values for gender are `petrovich.Androgynous`, `petrovich.Male` and `petrovich.Female`.


Full list of grammatical cases is in the table below.

| Case                      | Case (in Russian) | Question (in Russian)  |
|---------------------------|-------------------|------------------------|
| `petrovich.Nominative`    | Именительный      | Кто?                   |
| `petrovich.Genitive`      | Родительный       | Кого?                  |
| `petrovich.Dative`        | Дательный         | Кому?                  |
| `petrovich.Accusative`    | Винительный       | Кого?                  |
| `petrovich.Instrumental`  | Творительный      | Кем?                   |
| `petrovich.Prepositional` | Предложный        | О ком?                 |