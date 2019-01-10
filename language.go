package humanize

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Local type for the ordinals and times.
type Local string

// Ruleset for accessing rules
type Ruleset struct {
	Mags magnitudes `json:"magnitudes"`
	Inds indicators `json:"indicators"`
	Ords [][]string `json:"ordinals"`
}

type magnitudes struct {
	Now    string
	Second string
	Minute string
	Hour   string
	Day    string
	Week   string
	Month  string
	Year   string

	Seconds string
	Minutes string
	Hours   string
	Days    string
	Weeks   string
	Months  string
	Years   string

	Longtime string
}

type indicator struct {
	Word string
	Fix  string
}

type indicators struct {
	Before indicator
	Later  indicator
}

// Local for constant language values
const (
	English       Local = "en_US"
	Turkish       Local = "tr_TR"
	Uninitialized Local = ""
)

var active = Uninitialized
var ruleset = Ruleset{}
var rulesets = map[string]Ruleset{}

// ValidateLanguage for output
// Must be called before Time or Ordinal function
// If implemented to other functions, call it for not getting
// an error. This function automatically chooses language to English.
func ValidateLanguage() {
	if active == Uninitialized {
		SetLanguage(English)
	}
}

// GetLanguage of the humanizing option.
func GetLanguage() Local {
	return active
}

// GetRuleset returns current ruleset option
func GetRuleset() Ruleset {
	return ruleset
}

func GetLocalRuleset(l string) Ruleset {
	return rulesets[l]
}

// SetLanguage of the humanizing option.
func SetLanguage(l Local) {
	active = l
	parseRuleset(l)
	UpdateMagnitudes()
}

func LoadLanguages(l ...Local) {
	parseRulesets(l)
}

func parseRuleset(l Local) {
	fmt.Println("Reading", "locals/"+string(l)+".json")
	f, err := ioutil.ReadFile("locals/" + string(l) + ".json")
	if err == nil {
		ruleset = Ruleset{}
		err := json.Unmarshal(f, &ruleset)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println("Error ! Can not read file.")
	}
}

func parseRulesets(l []Local) {
	rsts := make(map[string]Ruleset)
	for _, local := range l {
		fmt.Println("Reading", "locals/"+string(local)+".json")
		f, err := ioutil.ReadFile("locals/" + string(local) + ".json")
		if err == nil {
			r := Ruleset{}
			err := json.Unmarshal(f, &r)
			if err == nil {
				rsts[string(local)] = r 
			} else {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Error ! Can not read file.")
			fmt.Println(err)
		}
	}
	rulesets = rsts

}
