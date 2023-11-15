package gofakeit

import (
	"bytes"
	"fmt"
	"math/rand"
	"reflect"
	"strconv"
	"strings"
	"text/template"
)

// interface for template options
type iOptions interface {
	GetData() interface{}
	GetFuncs() *template.FuncMap
	SetFuncs(f *template.FuncMap)
}

// baseTemplateOptions implements iOptions interface Funcs accessors
type baseTemplateOptions struct {
	funcs template.FuncMap
	Data  interface{} //Data to be used in template
}

// getValue returns the data interface
func (b baseTemplateOptions) GetData() interface{} {
	return b.Data
}

// GetFuncs returns the template.FuncMap
func (b baseTemplateOptions) GetFuncs() *template.FuncMap {
	return &b.funcs
}

// SetFuncs sets the template.FuncMap
func (b *baseTemplateOptions) SetFuncs(f *template.FuncMap) {
	b.funcs = *f
}

// TemplateOptions defines values needed for template document generation
type TemplateOptions struct {
	Data interface{} `json:"data" xml:"data" fake:"-"` // number of lines to generate this is passed to the template
	baseTemplateOptions
}

// getValue returns the data interface
func (f TemplateOptions) GetData() interface{} {
	return f.Data
}

// TemplateOptions defines values needed for template document generation
type EmailOptions struct {
	Sections_count int `json:"sections" xml:"sections" fake:"{number:1,10}"` // number of lines to generate this is passed to the template
	baseTemplateOptions
}

// getValue returns the data interface
func (f EmailOptions) GetData() interface{} {
	return f.Sections_count
}

// MarkdownOptions defines values needed for markdown document generation
type MarkdownOptions struct {
	Sections_count int `json:"sections" xml:"sections" fake:"{number:1,10}"` // number of lines to generate this is passed to the template
	baseTemplateOptions
}

// getValue returns the data interface
func (f MarkdownOptions) GetData() interface{} {
	return f.Sections_count
}

// Used with CreateListResult ListResult
type intRangeResult struct {
	Value int   // Current value
	Range []int //Stores a list of values
}

// Template for email text
const template_email_text = `{{RandomString (ListS "Hi" "Hello" "Dear" "Good morning" "Good afternoon" "Good evening" )}}, {{$saved_to:=Person}}{{$saved_to.FirstName}}\n
{{$res:=CreateListResult 1 7}}{{range $y := IntRange 1 (.GetData)}}{{$res = ListResult ($res) 1 7 true}}{{if eq $res.Value 1}}{{Paragraph (Number 1 3) (Number 3 5) (Number 10 20) "\n\n"}}
{{end}}{{if eq $res.Value 2}}{{Question}}\n{{end}}{{if eq $res.Value 3}}{{Quote}}\n{{end}}{{if eq $res.Value 4}}{{HipsterParagraph (Number 1 3) (Number 3 5) (Number 10 20) "\n\n"}}
{{end}}{{if eq $res.Value 5}}{{RandomString (ListS (Concat " " "Have you seen" (MovieName) "it is" (AdjectiveDescriptive)) (Concat " " "Have you seen the" (MovieName) "?") (Concat "Do you want to watch" (MovieName) (AdverbTimeDefinite) "?") )}}
{{end}}{{if eq $res.Value 6}}{{Concat " " (PronounPersonal) (VerbHelping) (VerbIntransitive) (AdverbTimeDefinite) (PronounDemonstrative) (AdverbFrequencyIndefinite) "Happens" (AdverbTimeDefinite)}}\n{{end}}{{if eq $res.Value 7}}{{HipsterSentence (Number 10 20)}}{{end}}{{end}}
{{$saved_from:=Person}}{{RandomString (ListS "Regards" "Sincerely" "Best wishes" "Cheers" "Take care" "Best" "Thank you" "I appreciate your help" "I appreciate your feedback"  "I appreciate your input")}}\n{{$saved_from.FirstName}}{{if Bool}} {{$saved_from.LastName}}{{end}}\n{{if Bool}}{{if Bool}}\nCompany: {{$saved_from.Job.Company}}\nJob Role: {{$saved_from.Job.Title}}\n{{end}}
{{if Bool}}Address: {{$saved_from.Address.Address}}\nCity: {{$saved_from.Address.City}}\nState: {{$saved_from.Address.State}}\nZip: {{$saved_from.Address.Zip}}{{end}}Phone: {{$saved_from.Contact.Phone}}\nEmail: {{$saved_from.Contact.Email}}{{end}}`

// Template for Markdown
const template_markdown = `{{$res:=CreateListResult 1 5}}
{{range $y := IntRange 1 (.GetData)}}{{$res = ListResult ($res) 1 5 true}}{{if eq $res.Value 1}}# Paragraph\n\n{{Paragraph (Number 1 5) (Number 1 5) (Number 1 30) "\n\n"}}\n\n---
{{end}}{{if eq $res.Value 2}}# Block Quote\n\n{{Paragraph (Number 1 5) (Number 1 5) (Number 1 30) "\n\n"}}\n\n---
{{end}}{{if eq $res.Value 3}}## Details\n\n{{Paragraph (Number 1 5) (Number 1 5) (Number 1 30) "\n\n"}}\n\n<details>\n<summary>{{SentenceSimple}} </summary>\n\n{{Paragraph (Number 3 5) (Number 1 5) (Number 1 30) "\n\n"}}.\n\n</details>\n\n---
{{end}}{{if eq $res.Value 4}}## Url\n\nGolang you will need to install\n\n{{range $y := IntRange 1 (Number 1 10)}}[{{$saved_url:=URL}}{{$saved_url}}]({{$saved_url}})\n{{end}}\n---
{{end}}{{if eq $res.Value 5}}## LISTS\n\n{{Paragraph (Number 1 5) (Number 1 5) (Number 1 30) "\n\n"}}.\n\n{{range $y := IntRange 1 (Number 1 10)}}1. {{PhraseVerb}}\n{{end}}\n---{{end}}\n{{end}}`

// function to fix the new line escape characters
func fixString(str string) string {
	str = strings.ReplaceAll(str, "\\n", "\n")
	//str = strings.ReplaceAll(str, `\\"`, "\"")
	return str
}

// Template generates an document based on the the supplied template
// A nil TemplateOptions returns a document.
func Template(template string, co *TemplateOptions) ([]byte, error) {
	if template == "" {
		return nil, fmt.Errorf("template parameter is empty")
	}
	template_result, err := templateFunc(globalFaker, template, co)
	return []byte(fixString(string(template_result))), err
}

// Template generates an document or an array of objects in json format
// A nil TemplateOptions returns a randomly structured CSV.
func (f *Faker) Template(template string, co *TemplateOptions) ([]byte, error) {
	if template == "" {
		return nil, fmt.Errorf("template parameter is empty")
	}
	template_result, err := templateFunc(f, template, co)
	return []byte(fixString(string(template_result))), err

}

// Template will return a single random Markdown template document
func Markdown(co *MarkdownOptions) (string, error) {
	template_result, err := templateFunc(globalFaker, template_markdown, co)
	return fixString(string(template_result)), err
}

// Template will return a single random Markdown template document
func (f *Faker) Markdown(co *MarkdownOptions) (string, error) {
	template_result, err := templateFunc(f, template_markdown, co)
	return fixString(string(template_result)), err
}

// Template will return a single random text email template document
func EmailText(co *EmailOptions) (string, error) {
	template_result, err := templateFunc(globalFaker, template_email_text, co)
	return fixString(string(template_result)), err
}

// Template will return a single random text email template document
func (f *Faker) EmailText(co *EmailOptions) (string, error) {
	template_result, err := templateFunc(globalFaker, template_email_text, co)
	return fixString(string(template_result)), err
}

// function to build the function map for the template engine from the global faker
func templateFuncMap(fm template.FuncMap) *template.FuncMap {

	funcMap := template.FuncMap{} // create a new function map

	// functions that wont work with template engine
	incompatible := map[string]string{
		"ShuffleAnySlice": "ShuffleAnySlice",
		"ShuffleInts":     "ShuffleInts",
		"ShuffleStrings":  "ShuffleStrings",
		"Slice":           "Slice",
		"Struct":          "Struct",
		"RandomMapKey":    "RandomMapKey",
		"SQL":             "SQL",
		"SVG":             "SVG",
	}

	// Build the function map from the globalFaker
	v := reflect.ValueOf(globalFaker)
	// loop through the methods
	for i := 0; i < v.NumMethod(); i++ {
		method := v.Method(i)
		// get the method name lowercase
		method_name := v.Type().Method(i).Name

		//check if method_name is in the incompatable list
		if _, ok := incompatible[method_name]; ok {
			continue
		}

		funcMap[method_name] = method.Interface()
	}

	// add the custom functions

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
	funcMap["ListResult"] = func(r intRangeResult, min int, max int, shuffle bool) intRangeResult {
		var new_list []int

		// if the remove value is in the list remove it
		if len(r.Range) > 0 {
			for _, v := range r.Range {
				if v != r.Value {
					new_list = append(new_list, v)
				}
			}
		}

		// See if the list is empty
		if len(r.Range) == 0 {
			// create a new slice of ints
			n := max - min + 1
			for i := 0; i < n; i++ {
				new_list = append(new_list, min+i)
			}
		}

		// Randomize the slice
		if shuffle {
			shuffleInts(globalFaker.Rand, new_list)
		}
		// return the first value and the rest of the list
		return intRangeResult{Value: new_list[0], Range: new_list[1:]}
	}

	// function  to replace all values in string
	funcMap["Replace"] = strings.ReplaceAll

	// function to concatenate strings
	funcMap["Concat"] = func(sep string, args ...string) string {
		return strings.Join(args, sep)
	}
	// function to make string upper case
	funcMap["Upper"] = strings.ToUpper

	// function to make string lower case
	funcMap["LCase"] = strings.ToLower

	// function to enable passing slice of interface to functions
	funcMap["ListI"] = func(args ...interface{}) []interface{} {
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
	funcMap["Int"] = func(args interface{}) (int, error) {
		switch v := args.(type) {
		case string:
			i, err := strconv.Atoi(v)
			if err != nil {
				// ... handle error
				return 0, err
			}
			return i, nil
		case int:
			return v, nil
		// Add whatever other types you need
		default:
			return 0, fmt.Errorf("int: unsupported type %T", v)
		}
	}

	// merge the function maps check if the key exists and if it does add a user to the start

	for k, v := range fm {
		if _, ok := funcMap[k]; ok {
			funcMap[fmt.Sprintf("user%s", k)] = v
		} else {
			funcMap[k] = v
		}
	}

	return &funcMap
}

func isNil(input interface{}) bool {
	if input == nil {
		return true
	}
	kind := reflect.ValueOf(input).Kind()
	switch kind {
	case reflect.Ptr:
		return reflect.ValueOf(input).IsNil()
	default:
		return false
	}
}

// function to build the function map for the template engine from the global faker
func templateFunc(f *Faker, tpl string, co iOptions) ([]byte, error) {
	if isNil(co) {
		// We didn't get a CSVOptions, so create a new random one
		//err := f.Struct(&co)
		//if err != nil {
		//	return nil, err
		//}
		co = &TemplateOptions{Data: nil}
	}

	// Check if we have a template else use email text
	if tpl == "" {
		return []byte(""), nil
	}

	// Merge the function user and inbuilt maps
	co.SetFuncs(templateFuncMap(*co.GetFuncs()))

	// Create a new template and parse
	template_gen, err := template.New("CodeRun").Funcs(*co.GetFuncs()).Parse(tpl)
	if err != nil {
		return nil, err
	}

	b := &bytes.Buffer{}
	err = template_gen.Execute(b, &co)
	if err != nil {
		return nil, err
	}

	// Return the result
	return b.Bytes(), nil

}

// addTemplateLookup will add the template functions to the global lookup
func addTemplateLookup() {
	AddFuncLookup("template", Info{
		Display:     "Template",
		Category:    "template",
		Description: "Generates document from template",
		Example: `
			Template
			{{Name}} {{LastName}}
			
			:output
			Markus Moen
		`,
		Output:      "[]byte",
		ContentType: "text/plain",
		Params: []Param{
			{Field: "template", Display: "Template", Type: "string", Description: "Golang template to generate the document from"},
			{Field: "data", Display: "Custom Data", Type: "string", Default: "", Optional: true, Description: "Custom data to pass to the template"},
		},
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (interface{}, error) {
			co := TemplateOptions{}

			//template to use
			tpl, err := info.GetString(m, "template")
			if err != nil {
				return nil, err
			}

			//the template type to use
			_, val, _ := info.GetField(m, "data")
			if val != nil {
				if len(val) > 0 {
					co.Data = val[0]
				}
			}
			f := &Faker{Rand: r}
			templateOut, err := templateFunc(f, tpl, &co)
			if err != nil {
				return nil, err
			}

			return fixString(string(templateOut)), nil
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
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (interface{}, error) {
			sections, err := info.GetInt(m, "sections_count")
			if err != nil {
				sections = 1
			}
			template_result, err := templateFunc(globalFaker, template_email_text, &EmailOptions{Sections_count: sections})
			return fixString(string(template_result)), err
		},
	})

	AddFuncLookup("markdown", Info{
		Display:     "Random markdown document.",
		Category:    "template",
		Description: "Generates random markdown document",
		Example:     "",
		Output:      "string",
		Params: []Param{
			{Field: "sections_count", Display: "Body Sections", Type: "int", Default: "1", Optional: true, Description: "Number of content sections to generate"},
		},
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (interface{}, error) {
			sections, err := info.GetInt(m, "sections_count")
			if err != nil {
				sections = 1
			}
			template_result, err := templateFunc(globalFaker, template_markdown, &MarkdownOptions{Sections_count: sections})
			return fixString(string(template_result)), err
		},
	})

}
