package gofakeit

import (
	"fmt"
	"strings"
	"testing"

	"github.com/brianvoe/gofakeit/v6/data"
)

// Template examples and tests

func ExampleTemplate() {
	// Make sure we get the same results every time
	Seed(11)
	globalFaker.Rand.Seed(11)

	value, err := Template("{{range $y := IntRange 1 .Lines}}\n{{FirstName}} {{LastName}}{{end}}", 4)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(value))

	// Output:
	// Markus Moen
	// Alayna Wuckert
	// Lura Lockman
	// Sylvan Mraz

}

func ExampleFaker_Template() {
	// Make sure we get the same results every time
	f := New(11)
	globalFaker.Rand.Seed(11)
	value, err := f.Template("{{range $y := IntRange 1 .Lines}}\n{{FirstName}} {{LastName}}{{end}}", 3)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(value))

	// Output:
	// Markus Moen
	// Alayna Wuckert
	// Lura Lockman

}

func TestTemplateLookup(t *testing.T) {
	faker := New(11)
	globalFaker.Rand.Seed(11)
	info := GetFuncLookup("template")

	m := MapParams{
		"template": {"{{range $y := IntRange 1 .Lines}}{{FirstName}} {{LastName}}\n{{end}}"},
		"lines":    {"5"},
	}

	value, err := info.Generate(faker.Rand, &m, info)
	if err != nil {
		t.Fatal(err.Error())
	}

	if value != "Markus Moen\nAlayna Wuckert\nLura Lockman\nSylvan Mraz\nPaolo Rutherford\n" {
		t.Error("Expected `Markus Moen Pagac`, got ", value)
	}

}

func TestTemplateNoTemplateParam(t *testing.T) {
	f := New(11)
	globalFaker.Rand.Seed(11)
	value, err := f.Template("", 3)

	if err != nil {
		t.Error("Error with template ", err)
	}
	if string(value) == "" {
		t.Error("Expected a document, got nothing")
	}
}

func ExampleTemplateDocument() {
	// Make sure we get the same results every time
	Seed(11)
	globalFaker.Rand.Seed(11)

	value, err := TemplateDocument()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(value))

	// Output:
	// From: Alayna Wuckert<lurlinetromp@fay.io>
	// To: Hailie Erdman<samsonrippin@emard.info>
	// Subject: winner, winner, chicken dinner
	// Date: 1958-05-27 19:01:50.637039849 +0000 UTC
	// Content-Type: text/html; charset=UTF-8
	//
	// <!DOCTYPE html>
	// <html>
	// <head>
	//     <meta charset="UTF-8">
	//     <meta name="viewport" content="width=device-width" initial-scale="1">
	//     <!--[if !mso]>
	// <meta http-equiv="X-UA-Compatible" content="IE=edge">
	// <![endif]-->
	//     <meta name="x-apple-disable-message-reformatting">
	//     <title></title>
	//     <!--[if mso]>
	// <style>
	//     * { font-family: sans-serif !important; }
	// </style>
	// <![endif]-->
	//     <!--[if !mso]><!-->
	//     <!-- Insert font reference, e.g. <link href="https://fonts.googleapis.com/css?family=Source+Sans+Pro:400,700" rel="stylesheet"> -->
	//     <!--<![endif]-->
	//     <style>
	//         *,
	//         *:after,
	//         *:before {
	//             -webkit-box-sizing: border-box;
	//             -moz-box-sizing: border-box;
	//             box-sizing: border-box;
	//         }
	//
	//         * {
	//             -ms-text-size-adjust: 100%;
	//             -webkit-text-size-adjust: 100%;
	//         }
	//
	//         html,
	//         body,
	//         .document {
	//             font-family: sans-serif;
	//             width: 100% !important;
	//             height: 100% !important;
	//             margin: 0;
	//             padding: 0;
	//         }
	//
	//         body {
	//             -webkit-font-smoothing: antialiased;
	//             -moz-osx-font-smoothing: grayscale;
	//             text-rendering: optimizeLegibility;
	//         }
	//
	//         div[style*="margin: 16px 0"] {
	//             margin: 0 !important;
	//         }
	//
	//         table,
	//         td {
	//             mso-table-lspace: 0pt;
	//             mso-table-rspace: 0pt;
	//         }
	//
	//         table {
	//             border-spacing: 0;
	//             border-collapse: collapse;
	//             table-layout: fixed;
	//             margin: 0 auto;
	//         }
	//
	//         img {
	//             -ms-interpolation-mode: bicubic;
	//             max-width: 100%;
	//             border: 0;
	//         }
	//
	//         *[x-apple-data-detectors] {
	//             color: inherit !important;
	//             text-decoration: none !important;
	//         }
	//
	//         .x-gmail-data-detectors,
	//         .x-gmail-data-detectors *,
	//         .aBn {
	//             border-bottom: 0 !important;
	//             cursor: default !important;
	//         }
	//
	//         .btn {
	//             -webkit-transition: all 200ms ease;
	//             transition: all 200ms ease;
	//         }
	//
	//         .btn:hover {
	//             background-color: dodgerblue;
	//         }
	//
	//         @media screen and (max-width: 750px) {
	//             .container {
	//                 width: 100%;
	//                 margin: auto;
	//             }
	//
	//             .stack {
	//                 display: block;
	//                 width: 100%;
	//                 max-width: 100%;
	//             }
	//         }
	//     </style>
	// </head>
	// <body>
	//     <div style="display: none; max-height: 0px; overflow: hidden;">
	//         you don't dip your pen in the company inkwell
	//     </div>
	//     <div style="display: none; max-height: 30px; overflow: hidden;">
	//         &nbsp;&zwnj;&nbsp;&zwnj;&nbsp;&zwnj;&nbsp;&zwnj;&nbsp;&zwnj;&nbsp;&zwnj;&nbsp;&zwnj;&nbsp;&zwnj;&nbsp;&zwnj;&nbsp;&zwnj;&nbsp;&zwnj;&nbsp;&zwnj;&nbsp;&zwnj;&nbsp;&zwnj;&nbsp;&zwnj;&nbsp;&zwnj;&nbsp;&zwnj;&nbsp;&zwnj;&nbsp;&zwnj;&nbsp;&zwnj;&nbsp;&zwnj;&nbsp;&zwnj;&nbsp;&zwnj;&nbsp;&zwnj;&nbsp;&zwnj;&nbsp;&zwnj;&nbsp;
	//     </div>
	//     <table role="presentation" aria-hidden="true" cellspacing="0" cellpadding="0" border="0" align="left" class="document">
	//         <tr>
	//             <td valign="top">
	//                 <table role="presentation" aria-hidden="true" cellspacing="0" cellpadding="0" border="0" align="left" width="750" class="container">
	//                     <tr>
	//                         <td>
	//                             <table role="presentation" aria-hidden="true" cellspacing="0" cellpadding="0" border="0" align="center" width="100%">
	//                                 <tr>
	//                                     <td>
	//                                         Hi Hailie
	//                                         <br>
	//                                         <br>
	//                                         Troop yesterday whom him plain caravan nearby reluctantly anyone words at victorious boxers each does auspicious yourselves first soup tomorrow. This that must conclude anyway some yearly who cough laugh himself both yet rarely me dolphin intensely block would leap. Plane us first then down them eager would hundred super throughout animal yet themselves been group flock shake part purchase. Up usually it her none it hers boat what their there Turkmen moreover one Lebanese to brace these shower in. It everybody should whatever those uninterested several nobody yours thankful innocent power to any from its few luxury none boy.\r\rReligion themselves it there might must near theirs mine thing tonight outside rapidly spotted solemnly finally been did confusion these. Were nobody early silently company quarterly American there time many French a each anybody rather paint ours did tonight remove. First including chair your at in by cackle whose they yearly wisdom nightly nightly when finally without oxygen scold what. Silence go time behind example me soon grade purely that heavy annually our whoever your eventually yearly gleaming theirs child. Annually ours problem slavery someone brother instance could movement otherwise way now disturbed to sandwich someone its honour what whichever.\r\rContrary from belief this upon at most homeless elsewhere has yearly under since where Californian all in today generally myself. After stupid highly heavy here lately his who generally from substantial himself poised cost formerly but spoon words Egyptian i.e.. Me tonight place few why from somebody hungrily mine were soon hail then you themselves drab when behind case ours. Cost couple consequently in those daily had anywhere anything what in bouquet which annually as Cypriot this our sand tightly. We first their staff invention however whoever itself over this pair smoke yourself so circumstances despite could project did embarrassed.
	//                                         <br>
	//                                         <br>
	//                                         Regards Alayna Wuckert
	//                                     </td>
	//                                 </tr>
	//                             </table>
	//                         </td>
	//                     </tr>
	//                 </table>
	//             </td>
	//         </tr>
	//         <tr>
	//             <td valign="bottom">
	//                 <table role="presentation" aria-hidden="true" cellspacing="0" cellpadding="0" border="0" align="left" width="750" class="container">
	//                     <tr>
	//                         <td>
	//                             9063 East Riverstad, Plano, Vermont 59145
	//                         </td>
	//                     </tr>
	//                     <tr>
	//                         <td>
	//                             Plano
	//                         </td>
	//                     </tr>
	//                     <tr>
	//                         <td>
	//                             Vermont
	//                         </td>
	//                     </tr>
	//                     <tr>
	//                         <td>
	//                             59145
	//                         </td>
	//                     </tr>
	//                     <tr>
	//                         <td>
	//                             2320202761
	//                         </td>
	//                     </tr>
	//                     <tr>
	//                         <td>
	//                             <a href="mail:lurlinetromp@fay.io">lurlinetromp@fay.io</a>"
	//                         </td>
	//                     </tr>
	//                 </table>
	//             </td>
	//             <td valign="bottom">
	//                 <!-- Unsubscribe link -->
	//                 <unsubscribe><a href="https://www.districtusers.io/grow/recontextualize">Unsubscribe</a></unsubscribe>
	//             </td>
	//         </tr>
	//     </table>
	// </body>
	// </html>
}

func ExampleFaker_TemplateDocument() {
	// Make sure we get the same results every time
	f := New(11)
	globalFaker.Rand.Seed(11)
	value, err := f.TemplateDocument()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(value))

	// Output:
	// <!DOCTYPE html>
	// <html lang="en">
	// <head>
	//     <meta charset="UTF-8" />
	//     <title>The tribe indeed swiftly laugh.</title>
	//     <meta name="viewport" content="width=device-width,initial-scale=1" />
	//     <meta name="description" content="" />
	//     <link href="https://cdn.jsdelivr.net/npm/modern-normalize@v2.0.0/modern-normalize.min.css" rel="stylesheet">
	//     <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.2/dist/css/bootstrap.min.css" rel="stylesheet" crossorigin="anonymous">
	//     <link rel="stylesheet" type="text/css" href="style.css" />
	//     <link rel="icon" type="image/svg+xml" href="data:image/svg;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHZpZXdCb3g9IjAgMCAxNiAxNiIgd2lkdGg9IjE2IiBoZWlnaHQ9IjE2Ij48cmVjdCB4PSIwIiB5PSIwIiB3aWR0aD0iMTAwJSIgaGVpZ2h0PSIxMDAlIiBmaWxsPSIjRkZGRkZGIiAvPjwvc3ZnPg==">
	// </head>
	// <body>
	//     <body>
	//         <nav class="navbar navbar-expand-lg bg-body-tertiary">
	//             <div class="container-fluid">
	//                 <a class="navbar-brand" href="#">Navbar</a>
	//                 <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarNavDropdown" aria-controls="navbarNavDropdown" aria-expanded="false" aria-label="Toggle navigation">
	//                     <span class="navbar-toggler-icon"></span>
	//                 </button>
	//                 <div class="collapse navbar-collapse" id="navbarNavDropdown">
	//                     <ul class="navbar-nav">
	//                         <li class="nav-item">
	//                             <a class="nav-link active" aria-current="page" href="#">Home</a>
	//                         </li>
	//                         <li class="nav-item">
	//                             <a class="nav-link" href="#">Features</a>
	//                         </li>
	//                         <li class="nav-item">
	//                             <a class="nav-link" href="#">Pricing</a>
	//                         </li>
	//                         <li class="nav-item dropdown">
	//                             <a class="nav-link dropdown-toggle" href="#" role="button" data-bs-toggle="dropdown" aria-expanded="false">
	//                                 Dropdown link
	//                             </a>
	//                             <ul class="dropdown-menu">
	//                                 <li><a class="dropdown-item" href="#">Action</a></li>
	//                                 <li><a class="dropdown-item" href="#">Another action</a></li>
	//                                 <li><a class="dropdown-item" href="#">Something else here</a></li>
	//                             </ul>
	//                         </li>
	//                     </ul>
	//                 </div>
	//             </div>
	//         </nav>
	//         <br>
	//         <h1>Section 1</h1><h2>Carousel Layout</h2><br>
	// <div id="carouselExampleCaptions" class="carousel slide">
	//     <div class="carousel-indicators">
	//         <button type="button" data-bs-target="#carouselExampleCaptions" data-bs-slide-to="0" class="active" aria-current="true" aria-label="Slide 1"></button>
	//         <button type="button" data-bs-target="#carouselExampleCaptions" data-bs-slide-to="1" aria-label="Slide 2"></button>
	//         <button type="button" data-bs-target="#carouselExampleCaptions" data-bs-slide-to="2" aria-label="Slide 3"></button>
	//     </div>
	//     <div class="carousel-inner">
	//         <div class="carousel-item active">
	//             <img src="data:image/png;base64,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" class="d-block w-100" alt="">
	//             <div class="carousel-caption d-none d-md-block">
	//                 <h5>The week greatly well buy a wild trip.</h5>
	//                 <p>Who additionally him now did it many being loss have. Himself that under alone pack about this motionless other child. Tomorrow defiant understimate ours theirs this virtually cooperative labour Uzbek. Choir in annually since them soon time firstly between here. Would nightly cash hourly neither hers hospitality what me occasionally.</p>
	//             </div>
	//         </div>
	//         <div class="carousel-item">
	//             <img src="data:image/png;base64,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" class="d-block w-100" alt="">
	//             <div class="carousel-caption d-none d-md-block">
	//                 <h5>Tender week watch class softly.</h5>
	//                 <p>Here might its knit think still his Burmese glasses where. E.g. range troupe off can gang now how Ecuadorian this. Her disregard whom until these handle crowded few many cut. Do tense there whom this your his from last his. Bravely e.g. line whichever down sit who whose we advertising.</p>
	//             </div>
	//         </div>
	//         <div class="carousel-item">
	//             <img src="data:image/png;base64,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" class="d-block w-100" alt="">
	//             <div class="carousel-caption d-none d-md-block">
	//                 <h5>Bored problem extremely frankly play poorly.</h5>
	//                 <p>Myself of ourselves we then what sunglasses of fiction anybody. Ever American this all accordingly instance may provided couple Cormoran. Occasionally over agree stack great anyone recently after how then. From airport yesterday including fact words while than virtually may. Whomever face eye how in where those those has these.</p>
	//             </div>
	//         </div>
	//     </div>
	//     <button class="carousel-control-prev" type="button" data-bs-target="#carouselExampleCaptions" data-bs-slide="prev">
	//         <span class="carousel-control-prev-icon" aria-hidden="true"></span>
	//         <span class="visually-hidden">Previous</span>
	//     </button>
	//     <button class="carousel-control-next" type="button" data-bs-target="#carouselExampleCaptions" data-bs-slide="next">
	//         <span class="carousel-control-next-icon" aria-hidden="true"></span>
	//         <span class="visually-hidden">Next</span>
	//     </button>
	// </div>
	// <hr><br><br>
	//         <br>
	//         <footer>
	//             <small>©
	//                 <script>document.write(new Date().getFullYear())</script> PYA Analytics. All Rights Reserved.
	//             </small>
	//         </footer>
	//         <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.3.2/dist/js/bootstrap.bundle.min.js" crossorigin="anonymous"></script>
	//         <script src="app.js"></script>
	//     </body>
	// </html>

}

func TestTemplateDocumentLookup(t *testing.T) {
	faker := New(11)
	globalFaker.Rand.Seed(11)
	info := GetFuncLookup("template_document")

	m := MapParams{
		"lines": {"5"},
	}

	value, err := info.Generate(faker.Rand, &m, info)
	if err != nil {
		t.Fatal(err.Error())
	}

	if !strings.Contains(value.(string), "<title>The tribe indeed swiftly laugh.</title>") {
		t.Error("Expected `<title>The tribe indeed swiftly laugh.</title>`, got ", value)
	}

}

// TemplateEmail examples and tests

func ExampleTemplateEmail() {
	// Make sure we get the same results every time
	Seed(11)
	globalFaker.Rand.Seed(11)

	value, err := TemplateEmail()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(value))

	// Output:
	// From: Marcel Pagac<lurlinetromp@fay.io>
	// To: Hailie Erdman<samsonrippin@emard.info>
	// Subject: winner, winner, chicken dinner
	// Date: 1958-05-27 19:01:50.637039849 +0000 UTC
	// Content-Type: text/plain; charset=UTF-8;
	//
	// Hi Hailie
	//
	// Software quarterly rather punctuation yellow where several his orchard to frequently hence victorious boxers each does auspicious yourselves first soup. Tomorrow this that must conclude anyway some yearly who cough laugh himself both yet rarely me dolphin intensely block would. Leap plane us first then down them eager would hundred super throughout animal yet themselves been group flock shake part. Purchase up usually it her none it hers boat what their there Turkmen moreover one Lebanese to brace these shower. In it everybody should whatever those uninterested several nobody yours thankful innocent power to any from its few luxury none.\r\rBoy religion themselves it there might must near theirs mine thing tonight outside rapidly spotted solemnly finally been did confusion. These were nobody early silently company quarterly American there time many French a each anybody rather paint ours did tonight. Remove first including chair your at in by cackle whose they yearly wisdom nightly nightly when finally without oxygen scold. What silence go time behind example me soon grade purely that heavy annually our whoever your eventually yearly gleaming theirs. Child annually ours problem slavery someone brother instance could movement otherwise way now disturbed to sandwich someone its honour what.\r\rWhichever contrary from belief this upon at most homeless elsewhere has yearly under since where Californian all in today generally. Myself after stupid highly heavy here lately his who generally from substantial himself poised cost formerly but spoon words Egyptian. I.e. me tonight place few why from somebody hungrily mine were soon hail then you themselves drab when behind case. Ours cost couple consequently in those daily had anywhere anything what in bouquet which annually as Cypriot this our sand. Tightly we first their staff invention however whoever itself over this pair smoke yourself so circumstances despite could project did.
	//
	// Regards Marcel Pagac

}

func ExampleFaker_TemplateEmail() {
	// Make sure we get the same results every time
	f := New(11)
	globalFaker.Rand.Seed(11)
	value, err := f.TemplateEmail()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(value))

	// Output:
	// From: Markus Moen<lamarkoelpin@heaney.biz>
	// To: Matteo Krajcik<samsonrippin@emard.info>
	// Subject: winner, winner, chicken dinner
	// Date: 1958-05-27 19:01:50.637039849 +0000 UTC
	// Content-Type: text/plain; charset=UTF-8;
	//
	// Hi Matteo
	//
	// Software quarterly rather punctuation yellow where several his orchard to frequently hence victorious boxers each does auspicious yourselves first soup. Tomorrow this that must conclude anyway some yearly who cough laugh himself both yet rarely me dolphin intensely block would. Leap plane us first then down them eager would hundred super throughout animal yet themselves been group flock shake part. Purchase up usually it her none it hers boat what their there Turkmen moreover one Lebanese to brace these shower. In it everybody should whatever those uninterested several nobody yours thankful innocent power to any from its few luxury none.\r\rBoy religion themselves it there might must near theirs mine thing tonight outside rapidly spotted solemnly finally been did confusion. These were nobody early silently company quarterly American there time many French a each anybody rather paint ours did tonight. Remove first including chair your at in by cackle whose they yearly wisdom nightly nightly when finally without oxygen scold. What silence go time behind example me soon grade purely that heavy annually our whoever your eventually yearly gleaming theirs. Child annually ours problem slavery someone brother instance could movement otherwise way now disturbed to sandwich someone its honour what.\r\rWhichever contrary from belief this upon at most homeless elsewhere has yearly under since where Californian all in today generally. Myself after stupid highly heavy here lately his who generally from substantial himself poised cost formerly but spoon words Egyptian. I.e. me tonight place few why from somebody hungrily mine were soon hail then you themselves drab when behind case. Ours cost couple consequently in those daily had anywhere anything what in bouquet which annually as Cypriot this our sand. Tightly we first their staff invention however whoever itself over this pair smoke yourself so circumstances despite could project did.
	//
	// Regards Markus Moen

}

func TestTemplateEmailLookup(t *testing.T) {
	faker := New(11)
	globalFaker.Rand.Seed(11)
	info := GetFuncLookup("template_email")

	m := MapParams{}

	value, err := info.Generate(faker.Rand, &m, info)
	if err != nil {
		t.Fatal(err.Error())
	}

	if !strings.Contains(value.(string), "To: Hailie Erdman<samsonrippin@emard.info>") {
		t.Error("Expected `To: Hailie Erdman<samsonrippin@emard.info>`, got ", value)
	}

}

func TestTemplateEmail(t *testing.T) {
	f := New(11)
	globalFaker.Rand.Seed(11)
	value, err := f.TemplateEmail()
	if err != nil {
		t.Error(err)
	}

	if string(value) == "" {
		t.Error("Expected a document, got nothing")
	}
}

// TemplateMarkdown examples and tests

func ExampleTemplateMarkdown() {
	// Make sure we get the same results every time
	Seed(11)
	globalFaker.Rand.Seed(11)

	value, err := TemplateMarkdown(3)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(value))

	// Output:
	// ## Images
	//
	// Generally half where everything niche backwards caused quarterly without week..
	//
	// <image src="data:image/png;base64,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"" width="200" height="200" alt/><em>Lilliputian</em>
	//
	// ---
	// ## Details
	//
	// Joy highly Intelligent nature frequently.
	//
	// <details>
	// <summary>The logic innocently win the tough freedom. </summary>
	//
	// ```
	// go install http://www.forwardinitiatives.io/extensible/rich
	// ```
	//
	// </details>
	//
	// ---
	// ## LISTS
	//
	// Cello empty whenever that love there nightly reluctantly fatally harvest..
	//
	// 1. wait a cautious cackle daringly from a wild life
	// 1. smell a purple string
	// 1. greatly painfully play an expensive friendship wisely up a muster obediently
	// 1. perfectly sing deeply powerfully rather fiercely
	// 1. jump too rapidly by the smoggy troop
	// 1. rather calmly sit positively wildly badly
	// 1. rapidly stack
	// 1. climb purely courageously thoroughly speedily
	//
	// ---
	//
}

func ExampleFaker_TemplateMarkdown() {
	// Make sure we get the same results every time
	f := New(11)
	globalFaker.Rand.Seed(11)
	value, err := f.TemplateMarkdown(4)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(value))

	// Output:
	// # Paragraph
	//
	// Fly had result everything niche backwards caused quarterly without week it hungry thing someone him regularly today whomever this revolt. Hence from his timing as quantity us these yours live these frantic not may another how this ours his them. Those whose them batch its Iraqi most that few abroad cheese this whereas next how there gorgeous genetics time choir. Fiction therefore yourselves am those infrequently heap software quarterly rather punctuation yellow where several his orchard to frequently hence victorious. Boxers each does auspicious yourselves first soup tomorrow this that must conclude anyway some yearly who cough laugh himself both.
	//
	// ---
	// ## Url
	//
	// Golang you will need to install
	//
	// [http://www.dynamicrobust.io/e-tailers/ubiquitous/cultivate/transition](http://www.dynamicrobust.io/e-tailers/ubiquitous/cultivate/transition)
	// [http://www.directsynergistic.net/back-end/revolutionary/matrix](http://www.directsynergistic.net/back-end/revolutionary/matrix)
	// [http://www.investorcross-media.io/revolutionary/target](http://www.investorcross-media.io/revolutionary/target)
	// [https://www.principalmarkets.org/orchestrate/interactive/rich](https://www.principalmarkets.org/orchestrate/interactive/rich)
	// [http://www.centraloptimize.biz/customized/synergize](http://www.centraloptimize.biz/customized/synergize)
	//
	// ---
	// ## Images
	//
	// Whose reel themselves which can vanish rather these each previously..
	//
	// <image src="data:image/png;base64,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"" width="200" height="200" alt/><em>Icelandic</em>
	//
	// ---
	// ## Url
	//
	// Golang you will need to install
	//
	// [http://www.dynamicenvisioneer.biz/e-business](http://www.dynamicenvisioneer.biz/e-business)
	// [https://www.forwardwebservices.org/incubate/visualize/next-generation](https://www.forwardwebservices.org/incubate/visualize/next-generation)
	// [https://www.chieftransform.org/e-services](https://www.chieftransform.org/e-services)
	// [http://www.districtuser-centric.io/eyeballs/seamless/synergize](http://www.districtuser-centric.io/eyeballs/seamless/synergize)
	// [http://www.dynamicopen-source.biz/grow](http://www.dynamicopen-source.biz/grow)
	// [http://www.globalfacilitate.org/deliver/turn-key/infrastructures](http://www.globalfacilitate.org/deliver/turn-key/infrastructures)
	// [http://www.forwardharness.org/innovate/initiatives](http://www.forwardharness.org/innovate/initiatives)
	// [https://www.internationalsticky.info/maximize/seize/mindshare](https://www.internationalsticky.info/maximize/seize/mindshare)
	// [http://www.chiefenhance.info/e-enable](http://www.chiefenhance.info/e-enable)
	// [http://www.corporatedrive.org/platforms](http://www.corporatedrive.org/platforms)
	//
	// ---
	//

}

func TestTemplateMarkdownLookup(t *testing.T) {
	faker := New(11)
	globalFaker.Rand.Seed(11)
	info := GetFuncLookup("template_markdown")

	m := MapParams{
		"sections": {"5"},
	}

	value, err := info.Generate(faker.Rand, &m, info)
	if err != nil {
		t.Fatal(err.Error())
	}

	if !strings.Contains(value.(string), "Generally half where everything niche") {
		t.Error("Expected `Generally half where everything niche`, got ", value)
	}

}

func TestTemplateMarkdown(t *testing.T) {
	f := New(11)
	globalFaker.Rand.Seed(11)
	value, err := f.TemplateMarkdown(10)
	if err != nil {
		t.Error(err)
	}

	if string(value) == "" {
		t.Error("Expected a document, got nothing")
	}
}

func ExampleTemplateMarkdownContent() {
	// Make sure we get the same results every time
	Seed(11)
	globalFaker.Rand.Seed(11)

	value, err := TemplateMarkdownContent()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(value))

	// Output:
	// # Paragraph
	//
	// Fly had result everything niche backwards caused quarterly without week it hungry thing someone him regularly today whomever this revolt. Hence from his timing as quantity us these yours live these frantic not may another how this ours his them. Those whose them batch its Iraqi most that few abroad cheese this whereas next how there gorgeous genetics time choir. Fiction therefore yourselves am those infrequently heap software quarterly rather punctuation yellow where several his orchard to frequently hence victorious. Boxers each does auspicious yourselves first soup tomorrow this that must conclude anyway some yearly who cough laugh himself both.
	//
	// ---

}

func ExampleFaker_TemplateMarkdownContent() {
	// Make sure we get the same results every time
	f := New(11)
	globalFaker.Rand.Seed(11)
	value, err := f.TemplateMarkdownContent()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(value))

	// Output:
	// # Paragraph
	//
	// There had result everything niche backwards caused quarterly without week it hungry thing someone him regularly today whomever this revolt. Hence from his timing as quantity us these yours live these frantic not may another how this ours his them. Those whose them batch its Iraqi most that few abroad cheese this whereas next how there gorgeous genetics time choir. Fiction therefore yourselves am those infrequently heap software quarterly rather punctuation yellow where several his orchard to frequently hence victorious. Boxers each does auspicious yourselves first soup tomorrow this that must conclude anyway some yearly who cough laugh himself both.
	//
	// ---

}

func TestTemplateMarkdownContentLookup(t *testing.T) {
	faker := New(11)
	globalFaker.Rand.Seed(11)
	info := GetFuncLookup("template_markdown_content")

	m := MapParams{}

	value, err := info.Generate(faker.Rand, &m, info)
	if err != nil {
		t.Fatal(err.Error())
	}

	if !strings.Contains(value.(string), "Fly had result everything niche backwards") {
		t.Error("Expected `Fly had result everything niche backwards`, got ", value)
	}

}

func TestTemplateMarkdownContent(t *testing.T) {
	f := New(11)
	globalFaker.Rand.Seed(11)
	value, err := f.TemplateMarkdownContent()
	if err != nil {
		t.Error(err)
	}

	if string(value) == "" {
		t.Error("Expected a document, got nothing")
	}
}

// TemplateHtml examples and tests

func ExampleTemplateHtml() {
	// Make sure we get the same results every time
	Seed(11)
	globalFaker.Rand.Seed(11)

	value, err := TemplateHtml(1)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(value))

	// Output:
	// <!DOCTYPE html>
	// <html lang="en">
	// <head>
	//     <meta charset="UTF-8" />
	//     <title>The tribe indeed swiftly laugh.</title>
	//     <meta name="viewport" content="width=device-width,initial-scale=1" />
	//     <meta name="description" content="" />
	//     <link href="https://cdn.jsdelivr.net/npm/modern-normalize@v2.0.0/modern-normalize.min.css" rel="stylesheet">
	//     <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.2/dist/css/bootstrap.min.css" rel="stylesheet" crossorigin="anonymous">
	//     <link rel="stylesheet" type="text/css" href="style.css" />
	//     <link rel="icon" type="image/svg+xml" href="data:image/svg;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHZpZXdCb3g9IjAgMCAxNiAxNiIgd2lkdGg9IjE2IiBoZWlnaHQ9IjE2Ij48cmVjdCB4PSIwIiB5PSIwIiB3aWR0aD0iMTAwJSIgaGVpZ2h0PSIxMDAlIiBmaWxsPSIjRkZGRkZGIiAvPjwvc3ZnPg==">
	// </head>
	// <body>
	//     <body>
	//         <nav class="navbar navbar-expand-lg bg-body-tertiary">
	//             <div class="container-fluid">
	//                 <a class="navbar-brand" href="#">Navbar</a>
	//                 <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarNavDropdown" aria-controls="navbarNavDropdown" aria-expanded="false" aria-label="Toggle navigation">
	//                     <span class="navbar-toggler-icon"></span>
	//                 </button>
	//                 <div class="collapse navbar-collapse" id="navbarNavDropdown">
	//                     <ul class="navbar-nav">
	//                         <li class="nav-item">
	//                             <a class="nav-link active" aria-current="page" href="#">Home</a>
	//                         </li>
	//                         <li class="nav-item">
	//                             <a class="nav-link" href="#">Features</a>
	//                         </li>
	//                         <li class="nav-item">
	//                             <a class="nav-link" href="#">Pricing</a>
	//                         </li>
	//                         <li class="nav-item dropdown">
	//                             <a class="nav-link dropdown-toggle" href="#" role="button" data-bs-toggle="dropdown" aria-expanded="false">
	//                                 Dropdown link
	//                             </a>
	//                             <ul class="dropdown-menu">
	//                                 <li><a class="dropdown-item" href="#">Action</a></li>
	//                                 <li><a class="dropdown-item" href="#">Another action</a></li>
	//                                 <li><a class="dropdown-item" href="#">Something else here</a></li>
	//                             </ul>
	//                         </li>
	//                     </ul>
	//                 </div>
	//             </div>
	//         </nav>
	//         <br>
	//         <h1>Section 1</h1><h2>Carousel Layout</h2><br>
	// <div id="carouselExampleCaptions" class="carousel slide">
	//     <div class="carousel-indicators">
	//         <button type="button" data-bs-target="#carouselExampleCaptions" data-bs-slide-to="0" class="active" aria-current="true" aria-label="Slide 1"></button>
	//         <button type="button" data-bs-target="#carouselExampleCaptions" data-bs-slide-to="1" aria-label="Slide 2"></button>
	//         <button type="button" data-bs-target="#carouselExampleCaptions" data-bs-slide-to="2" aria-label="Slide 3"></button>
	//     </div>
	//     <div class="carousel-inner">
	//         <div class="carousel-item active">
	//             <img src="data:image/png;base64,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" class="d-block w-100" alt="">
	//             <div class="carousel-caption d-none d-md-block">
	//                 <h5>The week greatly well buy a wild trip.</h5>
	//                 <p>Who additionally him now did it many being loss have. Himself that under alone pack about this motionless other child. Tomorrow defiant understimate ours theirs this virtually cooperative labour Uzbek. Choir in annually since them soon time firstly between here. Would nightly cash hourly neither hers hospitality what me occasionally.</p>
	//             </div>
	//         </div>
	//         <div class="carousel-item">
	//             <img src="data:image/png;base64,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" class="d-block w-100" alt="">
	//             <div class="carousel-caption d-none d-md-block">
	//                 <h5>Tender week watch class softly.</h5>
	//                 <p>Here might its knit think still his Burmese glasses where. E.g. range troupe off can gang now how Ecuadorian this. Her disregard whom until these handle crowded few many cut. Do tense there whom this your his from last his. Bravely e.g. line whichever down sit who whose we advertising.</p>
	//             </div>
	//         </div>
	//         <div class="carousel-item">
	//             <img src="data:image/png;base64,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" class="d-block w-100" alt="">
	//             <div class="carousel-caption d-none d-md-block">
	//                 <h5>Bored problem extremely frankly play poorly.</h5>
	//                 <p>Myself of ourselves we then what sunglasses of fiction anybody. Ever American this all accordingly instance may provided couple Cormoran. Occasionally over agree stack great anyone recently after how then. From airport yesterday including fact words while than virtually may. Whomever face eye how in where those those has these.</p>
	//             </div>
	//         </div>
	//     </div>
	//     <button class="carousel-control-prev" type="button" data-bs-target="#carouselExampleCaptions" data-bs-slide="prev">
	//         <span class="carousel-control-prev-icon" aria-hidden="true"></span>
	//         <span class="visually-hidden">Previous</span>
	//     </button>
	//     <button class="carousel-control-next" type="button" data-bs-target="#carouselExampleCaptions" data-bs-slide="next">
	//         <span class="carousel-control-next-icon" aria-hidden="true"></span>
	//         <span class="visually-hidden">Next</span>
	//     </button>
	// </div>
	// <hr><br><br>
	//         <br>
	//         <footer>
	//             <small>©
	//                 <script>document.write(new Date().getFullYear())</script> PYA Analytics. All Rights Reserved.
	//             </small>
	//         </footer>
	//         <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.3.2/dist/js/bootstrap.bundle.min.js" crossorigin="anonymous"></script>
	//         <script src="app.js"></script>
	//     </body>
	// </html>

}

func ExampleFaker_TemplateHtml() {
	// Make sure we get the same results every time
	f := New(11)
	globalFaker.Rand.Seed(11)
	value, err := f.TemplateHtml(1)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(value))

	// Output:
	// <!DOCTYPE html>
	// <html lang="en">
	// <head>
	//     <meta charset="UTF-8" />
	//     <title>The purple tribe indeed swiftly laugh.</title>
	//     <meta name="viewport" content="width=device-width,initial-scale=1" />
	//     <meta name="description" content="" />
	//     <link href="https://cdn.jsdelivr.net/npm/modern-normalize@v2.0.0/modern-normalize.min.css" rel="stylesheet">
	//     <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.2/dist/css/bootstrap.min.css" rel="stylesheet" crossorigin="anonymous">
	//     <link rel="stylesheet" type="text/css" href="style.css" />
	//     <link rel="icon" type="image/svg+xml" href="data:image/svg;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHZpZXdCb3g9IjAgMCAxNiAxNiIgd2lkdGg9IjE2IiBoZWlnaHQ9IjE2Ij48cmVjdCB4PSIwIiB5PSIwIiB3aWR0aD0iMTAwJSIgaGVpZ2h0PSIxMDAlIiBmaWxsPSIjRkZGRkZGIiAvPjwvc3ZnPg==">
	// </head>
	// <body>
	//     <body>
	//         <nav class="navbar navbar-expand-lg bg-body-tertiary">
	//             <div class="container-fluid">
	//                 <a class="navbar-brand" href="#">Navbar</a>
	//                 <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarNavDropdown" aria-controls="navbarNavDropdown" aria-expanded="false" aria-label="Toggle navigation">
	//                     <span class="navbar-toggler-icon"></span>
	//                 </button>
	//                 <div class="collapse navbar-collapse" id="navbarNavDropdown">
	//                     <ul class="navbar-nav">
	//                         <li class="nav-item">
	//                             <a class="nav-link active" aria-current="page" href="#">Home</a>
	//                         </li>
	//                         <li class="nav-item">
	//                             <a class="nav-link" href="#">Features</a>
	//                         </li>
	//                         <li class="nav-item">
	//                             <a class="nav-link" href="#">Pricing</a>
	//                         </li>
	//                         <li class="nav-item dropdown">
	//                             <a class="nav-link dropdown-toggle" href="#" role="button" data-bs-toggle="dropdown" aria-expanded="false">
	//                                 Dropdown link
	//                             </a>
	//                             <ul class="dropdown-menu">
	//                                 <li><a class="dropdown-item" href="#">Action</a></li>
	//                                 <li><a class="dropdown-item" href="#">Another action</a></li>
	//                                 <li><a class="dropdown-item" href="#">Something else here</a></li>
	//                             </ul>
	//                         </li>
	//                     </ul>
	//                 </div>
	//             </div>
	//         </nav>
	//         <br>
	//         <h1>Section 1</h1><h2>Carousel Layout</h2><br>
	// <div id="carouselExampleCaptions" class="carousel slide">
	//     <div class="carousel-indicators">
	//         <button type="button" data-bs-target="#carouselExampleCaptions" data-bs-slide-to="0" class="active" aria-current="true" aria-label="Slide 1"></button>
	//         <button type="button" data-bs-target="#carouselExampleCaptions" data-bs-slide-to="1" aria-label="Slide 2"></button>
	//         <button type="button" data-bs-target="#carouselExampleCaptions" data-bs-slide-to="2" aria-label="Slide 3"></button>
	//     </div>
	//     <div class="carousel-inner">
	//         <div class="carousel-item active">
	//             <img src="data:image/png;base64,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" class="d-block w-100" alt="">
	//             <div class="carousel-caption d-none d-md-block">
	//                 <h5>The week greatly well buy a wild trip.</h5>
	//                 <p>Who additionally him now did it many being loss have. Himself that under alone pack about this motionless other child. Tomorrow defiant understimate ours theirs this virtually cooperative labour Uzbek. Choir in annually since them soon time firstly between here. Would nightly cash hourly neither hers hospitality what me occasionally.</p>
	//             </div>
	//         </div>
	//         <div class="carousel-item">
	//             <img src="data:image/png;base64,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" class="d-block w-100" alt="">
	//             <div class="carousel-caption d-none d-md-block">
	//                 <h5>Tender week watch class softly.</h5>
	//                 <p>Here might its knit think still his Burmese glasses where. E.g. range troupe off can gang now how Ecuadorian this. Her disregard whom until these handle crowded few many cut. Do tense there whom this your his from last his. Bravely e.g. line whichever down sit who whose we advertising.</p>
	//             </div>
	//         </div>
	//         <div class="carousel-item">
	//             <img src="data:image/png;base64,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" class="d-block w-100" alt="">
	//             <div class="carousel-caption d-none d-md-block">
	//                 <h5>Bored problem extremely frankly play poorly.</h5>
	//                 <p>Myself of ourselves we then what sunglasses of fiction anybody. Ever American this all accordingly instance may provided couple Cormoran. Occasionally over agree stack great anyone recently after how then. From airport yesterday including fact words while than virtually may. Whomever face eye how in where those those has these.</p>
	//             </div>
	//         </div>
	//     </div>
	//     <button class="carousel-control-prev" type="button" data-bs-target="#carouselExampleCaptions" data-bs-slide="prev">
	//         <span class="carousel-control-prev-icon" aria-hidden="true"></span>
	//         <span class="visually-hidden">Previous</span>
	//     </button>
	//     <button class="carousel-control-next" type="button" data-bs-target="#carouselExampleCaptions" data-bs-slide="next">
	//         <span class="carousel-control-next-icon" aria-hidden="true"></span>
	//         <span class="visually-hidden">Next</span>
	//     </button>
	// </div>
	// <hr><br><br>
	//         <br>
	//         <footer>
	//             <small>©
	//                 <script>document.write(new Date().getFullYear())</script> PYA Analytics. All Rights Reserved.
	//             </small>
	//         </footer>
	//         <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.3.2/dist/js/bootstrap.bundle.min.js" crossorigin="anonymous"></script>
	//         <script src="app.js"></script>
	//     </body>
	// </html>
}

func TestTemplateHtml(t *testing.T) {
	f := New(11)
	globalFaker.Rand.Seed(11)
	value, err := f.TemplateHtml(2)
	if err != nil {
		t.Error(err)
	}

	if string(value) == "" {
		t.Error("Expected a document, got nothing")
	}
}

// TemplateHtmlContent examples and tests

func ExampleTemplateHtmlContent() {
	// Make sure we get the same results every time
	Seed(11)
	globalFaker.Rand.Seed(11)

	value, err := TemplateHtmlContent()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(value))

	// Output:
	// <h2>Person Profile</h2><br>
	// <div class="container">
	//     <div class="row">
	//         <div class="col-lg-12 mb-4 mb-sm-5">
	//             <div class="card card-style1 border-0">
	//                 <div class="card-body p-1-9 p-sm-2-3 p-md-6 p-lg-7">
	//                     <div class="row align-items-center">
	//                         <div class="col-lg-6 mb-4 mb-lg-0">
	//                             <img src="data:image/png;base64,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" width="50%" height="50%" alt="...">
	//                         </div>
	//                         <div class="col-lg-6 px-xl-10">
	//                             <div class="bg-secondary d-lg-inline-block py-1-9 px-1-9 px-sm-6 mb-1-9 rounded">
	//                                 <h3 class="h2 text-white mb-2">Vicente Mayert</h3>
	//                                 <span class="text-primary">Analyst</span>
	//                             </div>
	//                             <ul class="list-unstyled mb-1-9">
	//                                 <li class="mb-2 mb-xl-3 display-28"><span class="display-26 text-secondary me-2 font-weight-600">Position:</span> Analyst</li>
	//                                 <li class="mb-2 mb-xl-3 display-28"><span class="display-26 text-secondary me-2 font-weight-600">Level:</span> Analyst</li>
	//                                 <li class="mb-2 mb-xl-3 display-28"><span class="display-26 text-secondary me-2 font-weight-600">Email:</span> carriebaumbach@leffler.net</li>
	//                                 <li class="mb-2 mb-xl-3 display-28"><span class="display-26 text-secondary me-2 font-weight-600">Website:</span> www.microbilt-corporation.com</li>
	//                                 <li class="display-28"><span class="display-26 text-secondary me-2 font-weight-600">Phone:</span>3637235758</li>
	//                             </ul>
	//                             <ul class="social-icon-style1 list-unstyled mb-0 ps-0">
	//                                 <li><a href="#!"><i class="ti-twitter-alt"></i></a></li>
	//                                 <li><a href="#!"><i class="ti-facebook"></i></a></li>
	//                                 <li><a href="#!"><i class="ti-pinterest"></i></a></li>
	//                                 <li><a href="#!"><i class="ti-instagram"></i></a></li>
	//                             </ul>
	//                         </div>
	//                     </div>
	//                 </div>
	//             </div>
	//         </div>
	//         <div class="col-lg-12 mb-4 mb-sm-5">
	//             <div>
	//                 <span class="section-title text-primary mb-3 mb-sm-4">About Me</span>
	//                 <p>My us pack pack about this motionless other child tomorrow defiant understimate ours theirs this virtually cooperative labour Uzbek choir..</p>
	//                 <p class="mb-0">In annually since them soon time firstly between here would nightly cash hourly neither hers hospitality what me occasionally around..</p>
	//             </div>
	//         </div>
	//     </div>
	// </div>
	// <hr><br><br>

}

func ExampleFaker_TemplateHtmlContent() {

	// Make sure we get the same results every time
	f := New(11)
	globalFaker.Rand.Seed(11)
	value, err := f.TemplateHtmlContent()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(value))

	// Output:
	// <h2>Simple Content and Image</h2>
	// <p>Fly had result everything niche backwards caused quarterly without week it hungry thing someone him regularly today whomever this revolt. Hence from his timing as quantity us these yours live these frantic not may another how this ours his them. Those whose them batch its Iraqi most that few abroad cheese this whereas next how there gorgeous genetics time choir. Fiction therefore yourselves am those infrequently heap software quarterly rather punctuation yellow where several his orchard to frequently hence victorious. Boxers each does auspicious yourselves first soup tomorrow this that must conclude anyway some yearly who cough laugh himself both.<br>Yet rarely me dolphin intensely block would leap plane us first then down them eager would hundred super throughout animal. Yet themselves been group flock shake part purchase up usually it her none it hers boat what their there Turkmen. Moreover one Lebanese to brace these shower in it everybody should whatever those uninterested several nobody yours thankful innocent power. To any from its few luxury none boy religion themselves it there might must near theirs mine thing tonight outside. Rapidly spotted solemnly finally been did confusion these were nobody early silently company quarterly American there time many French a.<br>Each anybody rather paint ours did tonight remove first including chair your at in by cackle whose they yearly wisdom. Nightly nightly when finally without oxygen scold what silence go time behind example me soon grade purely that heavy annually. Our whoever your eventually yearly gleaming theirs child annually ours problem slavery someone brother instance could movement otherwise way now. Disturbed to sandwich someone its honour what whichever contrary from belief this upon at most homeless elsewhere has yearly under. Since where Californian all in today generally myself after stupid highly heavy here lately his who generally from substantial himself.</p>
	// <img src="data:image/png;base64,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" width="200" height="200" alt="Uzbek">
	// <hr><br><br>
}

func TestTemplateHtmlContent(t *testing.T) {
	f := New(11)
	globalFaker.Rand.Seed(11)
	value, err := f.TemplateHtmlContent()
	if err != nil {
		t.Error(err)
	}

	if string(value) == "" {
		t.Error("Expected a document, got nothing")
	}
}

// Template Test all Templates make sure no syntax errors
func TestTemplateAll(t *testing.T) {
	f := New(11)
	globalFaker.Rand.Seed(11)
	// combine all templates

	build_str := ""
	for k := range data.Template {
		for _, template := range data.Template[k] {
			value := strings.ReplaceAll(template, "'", "`")
			value = strings.ReplaceAll(value, "\\n", "\n")
			build_str = build_str + value
		}
	}

	value, err := f.Template(build_str, 1)
	if err != nil {
		t.Error(err)
	}

	if string(value) == "" {
		t.Error("Expected a document, got nothing")
	}
}

// Template Benchmarks

func BenchmarkTemplate100(b *testing.B) {
	f := New(11)
	globalFaker.Rand.Seed(11)

	for i := 0; i < 100; i++ {
		_, err := f.Template("{{range $y := IntRange 1 .Lines}}{{FirstName}} {{LastName}}\n{{end}}", 5)
		if err != nil {
			b.Fatal(err.Error())
		}
	}
}

func BenchmarkTemplateLookup1000(b *testing.B) {
	f := New(11)
	globalFaker.Rand.Seed(11)

	for i := 0; i < 1000; i++ {
		_, err := f.Template("{{range $y := IntRange 1 .Lines}}{{FirstName}} {{LastName}}\n{{end}}", 5)
		if err != nil {
			b.Fatal(err.Error())
		}

	}
}

func BenchmarkTemplateLookup10000(b *testing.B) {
	f := New(11)
	globalFaker.Rand.Seed(11)

	for i := 0; i < 10000; i++ {
		_, err := f.Template("{{range $y := IntRange 1 .Lines}}{{FirstName}} {{LastName}}\n{{end}}", 5)
		if err != nil {
			b.Fatal(err.Error())
		}

	}
}

func BenchmarkTemplateLookup100000(b *testing.B) {
	f := New(11)
	globalFaker.Rand.Seed(11)

	for i := 0; i < 100000; i++ {
		_, err := f.Template("{{range $y := IntRange 1 .Lines}}{{FirstName}} {{LastName}}\n{{end}}", 5)
		if err != nil {
			b.Fatal(err.Error())
		}

	}
}
