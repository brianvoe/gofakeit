package gofakeit

import (
	"bytes"
	"math/rand"
	"reflect"
	"strings"
	"text/template"
)

// TemplateOptions defines values needed for template document generation
type TemplateOptions struct {
	Template    string `json:"template" xml:"template" fake:"{template_document}"` // go lang template to use to generate the document email,fixed_width,markdown,html
	Lines       int    `json:"lines" xml:"lines" fake:"{number:1,10}"`             // number of lines to generate this is passed to the template
	functionMap template.FuncMap
}

// Used to pass data to the template
type templateData struct {
	Lines int // number of lines to generate this is passed to the template
}

// Used with CreateListResult ListResult
type intRangeResult struct {
	Value int   // Current value
	Range []int //Stores a list of values
}

// Template generates an document based on the the supplied template
// A nil TemplateOptions returns a document.
func Template(template string, lines int) ([]byte, error) {
	return templateFunc(globalFaker, &TemplateOptions{Template: template, Lines: lines})
}

// Template generates an document or an array of objects in json format
// A nil TemplateOptions returns a randomly structured CSV.
func (f *Faker) Template(template string, lines int) ([]byte, error) {
	return templateFunc(f, &TemplateOptions{Template: template, Lines: lines})
}

// Template Document will return a single random document
func TemplateDocument(sections int) (string, error) {
	return templateDocument(globalFaker, sections, []string{"template", "document"})
}

// Template will return a single random document
func (f *Faker) TemplateDocument(sections int) (string, error) {
	return templateDocument(f, sections, []string{"template", "document"})
}

// Template will return a single random Markdown template document
func TemplateMarkdown(sections int) (string, error) {
	return templateDocument(globalFaker, sections, []string{"template", "markdown"})
}

// Template will return a single random Markdown template document
func (f *Faker) TemplateMarkdown(sections int) (string, error) {
	return templateDocument(f, sections, []string{"template", "markdown"})
}

// Template will return a single random text email template document
func TemplateEmailText(sections int) (string, error) {
	return templateDocument(globalFaker, sections, []string{"template", "email"})
}

// Template will return a single random text email template document
func (f *Faker) TemplateEmailText(sections int) (string, error) {
	return templateDocument(f, sections, []string{"template", "email"})
}

// function to build the function map for the template engine from the global faker
func templateFuncMap() template.FuncMap {
	funcMap := template.FuncMap{}

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
	funcMap["Concat"] = func(args ...string) string {
		return strings.Join(args, " ")
	}
	// function to make string upper case
	funcMap["Upper"] = strings.ToUpper

	// function to make string lower case
	funcMap["Lower"] = strings.ToLower

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

	return funcMap
}

// generates a random template based of the data
func templateDocument(f *Faker, lines int, dataVal []string) (string, error) {

	random_template := fixString(getRandValue(f.Rand, dataVal))

	template_options := &TemplateOptions{Template: random_template}
	if lines > 0 {
		template_options.Lines = lines
	}
	document, err := templateFunc(f, template_options)
	if err != nil {
		return "", err
	}
	return string(document), nil
}

// function to fix the multiline template data ready for the template engine
func fixString(str string) string {
	str = strings.ReplaceAll(str, "'", "`")
	str = strings.ReplaceAll(str, "\\n", "\n")
	str = strings.ReplaceAll(str, "|n", "\\n")
	return str
}

// function to build the function map for the template engine from the global faker
func templateFunc(f *Faker, co *TemplateOptions) ([]byte, error) {
	if co == nil {
		// We didn't get a CSVOptions, so create a new random one
		err := f.Struct(&co)
		if err != nil {
			return nil, err
		}
	}

	// Check if we have a template else gentate a random one
	if co.Template == "" {
		co.Template = fixString(getRandValue(f.Rand, []string{"template", "document"}))
	}

	// Sort out the lines to pass to the template
	if co.Lines <= 0 {
		co.Lines = 1
	}
	td := templateData{Lines: co.Lines}

	// check if we have a function map
	if co.functionMap == nil {
		co.functionMap = templateFuncMap()
	}

	// Create a new template and parse
	template_gen, err := template.New("CodeRun").Funcs(co.functionMap).Parse(co.Template)
	if err != nil {
		return nil, err
	}

	b := &bytes.Buffer{}
	err = template_gen.Execute(b, td)
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
		Category:    "file",
		Description: "Generates document from template",
		Example: `
			Template
			{{range $y := IntRange 1 .Lines}}
			{{Name}} {{LastName}}{{end}}
			
			:output
			Markus Moen
			Alayna Wuckert
		`,
		Output:      "[]byte",
		ContentType: "text/plain",
		Params: []Param{
			{Field: "template", Display: "Template", Type: "string", Description: "Golang template to generate the document from", Optional: true},
			{Field: "lines", Display: "Body Sections", Type: "int", Optional: true, Description: "Number of content sections to generate"},
		},
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (interface{}, error) {
			co := TemplateOptions{}

			lines_optional := true
			template_optional := true

			param_values := info.Params
			for _, v := range param_values {
				switch v.Field {
				case "sections":
					lines_optional = v.Optional
				case "lines":
					template_optional = v.Optional
				}
			}

			//template to use
			template, err := info.GetString(m, "template")
			if err != nil && !template_optional {
				return nil, err
			}
			co.Template = template

			//the template type to use
			sections, err := info.GetInt(m, "lines")
			if err != nil && !lines_optional {
				return nil, err
			}
			co.Lines = sections

			f := &Faker{Rand: r}
			templateOut, err := templateFunc(f, &co)
			if err != nil {
				return nil, err
			}

			return string(templateOut), nil
		},
	})

	AddFuncLookup("template_document", Info{
		Display:     "Random Document",
		Category:    "template",
		Description: "Generates Random document.",
		Example:     "",
		Output:      "string",
		Params: []Param{
			{Field: "sections", Display: "Body Sections", Type: "int", Optional: true, Description: "Number of content sections to generate"},
		},
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (interface{}, error) {
			sections, err := info.GetInt(m, "sections")
			if err != nil {
				sections = -1
			}
			return templateDocument(globalFaker, sections, []string{"template", "document"})
		},
	})

	AddFuncLookup("template_email_text", Info{
		Display:     "Random text email Document",
		Category:    "template",
		Description: "Generates random email document.",
		Example:     "",
		Output:      "string",
		Params: []Param{
			{Field: "sections", Display: "Body Sections", Type: "int", Optional: true, Description: "Number of content sections to generate"},
		},
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (interface{}, error) {
			sections, err := info.GetInt(m, "sections")
			if err != nil {
				sections = -1
			}
			return templateDocument(globalFaker, sections, []string{"template", "email"})
		},
	})

	AddFuncLookup("template_markdown", Info{
		Display:     "Random markdown document.",
		Category:    "template",
		Description: "Generates random markdown document",
		Example:     "",
		Output:      "string",
		Params: []Param{
			{Field: "sections", Display: "Body Sections", Type: "int", Optional: true, Description: "Number of content sections to generate"},
		},
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (interface{}, error) {
			lines, err := info.GetInt(m, "sections")
			if err != nil {
				lines = -1
			}
			return templateDocument(globalFaker, lines, []string{"template", "markdown"})
		},
	})

}
