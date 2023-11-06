package gofakeit

import (
	"fmt"
	"testing"
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

// TemplateDocument examples and tests

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
	// # Black thing stack the hall.
	//
	// Late quarterly without week it hungry thing someone him regularly today whomever this revolt hence from his timing as quantity. Us these yours live these frantic not may another how this ours his them those whose them batch its Iraqi. Most that few abroad cheese this whereas next how there gorgeous genetics time choir fiction therefore yourselves am those infrequently. Heap software quarterly rather punctuation yellow where several his orchard to frequently hence victorious boxers each does auspicious yourselves first. Soup tomorrow this that must conclude anyway some yearly who cough laugh himself both yet rarely me dolphin intensely block.
	//
	// ---
	//
	// ## introduction
	//
	// Would leap plane us first then down them eager would.
	//
	// Hundred super throughout animal yet themselves been group flock shake.
	//
	// ---
	//
	// ## Warning
	//
	// Part purchase up usually it her none it hers boat.
	//
	// ---
	//
	// ## Requirements
	//
	// Golang you will need to install go 1.20
	//
	// [https://www.centralconvergence.net/one-to-one/mission-critical](https://www.centralconvergence.net/one-to-one/mission-critical)
	//
	// ## How to install
	//
	// Turkmen moreover one Lebanese to.
	//
	// <details>
	// <summary>1. A body far fiercely paint enormously bravely. </summary>
	//
	// ```
	// go install https://www.centralconvergence.net/one-to-one/mission-critical
	// ```
	//
	// </details>
	//
	// ---
	//
	// ## TODO
	//
	// Case uninterested several nobody yours thankful innocent power to any..
	//
	// 1. purely tightly
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
	// From: Marcel Pagac <marquesjakubowski@mraz.net>
	// To: Anibal Kozey <andrearmstrong@stanton.io>
	// Subject: hit me
	// Date: 1914-06-18 13:40:57.858870876 +0000 UTC
	// Content-Type: text/plain; charset=UTF-8;
	//
	// Hi Anibal Kozey
	//
	// Week it hungry thing someone him regularly today whomever this revolt hence from his timing as quantity us these yours. Live these frantic not may another how this ours his them those whose them batch its Iraqi most that few. Abroad cheese this whereas next how there gorgeous genetics time choir fiction therefore yourselves am those infrequently heap software quarterly. Rather punctuation yellow where several his orchard to frequently hence victorious boxers each does auspicious yourselves first soup tomorrow this. That must conclude anyway some yearly who cough laugh himself both yet rarely me dolphin intensely block would leap plane.
	//
	// Us first then down them eager would hundred super throughout animal yet themselves been group flock shake part purchase up. Usually it her none it hers boat what their there Turkmen moreover one Lebanese to brace these shower in it. Everybody should whatever those uninterested several nobody yours thankful innocent power to any from its few luxury none boy religion. Themselves it there might must near theirs mine thing tonight outside rapidly spotted solemnly finally been did confusion these were. Nobody early silently company quarterly American there time many French a each anybody rather paint ours did tonight remove first.
	//
	// Including chair your at in by cackle whose they yearly wisdom nightly nightly when finally without oxygen scold what silence. Go time behind example me soon grade purely that heavy annually our whoever your eventually yearly gleaming theirs child annually. Ours problem slavery someone brother instance could movement otherwise way now disturbed to sandwich someone its honour what whichever contrary. From belief this upon at most homeless elsewhere has yearly under since where Californian all in today generally myself after. Stupid highly heavy here lately his who generally from substantial himself poised cost formerly but spoon words Egyptian i.e. me.
	// Regards Marcel Pagac
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
	// From: Marcel Pagac <marquesjakubowski@mraz.net>
	// To: Anibal Kozey <andrearmstrong@stanton.io>
	// Subject: hit me
	// Date: 1914-06-18 13:40:57.858870876 +0000 UTC
	// Content-Type: text/plain; charset=UTF-8;
	//
	// Hi Anibal Kozey
	//
	// Week it hungry thing someone him regularly today whomever this revolt hence from his timing as quantity us these yours. Live these frantic not may another how this ours his them those whose them batch its Iraqi most that few. Abroad cheese this whereas next how there gorgeous genetics time choir fiction therefore yourselves am those infrequently heap software quarterly. Rather punctuation yellow where several his orchard to frequently hence victorious boxers each does auspicious yourselves first soup tomorrow this. That must conclude anyway some yearly who cough laugh himself both yet rarely me dolphin intensely block would leap plane.
	//
	// Us first then down them eager would hundred super throughout animal yet themselves been group flock shake part purchase up. Usually it her none it hers boat what their there Turkmen moreover one Lebanese to brace these shower in it. Everybody should whatever those uninterested several nobody yours thankful innocent power to any from its few luxury none boy religion. Themselves it there might must near theirs mine thing tonight outside rapidly spotted solemnly finally been did confusion these were. Nobody early silently company quarterly American there time many French a each anybody rather paint ours did tonight remove first.
	//
	// Including chair your at in by cackle whose they yearly wisdom nightly nightly when finally without oxygen scold what silence. Go time behind example me soon grade purely that heavy annually our whoever your eventually yearly gleaming theirs child annually. Ours problem slavery someone brother instance could movement otherwise way now disturbed to sandwich someone its honour what whichever contrary. From belief this upon at most homeless elsewhere has yearly under since where Californian all in today generally myself after. Stupid highly heavy here lately his who generally from substantial himself poised cost formerly but spoon words Egyptian i.e. me.
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
	// From: Markus Moen <luralockman@jakubowski.com>
	// To: Alayna Wuckert <paolorutherford@armstrong.org>
	// Subject: the rest is history
	// Date: 1926-07-07 22:25:40.559683317 +0000 UTC
	// Content-Type: text/plain; charset=UTF-8;
	//
	// Hi Alayna Wuckert
	//
	// Weight government hourly constantly yesterday someone him regularly today whomever this revolt hence from his timing as quantity us these. Yours live these frantic not may another how this ours his them those whose them batch its Iraqi most that. Few abroad cheese this whereas next how there gorgeous genetics time choir fiction therefore yourselves am those infrequently heap software. Quarterly rather punctuation yellow where several his orchard to frequently hence victorious boxers each does auspicious yourselves first soup tomorrow. This that must conclude anyway some yearly who cough laugh himself both yet rarely me dolphin intensely block would leap.
	//
	// Plane us first then down them eager would hundred super throughout animal yet themselves been group flock shake part purchase. Up usually it her none it hers boat what their there Turkmen moreover one Lebanese to brace these shower in. It everybody should whatever those uninterested several nobody yours thankful innocent power to any from its few luxury none boy. Religion themselves it there might must near theirs mine thing tonight outside rapidly spotted solemnly finally been did confusion these. Were nobody early silently company quarterly American there time many French a each anybody rather paint ours did tonight remove.
	//
	// First including chair your at in by cackle whose they yearly wisdom nightly nightly when finally without oxygen scold what. Silence go time behind example me soon grade purely that heavy annually our whoever your eventually yearly gleaming theirs child. Annually ours problem slavery someone brother instance could movement otherwise way now disturbed to sandwich someone its honour what whichever. Contrary from belief this upon at most homeless elsewhere has yearly under since where Californian all in today generally myself. After stupid highly heavy here lately his who generally from substantial himself poised cost formerly but spoon words Egyptian i.e..
	// Regards Markus Moen

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
	// # The tribe indeed swiftly laugh.
	//
	// Backwards caused quarterly without week it hungry thing someone him regularly today whomever this revolt hence from his timing as. Quantity us these yours live these frantic not may another how this ours his them those whose them batch its. Iraqi most that few abroad cheese this whereas next how there gorgeous genetics time choir fiction therefore yourselves am those. Infrequently heap software quarterly rather punctuation yellow where several his orchard to frequently hence victorious boxers each does auspicious yourselves. First soup tomorrow this that must conclude anyway some yearly who cough laugh himself both yet rarely me dolphin intensely.
	//
	// ---
	//
	// ## introduction
	//
	// Block would leap plane us first then down them eager.
	//
	// Would hundred super throughout animal yet themselves been group flock.
	//
	// ---
	//
	// ## Warning
	//
	// Shake part purchase up usually it her none it hers.
	//
	// ---
	//
	// ## Requirements
	//
	// Golang you will need to install go 1.20
	//
	// [http://www.directbleeding-edge.name/revolutionary/relationships](http://www.directbleeding-edge.name/revolutionary/relationships)
	//
	// ## How to install
	//
	// She Turkmen moreover one Lebanese.
	//
	// <details>
	// <summary>1. The comfortable captain how beautifully dance. </summary>
	//
	// ```
	// go install http://www.directbleeding-edge.name/revolutionary/relationships
	// ```
	//
	// </details>
	//
	// ---
	//
	// ## TODO
	//
	// Evil you these those uninterested several nobody yours thankful innocent..
	//
	// 1. honestly
	// 1. so equally
	// 1. doubtfully

}

func ExampleFaker_TemplateMarkdown() {
	// Make sure we get the same results every time
	f := New(11)
	globalFaker.Rand.Seed(11)
	value, err := f.TemplateMarkdown(3)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(value))

	// Output:
	// # The purple tribe indeed swiftly laugh.
	//
	// Backwards caused quarterly without week it hungry thing someone him regularly today whomever this revolt hence from his timing as. Quantity us these yours live these frantic not may another how this ours his them those whose them batch its. Iraqi most that few abroad cheese this whereas next how there gorgeous genetics time choir fiction therefore yourselves am those. Infrequently heap software quarterly rather punctuation yellow where several his orchard to frequently hence victorious boxers each does auspicious yourselves. First soup tomorrow this that must conclude anyway some yearly who cough laugh himself both yet rarely me dolphin intensely.
	//
	// ---
	//
	// ## introduction
	//
	// Block would leap plane us first then down them eager.
	//
	// Would hundred super throughout animal yet themselves been group flock.
	//
	// ---
	//
	// ## Warning
	//
	// Shake part purchase up usually it her none it hers.
	//
	// ---
	//
	// ## Requirements
	//
	// Golang you will need to install go 1.20
	//
	// [http://www.directbleeding-edge.name/revolutionary/relationships](http://www.directbleeding-edge.name/revolutionary/relationships)
	//
	// ## How to install
	//
	// She Turkmen moreover one Lebanese.
	//
	// <details>
	// <summary>1. The comfortable captain how beautifully dance. </summary>
	//
	// ```
	// go install http://www.directbleeding-edge.name/revolutionary/relationships
	// ```
	//
	// </details>
	//
	// ---
	//
	// ## TODO
	//
	// Evil you these those uninterested several nobody yours thankful innocent..
	//
	// 1. honestly
	// 1. so equally
	// 1. doubtfully

}

func TestTemplateMD(t *testing.T) {
	f := New(11)
	globalFaker.Rand.Seed(11)
	value, err := f.TemplateMarkdown(4)
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
	//   <meta charset="UTF-8" />
	//   <title>The tribe indeed swiftly laugh.</title>
	//   <meta name="viewport" content="width=device-width,initial-scale=1" />
	//   <meta name="description" content="" />
	//   <link href="https://cdn.jsdelivr.net/npm/modern-normalize@v2.0.0/modern-normalize.min.css" rel="stylesheet">
	//   <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.2/dist/css/bootstrap.min.css" rel="stylesheet" crossorigin="anonymous">
	//   <link rel="stylesheet" type="text/css" href="style.css" />
	//   <link rel="icon" type="image/svg+xml" href="<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 16 16" width="16" height="16"><rect x="0" y="0" width="100%" height="100%" fill="#403831" /><polygon points="7,13 13,7 6,1" fill="#30261c" /><polygon points="1,15 4,9 4,13" fill="#1f5f61" /><polygon points="3,7 4,10 2,10" fill="#30261c" /><polygon points="7,15 12,12 0,11" fill="#36544f" /><polygon points="7,5 15,13 2,6" fill="#0b8185" /><polygon points="0,14 4,5 11,14" fill="#1f5f61" /><polygon points="3,8 14,13 16,5" fill="#0b8185" /><polygon points="12,11 9,16 13,10" fill="#1f5f61" /><polygon points="5,11 3,2 12,4" fill="#1f5f61" /><polygon points="10,2 8,15 12,10" fill="#30261c" /><polygon points="1,11 16,15 8,12" fill="#1f5f61" /></svg>">
	//
	//   <meta name="theme-color" content="">
	//   <meta property="og:title" content="" />
	//   <meta property="og:description" content="" />
	//   <meta property="og:image" content="" />
	//   <meta name="twitter:card" content="" />
	//   <meta name="twitter:site" content="" />
	//   <meta name="twitter:title" content="" />
	//   <meta name="twitter:description" content="" />
	//   <meta name="twitter:image" content="" />
	//
	// </head>
	// <body>
	// <body>
	//   <nav class="navbar navbar-expand-lg bg-body-tertiary">
	//     <div class="container-fluid">
	//       <a class="navbar-brand" href="#">Navbar</a>
	//       <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarNavDropdown" aria-controls="navbarNavDropdown" aria-expanded="false" aria-label="Toggle navigation">
	//         <span class="navbar-toggler-icon"></span>
	//       </button>
	//       <div class="collapse navbar-collapse" id="navbarNavDropdown">
	//         <ul class="navbar-nav">
	//           <li class="nav-item">
	//             <a class="nav-link active" aria-current="page" href="#">Home</a>
	//           </li>
	//           <li class="nav-item">
	//             <a class="nav-link" href="#">Features</a>
	//           </li>
	//           <li class="nav-item">
	//             <a class="nav-link" href="#">Pricing</a>
	//           </li>
	//           <li class="nav-item dropdown">
	//             <a class="nav-link dropdown-toggle" href="#" role="button" data-bs-toggle="dropdown" aria-expanded="false">
	//               Dropdown link
	//             </a>
	//             <ul class="dropdown-menu">
	//               <li><a class="dropdown-item" href="#">Action</a></li>
	//               <li><a class="dropdown-item" href="#">Another action</a></li>
	//               <li><a class="dropdown-item" href="#">Something else here</a></li>
	//             </ul>
	//           </li>
	//         </ul>
	//       </div>
	//     </div>
	//   </nav>
	// <br>
	// <h1>Selfish pod hug noisily.</h1>
	// <p>How there gorgeous genetics time choir fiction therefore yourselves am those infrequently heap software quarterly rather punctuation yellow where several. His orchard to frequently hence victorious boxers each does auspicious yourselves first soup tomorrow this that must conclude anyway some. Yearly who cough laugh himself both yet rarely me dolphin intensely block would leap plane us first then down them. Eager would hundred super throughout animal yet themselves been group flock shake part purchase up usually it her none it. Hers boat what their there Turkmen moreover one Lebanese to brace these shower in it everybody should whatever those uninterested.<br>Several nobody yours thankful innocent power to any from its few luxury none boy religion themselves it there might must. Near theirs mine thing tonight outside rapidly spotted solemnly finally been did confusion these were nobody early silently company quarterly. American there time many French a each anybody rather paint ours did tonight remove first including chair your at in. By cackle whose they yearly wisdom nightly nightly when finally without oxygen scold what silence go time behind example me. Soon grade purely that heavy annually our whoever your eventually yearly gleaming theirs child annually ours problem slavery someone brother.<br>Instance could movement otherwise way now disturbed to sandwich someone its honour what whichever contrary from belief this upon at. Most homeless elsewhere has yearly under since where Californian all in today generally myself after stupid highly heavy here lately. His who generally from substantial himself poised cost formerly but spoon words Egyptian i.e. me tonight place few why from. Somebody hungrily mine were soon hail then you themselves drab when behind case ours cost couple consequently in those daily. Had anywhere anything what in bouquet which annually as Cypriot this our sand tightly we first their staff invention however.</p>
	//
	// <div id="carouselExampleCaptions" class="carousel slide">
	// <div class="carousel-indicators">
	//   <button type="button" data-bs-target="#carouselExampleCaptions" data-bs-slide-to="0" class="active" aria-current="true" aria-label="Slide 1"></button>
	//   <button type="button" data-bs-target="#carouselExampleCaptions" data-bs-slide-to="1" aria-label="Slide 2"></button>
	//   <button type="button" data-bs-target="#carouselExampleCaptions" data-bs-slide-to="2" aria-label="Slide 3"></button>
	// </div>
	// <div class="carousel-inner">
	//   <div class="carousel-item active">
	// 	<img src="data:image/png;base64,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" class="d-block w-100" alt="">
	// 	<div class="carousel-caption d-none d-md-block">
	// 	  <h5>A host absolutely honestly cook a wicked time.</h5>
	// 	  <p>Hand riches heap peep yours envy last cackle however guitar. These lastly besides bouquet that of will myself these those. Iraqi insufficient usually brother other case have those these my. Infrequently all truth occasionally army lastly these i.e. costume Chinese. Then lastly anxiously this Lincolnian man next board of monthly.</p>
	// 	</div>
	//   </div>
	//   <div class="carousel-item">
	// 	<img src="data:image/png;base64,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" class="d-block w-100" alt="">
	// 	<div class="carousel-caption d-none d-md-block">
	// 	<h5>A group sing beautiful valley straight on the class.</h5>
	// 	<p>Fight behind outfit house thoroughly we did now until collection. Band i.e. buy problem eye itchy batch warmth above out. Several will instance yesterday money every but cloud way a. Ever provided fast constantly it inquiring absolutely here anyone transform. Moreover that of stagger summation these to uptight double transportation.</p>
	// 	</div>
	//   </div>
	//   <div class="carousel-item">
	// 	<img src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAB4AAAAeCAIAAAC0Ujn1AAAKuklEQVR4nACqClX1BHEZQOv14Up7DpUbn88jWna7a+0ZpFRBPgkWncDoizVHfmZVRPD4yvcEWksHQ0oYlF5eRYBocsaw3zPZc80O+sUfkAMCXWxEpDSRPO+o6QTy4g9I0Grq3iM3XQEdfo7l41sRxcu5PBKOpP6+ALqknnvXj4MOxWO6UixIkN4rwzvGZxFGNwcnmkNK1OULJw6FulKxYsYWyTBAU+F4ZHEFxu3dolO6j0corZ7THBlDBHl4ScECGEACJBubnsQC91xqG8tj1zTf4UtVnAKKFWQ+O0dq1wA3y7ZF1H7XDJTys/tuS23Bi2W+e8cXYmca0XCF4S79ndefpSgXNZhp1P8YqkeEAZccwirKG+Fzt2WbOwbPAHStRyNGy4rlzanSeyvg/W8AdHuxL2dPv03S7ufjyakFuDU/0w1nAskWGv/xkomoXJwNvxlio5H7Ei8MJPmV55xEsDrRdbipauacU1H9h1lxDjGu4ga8cI1vYAIUHs31e62MDa8CpiMf8hmuXlAHeeBYtZQbD2gCT0P8jmyRk5ka960scPRHOd3cDslARdHknhy+jaULc4JFv/F3xuv/Rh32MioTGO+3YbHD4lsQj6LH+hr8yLUEnnHAgTWsAxlPC9X2uIOsV8fUy1/nIAFW4Oyrx/rkFv8pAmLOE+b55mdsIILD+0BTV1kn2Hs5VQw3g6brw0AlO1rs/lOY2CfDumgt7oNfADZ+Q7FcLfb9Q3ToA7bntr/7xt0pxjfytGNVwt7+5aGV1c6MrHZdqO+8t0Mq6pv8y+McnycUNkuMl5vmtGgIs6dIg4aTCe5nlUY/hDfgVevhVQ30TN9lYeRxsS0GOlFP4JIlv1qXLQGAS1DWRzVu1mB8tGOhM8DewjuIuNIASvkNxt9XPtb3RPkIo/vn09qAP/wHhPOJK+DCVqzR34kJf/EHxaCy4ZUDK3BZdTNm8AOgdKOHj5VnNa/RbTRarCJHuPsAZpKJbcMhGLjuch7i99IEVloKFmcbZLXcBIM2rPtK4+WE8TPrlCijL89jJv4Vo0rn6UD3RxHs+7QRKvV4+dyHTlgQ4ygOI29uDQ768WkJl84I8sJzF8gAJ8sTAgbCxQmJo5cciIxW+bRA/dQ+OwS6z/troje4DLL0rS4ZuGr9JzLzioK7QT8eXq9MC/2vNw+VWgYzr69n/1O90p8ZNTCn6utckY6vN4IxCEOc1iVx9RIkDHCyKACR9jY5y1SmSY9dsgEjHFlOHNtcwVulHF0K7elcTQ8cX82zq7sTRV5VdryomT8VqLfLyTk+C+d1VL2CyEmuo5HO7udrSdrQ/kKi1vmNowDuciyL8ahBYB0C7dQAFRyUgRv5PUZubv4+AIp5HcBAhuxM1pcREDVY1xUTIGP+7fDGdZuepAkl6KGC7r4KyjC1CDBdMyXOnqpkBMtbFSN8GUww2W2UVDFyvGXkVWq22eeAAFLmggMaBA86xGCMlwD7A8NnzgAZZF4mEftq5RVGjz/ummVhK++3x/NrNv2MME6a+h3o6fzNijt9s2jlQx0x7ww85J0y+Bjglw3keJSioODrf0gyiiDoylxFaBRRpgm3dgOYC3Cqpy2KRA9rGacMTdiBdEx+/5eoWERWT9gfCiUFHtJGmAo0ElmHIVOe8ATFrOZvSYDRNONUcRqCG8WmD+36R/rFTO3Ub3o3yNp4gMsNSMsaVc7qNfrvUmsEzhXmn0dlanfG3HlXaRG9WRMV6SYal/WKCMVl+oFNMfnHasRxRkKh2wE+BDHLELU/FPT87pyouWY38FRtTBDdqHs6Qvvu3oNvEmj5J7jIl0X3hXL6QCnJHaKgAuOHAWMMXqUJuCTzmBiY3KnsEw/zelnO4piOsLRqKVSQ01fNVDGVIA5tyiatS9TKvSOXPVsov7MLE4U3IjzFALiiZDp1H23voXTvJTOVoXiTtxbjVa+AzwjEvwILAuUBVaAh6eNExxGGcboniDFKIkH1yvgnPazI9N9v1wFWGwQhzCmReUDfoOHUTEpxVughNzl6/vF0ZNgJjnBcLgBmxB9DvCPeo98aot1X/pCcik6MNrjzdnoATnr64TjgiNxCSWC9rdx4IOO1pSBE6YhNbk9PIMsaqDHZndkjHICciG1YvTTN2oCS73iEZSG1cgN6+3DgJW7wgwPOrzMoMXSpHRwGu3ouIj/Jtu/TSony0EYAA+MfNuBa+b5l+QUN3/rkag65Tfs9WrmpyyJMmenr+3wRufgl3Wx2dk0HMSro130Q/nd8ngAvFmmt0cDx7c4qfWvP8cnXRysQmkS+s/htQHITeUuikggPOwudwQJxXDjE2XAokzsIUr6TZO34KIIy2EqNmu3sIrt7ZC2oqii/TIEIEA3DKdPImOaXhJFlGj7e15lIAprFR8bdEKrp2uUO+KIRLDxXU2LJasKYlx7CZSIT7/g6QOMAEic2Y/odivIIlrRmJqOP0uOg0ppfmbv80/bGpEO88R5RBbO/z5P2AeOk3iTTXZCu3jmVNXWzS0vqErEK+zs8GKgon9omT146vv7dlRQmveVlGjPNRdinYEYbAVOwX5D8c3DgsJppAvqjLC1nVT/588awCh3jakPR6NLQHoMP2T2tM+atgtY7Ps6V0Ed18huT6ErXU8ExnWbDT9Sh2QW/VlhbpQQ/FKJVg/tEEnlL1roUPfA1AgDiZE8y0xG2tk00TFwACeVi5QuDZp0z9L5IpY5GM5z8ssmFC+5eye9I7U5bBT3aclgNVCHY9bRsenyA5nIjxGH+5cw9dMf2t34rHgdBEDZHOOm2PByelKE6RPcCvVfde2l+Lg+5m0kWC+++MjNt8QM97heISVOjC/9jX5z7VrKmHyZRKViuPKepIbLhVjwH4t4ml4zgEEjqezIoyOBZ4sh7Puhtw1OxXwZGDGGdYjxX7dvH+zDGANHXj/7VXeUqrNi0qJhMJPYY7F3WC/SYyEgi+jpQB+U8gA/xgkqKpdf2W7qZ04ghilqzjLtOATDE0Z8dXjkfHGfcagL/82IzelAigdm/Jc8eyvpwTJmqRp+pWQMii3WrUOF6HQ5V5WXODTtapxvIS/MJD9yvWh9fBnHiZvrIt2WLPunVw0asoIMcQhmSDdZVWOI9ZYBdV/vdAJOhucv+E7XSwgxDQOYSmLrYBiuocJi1RXHGBcwAUb2GpyYjdTFC7BMufM/wULmiGu3RTn0R4/urfNPNgLk8iL8vtQEHVoKsCfnbhrTqV0hiKwQmjpjmH5TmRxQ8C5Onw1j0nkryZJwBTPPq/nbSJoBvLmcWBtrqAv2CEJ5XYgTpuAVAzRH1nEayc+ZMmabrUOla0TzQUeTGKh+VmNYfQX3ZDNUOITv9SSP13Eq96adWHkaPQVNW8qhmNexvAWjVxV5XOh03OHhpjtns9rBai4xC6gDeLPNnK2Tr6qdXL7IHKfuaVGfFHAANNMJlzsbF/Nu2LRlkwuSuoSSwHTOePfFm2y/XWZ4YhYUm8KD08euymfWtJdgcwKCSZfDiLBrJN5JF53vt9q8sfZj8aJAABXNMqMw0wvXzcXbjRYVPpcLld6/9HX+N0hjP91wcSRURtI3XPEOkw3Qsjha950w7uL+OiufvrWwAbWquGJ0ZFpFfelQA9M3p7mJL3gTYllsJTiVetYlLSp8kAQAA///NNT4K7NXcCwAAAABJRU5ErkJggg==" class="d-block w-100" alt="">
	// 	<div class="carousel-caption d-none d-md-block">
	// 	<h5>An elegant group stack terribly quickly off an impromptu congregation.</h5>
	// 	<p>There depend you backwards ours upshot had wait I watch. Head as those cut ingeniously there think on do had. Your himself its yet by had this flag one later. I behind upon later theirs that that everything mortally a. Would sensibly troop i.e. i.e. throw together covey place where.</p>
	// 	</div>
	//   </div>
	// </div>
	// <button class="carousel-control-prev" type="button" data-bs-target="#carouselExampleCaptions" data-bs-slide="prev">
	//   <span class="carousel-control-prev-icon" aria-hidden="true"></span>
	//   <span class="visually-hidden">Previous</span>
	// </button>
	// <button class="carousel-control-next" type="button" data-bs-target="#carouselExampleCaptions" data-bs-slide="next">
	//   <span class="carousel-control-next-icon" aria-hidden="true"></span>
	//   <span class="visually-hidden">Next</span>
	// </button>
	// </div>
	//
	// <br>
	// 	<footer>
	//     <small>© <script>document.write(new Date().getFullYear())</script> Your company name. All Rights Reserved.</small>
	//   </footer>
	//   <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.3.2/dist/js/bootstrap.bundle.min.js" crossorigin="anonymous"></script>
	//   <script src="app.js"></script>
	// </body>
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
	//   <meta charset="UTF-8" />
	//   <title>The purple tribe indeed swiftly laugh.</title>
	//   <meta name="viewport" content="width=device-width,initial-scale=1" />
	//   <meta name="description" content="" />
	//   <link href="https://cdn.jsdelivr.net/npm/modern-normalize@v2.0.0/modern-normalize.min.css" rel="stylesheet">
	//   <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.2/dist/css/bootstrap.min.css" rel="stylesheet" crossorigin="anonymous">
	//   <link rel="stylesheet" type="text/css" href="style.css" />
	//   <link rel="icon" type="image/svg+xml" href="<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 16 16" width="16" height="16"><rect x="0" y="0" width="100%" height="100%" fill="#403831" /><polygon points="7,13 13,7 6,1" fill="#30261c" /><polygon points="1,15 4,9 4,13" fill="#1f5f61" /><polygon points="3,7 4,10 2,10" fill="#30261c" /><polygon points="7,15 12,12 0,11" fill="#36544f" /><polygon points="7,5 15,13 2,6" fill="#0b8185" /><polygon points="0,14 4,5 11,14" fill="#1f5f61" /><polygon points="3,8 14,13 16,5" fill="#0b8185" /><polygon points="12,11 9,16 13,10" fill="#1f5f61" /><polygon points="5,11 3,2 12,4" fill="#1f5f61" /><polygon points="10,2 8,15 12,10" fill="#30261c" /><polygon points="1,11 16,15 8,12" fill="#1f5f61" /></svg>">
	//
	//   <meta name="theme-color" content="">
	//   <meta property="og:title" content="" />
	//   <meta property="og:description" content="" />
	//   <meta property="og:image" content="" />
	//   <meta name="twitter:card" content="" />
	//   <meta name="twitter:site" content="" />
	//   <meta name="twitter:title" content="" />
	//   <meta name="twitter:description" content="" />
	//   <meta name="twitter:image" content="" />
	//
	// </head>
	// <body>
	// <body>
	//   <nav class="navbar navbar-expand-lg bg-body-tertiary">
	//     <div class="container-fluid">
	//       <a class="navbar-brand" href="#">Navbar</a>
	//       <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarNavDropdown" aria-controls="navbarNavDropdown" aria-expanded="false" aria-label="Toggle navigation">
	//         <span class="navbar-toggler-icon"></span>
	//       </button>
	//       <div class="collapse navbar-collapse" id="navbarNavDropdown">
	//         <ul class="navbar-nav">
	//           <li class="nav-item">
	//             <a class="nav-link active" aria-current="page" href="#">Home</a>
	//           </li>
	//           <li class="nav-item">
	//             <a class="nav-link" href="#">Features</a>
	//           </li>
	//           <li class="nav-item">
	//             <a class="nav-link" href="#">Pricing</a>
	//           </li>
	//           <li class="nav-item dropdown">
	//             <a class="nav-link dropdown-toggle" href="#" role="button" data-bs-toggle="dropdown" aria-expanded="false">
	//               Dropdown link
	//             </a>
	//             <ul class="dropdown-menu">
	//               <li><a class="dropdown-item" href="#">Action</a></li>
	//               <li><a class="dropdown-item" href="#">Another action</a></li>
	//               <li><a class="dropdown-item" href="#">Something else here</a></li>
	//             </ul>
	//           </li>
	//         </ul>
	//       </div>
	//     </div>
	//   </nav>
	// <br>
	// <h1>Selfish pod hug noisily.</h1>
	// <p>How there gorgeous genetics time choir fiction therefore yourselves am those infrequently heap software quarterly rather punctuation yellow where several. His orchard to frequently hence victorious boxers each does auspicious yourselves first soup tomorrow this that must conclude anyway some. Yearly who cough laugh himself both yet rarely me dolphin intensely block would leap plane us first then down them. Eager would hundred super throughout animal yet themselves been group flock shake part purchase up usually it her none it. Hers boat what their there Turkmen moreover one Lebanese to brace these shower in it everybody should whatever those uninterested.<br>Several nobody yours thankful innocent power to any from its few luxury none boy religion themselves it there might must. Near theirs mine thing tonight outside rapidly spotted solemnly finally been did confusion these were nobody early silently company quarterly. American there time many French a each anybody rather paint ours did tonight remove first including chair your at in. By cackle whose they yearly wisdom nightly nightly when finally without oxygen scold what silence go time behind example me. Soon grade purely that heavy annually our whoever your eventually yearly gleaming theirs child annually ours problem slavery someone brother.<br>Instance could movement otherwise way now disturbed to sandwich someone its honour what whichever contrary from belief this upon at. Most homeless elsewhere has yearly under since where Californian all in today generally myself after stupid highly heavy here lately. His who generally from substantial himself poised cost formerly but spoon words Egyptian i.e. me tonight place few why from. Somebody hungrily mine were soon hail then you themselves drab when behind case ours cost couple consequently in those daily. Had anywhere anything what in bouquet which annually as Cypriot this our sand tightly we first their staff invention however.</p>
	//
	// <div id="carouselExampleCaptions" class="carousel slide">
	// <div class="carousel-indicators">
	//   <button type="button" data-bs-target="#carouselExampleCaptions" data-bs-slide-to="0" class="active" aria-current="true" aria-label="Slide 1"></button>
	//   <button type="button" data-bs-target="#carouselExampleCaptions" data-bs-slide-to="1" aria-label="Slide 2"></button>
	//   <button type="button" data-bs-target="#carouselExampleCaptions" data-bs-slide-to="2" aria-label="Slide 3"></button>
	// </div>
	// <div class="carousel-inner">
	//   <div class="carousel-item active">
	// 	<img src="data:image/png;base64,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" class="d-block w-100" alt="">
	// 	<div class="carousel-caption d-none d-md-block">
	// 	  <h5>A host absolutely honestly cook a wicked time.</h5>
	// 	  <p>Hand riches heap peep yours envy last cackle however guitar. These lastly besides bouquet that of will myself these those. Iraqi insufficient usually brother other case have those these my. Infrequently all truth occasionally army lastly these i.e. costume Chinese. Then lastly anxiously this Lincolnian man next board of monthly.</p>
	// 	</div>
	//   </div>
	//   <div class="carousel-item">
	// 	<img src="data:image/png;base64,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" class="d-block w-100" alt="">
	// 	<div class="carousel-caption d-none d-md-block">
	// 	<h5>A group sing beautiful valley straight on the class.</h5>
	// 	<p>Fight behind outfit house thoroughly we did now until collection. Band i.e. buy problem eye itchy batch warmth above out. Several will instance yesterday money every but cloud way a. Ever provided fast constantly it inquiring absolutely here anyone transform. Moreover that of stagger summation these to uptight double transportation.</p>
	// 	</div>
	//   </div>
	//   <div class="carousel-item">
	// 	<img src="data:image/png;base64,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" class="d-block w-100" alt="">
	// 	<div class="carousel-caption d-none d-md-block">
	// 	<h5>An elegant group stack terribly quickly off an impromptu congregation.</h5>
	// 	<p>There depend you backwards ours upshot had wait I watch. Head as those cut ingeniously there think on do had. Your himself its yet by had this flag one later. I behind upon later theirs that that everything mortally a. Would sensibly troop i.e. i.e. throw together covey place where.</p>
	// 	</div>
	//   </div>
	// </div>
	// <button class="carousel-control-prev" type="button" data-bs-target="#carouselExampleCaptions" data-bs-slide="prev">
	//   <span class="carousel-control-prev-icon" aria-hidden="true"></span>
	//   <span class="visually-hidden">Previous</span>
	// </button>
	// <button class="carousel-control-next" type="button" data-bs-target="#carouselExampleCaptions" data-bs-slide="next">
	//   <span class="carousel-control-next-icon" aria-hidden="true"></span>
	//   <span class="visually-hidden">Next</span>
	// </button>
	// </div>
	//
	// <br>
	// 	<footer>
	//     <small>© <script>document.write(new Date().getFullYear())</script> Your company name. All Rights Reserved.</small>
	//   </footer>
	//   <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.3.2/dist/js/bootstrap.bundle.min.js" crossorigin="anonymous"></script>
	//   <script src="app.js"></script>
	// </body>
	// </html>

}

func TestTemplateHTML(t *testing.T) {
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
	// <section id="another-page">
	// <h1>Another Page</h1>
	// <p>Generally half where everything niche backwards caused quarterly without week it hungry thing someone him regularly today whomever this revolt. Hence from his timing as quantity us these yours live these frantic not may another how this ours his them. Those whose them batch its Iraqi most that few abroad cheese this whereas next how there gorgeous genetics time choir. Fiction therefore yourselves am those infrequently heap software quarterly rather punctuation yellow where several his orchard to frequently hence victorious. Boxers each does auspicious yourselves first soup tomorrow this that must conclude anyway some yearly who cough laugh himself both.<br>Yet rarely me dolphin intensely block would leap plane us first then down them eager would hundred super throughout animal. Yet themselves been group flock shake part purchase up usually it her none it hers boat what their there Turkmen. Moreover one Lebanese to brace these shower in it everybody should whatever those uninterested several nobody yours thankful innocent power. To any from its few luxury none boy religion themselves it there might must near theirs mine thing tonight outside. Rapidly spotted solemnly finally been did confusion these were nobody early silently company quarterly American there time many French a.<br>Each anybody rather paint ours did tonight remove first including chair your at in by cackle whose they yearly wisdom. Nightly nightly when finally without oxygen scold what silence go time behind example me soon grade purely that heavy annually. Our whoever your eventually yearly gleaming theirs child annually ours problem slavery someone brother instance could movement otherwise way now. Disturbed to sandwich someone its honour what whichever contrary from belief this upon at most homeless elsewhere has yearly under. Since where Californian all in today generally myself after stupid highly heavy here lately his who generally from substantial himself.</p>
	// </section>
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
	// <h1>The tribe indeed swiftly laugh.</h1>
	// <p>Backwards caused quarterly without week it hungry thing someone him regularly today whomever this revolt hence from his timing as. Quantity us these yours live these frantic not may another how this ours his them those whose them batch its. Iraqi most that few abroad cheese this whereas next how there gorgeous genetics time choir fiction therefore yourselves am those. Infrequently heap software quarterly rather punctuation yellow where several his orchard to frequently hence victorious boxers each does auspicious yourselves. First soup tomorrow this that must conclude anyway some yearly who cough laugh himself both yet rarely me dolphin intensely.<br>Block would leap plane us first then down them eager would hundred super throughout animal yet themselves been group flock. Shake part purchase up usually it her none it hers boat what their there Turkmen moreover one Lebanese to brace. These shower in it everybody should whatever those uninterested several nobody yours thankful innocent power to any from its few. Luxury none boy religion themselves it there might must near theirs mine thing tonight outside rapidly spotted solemnly finally been. Did confusion these were nobody early silently company quarterly American there time many French a each anybody rather paint ours.<br>Did tonight remove first including chair your at in by cackle whose they yearly wisdom nightly nightly when finally without. Oxygen scold what silence go time behind example me soon grade purely that heavy annually our whoever your eventually yearly. Gleaming theirs child annually ours problem slavery someone brother instance could movement otherwise way now disturbed to sandwich someone its. Honour what whichever contrary from belief this upon at most homeless elsewhere has yearly under since where Californian all in. Today generally myself after stupid highly heavy here lately his who generally from substantial himself poised cost formerly but spoon.</p>
	// <img src="data:image/png;base64,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" alt="">
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
