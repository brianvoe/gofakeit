package data

// Email Template
const email_template1 = `{{RandomString (ListS 'Hi' 'Hello' 'Dear' 'Good morning' 'Good afternoon' 'Good evening' )}}, {{$saved_to:=Person}}{{$saved_to.FirstName}}\n
{{$res:=CreateListResult 1 7}}{{range $y := IntRange 1 .Lines}}{{$res = ListResult ($res) 1 7 true}}{{if eq $res.Value 1}}{{Paragraph (Number 1 3) (Number 3 5) (Number 10 20) '\n\n'}}
{{end}}{{if eq $res.Value 2}}{{Question}}\n{{end}}{{if eq $res.Value 3}}{{Quote}}\n{{end}}{{if eq $res.Value 4}}{{HipsterParagraph (Number 1 3) (Number 3 5) (Number 10 20) '\n\n'}}
{{end}}{{if eq $res.Value 5}}{{RandomString (ListS (conc 'Have you seen' (MovieName) 'it is' (AdjectiveDescriptive)) (conc 'Have you seen the' (MovieName) '?') (conc 'Do you want to watch' (MovieName) (AdverbTimeDefinite) '?') )}}
{{end}}{{if eq $res.Value 6}}{{conc (PronounPersonal) (VerbHelping) (VerbIntransitive) (AdverbTimeDefinite) (PronounDemonstrative) (AdverbFrequencyIndefinite) 'Happens' (AdverbTimeDefinite)}}\n{{end}}{{if eq $res.Value 7}}{{HipsterSentence (Number 10 20)}}{{end}}{{end}}
{{$saved_from:=Person}}{{RandomString (ListS 'Regards' 'Sincerely' 'Best wishes' 'Cheers' 'Take care' 'Best' 'Thank you' 'I appreciate your help' 'I appreciate your feedback'  'I appreciate your input')}}\n{{$saved_from.FirstName}}{{if Bool}} {{$saved_from.LastName}}{{end}}\n{{if Bool}}{{if Bool}}\nCompany: {{$saved_from.Job.Company}}\nJob Role: {{$saved_from.Job.Title}}\n{{end}}
{{if Bool}}Address: {{$saved_from.Address.Address}}\nCity: {{$saved_from.Address.City}}\nState: {{$saved_from.Address.State}}\nZip: {{$saved_from.Address.Zip}}{{end}}Phone: {{$saved_from.Contact.Phone}}\nEmail: {{$saved_from.Contact.Email}}{{end}}`

// Markdown Template
const markdown_body1 = `{{$res:=CreateListResult 1 7}}
{{range $y := IntRange 1 .Lines}}{{$res = ListResult ($res) 1 7 true}}{{if eq $res.Value 1}}# Paragraph

{{Paragraph (Number 1 5) (Number 1 5) (Number 1 30) '\n\n'}}

---
{{end}}{{if eq $res.Value 2}}# Block Quote

{{Paragraph (Number 1 5) (Number 1 5) (Number 1 30) '\n\n'}}

---
{{end}}{{if eq $res.Value 3}}## Details

{{Paragraph (Number 1 5) (Number 1 5) (Number 1 30) '\n\n'}}

<details>
<summary>{{SentenceSimple}} </summary>

'''
go install {{URL}}
'''

</details>

---
{{end}}{{if eq $res.Value 4}}## Url

Golang you will need to install

{{range $y := IntRange 1 (Number 1 10)}}[{{$saved_url:=URL}}{{$saved_url}}]({{$saved_url}})
{{end}}
---
{{end}}{{if eq $res.Value 5}}## LISTS

{{Paragraph (Number 1 5) (Number 1 5) (Number 1 30) '\n\n'}}.

{{range $y := IntRange 1 (Number 1 10)}}1. {{PhraseVerb}}
{{end}}
---
{{end}}{{if eq $res.Value 6}}## Images

{{Paragraph (Number 1 5) (Number 1 5) (Number 10 30) '\n\n'}}.

{{range $y := IntRange 1 (Number 1 2)}}<image src="data:image/png;base64,{{Base64Enc (ImagePng 50 50)}}"" width="200" height="200" alt/><em>{{AdjectiveProper}}</em>
{{end}}
---
{{end}}{{if eq $res.Value 7}}## Code Examples {{$data_type:=RandomString (ListS 'Int' 'Float32' 'Float64' 'Uint32' 'Int63' )}}

'''
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// GenerateRandom{{$data_type}} generates a random {{lc ($data_type)}} between min and max (inclusive).
//
// Parameters:
// min ({{lc ($data_type)}}): The minimum value of the random {{lc ($data_type)}}.
// max ({{lc ($data_type)}}): The maximum value of the random {{lc ($data_type)}}.
//
// Returns:
// int: A random {{lc ($data_type)}} between min and max (inclusive).
func GenerateRandom{{$data_type}}(min, max {{lc ($data_type)}}) {{lc ($data_type)}} {
	rand.Seed(time.Now().UnixNano())
	return min + rand.{{$data_type}}()*(max-min)
}

func main() {
	random{{$data_type}} := GenerateRandom{{$data_type}} (1, 10)
	fmt.Printf("Random {{$data_type}}: %d'|n'", random{{$data_type}})
}

'''

---{{end}}
{{end}}`

// template contains golang templates
var Template = map[string][]string{
	"email":    {email_template1},
	"markdown": {markdown_body1},
	"document": {email_template1, markdown_body1},
}
