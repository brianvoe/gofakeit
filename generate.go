package gofakeit

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"regexp/syntax"
	"strings"
)

// Generate fake information from given string.
// Replaceable values should be within {}
//
// Functions
// Ex: {firstname} - billy
// Ex: {sentence:3} - Record river mind.
// Ex: {number:1,10} - 4
// Ex: {uuid} - 590c1440-9888-45b0-bd51-a817ee07c3f2
//
// Letters/Numbers
// Ex: ### - 481 - random numbers
// Ex: ??? - fda - random letters
//
// For a complete list of runnable functions use FuncsLookup
func Generate(dataVal string) string {
	// Replace # with numbers and ? with letters
	dataVal = replaceWithNumbers(dataVal)
	dataVal = replaceWithLetters(dataVal)

	// Identify items between brackets: {person.first}
	for strings.Count(dataVal, "{") > 0 && strings.Count(dataVal, "}") > 0 {
		fParts := dataVal[(strings.Index(dataVal, "{") + 1):strings.Index(dataVal, "}")]

		// Check if has params separated by :
		fNameSplit := strings.SplitN(fParts, ":", 2)
		fName := ""
		fParams := ""
		if len(fNameSplit) >= 1 {
			fName = fNameSplit[0]
		}
		if len(fNameSplit) >= 2 {
			fParams = fNameSplit[1]
		}

		// Check to see if its a replacable lookup function
		if info := GetFuncLookup(fName); info != nil {
			// Get parameters, make sure params and the split both have values
			var mapParams map[string][]string
			paramsLen := len(info.Params)
			if paramsLen > 0 && fParams != "" {
				splitVals := funcLookupSplit(fParams)
				for i := 0; i < len(splitVals); i++ {
					if paramsLen-1 >= i {
						if mapParams == nil {
							mapParams = make(map[string][]string)
						}
						if strings.HasPrefix(splitVals[i], "[") {
							mapParams[info.Params[i].Field] = funcLookupSplit(strings.TrimRight(strings.TrimLeft(splitVals[i], "["), "]"))
						} else {
							mapParams[info.Params[i].Field] = []string{splitVals[i]}
						}
					}
				}
			}

			// Call function
			fValue, err := info.Call(&mapParams, info)
			if err != nil {
				// If we came across an error just dont replace value
				dataVal = strings.Replace(dataVal, "{"+fParts+"}", err.Error(), 1)
				continue
			}
			dataVal = strings.Replace(dataVal, "{"+fParts+"}", fmt.Sprintf("%v", fValue), 1)
			continue
		}

		// Couldnt find anything - set to n/a
		dataVal = strings.Replace(dataVal, "{"+fParts+"}", "n/a", 1)
		continue
	}

	return dataVal
}

// Regex will generate a string based upon a RE2 syntax
func Regex(regexStr string) string {
	re, err := syntax.Parse(regexStr, syntax.Perl)
	if err != nil {
		return "Could not parse regex string"
	}

	return regexGenerate(re)
}

func regexGenerate(re *syntax.Regexp) string {
	op := re.Op
	switch op {
	case syntax.OpNoMatch: // matches no strings
		// Do Nothing
	case syntax.OpEmptyMatch: // matches empty string
		return ""
	case syntax.OpLiteral: // matches Runes sequence
		var b strings.Builder
		for _, r := range re.Rune {
			b.WriteRune(r)
		}
		return b.String()
	case syntax.OpCharClass: // matches Runes interpreted as range pair list
		// number of possible chars
		sum := 0
		for i := 0; i < len(re.Rune); i += 2 {
			sum += int(re.Rune[i+1]-re.Rune[i]) + 1
			if re.Rune[i+1] == 0x10ffff { // rune range end
				sum = -1
				break
			}
		}

		// pick random char in range (inverse match group)
		if sum == -1 {
			chars := []uint8{}
			for j := 0; j < len(allStr); j++ {
				c := allStr[j]

				// Check c in range
				for i := 0; i < len(re.Rune); i += 2 {
					if rune(c) >= re.Rune[i] && rune(c) <= re.Rune[i+1] {
						chars = append(chars, c)
						break
					}
				}
			}
			if len(chars) > 0 {
				return string([]byte{chars[rand.Intn(len(chars))]})
			}
		}

		r := rand.Intn(int(sum))
		var ru rune
		sum = 0
		for i := 0; i < len(re.Rune); i += 2 {
			gap := int(re.Rune[i+1]-re.Rune[i]) + 1
			if sum+gap > r {
				ru = re.Rune[i] + rune(r-sum)
				break
			}
			sum += gap
		}

		return string(ru)
	case syntax.OpAnyCharNotNL, syntax.OpAnyChar: // matches any character(and except newline)
		return randCharacter(allStr)
	case syntax.OpBeginLine: // matches empty string at beginning of line
	case syntax.OpEndLine: // matches empty string at end of line
	case syntax.OpBeginText: // matches empty string at beginning of text
	case syntax.OpEndText: // matches empty string at end of text
	case syntax.OpWordBoundary: // matches word boundary `\b`
	case syntax.OpNoWordBoundary: // matches word non-boundary `\B`
	case syntax.OpCapture: // capturing subexpression with index Cap, optional name Name
		return regexGenerate(re.Sub0[0])
	case syntax.OpStar: // matches Sub[0] zero or more times
		var b strings.Builder
		for i := 0; i < Number(0, 10); i++ {
			for _, r := range re.Sub {
				b.WriteString(regexGenerate(r))
			}
		}
		return b.String()
	case syntax.OpPlus: // matches Sub[0] one or more times
		var b strings.Builder
		for i := 0; i < Number(1, 10); i++ {
			for _, r := range re.Sub {
				b.WriteString(regexGenerate(r))
			}
		}
		return b.String()
	case syntax.OpQuest: // matches Sub[0] zero or one times
		var b strings.Builder
		for i := 0; i < Number(0, 1); i++ {
			for _, r := range re.Sub {
				b.WriteString(regexGenerate(r))
			}
		}
		return b.String()
	case syntax.OpRepeat: // matches Sub[0] at least Min times, at most Max (Max == -1 is no limit)
		var b strings.Builder
		count := 0
		re.Max = int(math.Min(float64(re.Max), float64(10)))
		if re.Max > re.Min {
			count = rand.Intn(re.Max - re.Min + 1)
		}
		for i := 0; i < re.Min || i < (re.Min+count); i++ {
			for _, r := range re.Sub {
				b.WriteString(regexGenerate(r))
			}
		}
		return b.String()
	case syntax.OpConcat: // matches concatenation of Subs
		var b strings.Builder
		for _, r := range re.Sub {
			b.WriteString(regexGenerate(r))
		}
		return b.String()
	case syntax.OpAlternate: // matches alternation of Subs
		return regexGenerate(re.Sub[Number(0, len(re.Sub)-1)])
	}

	return ""
}

// Map will generate a random set of map data
func Map() map[string]interface{} {
	m := map[string]interface{}{}

	randWordType := func() string {
		s := RandomString([]string{"lorem", "bs", "job", "name", "address"})
		switch s {
		case "bs":
			return BS()
		case "job":
			return JobTitle()
		case "name":
			return Name()
		case "address":
			return Street() + ", " + City() + ", " + State() + " " + Zip()
		}
		return Word()
	}

	randSlice := func() []string {
		var sl []string
		for ii := 0; ii < Number(3, 10); ii++ {
			sl = append(sl, Word())
		}
		return sl
	}

	for i := 0; i < Number(3, 10); i++ {
		t := RandomString([]string{"string", "int", "float", "slice", "map"})
		switch t {
		case "string":
			m[Word()] = randWordType()
		case "int":
			m[Word()] = Number(1, 10000000)
		case "float":
			m[Word()] = Float32Range(1, 1000000)
		case "slice":
			m[Word()] = randSlice()
		case "map":
			mm := map[string]interface{}{}
			tt := RandomString([]string{"string", "int", "float", "slice"})
			switch tt {
			case "string":
				mm[Word()] = randWordType()
			case "int":
				mm[Word()] = Number(1, 10000000)
			case "float":
				mm[Word()] = Float32Range(1, 1000000)
			case "slice":
				mm[Word()] = randSlice()
			}
			m[Word()] = mm
		}
	}

	return m
}

func addGenerateLookup() {
	AddFuncLookup("generate", Info{
		Display:     "Generate",
		Category:    "generate",
		Description: "Random string generated from string value based upon available data sets",
		Example:     "{firstname} {lastname} {email} - Markus Moen markusmoen@pagac.net",
		Output:      "string",
		Params: []Param{
			{Field: "str", Display: "String", Type: "string", Description: "String value to generate from"},
		},
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			str, err := info.GetString(m, "str")
			if err != nil {
				return nil, err
			}

			// Limit the length of the string passed
			if len(str) > 1000 {
				return nil, errors.New("String length is too large. Limit to 1000 characters")
			}

			return Generate(str), nil
		},
	})

	AddFuncLookup("regex", Info{
		Display:     "Regex",
		Category:    "generate",
		Description: "Random string generated from regex RE2 syntax string",
		Example:     "[abcdef]{5} - affec",
		Output:      "string",
		Params: []Param{
			{Field: "str", Display: "String", Type: "string", Description: "Regex RE2 syntax string"},
		},
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			str, err := info.GetString(m, "str")
			if err != nil {
				return nil, err
			}

			// Limit the length of the string passed
			if len(str) > 500 {
				return nil, errors.New("String length is too large. Limit to 500 characters")
			}

			return Regex(str), nil
		},
	})
}
