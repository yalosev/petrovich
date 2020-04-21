// DO NOT EDIT!
// Code generated from cmd/generator.go

package rules

// AllRules presents all geenerated petrovich rules
var AllRules = Rules{
	Firstname: RulesGroup{
		Exceptions: []Rule{
			{
				Gender: Male,
				Test: []string{
					"лев",
				},
				Mods: []string{
					"--ьва",
					"--ьву",
					"--ьва",
					"--ьвом",
					"--ьве",
				},
			},
			{
				Gender: Male,
				Test: []string{
					"пётр",
				},
				Mods: []string{
					"---етра",
					"---етру",
					"---етра",
					"---етром",
					"---етре",
				},
			},
			{
				Gender: Male,
				Test: []string{
					"павел",
				},
				Mods: []string{
					"--ла",
					"--лу",
					"--ла",
					"--лом",
					"--ле",
				},
			},
			{
				Gender: Male,
				Test: []string{
					"яша",
				},
				Mods: []string{
					"-и",
					"-е",
					"-у",
					"-ей",
					"-е",
				},
			},
			{
				Gender: Male,
				Test: []string{
					"илья",
				},
				Mods: []string{
					"-и",
					"-е",
					"-ю",
					"-ёй",
					"-е",
				},
			},
			{
				Gender: Male,
				Test: []string{
					"шота",
				},
				Mods: []string{
					".",
					".",
					".",
					".",
					".",
				},
			},
			{
				Gender: Female,
				Test: []string{
					"агидель",
					"жизель",
					"нинель",
					"рашель",
					"рахиль",
				},
				Mods: []string{
					"-и",
					"-и",
					".",
					"ю",
					"-и",
				},
			},
		},
		Suffixes: []Rule{
			{
				Gender: Androgynous,
				Test: []string{
					"е",
					"ё",
					"и",
					"о",
					"у",
					"ы",
					"э",
					"ю",
				},
				Mods: []string{
					".",
					".",
					".",
					".",
					".",
				},
			},
			{
				Gender: Male,
				Test: []string{
					"уа",
					"иа",
				},
				Mods: []string{
					".",
					".",
					".",
					".",
					".",
				},
			},
			{
				Gender: Female,
				Test: []string{
					"б",
					"в",
					"г",
					"д",
					"ж",
					"з",
					"й",
					"к",
					"л",
					"м",
					"н",
					"п",
					"р",
					"с",
					"т",
					"ф",
					"х",
					"ц",
					"ч",
					"ш",
					"щ",
					"ъ",
					"иа",
					"ль",
				},
				Mods: []string{
					".",
					".",
					".",
					".",
					".",
				},
			},
			{
				Gender: Female,
				Test: []string{
					"ь",
				},
				Mods: []string{
					"-и",
					"-и",
					".",
					"ю",
					"-и",
				},
			},
			{
				Gender: Male,
				Test: []string{
					"ь",
				},
				Mods: []string{
					"-я",
					"-ю",
					"-я",
					"-ем",
					"-е",
				},
			},
			{
				Gender: Androgynous,
				Test: []string{
					"га",
					"ка",
					"ха",
					"ча",
					"ща",
					"жа",
				},
				Mods: []string{
					"-и",
					"-е",
					"-у",
					"-ой",
					"-е",
				},
			},
			{
				Gender: Female,
				Test: []string{
					"ша",
				},
				Mods: []string{
					"-и",
					"-е",
					"-у",
					"-ей",
					"-е",
				},
			},
			{
				Gender: Male,
				Test: []string{
					"ша",
					"ча",
					"жа",
				},
				Mods: []string{
					"-и",
					"-е",
					"-у",
					"-ей",
					"-е",
				},
			},
			{
				Gender: Androgynous,
				Test: []string{
					"а",
				},
				Mods: []string{
					"-ы",
					"-е",
					"-у",
					"-ой",
					"-е",
				},
			},
			{
				Gender: Female,
				Test: []string{
					"ка",
					"га",
					"ха",
				},
				Mods: []string{
					"-и",
					"-е",
					"-у",
					"-ой",
					"-е",
				},
			},
			{
				Gender: Female,
				Test: []string{
					"ца",
				},
				Mods: []string{
					"-ы",
					"-е",
					"-у",
					"-ей",
					"-е",
				},
			},
			{
				Gender: Female,
				Test: []string{
					"а",
				},
				Mods: []string{
					"-ы",
					"-е",
					"-у",
					"-ой",
					"-е",
				},
			},
			{
				Gender: Female,
				Test: []string{
					"ия",
				},
				Mods: []string{
					"-и",
					"-и",
					"-ю",
					"-ей",
					"-и",
				},
			},
			{
				Gender: Androgynous,
				Test: []string{
					"я",
				},
				Mods: []string{
					"-и",
					"-е",
					"-ю",
					"-ей",
					"-е",
				},
			},
			{
				Gender: Male,
				Test: []string{
					"ий",
				},
				Mods: []string{
					"-я",
					"-ю",
					"-я",
					"-ем",
					"-и",
				},
			},
			{
				Gender: Male,
				Test: []string{
					"ей",
					"й",
				},
				Mods: []string{
					"-я",
					"-ю",
					"-я",
					"-ем",
					"-е",
				},
			},
			{
				Gender: Male,
				Test: []string{
					"ш",
					"ж",
				},
				Mods: []string{
					"а",
					"у",
					"а",
					"ем",
					"е",
				},
			},
			{
				Gender: Male,
				Test: []string{
					"б",
					"в",
					"г",
					"д",
					"ж",
					"з",
					"к",
					"л",
					"м",
					"н",
					"п",
					"р",
					"с",
					"т",
					"ф",
					"х",
					"ц",
					"ч",
				},
				Mods: []string{
					"а",
					"у",
					"а",
					"ом",
					"е",
				},
			},
			{
				Gender: Androgynous,
				Test: []string{
					"ния",
					"рия",
					"вия",
				},
				Mods: []string{
					"-и",
					"-и",
					"-ю",
					"-ем",
					"-ем",
				},
			},
		},
	},
	Middlename: RulesGroup{
		Exceptions: []Rule{
			{
				Gender: Androgynous,
				Test: []string{
					"борух",
				},
				Mods: []string{
					".",
					".",
					".",
					".",
					".",
				},
			},
		},
		Suffixes: []Rule{
			{
				Gender: Male,
				Test: []string{
					"мич",
					"ьич",
					"кич",
				},
				Mods: []string{
					"а",
					"у",
					"а",
					"ом",
					"е",
				},
			},
			{
				Gender: Male,
				Test: []string{
					"ич",
				},
				Mods: []string{
					"а",
					"у",
					"а",
					"ем",
					"е",
				},
			},
			{
				Gender: Female,
				Test: []string{
					"на",
				},
				Mods: []string{
					"-ы",
					"-е",
					"-у",
					"-ой",
					"-е",
				},
			},
		},
	},
	Lastname: RulesGroup{
		Exceptions: []Rule{
			{
				Gender: Androgynous,
				Test: []string{
					"бонч",
					"абдул",
					"белиц",
					"гасан",
					"дюссар",
					"дюмон",
					"книппер",
					"корвин",
					"ван",
					"шолом",
					"тер",
					"призван",
					"мелик",
					"вар",
					"фон",
				},
				Mods: []string{
					".",
					".",
					".",
					".",
					".",
				},
			},
			{
				Gender: Androgynous,
				Test: []string{
					"дюма",
					"тома",
					"дега",
					"люка",
					"ферма",
					"гамарра",
					"петипа",
					"шандра",
					"скаля",
					"каруана",
				},
				Mods: []string{
					".",
					".",
					".",
					".",
					".",
				},
			},
			{
				Gender: Androgynous,
				Test: []string{
					"гусь",
					"ремень",
					"камень",
					"онук",
					"богода",
					"нечипас",
					"долгопалец",
					"маненок",
					"рева",
					"кива",
				},
				Mods: []string{
					".",
					".",
					".",
					".",
					".",
				},
			},
			{
				Gender: Male,
				Test: []string{
					"вий",
					"сой",
					"цой",
					"хой",
				},
				Mods: []string{
					"-я",
					"-ю",
					"-я",
					"-ем",
					"-е",
				},
			},
			{
				Gender: Male,
				Test: []string{
					"грин",
					"дарвин",
					"регин",
					"цин",
				},
				Mods: []string{
					"а",
					"у",
					"а",
					"ом",
					"е",
				},
			},
		},
		Suffixes: []Rule{
			{
				Gender: Female,
				Test: []string{
					"б",
					"в",
					"г",
					"д",
					"ж",
					"з",
					"й",
					"к",
					"л",
					"м",
					"н",
					"п",
					"р",
					"с",
					"т",
					"ф",
					"х",
					"ц",
					"ч",
					"ш",
					"щ",
					"ъ",
					"ь",
				},
				Mods: []string{
					".",
					".",
					".",
					".",
					".",
				},
			},
			{
				Gender: Androgynous,
				Test: []string{
					"орота",
				},
				Mods: []string{
					".",
					".",
					".",
					".",
					".",
				},
			},
			{
				Gender: Female,
				Test: []string{
					"ска",
					"цка",
				},
				Mods: []string{
					"-ой",
					"-ой",
					"-ую",
					"-ой",
					"-ой",
				},
			},
			{
				Gender: Female,
				Test: []string{
					"чая",
				},
				Mods: []string{
					"--ей",
					"--ей",
					"--ую",
					"--ей",
					"--ей",
				},
			},
			{
				Gender: Male,
				Test: []string{
					"чий",
				},
				Mods: []string{
					"--его",
					"--ему",
					"--его",
					"--им",
					"--ем",
				},
			},
			{
				Gender: Female,
				Test: []string{
					"цкая",
					"ская",
					"ная",
					"ая",
				},
				Mods: []string{
					"--ой",
					"--ой",
					"--ую",
					"--ой",
					"--ой",
				},
			},
			{
				Gender: Female,
				Test: []string{
					"яя",
				},
				Mods: []string{
					"--ей",
					"--ей",
					"--юю",
					"--ей",
					"--ей",
				},
			},
			{
				Gender: Male,
				Test: []string{
					"иной",
					"уй",
				},
				Mods: []string{
					"-я",
					"-ю",
					"-я",
					"-ем",
					"-е",
				},
			},
			{
				Gender: Androgynous,
				Test: []string{
					"ца",
				},
				Mods: []string{
					"-ы",
					"-е",
					"-у",
					"-ей",
					"-е",
				},
			},
			{
				Gender: Male,
				Test: []string{
					"рих",
				},
				Mods: []string{
					"а",
					"у",
					"а",
					"ом",
					"е",
				},
			},
			{
				Gender: Androgynous,
				Test: []string{
					"ия",
				},
				Mods: []string{
					"-и",
					"-и",
					"-ю",
					"-ей",
					"-и",
				},
			},
			{
				Gender: Androgynous,
				Test: []string{
					"иа",
					"аа",
					"оа",
					"уа",
					"ыа",
					"еа",
					"юа",
					"эа",
				},
				Mods: []string{
					".",
					".",
					".",
					".",
					".",
				},
			},
			{
				Gender: Androgynous,
				Test: []string{
					"о",
					"е",
					"э",
					"и",
					"ы",
					"у",
					"ю",
				},
				Mods: []string{
					".",
					".",
					".",
					".",
					".",
				},
			},
			{
				Gender: Male,
				Test: []string{
					"их",
					"ых",
				},
				Mods: []string{
					".",
					".",
					".",
					".",
					".",
				},
			},
			{
				Gender: Female,
				Test: []string{
					"ова",
					"ева",
					"на",
					"ёва",
				},
				Mods: []string{
					"-ой",
					"-ой",
					"-у",
					"-ой",
					"-ой",
				},
			},
			{
				Gender: Androgynous,
				Test: []string{
					"га",
					"ка",
					"ха",
					"ча",
					"ща",
					"жа",
					"ша",
				},
				Mods: []string{
					"-и",
					"-е",
					"-у",
					"-ой",
					"-е",
				},
			},
			{
				Gender: Androgynous,
				Test: []string{
					"а",
				},
				Mods: []string{
					"-ы",
					"-е",
					"-у",
					"-ой",
					"-е",
				},
			},
			{
				Gender: Male,
				Test: []string{
					"ь",
				},
				Mods: []string{
					"-я",
					"-ю",
					"-я",
					"-ем",
					"-е",
				},
			},
			{
				Gender: Androgynous,
				Test: []string{
					"я",
				},
				Mods: []string{
					"-и",
					"-е",
					"-ю",
					"-ей",
					"-е",
				},
			},
			{
				Gender: Male,
				Test: []string{
					"обей",
				},
				Mods: []string{
					"--ья",
					"--ью",
					"--ья",
					"--ьем",
					"--ье",
				},
			},
			{
				Gender: Male,
				Test: []string{
					"ей",
				},
				Mods: []string{
					"-я",
					"-ю",
					"-я",
					"-ем",
					"-е",
				},
			},
			{
				Gender: Male,
				Test: []string{
					"ян",
					"ан",
					"йн",
				},
				Mods: []string{
					"а",
					"у",
					"а",
					"ом",
					"е",
				},
			},
			{
				Gender: Male,
				Test: []string{
					"ынец",
				},
				Mods: []string{
					"--ца",
					"--цу",
					"--ца",
					"--цом",
					"--це",
				},
			},
			{
				Gender: Male,
				Test: []string{
					"нец",
					"робец",
				},
				Mods: []string{
					"--ца",
					"--цу",
					"--ца",
					"--цем",
					"--це",
				},
			},
			{
				Gender: Male,
				Test: []string{
					"ай",
				},
				Mods: []string{
					"-я",
					"-ю",
					"-я",
					"-ем",
					"-е",
				},
			},
			{
				Gender: Male,
				Test: []string{
					"гой",
					"кой",
				},
				Mods: []string{
					"-го",
					"-му",
					"-го",
					"--им",
					"-м",
				},
			},
			{
				Gender: Male,
				Test: []string{
					"ой",
				},
				Mods: []string{
					"-го",
					"-му",
					"-го",
					"--ым",
					"-м",
				},
			},
			{
				Gender: Male,
				Test: []string{
					"ах",
					"ив",
					"шток",
				},
				Mods: []string{
					"а",
					"у",
					"а",
					"ом",
					"е",
				},
			},
			{
				Gender: Male,
				Test: []string{
					"ший",
					"щий",
					"жий",
					"ний",
				},
				Mods: []string{
					"--его",
					"--ему",
					"--его",
					"-м",
					"--ем",
				},
			},
			{
				Gender: Male,
				Test: []string{
					"ый",
					"кий",
					"хий",
				},
				Mods: []string{
					"--ого",
					"--ому",
					"--ого",
					"-м",
					"--ом",
				},
			},
			{
				Gender: Male,
				Test: []string{
					"ий",
				},
				Mods: []string{
					"-я",
					"-ю",
					"-я",
					"-ем",
					"-и",
				},
			},
			{
				Gender: Male,
				Test: []string{
					"ок",
				},
				Mods: []string{
					"--ка",
					"--ку",
					"--ка",
					"--ком",
					"--ке",
				},
			},
			{
				Gender: Male,
				Test: []string{
					"обец",
					"швец",
					"ьвец",
				},
				Mods: []string{
					"а",
					"у",
					"а",
					"ем",
					"е",
				},
			},
			{
				Gender: Male,
				Test: []string{
					"аец",
					"иец",
					"еец",
				},
				Mods: []string{
					"--йца",
					"--йцу",
					"--йца",
					"--йцем",
					"--йце",
				},
			},
			{
				Gender: Male,
				Test: []string{
					"опец",
				},
				Mods: []string{
					"--ца",
					"--цу",
					"--ца",
					"--цем",
					"--це",
				},
			},
			{
				Gender: Male,
				Test: []string{
					"вец",
					"убец",
					"ырец",
				},
				Mods: []string{
					"--ца",
					"--цу",
					"--ца",
					"--цом",
					"--це",
				},
			},
			{
				Gender: Male,
				Test: []string{
					"ц",
					"ч",
					"ш",
					"щ",
				},
				Mods: []string{
					"а",
					"у",
					"а",
					"ем",
					"е",
				},
			},
			{
				Gender: Male,
				Test: []string{
					"ен",
					"нн",
					"он",
					"ун",
					"б",
					"г",
					"д",
					"ж",
					"з",
					"к",
					"л",
					"м",
					"п",
					"р",
					"с",
					"т",
					"ф",
					"х",
				},
				Mods: []string{
					"а",
					"у",
					"а",
					"ом",
					"е",
				},
			},
			{
				Gender: Male,
				Test: []string{
					"в",
					"н",
				},
				Mods: []string{
					"а",
					"у",
					"а",
					"ым",
					"е",
				},
			},
		},
	},
}
