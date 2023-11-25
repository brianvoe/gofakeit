package gofakeit

import (
	"bytes"
	"fmt"

	"math/rand"
	"reflect"
	"strconv"
	"strings"
	template_engine "text/template"
	"time"
)

// TemplateOptions defines values needed for template document generation
type TemplateOptions struct {
	Funcs template_engine.FuncMap `fake:"-"`
	Data  any                     `json:"data" xml:"data" fake:"-"`
}

// Used with CreateListResult ListResult
type intRangeResult struct {
	Value int   // Current value
	Range []int //Stores a list of values
}

// Template generates an document based on the the supplied template
func Template(template string, co *TemplateOptions) (string, error) {
	if co == nil {
		co = &TemplateOptions{}
		globalFaker.Struct(co)
	}
	return templateFunc(template, templateFuncMap(globalFaker.Rand, &co.Funcs), co)
}

// Template generates an document based on the the supplied template
func (f *Faker) Template(template string, co *TemplateOptions) (string, error) {
	if co == nil {
		co = &TemplateOptions{}
		f.Struct(co)
	}
	return templateFunc(template, templateFuncMap(f.Rand, &co.Funcs), co)
}

// MarkdownOptions defines values needed for markdown document generation
type MarkdownOptions struct {
	Funcs template_engine.FuncMap `fake:"-"`
	//Data          any `json:"data" xml:"data"  fake:"-"`
	SectionsCount int `json:"sections" xml:"sections" fake:"{number:1,10}"`
}

// Template for Markdown
const template_markdown = `{{$res:=CreateListResult 1 5}}
{{range $y := IntRange 1 (.SectionsCount)}}{{$res = ListResult ($res) 1 5 true}}{{if eq $res.Value 1}}# Paragraph\n\n{{Paragraph (Number 1 5) (Number 1 5) (Number 1 30) "\n\n"}}\n\n---
{{end}}{{if eq $res.Value 2}}# Block Quote\n\n{{Paragraph (Number 1 5) (Number 1 5) (Number 1 30) "\n\n"}}\n\n---
{{end}}{{if eq $res.Value 3}}## Details\n\n{{Paragraph (Number 1 5) (Number 1 5) (Number 1 30) "\n\n"}}\n\n<details>\n<summary>{{SentenceSimple}} </summary>\n\n{{Paragraph (Number 3 5) (Number 1 5) (Number 1 30) "\n\n"}}.\n\n</details>\n\n---
{{end}}{{if eq $res.Value 4}}## Url\n\nGolang you will need to install\n\n{{range $y := IntRange 1 (Number 1 10)}}[{{$saved_url:=URL}}{{$saved_url}}]({{$saved_url}})\n{{end}}\n---
{{end}}{{if eq $res.Value 5}}## LISTS\n\n{{Paragraph (Number 1 5) (Number 1 5) (Number 1 30) "\n\n"}}.\n\n{{range $y := IntRange 1 (Number 1 10)}}1. {{PhraseVerb}}\n{{end}}\n---{{end}}\n{{end}}`

// Markdown will return a single random Markdown template document
func Markdown(co *MarkdownOptions) (string, error) {
	if co == nil {
		co = &MarkdownOptions{}
		globalFaker.Struct(co)
	}
	return templateFunc(template_markdown, templateFuncMap(globalFaker.Rand, &co.Funcs), co)
}

// Markdown will return a single random Markdown template document
func (f *Faker) Markdown(co *MarkdownOptions) (string, error) {
	if co == nil {
		co = &MarkdownOptions{}
		f.Struct(co)
	}
	return templateFunc(template_markdown, templateFuncMap(f.Rand, &co.Funcs), co)
}

// EmailOptions defines values needed for email document generation
type EmailOptions struct {
	Funcs template_engine.FuncMap `fake:"-"`
	//Data          any `json:"data" xml:"data" fake:"-"`
	SectionsCount int `json:"sections" xml:"sections" fake:"{number:1,10}"`
}

// Template for email text
const template_email_text = `{{RandomString (ListS "Hi" "Hello" "Dear" "Good morning" "Good afternoon" "Good evening" )}}, {{$saved_to:=Person}}{{$saved_to.FirstName}}\n
{{$res:=CreateListResult 1 7}}{{range $y := IntRange 1 (.SectionsCount)}}{{$res = ListResult ($res) 1 7 true}}{{if eq $res.Value 1}}{{Paragraph (Number 1 3) (Number 3 5) (Number 10 20) "\n\n"}}
{{end}}{{if eq $res.Value 2}}{{Question}}\n{{end}}{{if eq $res.Value 3}}{{Quote}}\n{{end}}{{if eq $res.Value 4}}{{HipsterParagraph (Number 1 3) (Number 3 5) (Number 10 20) "\n\n"}}
{{end}}{{if eq $res.Value 5}}{{RandomString (ListS (Concat " " "Have you seen" (MovieName) "it is" (AdjectiveDescriptive)) (Concat " " "Have you seen the" (MovieName) "?") (Concat "Do you want to watch" (MovieName) (AdverbTimeDefinite) "?") )}}
{{end}}{{if eq $res.Value 6}}{{Concat " " (PronounPersonal) (VerbHelping) (VerbIntransitive) (AdverbTimeDefinite) (PronounDemonstrative) (AdverbFrequencyIndefinite) "Happens" (AdverbTimeDefinite)}}\n{{end}}{{if eq $res.Value 7}}{{HipsterSentence (Number 10 20)}}{{end}}{{end}}
{{$saved_from:=Person}}{{RandomString (ListS "Regards" "Sincerely" "Best wishes" "Cheers" "Take care" "Best" "Thank you" "I appreciate your help" "I appreciate your feedback"  "I appreciate your input")}}\n{{$saved_from.FirstName}}{{if Bool}} {{$saved_from.LastName}}{{end}}\n{{if Bool}}{{if Bool}}\nCompany: {{$saved_from.Job.Company}}\nJob Role: {{$saved_from.Job.Title}}\n{{end}}
{{if Bool}}Address: {{$saved_from.Address.Address}}\nCity: {{$saved_from.Address.City}}\nState: {{$saved_from.Address.State}}\nZip: {{$saved_from.Address.Zip}}{{end}}Phone: {{$saved_from.Contact.Phone}}\nEmail: {{$saved_from.Contact.Email}}{{end}}`

// EmailText will return a single random text email template document
func EmailText(co *EmailOptions) (string, error) {
	if co == nil {
		co = &EmailOptions{}
		globalFaker.Struct(co)
	}
	return templateFunc(template_email_text, templateFuncMap(globalFaker.Rand, &co.Funcs), co)
}

// EmailText will return a single random text email template document
func (f *Faker) EmailText(co *EmailOptions) (string, error) {
	if co == nil {
		co = &EmailOptions{}
		f.Struct(co)
	}
	return templateFunc(template_email_text, templateFuncMap(f.Rand, &co.Funcs), co)
}

// functions that wont work with template engine
var templateExclusion = []string{
	"RandomMapKey",
	"Slice",
	"Struct",
	"SQL",
	"Generate",
	"Template",
}

// Build the template_engine.FuncMap for the template engine
func templateFuncMap(r *rand.Rand, fm *template_engine.FuncMap) *template_engine.FuncMap {

	// create a new function map
	funcMap := template_engine.FuncMap{}

	// build the function map from a faker using their rand
	f := &Faker{Rand: r}

	v := reflect.ValueOf(f)

	// loop through the methods
	for i := 0; i < v.NumMethod(); i++ {
		// check if the method is in the exclusion list
		if stringInSlice(v.Type().Method(i).Name, templateExclusion) {
			continue
		}

		// Check if method has return values
		// If not don't add to function map
		if v.Type().Method(i).Type.NumOut() == 0 {
			continue
		}

		// add the method to the function map
		funcMap[v.Type().Method(i).Name] = v.Method(i).Interface()
	}

	// add the custom functions
	funcMap["String"] = func(value interface{}) string {
		return fmt.Sprintf("%s", value)
	}

	// function to generate a range of integers
	funcMap["IntRange"] = func(start, end int) []int { // function to generate a range of integers
		n := end - start + 1
		result := make([]int, n)
		for i := 0; i < n; i++ {
			result[i] = start + i
		}
		return result
	}

	// function to generate a range of integers and store in intRangeResult
	funcMap["CreateListResult"] = func(start, end int) intRangeResult {
		n := end - start + 1
		result := make([]int, n)
		for i := 0; i < n; i++ {
			result[i] = start + i
		}
		return intRangeResult{Value: start - 1, Range: result}
	}

	// function to remove value from list
	funcMap["ListResult"] = func(rng intRangeResult, min int, max int, shuffle bool) intRangeResult {
		var new_list []int

		// if the remove value is in the list remove it
		if len(rng.Range) > 0 {
			for _, v := range rng.Range {
				if v != rng.Value {
					new_list = append(new_list, v)
				}
			}
		}

		// see if the list is empty
		if len(rng.Range) == 0 {
			// create a new slice of ints
			n := max - min + 1
			for i := 0; i < n; i++ {
				new_list = append(new_list, min+i)
			}
		}

		// randomize the slice
		if shuffle {
			shuffleInts(r, new_list)
		}
		// return the first value and the rest of the list
		return intRangeResult{Value: new_list[0], Range: new_list[1:]}
	}

	// function to replace all values in string
	funcMap["Replace"] = strings.ReplaceAll

	// function to concatenate strings
	funcMap["Concat"] = func(sep string, args ...string) string {
		return strings.Join(args, sep)
	}

	// function to concatenate strings
	funcMap["DateS"] = func(date_string string) time.Time {
		date, err := time.Parse("2006-01-02", date_string)
		if err != nil {
			return time.Now()
		}
		return date
	}

	// function to make string upper case
	funcMap["Upper"] = strings.ToUpper

	// function to make string lower case
	funcMap["LCase"] = strings.ToLower

	// function to enable passing slice of interface to functions
	funcMap["ListI"] = func(args ...any) []any {
		return args
	}

	// function to enable passing slice of string to functions
	funcMap["ListS"] = func(args ...string) []string {
		return args
	}

	// function to enable passing slice of uint to functions
	funcMap["ListUInt"] = func(args ...uint) []uint {
		return args
	}

	// function to enable passing slice of int to functions
	funcMap["ListInt"] = func(args ...int) []int {
		return args
	}

	// function to enable passing slice of int to functions
	funcMap["ListF32"] = func(args ...float32) []float32 {
		return args
	}

	// function to enable passing slice of int to functions
	funcMap["Int"] = func(args any) (int, error) {
		switch v := args.(type) {
		case string:
			i, err := strconv.Atoi(v)
			if err != nil {
				// ... handle error
				return 0, err
			}
			return i, nil
		case float64:
			return int(v), nil
		case float32:
			return int(v), nil
		case int:
			return v, nil

		// Add whatever other types you need
		default:
			return 0, fmt.Errorf("int: unsupported type %T", v)
		}
	}

	// add custom functions
	// override the function if it is passed in
	if fm != nil {
		if reflect.ValueOf(fm).IsNil() {
			return &funcMap
		}
		for k, v := range *fm {
			funcMap[k] = v
		}
	}

	return &funcMap
}

// function to build the function map for the template engine from the global faker
func templateFunc(temp string, funcs *template_engine.FuncMap, data any) (string, error) {
	if temp == "" {
		return "", fmt.Errorf("template parameter is empty")
	}

	// Check if we have a template else just return
	if temp == "" {
		return "", nil
	}

	// Create a new template and parse
	template_gen, err := template_engine.New("CodeRun").Funcs(*funcs).Parse(temp)
	if err != nil {
		return "", err
	}

	b := &bytes.Buffer{}
	err = template_gen.Execute(b, data)
	if err != nil {
		return "", err
	}

	// Return the result
	return strings.ReplaceAll(b.String(), "\\n", "\n"), nil

}

// addTemplateLookup will add the template functions to the global lookup
func addTemplateLookup() {
	AddFuncLookup("template", Info{
		Display:     "Template",
		Category:    "template",
		Description: "Generates document from template",
		Example: `
			{{name}}
			
			// output
			Markus Moen
		`,
		Output:      "string",
		ContentType: "text/plain",
		Params: []Param{
			{Field: "template", Display: "Template", Type: "string", Description: "Golang template to generate the document from"},
			{Field: "data", Display: "Custom Data", Type: "string", Default: "", Optional: true, Description: "Custom data to pass to the template"},
		},
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (any, error) {
			tpl, err := info.GetString(m, "template")
			if err != nil {
				return nil, err
			}

			data, err := info.GetAny(m, "data")
			if err != nil {
				return nil, err
			}

			templateOut, err := templateFunc(tpl, templateFuncMap(r, nil), &TemplateOptions{Data: data})
			if err != nil {
				return nil, err
			}

			return string(templateOut), nil
		},
	})

	AddFuncLookup("markdown", Info{
		Display:     "Random markdown document",
		Category:    "template",
		Description: "Generates random markdown document",
		Example:     "",
		Output:      "string",
		Params: []Param{
			{Field: "sections_count", Display: "Body Sections", Type: "int", Default: "1", Optional: true, Description: "Number of content sections to generate"},
		},
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (any, error) {
			sections, err := info.GetInt(m, "sections_count")
			if err != nil {
				sections = 1
			}

			template_result, err := templateFunc(template_markdown, templateFuncMap(r, nil), &MarkdownOptions{SectionsCount: sections})
			return string(template_result), err
		},
	})

	AddFuncLookup("email_text", Info{
		Display:     "Random text email Document",
		Category:    "template",
		Description: "Generates random email document.",
		Example:     "",
		Output:      "string",
		Params: []Param{
			{Field: "sections_count", Display: "Body Sections", Type: "int", Default: "1", Optional: true, Description: "Number of content sections to generate"},
		},
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (any, error) {
			sections, err := info.GetInt(m, "sections_count")
			if err != nil {
				sections = 1
			}

			template_result, err := templateFunc(template_email_text, templateFuncMap(r, nil), &EmailOptions{SectionsCount: sections})
			return string(template_result), err
		},
	})
}
