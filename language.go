package humanize

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
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

func LoadLanguages(l ...Local) error {
	err := parseRulesets(l)
	return err
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

func loadFile(path string) ([]byte, error) {
	match, _ := regexp.MatchString("^[a-z]{2}_[A-Z]{2}$", path)
	if match {
		fmt.Println("Reading", "locals/"+string(path)+".json")
		return ioutil.ReadFile("locals/" + string(path) + ".json")
	}
	fmt.Println("Reading", path)
	return ioutil.ReadFile(path)
}

func parseLocalePath(path string) string {
	match, _ := regexp.MatchString("^[a-z]{2}_[A-Z]{2}$", path)
	if match {
		return path
	}
	s := strings.Split(path, "/")
	tag := strings.Split(s[len(s)-1], ".")
	return tag[0]
}

func parseRulesets(l []Local) error {
	rsts := make(map[string]Ruleset)
	for _, local := range l {
		f, err := loadFile(string(local))
		if err == nil {
			r := Ruleset{}
			err := json.Unmarshal(f, &r)
			if err == nil {
				rsts[parseLocalePath(string(local))] = r 
			} else {
				return err
			}
		} else {
			return err
		}
	}
	rulesets = rsts
	return nil

}
