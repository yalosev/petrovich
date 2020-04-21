package rules

type Rule struct {
	Gender Gender
	Test   []string
	Mods   []string
}

type RulesSet []Rule

type RulesGroup struct {
	Exceptions RulesSet
	Suffixes   RulesSet
}

// верхний уровень
type Rules struct {
	Lastname   RulesGroup
	Firstname  RulesGroup
	Middlename RulesGroup
}
