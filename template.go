package gofakeit

import (
	"bytes"
	b64 "encoding/base64"
	"fmt"
	"math"
	"math/rand"
	"reflect"
	"strings"
	"text/template"

	"github.com/brianvoe/gofakeit/v6/data"
)

// CSVOptions defines values needed for csv generation
type TemplateOptions struct {
	Template     string           `json:"template" xml:"template" fake:"{template_document}"` // go lang template to use to generate the document
	TemplateType string           `json:"template_type" xml:"template_type" fake:"-"`         // Rather than pass a template you can pass the type of template you want to generate email,fixed_format,markdown,html
	Lines        int              `json:"lines" xml:"lines" fake:"{number:1,10}"`             // number of lines to generate this is passed to the template
	FunctionMap  template.FuncMap `json:"function_map" xml:"function_map" fake:"-"`
}

// Used to pass data to the template
type templateData struct {
	Lines int // number of lines to generate this is passed to the template
}

// Template generates an object based on the the supplied template
// A nil TemplateOptions returns a document.
func Template(co *TemplateOptions) ([]byte, error) { return templateFunc(globalFaker, co) }

// Template generates an object or an array of objects in json format
// A nil TemplateOptions returns a randomly structured CSV.
func (f *Faker) Template(co *TemplateOptions) ([]byte, error) {
	return templateFunc(f, co)
}

func TemplateDocument() string { return templateDocument(globalFaker.Rand) }

// Template will return a single template document
func (f *Faker) TemplateDocument() string { return templateDocument(f.Rand) }

// generates a random template
func templateDocument(r *rand.Rand) string { return getRandValue(r, []string{"template", "document"}) }

// function to build the function map for the template engine from the global faker
func BuildFunctionMap() template.FuncMap {
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
	funcMap["Pad"] = strPad                            // function to pad a string with a character to a certain length and side
	funcMap["IntRange"] = func(start, end int) []int { // function to generate a range of integers
		n := end - start + 1
		result := make([]int, n)
		for i := 0; i < n; i++ {
			result[i] = start + i
		}
		return result
	}
	funcMap["Base64Enc"] = base64EncString
	return funcMap
}

// function to truncate a string
func truncateText(s string, max int) string {
	if max < len(s) && max > 0 {
		return s[:max]
	}
	return s
}

// Function to pad a string
func strPad(input string, padLength int, padString string, padType string, trim bool) string {
	var output string

	inputLength := len(input)
	padStringLength := len(padString)

	if inputLength >= padLength {
		if trim {
			return truncateText(input, padLength)
		}
		return input
	}

	repeat := math.Ceil(float64(1) + (float64(padLength-padStringLength))/float64(padStringLength))

	switch padType {
	case "RIGHT":
		output = input + strings.Repeat(padString, int(repeat))
		output = output[:padLength]
	case "LEFT":
		output = strings.Repeat(padString, int(repeat)) + input
		output = output[len(output)-padLength:]
	case "BOTH":
		length := (float64(padLength - inputLength)) / float64(2)
		repeat = math.Ceil(length / float64(padStringLength))
		output = strings.Repeat(padString, int(repeat))[:int(math.Floor(float64(length)))] + input + strings.Repeat(padString, int(repeat))[:int(math.Ceil(float64(length)))]
	}

	return output
}

// function to base64 encode a string
func base64EncString(value interface{}) (string, error) {
	switch v := value.(type) {
	case []byte:
		return b64.StdEncoding.EncodeToString(v), nil
	case string:
		return b64.StdEncoding.EncodeToString([]byte(v)), nil
	default:
		return "", fmt.Errorf("value must be a string or []byte")
	}
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

	// Check if we have a template
	if co.Template == "" && co.TemplateType == "" {
		co.Template = getRandValue(f.Rand, []string{"template", "document"})
	} else if co.Template == "" && co.TemplateType != "" {
		co.Template = data.GetTemplateData(co.TemplateType)
	}

	//Sort out the lines to pass to the template
	if co.Lines <= 0 {
		co.Lines = 1
	}
	td := templateData{Lines: co.Lines}

	//check if we have a function map
	if co.FunctionMap == nil {
		co.FunctionMap = BuildFunctionMap()
	}

	//Create a new template and parse
	template_gen, err := template.New("CodeRun").Funcs(co.FunctionMap).Parse(co.Template)
	if err != nil {
		return nil, err
	}

	//**************************************
	//Run the template to verify the output.
	//**************************************
	b := &bytes.Buffer{}
	err = template_gen.Execute(b, td)
	if err != nil {
		return nil, err
	}

	//******************
	//Return the result
	//******************
	return b.Bytes(), nil

}

func addFileTemplateLookup() {
	AddFuncLookup("template", Info{
		Display:     "Template",
		Category:    "file",
		Description: "Generates document from template",
		Example: `
			Markus Moen Pagac
		`,
		Output:      "[]byte",
		ContentType: "text/html",
		Params: []Param{
			{Field: "lines", Display: "Lines", Type: "int", Default: "1", Description: "Used for templates that generate multiple lines", Optional: true},
			{Field: "template", Display: "Template", Type: "string", Default: "", Description: "Document Template", Optional: true},
			{Field: "template_type", Display: "Template Type", Type: "string", Default: "", Optional: true, Description: "A predefined template type to use email,fixed_format,markdown,html"},
		},
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (interface{}, error) {
			co := TemplateOptions{}

			lines_optional := true
			template_optional := true
			template_type_optional := true

			param_values := info.Params
			for _, v := range param_values {
				switch v.Field {
				case "lines":
					lines_optional = v.Optional
				case "template":
					template_optional = v.Optional
				case "template_type":
					template_type_optional = v.Optional
				}

			}

			//template to use
			template, err := info.GetString(m, "template")
			if err != nil && !template_optional {
				return nil, err
			}
			co.Template = template

			//the template type to use

			template_type, err := info.GetString(m, "template_type")
			if err != nil && !template_type_optional {
				return nil, err
			}
			co.TemplateType = template_type

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
		Display:     "Template Document",
		Category:    "file",
		Description: "Generates Random document template",
		Example:     "{{gen `name`}} {{gen `lastname`}}",
		Output:      "string",
		Params:      []Param{},
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (interface{}, error) {
			value := getRandValue(r, []string{"template", "document"})
			return value, nil
		},
	})
}
