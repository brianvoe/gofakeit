package data

import "strings"

//***************
// Email Template
//***************
const EMAIL_TEMPLATE = `{{$saved_from_name:=Name}}{{$saved_to_name:=Name}}
From: {{$saved_from_name}} <{{Email}}>
To: {{$saved_to_name}} <{{Email}}>
Subject: {{Phrase}}
Date: {{Date}}
Content-Type: text/plain; charset=UTF-8;

Hi {{$saved_to_name}}

{{Paragraph 3 5 20 '\n\n'}}
Regards {{$saved_from_name}}`

//**********************
// Fixed Format Template
//**********************
const FIXED_FORMAT_TEMPLATE = `{{Pad 'FIRSTNAME' 10 ' ' 'RIGHT' true}}{{Pad 'LASTNAME' 10 ' ' 'RIGHT' true}}{{Pad 'Gender' 8 ' ' 'RIGHT' true}}{{Pad 'Job' 15 ' ' 'RIGHT' true}}{{Pad 'Hobby' 20 ' ' 'RIGHT' true}}{{Pad 'CreditCard' 20 ' ' 'RIGHT' true}}{{'Address'}}
{{range $y := IntRange 1 .Lines}}{{$person:=Person}}{{Pad ($person.FirstName) 10 ' ' 'RIGHT' true}}{{Pad ($person.LastName) 10 ' ' 'RIGHT' true}}{{Pad ($person.Gender) 8 ' ' 'RIGHT' true}}{{Pad ($person.Job.Title) 15 ' ' 'RIGHT' true}}{{Pad ($person.Hobby) 20 ' ' 'RIGHT' true}}{{Pad ($person.CreditCard.Number) 20 ' ' 'RIGHT' true}}{{$person.Address.Address}}
{{end}}`

//******************
// Markdown Template
//******************
const MARKDOWN_TEMPLATE = `# {{SentenceSimple}}  

{{Paragraph 1 5 20 '\n\n'}}

---

## introduction

{{Paragraph 2 1 10 '\n\n'}}

---

## Warning  

{{Paragraph 1 1 10 '\n\n'}}

---

## Requirements

Golang you will need to install go 1.20

[{{$saved_url:=URL}}{{$saved_url}}]({{$saved_url}})

## How to install

{{Paragraph 1 1 5 '\n\n'}}

<details>
<summary>1. {{SentenceSimple}} </summary>

'''
go install {{$saved_url}}
'''

</details>

---

## TODO

{{Paragraph 1 1 10 '\n\n'}}.

{{range $y := IntRange 1 .Lines}}
1. {{PhraseAdverb}}
{{end}}`

//**************
// HTML Template
//**************
const HTML_TEMPLATE = `<!DOCTYPE html>
<html>
	<body>
		<img src="data:image/png;base64, {{Base64Enc (ImagePng 100 100)}}" alt="Red dot" />
		<h1 style="font-family:verdana;">{{SentenceSimple}}</h1>
		<p style="font-family:courier;">{{Paragraph 1 5 20 '\n\n'}}</p>
		{{range $y := IntRange 1 .Lines}}
		<h1 style="font-family:verdana;">{{SentenceSimple}}</h1>
		<p style="font-family:courier;">{{Paragraph 1 5 20 '\n\n'}}</p>
		{{end}}
	</body>
</html>`

// replaceSingleQuote replaces single quote with multiline quote
func GetTemplateData(template string) string {

	switch template {
	case "email":
		return strings.ReplaceAll(EMAIL_TEMPLATE, "'", "`")
	case "fixed_format":
		return strings.ReplaceAll(FIXED_FORMAT_TEMPLATE, "'", "`")
	case "markdown":
		return strings.ReplaceAll(MARKDOWN_TEMPLATE, "'", "`")
	case "html":
		return strings.ReplaceAll(HTML_TEMPLATE, "'", "`")
	default:
		return strings.ReplaceAll(EMAIL_TEMPLATE, "'", "`")
	}
}

// template contains golang templates
var Template = map[string][]string{
	"document": {GetTemplateData("email"), GetTemplateData("fixed_format"), GetTemplateData("markdown"), GetTemplateData("html")},
}
