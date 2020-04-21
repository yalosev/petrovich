package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/yalosev/petrovich/rules"
)

func main() {
	generateRules()
	generateDict()
}

type Row struct {
	Name string `json:"Name"`
	Sex  string `json:"Sex"`
}

type NamesDict []Row

func generateDict() {
	f, err := ioutil.ReadFile("cmd/russian_names.json")
	if err != nil {
		panic(err)
	}

	f = bytes.TrimPrefix(f, []byte("\xef\xbb\xbf")) // отпилить BOM

	var rows NamesDict

	err = json.Unmarshal(f, &rows)
	if err != nil {
		panic(err)
	}

	out, err := os.Create("rules/generated_names.go")
	if err != nil {
		panic(err)
	}

	defer out.Close()

	_, _ = fmt.Fprint(out, "// DO NOT EDIT!\n// Code generated from cmd/generator.go\n\npackage rules\n\n")

	fmt.Fprint(out, "var AllNames = map[string]Gender{\n")

	for _, row := range rows {
		var genderType string
		switch row.Sex {
		case "Ж":
			genderType = fmt.Sprintf("%s", strings.Title(rules.Female.String()))

		case "М":
			genderType = fmt.Sprintf("%s", strings.Title(rules.Male.String()))

		default:
			genderType = fmt.Sprintf("%s", strings.Title(rules.Androgynous.String()))
		}

		fmt.Fprintf(out, "\t%q: %s,\n", strings.ToLower(row.Name), genderType)
	}

	fmt.Fprint(out, "}\n")
}

func generateRules() {
	f, err := ioutil.ReadFile("petrovich-rules/rules.json")
	if err != nil {
		panic(err)
	}

	var petrovich rules.Rules

	err = json.Unmarshal(f, &petrovich)
	if err != nil {
		panic(err)
	}

	out, err := os.Create("rules/generated_rules.go")
	if err != nil {
		panic(err)
	}

	defer out.Close()

	_, _ = fmt.Fprint(out, "// DO NOT EDIT!\n// Code generated from cmd/generator.go\n\npackage rules\n\n")

	fmt.Fprint(out, "var AllRules = Rules{\n\tFirstname: RulesGroup{\n")
	printRulesGroup(out, petrovich.Firstname)

	fmt.Fprint(out, "\t},\n\tMiddlename: RulesGroup{\n")
	printRulesGroup(out, petrovich.Middlename)

	fmt.Fprint(out, "\t},\n\tLastname: RulesGroup{\n")
	printRulesGroup(out, petrovich.Lastname)

	fmt.Fprint(out, "\t},\n}\n")
}

func printRulesGroup(o io.Writer, g rules.RulesGroup) {
	fmt.Fprint(o, "\t\tExceptions: []Rule{\n")
	printRulesSet(o, g.Exceptions)
	fmt.Fprint(o, "\t\t},\n")

	fmt.Fprint(o, "\t\tSuffixes: []Rule{\n")
	printRulesSet(o, g.Suffixes)
	fmt.Fprint(o, "\t\t},\n")
}

func printRulesSet(o io.Writer, s rules.RulesSet) {
	for _, r := range s {
		genderType := fmt.Sprintf("%s", strings.Title(r.Gender.String()))

		fmt.Fprintf(o, "\t\t\t{\n\t\t\t\tGender: %s,\n\t\t\t\tTest: []string{\n", genderType)
		for _, t := range r.Test {
			fmt.Fprintf(o, "\t\t\t\t\t\"%s\",\n", t)
		}
		fmt.Fprint(o, "\t\t\t\t},\n\t\t\t\tMods: []string{\n")
		for _, t := range r.Mods {
			fmt.Fprintf(o, "\t\t\t\t\t\"%s\",\n", t)
		}
		fmt.Fprint(o, "\t\t\t\t},\n\t\t\t},\n")
	}
}
