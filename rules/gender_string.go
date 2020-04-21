// Code generated by "enumer -json -transform=snake -type=Gender -output=gender_string.go gender.go"; DO NOT EDIT.

//
package rules

import (
	"encoding/json"
	"fmt"
)

const _GenderName = "androgynousmalefemale"

var _GenderIndex = [...]uint8{0, 11, 15, 21}

func (i Gender) String() string {
	if i < 0 || i >= Gender(len(_GenderIndex)-1) {
		return fmt.Sprintf("Gender(%d)", i)
	}
	return _GenderName[_GenderIndex[i]:_GenderIndex[i+1]]
}

var _GenderValues = []Gender{0, 1, 2}

var _GenderNameToValueMap = map[string]Gender{
	_GenderName[0:11]:  0,
	_GenderName[11:15]: 1,
	_GenderName[15:21]: 2,
}

// GenderString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func GenderString(s string) (Gender, error) {
	if val, ok := _GenderNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Gender values", s)
}

// GenderValues returns all values of the enum
func GenderValues() []Gender {
	return _GenderValues
}

// IsAGender returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Gender) IsAGender() bool {
	for _, v := range _GenderValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for Gender
func (i Gender) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for Gender
func (i *Gender) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("Gender should be a string, got %s", data)
	}

	var err error
	*i, err = GenderString(s)
	return err
}
