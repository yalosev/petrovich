package rules

// For generation

// Rule for petrovich-rules generation
type Rule struct {
	Gender Gender
	Test   []string
	Mods   []string
}

// RulesSet set of generated rules
type RulesSet []Rule

// RulesGroup group of rules
type RulesGroup struct {
	Exceptions RulesSet
	Suffixes   RulesSet
}

// Rules top level petrovich-rules
type Rules struct {
	Lastname   RulesGroup
	Firstname  RulesGroup
	Middlename RulesGroup
}
