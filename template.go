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
func TemplateDocument() (string, error) {
	return templateDocument(globalFaker, Number(1, 5), []string{"template", "document"})
}

// Template will return a single random document
func (f *Faker) TemplateDocument() (string, error) {
	return templateDocument(f, Number(1, 5), []string{"template", "document"})
}

// Template will return a single random HTML document
func TemplateHtmlContent() (string, error) {
	return templateDocument(globalFaker, Number(1, 5), []string{"template", "html_content"})
}

// Template will return a single random HTML document
func (f *Faker) TemplateHtmlContent() (string, error) {
	return templateDocument(f, Number(1, 5), []string{"template", "html_content"})
}

// Template will return a single random HTML document
func TemplateHtml(sections int) (string, error) {
	return templateDocument(globalFaker, sections, []string{"template", "html"})
}

// Template will return a single random HTML document
func (f *Faker) TemplateHtml(sections int) (string, error) {
	return templateDocument(f, sections, []string{"template", "html"})
}

// Template will return a single random Markdown template document
func TemplateMarkdown(lines int) (string, error) {
	return templateDocument(globalFaker, lines, []string{"template", "markdown"})
}

// Template will return a single random Markdown template document
func (f *Faker) TemplateMarkdown(lines int) (string, error) {
	return templateDocument(f, lines, []string{"template", "markdown"})
}

// Template will return a single random email template document
func TemplateEmail() (string, error) {
	return templateDocument(globalFaker, -1, []string{"template", "email"})
}

// Template will return a single random email template document
func (f *Faker) TemplateEmail() (string, error) {
	return templateDocument(f, -1, []string{"template", "email"})
}

// function to build the function map for the template engine from the global faker
func createGlobalFakerFunctionMap() template.FuncMap {
	funcMap := template.FuncMap{}

	//Build the function map from the globalFaker
	v := reflect.ValueOf(globalFaker)
	//loop through the methods
	for i := 0; i < v.NumMethod(); i++ {
		method := v.Method(i)
		//get the method name lowercase
		method_name := v.Type().Method(i).Name
		//get the arguments for the method
		methodArgs := make([]reflect.Value, method.Type().NumIn())
		//Check to see if the method has any arguments that are not supported by templates
		has_invalid_arg := false
		for m := range methodArgs {
			if method.Type().In(m).Kind() == reflect.Array || method.Type().In(m).Kind() == reflect.Slice || method.Type().In(m).Kind() == reflect.Interface {
				has_invalid_arg = true
				break
			}
		}
		//if the method has no invalid arguments add it to the funcMap
		if !has_invalid_arg {
			funcMap[method_name] = method.Interface()
		}
	}

	//add the custom functions

	//function to generate a range of integers
	funcMap["IntRange"] = func(start, end int) []int { // function to generate a range of integers
		n := end - start + 1
		result := make([]int, n)
		for i := 0; i < n; i++ {
			result[i] = start + i
		}
		return result
	}

	//function to generate a base64 encode string useful for images
	funcMap["Base64Enc"] = base64EncString

	//function  to replace all values in string
	funcMap["replace"] = strings.ReplaceAll

	//function to make string lower case
	funcMap["lc"] = strings.ToLower

	//function to make string upper case
	funcMap["uc"] = strings.ToUpper

	//function to generate a random SVG this is because template engine cant handle the struct
	funcMap["SVG"] = func(width int, height int) string {
		return globalFaker.Svg(&SVGOptions{Width: width, Height: height, Type: "svg", Colors: []string{"#000000", "#FFFFFF"}})
	}

	//function to generate a random sql this is because template engine cant handle the struct
	funcMap["Sql"] = func() (string, error) {
		return globalFaker.SQL(&SQLOptions{
			Table: "people",
			Count: 2,
			Fields: []Field{
				{Name: "id", Function: "autoincrement"},
				{Name: "first_name", Function: "firstname"},
				{Name: "price", Function: "price"},
				{Name: "age", Function: "number", Params: MapParams{"min": {"1"}, "max": {"99"}}},
				{Name: "created_at", Function: "date", Params: MapParams{"format": {"2006-01-02 15:04:05"}}},
			}})
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

// function to fix the string for the template engine
func fixString(str string) string {
	str = strings.ReplaceAll(str, "'", "`")
	str = strings.ReplaceAll(str, "\\n", "\n")
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
		co.functionMap = createGlobalFakerFunctionMap()
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
			{{Name}} {{LastName}}

			:output
			Markus Moen
		`,
		Output:      "[]byte",
		ContentType: "text/plain",
		Params: []Param{
			{Field: "template", Display: "Template", Type: "string", Description: "Golang template or pass email, html or markdown for example document or leave blank for random document", Optional: true},
			{Field: "lines", Display: "Lines", Type: "int", Description: "Used for templates that generate multiple lines can use .Lines to access value in template", Optional: true},
		},
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (interface{}, error) {
			co := TemplateOptions{}

			lines_optional := true
			template_optional := true

			param_values := info.Params
			for _, v := range param_values {
				switch v.Field {
				case "lines":
					lines_optional = v.Optional
				case "template":
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
			lines, err := info.GetInt(m, "lines")
			if err != nil && !lines_optional {
				return nil, err
			}
			co.Lines = lines

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
		Category:    "file",
		Description: "Generates Random document",
		Example:     "",
		Output:      "string",
		Params: []Param{
			{Field: "lines", Display: "Lines", Type: "int", Optional: true, Description: "Used for templates that generate multiple lines can use .Lines to access value in template"},
		},
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (interface{}, error) {
			lines, err := info.GetInt(m, "lines")
			if err != nil {
				lines = -1
			}
			return templateDocument(globalFaker, lines, []string{"template", "document"})
		},
	})

	AddFuncLookup("template_email", Info{
		Display:     "Random email Document",
		Category:    "file",
		Description: "Generates random email document",
		Example:     "",
		Output:      "string",
		Params:      []Param{},
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (interface{}, error) {
			return templateDocument(globalFaker, -1, []string{"template", "email"})
		},
	})

	AddFuncLookup("template_html", Info{
		Display:     "Random html Document",
		Category:    "file",
		Description: "Generates random html document",
		Example:     "",
		Output:      "string",
		Params: []Param{
			{Field: "sections", Display: "Body Sections", Type: "int", Default: "2", Description: "Number of content sections to generate"},
		},
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (interface{}, error) {
			lines, err := info.GetInt(m, "sections")
			if err != nil {
				lines = 1
			}
			return templateDocument(globalFaker, lines, []string{"template", "html"})
		},
	})

	AddFuncLookup("template_html_content", Info{
		Display:     "Random html body content Document",
		Category:    "file",
		Description: "Generates random html body content document",
		Example:     "",
		Output:      "string",
		Params:      []Param{},
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (interface{}, error) {
			return templateDocument(globalFaker, -1, []string{"template", "html_content"})
		},
	})

	AddFuncLookup("template_markdown", Info{
		Display:     "Random markdown document",
		Category:    "file",
		Description: "Generates random markdown document",
		Example:     "",
		Output:      "string",
		Params: []Param{
			{Field: "lines", Display: "Lines", Type: "int", Optional: true, Description: "Used for templates that generate multiple lines can use .Lines to access value in template"},
		},
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (interface{}, error) {
			lines, err := info.GetInt(m, "lines")
			if err != nil {
				lines = -1
			}
			return templateDocument(globalFaker, lines, []string{"template", "markdown"})
		},
	})
}
