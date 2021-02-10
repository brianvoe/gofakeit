package gofakeit

import (
	"errors"
	"fmt"
	"math"
	rand "math/rand"
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
func Generate(dataVal string) string { return generate(globalFaker.Rand, dataVal) }

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
func (f *Faker) Generate(dataVal string) string { return generate(f.Rand, dataVal) }

func generate(r *rand.Rand, dataVal string) string {
	// Replace # with numbers and ? with letters
	dataVal = replaceWithNumbers(r, dataVal)
	dataVal = replaceWithLetters(r, dataVal)

	// Variables to identify the index in which it exists
	startCurly := -1
	startCurlyIgnore := []int{}
	endCurly := -1
	endCurlyIgnore := []int{}

	// Loop through string characters
	for i := 0; i < len(dataVal); i++ {
		// Check for ignores if equal skip
		shouldSkip := false
		for _, igs := range startCurlyIgnore {
			if i == igs {
				shouldSkip = true
			}
		}
		for _, ige := range endCurlyIgnore {
			if i == ige {
				shouldSkip = true
			}
		}
		if shouldSkip {
			continue
		}

		// Identify items between brackets. Ex: {firstname}
		if string(dataVal[i]) == "{" {
			startCurly = i
			continue
		}
		if startCurly != -1 && string(dataVal[i]) == "}" {
			endCurly = i
		}
		if startCurly == -1 || endCurly == -1 {
			continue
		}

		// Get the value between brackets
		fParts := dataVal[startCurly+1 : endCurly]

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

		// Check to see if its a replaceable lookup function
		if info := GetFuncLookup(fName); info != nil {
			// Get parameters, make sure params and the split both have values
			var mapParams *MapParams
			paramsLen := len(info.Params)

			// If just one param and its a string simply just pass it
			if paramsLen == 1 && info.Params[0].Type == "string" {
				if mapParams == nil {
					mapParams = NewMapParams()
				}
				mapParams.Add(info.Params[0].Field, fParams)
			} else if paramsLen > 0 && fParams != "" {
				splitVals := funcLookupSplit(fParams)
				for ii := 0; ii < len(splitVals); ii++ {
					if paramsLen-1 >= ii {
						if mapParams == nil {
							mapParams = NewMapParams()
						}
						if strings.HasPrefix(splitVals[ii], "[") {
							lookupSplits := funcLookupSplit(strings.TrimRight(strings.TrimLeft(splitVals[ii], "["), "]"))
							for _, v := range lookupSplits {
								mapParams.Add(info.Params[ii].Field, v)
							}
						} else {
							mapParams.Add(info.Params[ii].Field, splitVals[ii])
						}
					}
				}
			}

			// Call function
			fValue, err := info.Generate(r, mapParams, info)
			if err != nil {
				// If we came across an error just dont replace value
				dataVal = strings.Replace(dataVal, "{"+fParts+"}", err.Error(), 1)
			} else {
				// Successfully found, run replace with new value
				dataVal = strings.Replace(dataVal, "{"+fParts+"}", fmt.Sprintf("%v", fValue), 1)
			}

			// Reset the curly index back to -1 and reset ignores
			startCurly = -1
			startCurlyIgnore = []int{}
			endCurly = -1
			endCurlyIgnore = []int{}
			i = -1 // Reset back to the start of the string
			continue
		}

		// Couldnt find anything - mark curly brackets to skip and rerun
		startCurlyIgnore = append(startCurlyIgnore, startCurly)
		endCurlyIgnore = append(endCurlyIgnore, endCurly)

		// Reset the curly index back to -1
		startCurly = -1
		endCurly = -1
		i = -1 // Reset back to the start of the string
		continue
	}

	return dataVal
}

// Regex will generate a string based upon a RE2 syntax
func Regex(regexStr string) string { return regex(globalFaker.Rand, regexStr) }

// Regex will generate a string based upon a RE2 syntax
func (f *Faker) Regex(regexStr string) string { return regex(f.Rand, regexStr) }

func regex(r *rand.Rand, regexStr string) string {
	re, err := syntax.Parse(regexStr, syntax.Perl)
	if err != nil {
		return "Could not parse regex string"
	}

	return regexGenerate(r, re)
}

func regexGenerate(ra *rand.Rand, re *syntax.Regexp) string {
	op := re.Op
	switch op {
	case syntax.OpNoMatch: // matches no strings
		// Do Nothing
	case syntax.OpEmptyMatch: // matches empty string
		return ""
	case syntax.OpLiteral: // matches Runes sequence
		var b strings.Builder
		for _, ru := range re.Rune {
			b.WriteRune(ru)
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
				return string([]byte{chars[ra.Intn(len(chars))]})
			}
		}

		r := ra.Intn(int(sum))
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
		return randCharacter(ra, allStr)
	case syntax.OpBeginLine: // matches empty string at beginning of line
	case syntax.OpEndLine: // matches empty string at end of line
	case syntax.OpBeginText: // matches empty string at beginning of text
	case syntax.OpEndText: // matches empty string at end of text
	case syntax.OpWordBoundary: // matches word boundary `\b`
	case syntax.OpNoWordBoundary: // matches word non-boundary `\B`
	case syntax.OpCapture: // capturing subexpression with index Cap, optional name Name
		return regexGenerate(ra, re.Sub0[0])
	case syntax.OpStar: // matches Sub[0] zero or more times
		var b strings.Builder
		for i := 0; i < number(ra, 0, 10); i++ {
			for _, rs := range re.Sub {
				b.WriteString(regexGenerate(ra, rs))
			}
		}
		return b.String()
	case syntax.OpPlus: // matches Sub[0] one or more times
		var b strings.Builder
		for i := 0; i < number(ra, 1, 10); i++ {
			for _, rs := range re.Sub {
				b.WriteString(regexGenerate(ra, rs))
			}
		}
		return b.String()
	case syntax.OpQuest: // matches Sub[0] zero or one times
		var b strings.Builder
		for i := 0; i < number(ra, 0, 1); i++ {
			for _, rs := range re.Sub {
				b.WriteString(regexGenerate(ra, rs))
			}
		}
		return b.String()
	case syntax.OpRepeat: // matches Sub[0] at least Min times, at most Max (Max == -1 is no limit)
		var b strings.Builder
		count := 0
		re.Max = int(math.Min(float64(re.Max), float64(10)))
		if re.Max > re.Min {
			count = ra.Intn(re.Max - re.Min + 1)
		}
		for i := 0; i < re.Min || i < (re.Min+count); i++ {
			for _, rs := range re.Sub {
				b.WriteString(regexGenerate(ra, rs))
			}
		}
		return b.String()
	case syntax.OpConcat: // matches concatenation of Subs
		var b strings.Builder
		for _, rs := range re.Sub {
			b.WriteString(regexGenerate(ra, rs))
		}
		return b.String()
	case syntax.OpAlternate: // matches alternation of Subs
		return regexGenerate(ra, re.Sub[number(ra, 0, len(re.Sub)-1)])
	}

	return ""
}

// Map will generate a random set of map data
func Map() map[string]interface{} { return mapFunc(globalFaker.Rand) }

// Map will generate a random set of map data
func (f *Faker) Map() map[string]interface{} { return mapFunc(f.Rand) }

func mapFunc(r *rand.Rand) map[string]interface{} {
	m := map[string]interface{}{}

	randWordType := func() string {
		s := randomString(r, []string{"lorem", "bs", "job", "name", "address"})
		switch s {
		case "bs":
			return bs(r)
		case "job":
			return jobTitle(r)
		case "name":
			return name(r)
		case "address":
			return street(r) + ", " + city(r) + ", " + state(r) + " " + zip(r)
		}
		return word(r)
	}

	randSlice := func() []string {
		var sl []string
		for ii := 0; ii < number(r, 3, 10); ii++ {
			sl = append(sl, word(r))
		}
		return sl
	}

	for i := 0; i < number(r, 3, 10); i++ {
		t := randomString(r, []string{"string", "int", "float", "slice", "map"})
		switch t {
		case "string":
			m[word(r)] = randWordType()
		case "int":
			m[word(r)] = number(r, 1, 10000000)
		case "float":
			m[word(r)] = float32Range(r, 1, 1000000)
		case "slice":
			m[word(r)] = randSlice()
		case "map":
			mm := map[string]interface{}{}
			tt := randomString(r, []string{"string", "int", "float", "slice"})
			switch tt {
			case "string":
				mm[word(r)] = randWordType()
			case "int":
				mm[word(r)] = number(r, 1, 10000000)
			case "float":
				mm[word(r)] = float32Range(r, 1, 1000000)
			case "slice":
				mm[word(r)] = randSlice()
			}
			m[word(r)] = mm
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
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (interface{}, error) {
			str, err := info.GetString(m, "str")
			if err != nil {
				return nil, err
			}

			// Limit the length of the string passed
			if len(str) > 1000 {
				return nil, errors.New("String length is too large. Limit to 1000 characters")
			}

			return generate(r, str), nil
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
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (interface{}, error) {
			str, err := info.GetString(m, "str")
			if err != nil {
				return nil, err
			}

			// Limit the length of the string passed
			if len(str) > 500 {
				return nil, errors.New("String length is too large. Limit to 500 characters")
			}

			return regex(r, str), nil
		},
	})
}
