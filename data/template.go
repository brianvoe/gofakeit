package data

// Email Template
const email_template1 = `{{$saved_from_name:=Name}}{{$saved_to_name:=Name}}From: {{$saved_from_name}} <{{Email}}>
To: {{$saved_to_name}} <{{Email}}>
Subject: {{Phrase}}
Date: {{Date}}
Content-Type: text/plain; charset=UTF-8;

Hi {{$saved_to_name}}

{{Paragraph 3 5 20 '\n\n'}}
Regards {{$saved_from_name}}`

// Markdown Template
const markdown_template1 = `# {{SentenceSimple}}

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
1. {{PhraseAdverb}}{{end}}`

// HTML Template
const html_body1 = `<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8" />
  <title>{{SentenceSimple}}</title>
  <meta name="viewport" content="width=device-width,initial-scale=1" />
  <meta name="description" content="" />
  <link href="https://cdn.jsdelivr.net/npm/modern-normalize@v2.0.0/modern-normalize.min.css" rel="stylesheet">
  <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.2/dist/css/bootstrap.min.css" rel="stylesheet" crossorigin="anonymous">
  <link rel="stylesheet" type="text/css" href="style.css" />
  <link rel="icon" type="image/svg+xml" href="data:image/svg;base64,{{Base64Enc (SVG 16 16)}}">

  <meta name="theme-color" content="">
  <meta property="og:title" content="" />
  <meta property="og:description" content="" />
  <meta property="og:image" content="" />
  <meta name="twitter:card" content="" />
  <meta name="twitter:site" content="" />
  <meta name="twitter:title" content="" />
  <meta name="twitter:description" content="" />
  <meta name="twitter:image" content="" />

</head>
<body>
<body>
  <nav class="navbar navbar-expand-lg bg-body-tertiary">
    <div class="container-fluid">
      <a class="navbar-brand" href="#">Navbar</a>
      <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarNavDropdown" aria-controls="navbarNavDropdown" aria-expanded="false" aria-label="Toggle navigation">
        <span class="navbar-toggler-icon"></span>
      </button>
      <div class="collapse navbar-collapse" id="navbarNavDropdown">
        <ul class="navbar-nav">
          <li class="nav-item">
            <a class="nav-link active" aria-current="page" href="#">Home</a>
          </li>
          <li class="nav-item">
            <a class="nav-link" href="#">Features</a>
          </li>
          <li class="nav-item">
            <a class="nav-link" href="#">Pricing</a>
          </li>
          <li class="nav-item dropdown">
            <a class="nav-link dropdown-toggle" href="#" role="button" data-bs-toggle="dropdown" aria-expanded="false">
              Dropdown link
            </a>
            <ul class="dropdown-menu">
              <li><a class="dropdown-item" href="#">Action</a></li>
              <li><a class="dropdown-item" href="#">Another action</a></li>
              <li><a class="dropdown-item" href="#">Something else here</a></li>
            </ul>
          </li>
        </ul>
      </div>
    </div>
  </nav>
<br>
{{range $y := IntRange 1 .Lines}}
<h1>Section {{$y}}</h1>{{TemplateHtmlContent}}
{{end}}
<br>
	<footer>
    <small>© <script>document.write(new Date().getFullYear())</script> Your company name. All Rights Reserved.</small>
  </footer>
  <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.3.2/dist/js/bootstrap.bundle.min.js" crossorigin="anonymous"></script>
  <script src="app.js"></script>
</body>
</html>`

const html_content1 = `<h2>Simple Content and Image</h2>
<p>{{Paragraph 3 5 20 '<br>'}}</p>
<img src="data:image/png;base64,{{Base64Enc (ImagePng 30 30)}}" width="200" height="200" alt=""><hr><br><br>`

const html_content2 = `{{$person:=Person}}<h2>Form Populated With a Persons Details</h2><br><form action="#">
<label for="firstName" class="form-label">First name:</label><br>
<input type="text" class="form-control" name="firstName" value="{{$person.FirstName}}"><br>
<label for="lastName" class="form-label">Last name:</label><br>
<input type="text" class="form-control" name="lastName" value="{{$person.LastName}}""><br>
<label for="email" class="form-label">Email:</label><br>
<input type="email" class="form-control" name="email" value="{{$person.Contact.Email}}"><br>
<label for="phone" class="form-label">Phone:</label><br>
<input type="tel" class="form-control" name="phone" value="{{$person.Contact.Phone}}"><br>
<label for="company" class="form-label">Company:</label><br>
<input type="text" class="form-control" name="company" value="{{$person.Job.Company}}"><br>
<input type="submit" class="btn btn-primary" value="Submit">
</form><hr><br><br>`

const html_content3 = `<h2>Accordion Layout</h2><br><div class="accordion" id="accordionExample">
<div class="accordion-item">
  <h2 class="accordion-header" id="headingOne">
	<button class="accordion-button" type="button" data-bs-toggle="collapse" data-bs-target="#collapseOne" aria-expanded="true" aria-controls="collapseOne">
	{{SentenceSimple}} Heading 1
	</button>
  </h2>
  <div id="collapseOne" class="accordion-collapse collapse show" aria-labelledby="headingOne" data-bs-parent="#accordionExample">
	<div class="accordion-body">
	{{SentenceSimple}} Content 1
	</div>
  </div>
</div>
<div class="accordion-item">
  <h2 class="accordion-header" id="headingTwo">
	<button class="accordion-button collapsed" type="button" data-bs-toggle="collapse" data-bs-target="#collapseTwo" aria-expanded="false" aria-controls="collapseTwo">
	{{SentenceSimple}} Heading 2
	</button>
  </h2>
  <div id="collapseTwo" class="accordion-collapse collapse" aria-labelledby="headingTwo" data-bs-parent="#accordionExample">
	<div class="accordion-body">
	{{SentenceSimple}} Content 2
	</div>
  </div>
</div>
<div class="accordion-item">
  <h2 class="accordion-header" id="headingThree">
	<button class="accordion-button collapsed" type="button" data-bs-toggle="collapse" data-bs-target="#collapseThree" aria-expanded="false" aria-controls="collapseThree">
	{{SentenceSimple}} Heading 3
	</button>
  </h2>
  <div id="collapseThree" class="accordion-collapse collapse" aria-labelledby="headingThree" data-bs-parent="#accordionExample">
	<div class="accordion-body">
	{{SentenceSimple}} Content 3
	</div>
  </div>
</div>
</div><hr><br><br>`

const html_content4 = `<h2>Carousel Layout</h2><br><div id="carouselExampleCaptions" class="carousel slide">
<div class="carousel-indicators">
  <button type="button" data-bs-target="#carouselExampleCaptions" data-bs-slide-to="0" class="active" aria-current="true" aria-label="Slide 1"></button>
  <button type="button" data-bs-target="#carouselExampleCaptions" data-bs-slide-to="1" aria-label="Slide 2"></button>
  <button type="button" data-bs-target="#carouselExampleCaptions" data-bs-slide-to="2" aria-label="Slide 3"></button>
</div>
<div class="carousel-inner">
  <div class="carousel-item active">
	<img src="data:image/png;base64,{{Base64Enc (ImagePng 30 30)}}" class="d-block w-100" alt="">
	<div class="carousel-caption d-none d-md-block">
	  <h5>{{SentenceSimple}}</h5>
	  <p>{{Paragraph 1 5 10 '<br>'}}</p>
	</div>
  </div>
  <div class="carousel-item">
	<img src="data:image/png;base64,{{Base64Enc (ImagePng 30 30)}}" class="d-block w-100" alt="">
	<div class="carousel-caption d-none d-md-block">
	<h5>{{SentenceSimple}}</h5>
	<p>{{Paragraph 1 5 10 '<br>'}}</p>
	</div>
  </div>
  <div class="carousel-item">
	<img src="data:image/png;base64,{{Base64Enc (ImagePng 30 30)}}" class="d-block w-100" alt="">
	<div class="carousel-caption d-none d-md-block">
	<h5>{{SentenceSimple}}</h5>
	<p>{{Paragraph 1 5 10 '<br>'}}</p>
	</div>
  </div>
</div>
<button class="carousel-control-prev" type="button" data-bs-target="#carouselExampleCaptions" data-bs-slide="prev">
  <span class="carousel-control-prev-icon" aria-hidden="true"></span>
  <span class="visually-hidden">Previous</span>
</button>
<button class="carousel-control-next" type="button" data-bs-target="#carouselExampleCaptions" data-bs-slide="next">
  <span class="carousel-control-next-icon" aria-hidden="true"></span>
  <span class="visually-hidden">Next</span>
</button>
</div><hr><br><br>`

const html_content5 = `<h2>Table data Layout</h2><br><table class="table">
<tr>
  <th>Animal</th>
  <th>Name</td>
  <th>Age</th>
</tr>
{{range $y := IntRange 1 (Number 1 20)}}<tr>
  <td>{{Animal}}</td>
  <td>{{PetName}}</td>
  <td>{{Number 1 30}}</td>
</tr>{{end}}
</table><hr><br><br>`

const html_content6 = `<h2>Person Profile</h2><br>
<div class="container">
    <div class="row">
        <div class="col-lg-12 mb-4 mb-sm-5">
            <div class="card card-style1 border-0">
                <div class="card-body p-1-9 p-sm-2-3 p-md-6 p-lg-7">
                    <div class="row align-items-center">
                        <div class="col-lg-6 mb-4 mb-lg-0">
                            <img src="data:image/png;base64,{{Base64Enc (ImagePng 30 30)}}" width="50%" height="50%" alt="...">
                        </div>
                        <div class="col-lg-6 px-xl-10">
                            <div class="bg-secondary d-lg-inline-block py-1-9 px-1-9 px-sm-6 mb-1-9 rounded">
                                <h3 class="h2 text-white mb-2">{{$person:=Person}}{{$person.FirstName}} {{$person.LastName}}</h3>
                                <span class="text-primary">{{$person.Job.Title}}</span>
                            </div>
                            <ul class="list-unstyled mb-1-9">
                                <li class="mb-2 mb-xl-3 display-28"><span class="display-26 text-secondary me-2 font-weight-600">Position:</span> {{$person.Job.Title}}</li>
                                <li class="mb-2 mb-xl-3 display-28"><span class="display-26 text-secondary me-2 font-weight-600">Level:</span> {{$person.Job.Title}}</li>
                                <li class="mb-2 mb-xl-3 display-28"><span class="display-26 text-secondary me-2 font-weight-600">Email:</span> {{$person.Contact.Email}}</li>
                                <li class="mb-2 mb-xl-3 display-28"><span class="display-26 text-secondary me-2 font-weight-600">Website:</span> www.{{lc (replace $person.Job.Company ' ' '-')}}.com</li>
                                <li class="display-28"><span class="display-26 text-secondary me-2 font-weight-600">Phone:</span>{{$person.Contact.Phone}}</li>
                            </ul>
                            <ul class="social-icon-style1 list-unstyled mb-0 ps-0">
                                <li><a href="#!"><i class="ti-twitter-alt"></i></a></li>
                                <li><a href="#!"><i class="ti-facebook"></i></a></li>
                                <li><a href="#!"><i class="ti-pinterest"></i></a></li>
                                <li><a href="#!"><i class="ti-instagram"></i></a></li>
                            </ul>
                        </div>
                    </div>
                </div>
            </div>
        </div>
        <div class="col-lg-12 mb-4 mb-sm-5">
            <div>
                <span class="section-title text-primary mb-3 mb-sm-4">About Me</span>
                <p>{{Paragraph 1 1 20 '<br>'}}.</p>
                <p class="mb-0">{{Paragraph 1 1 20 '<br>'}}.</p>
            </div>
        </div>
       
    </div>
</div><hr><br><br>`

// template contains golang templates
var Template = map[string][]string{
	"email":        {email_template1},
	"html":         {html_body1},
	"html_content": {html_content1, html_content2, html_content3, html_content4, html_content5, html_content6},
	"markdown":     {markdown_template1},
	"document":     {email_template1, html_body1, markdown_template1},
}
