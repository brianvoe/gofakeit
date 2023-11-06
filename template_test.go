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
	//   <link rel="icon" type="image/svg+xml" href="data:image/svg;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHZpZXdCb3g9IjAgMCAxNiAxNiIgd2lkdGg9IjE2IiBoZWlnaHQ9IjE2Ij48cmVjdCB4PSIwIiB5PSIwIiB3aWR0aD0iMTAwJSIgaGVpZ2h0PSIxMDAlIiBmaWxsPSIjRkZGRkZGIiAvPjwvc3ZnPg==">
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
	//
	// <h1>Section 1</h1><h2>Carousel Layout</h2><br><div id="carouselExampleCaptions" class="carousel slide">
	// <div class="carousel-indicators">
	//   <button type="button" data-bs-target="#carouselExampleCaptions" data-bs-slide-to="0" class="active" aria-current="true" aria-label="Slide 1"></button>
	//   <button type="button" data-bs-target="#carouselExampleCaptions" data-bs-slide-to="1" aria-label="Slide 2"></button>
	//   <button type="button" data-bs-target="#carouselExampleCaptions" data-bs-slide-to="2" aria-label="Slide 3"></button>
	// </div>
	// <div class="carousel-inner">
	//   <div class="carousel-item active">
	// 	<img src="data:image/png;base64,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" class="d-block w-100" alt="">
	// 	<div class="carousel-caption d-none d-md-block">
	// 	  <h5>The week greatly well buy a wild trip.</h5>
	// 	  <p>Who additionally him now did it many being loss have. Himself that under alone pack about this motionless other child. Tomorrow defiant understimate ours theirs this virtually cooperative labour Uzbek. Choir in annually since them soon time firstly between here. Would nightly cash hourly neither hers hospitality what me occasionally.</p>
	// 	</div>
	//   </div>
	//   <div class="carousel-item">
	// 	<img src="data:image/png;base64,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" class="d-block w-100" alt="">
	// 	<div class="carousel-caption d-none d-md-block">
	// 	<h5>Tender week watch class softly.</h5>
	// 	<p>Here might its knit think still his Burmese glasses where. E.g. range troupe off can gang now how Ecuadorian this. Her disregard whom until these handle crowded few many cut. Do tense there whom this your his from last his. Bravely e.g. line whichever down sit who whose we advertising.</p>
	// 	</div>
	//   </div>
	//   <div class="carousel-item">
	// 	<img src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAB4AAAAeCAIAAAC0Ujn1AAAKuklEQVR4nACqClX1AglqpfC/qAK38EnfQ6f1+/x72yZdhCvqivkmvd4s82crZOvqp1cvsgcp+5pUZ8UcAA00wmXOxsX827YtGWTC5K6hJLAdM5498WbbL9dZnhiFhSbwoPTx67KZ9QMIYBYXef7+gCWAhSXc5O0GxxEkPmupP7zBATB0FIhyfdz5Gobl5BEJNmA2VorCQIvbDq6RciYs0kdz0fdeIl2LrCdNRMHevi9t1MXZbSIiRWZVgTmdvDmJHEgBnBUqF+sZpgKaGdAwR+724z3qwRbs/HYFGJBi6/XhSnsOlRufzyNadrtr7RmkVEE+CRadwOiLNUd+ZlVE8PjK9wRaSwdDShiUXl5FgGhyxrDfM9lzzQ76xR+QANeo+iHQlIpm1tQR6Q5GHDEjOSWS62IrEh1+jgJh6RMmtMxixloGxBgGfryk+ZMzfKH431tKC6Pa6c6dJJQENdo7PAHVf0upZFbQctuKxIzsiqK1uuIIm1psDADUFFwYXAYhpiqZXduhxiz5fAIQlYVT5jFBmSmgJesKgh7nLSkxOqP5UdNYpoOol7rcP0kySkJukC6iG/ugmCeNNqpMQkDWDiLRl4k98d5dXA+D47d/3zr/lCMA+5GCbUyU5IFogH/5GsfLjndikPMNuw0sdK1HI0bLiuXNqdJ7K+D9bwB0e7EvZ0+/TdLu5+PJqQW4NT/TDWcCyRYa//GSiahcnA2/GWKjkfsSLwwk+ZXnnESwAkJoywkuoEZ2pTnWwYBMjo/7PgqQ2bL79RQezfV7rYwNrwKmIx/yGa5eUAd54Fi1lBsPaAJPQ/yObJGTmRr3rSxw9Ec53dwOyUBF0eSeHL6NpQtzgkW/8XfG6wTGfqaZ18HG8fHtDkO1fCh55icE5rnNywC5acCBNawDGU8L1fa4g6xXx9TLX+cgAVbg7KvH+uQW/ykCYs4T5vnmZ2wggsP7QFNXWSfYezlVDDeDpuvDQCU7WuwAXTpmoKJIHubfC+Z9kxzJ8ByTrHbeOm26yQUgcDYWY7E4w3Eyxbg8bWhP/gZ3ns0N6Soqu+c5/uYIfonsP21HtK2W2OYAN9TwnXmG+8JCMz5BcJ3CyMQEA3RsAQgUxQ+n3v8NG8zi2+2MXnu6DhD7035T06gNpdZHNW7WYHy0Y6EzwN7CO4i40gBK+Q3G31c+1vdE+Qij++fT2oA//AeE84kr4MJWrNHfiQl/8QfFoLLhlQMrcALtBeS0pN5lPOeqdiFDAmQK7PlcdJrOBtbmRzkXMZxUUAkyApoWg/yXScfPngYdos6wqkkB5IdBishHNTQDVxIev9YOapUCi4eGK+sTHVe+QYvmvVIDw8xVFOUCdvactkKP0wzWEbbuN05vpMcxx/B0WucLBsLFCYmjlxyIjFb5tED91D47BLrP+2uiN7gMsvStLhm4av0nMvOKgrtBPx5er0wL/a83D5VaBjOvr2f/U73Snxk1AKAVGnbwe5MAE1cMN6+baNm6iysfZUS9lJH2NjnLVKZJj12yASMcWU4c21zBW6UcXQrt6VxNDxxfzbOruxNFXlV2vKiZPxWot8vJOT4L53VUvYLISa6jkc7u5wBHAFMxRfqxRoj/nFs5Uje7Eid+VpOEBtkVHJSBG/k9Rm5u/j4AinkdwECG7EzWlxEQNVjXFRMgY/7t8MZ1m56kCSXooYLuvgrKMLUIMF0zJc6eqmQEy1sVI3wE+lZNI1wPHbn25rQwkgf/wTnD+6FZ2p0PD8LEYIyXAPsDw2fOABlkXiYR+2rlFUaPP+6aZWEr77fH82s2/YwwTpr6Hejp/M2KO32zaOVDHTHvDDzknTL4GOCXA6Cdm0B3uSJLOek5qbQvcg3MaaR8msIAj5ffTKqnLYpED2sZpwxN2IF0TH7/l6hYRFZP2B8KJQUe0kaYCjQSWYchU57wBMWs5m9JgNE041RxGoIbxaYP7fpH+gQeEFqy410DutgRw2GzR9oSckVkxIsrMjy/meafR2Vqd8bceVdpEb1ZExXpJhqX9YoIxWX6gU0x+cdqxHFGQqHbAT4EMcsQtT8U9PzunKi5ZjfwVG1MEN2oezoBIuLXDdtF7er334QtHP2yAsqvRxnx+NhEA++tFszCoK4gfWM3Rt4BFstMQC2k4Q/vR4WWFl3Gph9XkQErIAptmNn+r3FMFrKOycGOxDvnjRWLe5ZA6567JFieAuyoF3PH/LyirX8YybxqpieNnZIagPPw0gsC5QFVoCHp40THEYZxuieIMUoiQfXK+Cc9rMj032/XAVYbBCHMKZF5QN+g4dRMSnFW6CE3OXr+8XRk2AmOcFwuAALlpw3WzIEVCfwWsp0PwFkPYW/uMo24NbboppJvRftWp/p36hBT3iEPA5sxeXbpKtn1+ZHwYWMrxamoKMU8FKyge1PYqtvqmKmZxW+bUmjCWOrWySmA/+ZnnJYDvrJY/HqF3ELamFJ73onoeswCcO/EKEq/vd8i4Fr5vmX5BQ3f+uRqDrlN+z1auanLIkyZ6ev7fBG5+CXdbHZ2TQcxKujXfRD+d3yeAC8Waa3RwPHtzip9a8/xAjSnSg8CLYAmDWvXBKpVUu+QuC9Mx+3L6nFcOMTZcCiTOwhSvpNk7fgogjLYSo2a7ewiu3tkLaiqKL9MgQgQDcMp08iY5peEkWUaPt7XmUgCmsVHxt0Qquna5QMNMyny987sNrZNnLHaOv7fIf4zVw/rYbGpT0pOqOkvGZ8u60ba2/SYSgsU/aa1C7xKQ/3eTAhU7q9OmUdcJCc6YN12Uz6sCzFL8AqlE1kjDBu3UnDqn8dDGN4AARo367LkPE3kOMxpWwTAAgtNfGE/8f2NU7Bf46zSU4yC7fWE55iwFP8FU/j4GagCNotseVxUSyxyzjtLC+h+8ZUAx9A+lWUO3NoA923oQUQ7AnXYaDgnPNkAAAHnIwFWFHYW4+tlWIZoFsIxxqBofuG3wOJkTzLTEba2TTRMXAAJ5WLlC4NmnTP0vkiljkYznPyyyYUL7l7J70jtTlsFPdpyWA1UIdj1tGx6fIDmciPEYf7lzACFXzQWEk0e6RKa4dpkGnf7tukVPADJieafuyytPI/kxQbPlXIL+KOUGHh0adohC0aR+DFRMv9bTsTbvZR970BxRfyXrOb7JDljkCi609oDBlyQLlye9onGxSUAq7Yw4scO6FPvUnI9pCrjYBI7lLXu+DLT0deP/tVd5Sqs2LSomEwk9hjsXdYL9JjISCL6OlAH5TyAD/GCSoql1/ZbupnTiCGKWrOMu04BMMTRnx1eOR8cZ9xqAyFrU0/2fyJHoNf60Dbza3Ng+fDJENdIX80/XatQ4XodDlXlZc4NO1qnG8hL8wkP3K9aH18GceJm+si3ZYs+6dXDRqyggxxCGZIN1lVY4j1lgF1X+90Ak6G5ywAuONx8BqtNI9EdHQfVgT1BA66L1yIgy2tRvYanJiN1MULsEy58z/BQuaIa7dFOfRHj+6t8082AuTyIvy+1AQdWgqwJ+duGtOpXSGIrBCaOmOYflOZHFDwLk6cBmDZQi+4D9ZeZINJkDfnL9IjyMRJxnzcNRegL9z7vNJ11eDkBnHGRCaeJas5V9C/32O0b7E6irNxIQ9Vh5MyBSDtwC6xE46o3uYwL+4TRwC31MDUjNUcHGY+uAQAA///GMjiPIp1C6gAAAABJRU5ErkJggg==" class="d-block w-100" alt="">
	// 	<div class="carousel-caption d-none d-md-block">
	// 	<h5>Bored problem extremely frankly play poorly.</h5>
	// 	<p>Myself of ourselves we then what sunglasses of fiction anybody. Ever American this all accordingly instance may provided couple Cormoran. Occasionally over agree stack great anyone recently after how then. From airport yesterday including fact words while than virtually may. Whomever face eye how in where those those has these.</p>
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
	// </div><hr><br><br>
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
	//   <link rel="icon" type="image/svg+xml" href="data:image/svg;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHZpZXdCb3g9IjAgMCAxNiAxNiIgd2lkdGg9IjE2IiBoZWlnaHQ9IjE2Ij48cmVjdCB4PSIwIiB5PSIwIiB3aWR0aD0iMTAwJSIgaGVpZ2h0PSIxMDAlIiBmaWxsPSIjRkZGRkZGIiAvPjwvc3ZnPg==">
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
	//
	// <h1>Section 1</h1><h2>Carousel Layout</h2><br><div id="carouselExampleCaptions" class="carousel slide">
	// <div class="carousel-indicators">
	//   <button type="button" data-bs-target="#carouselExampleCaptions" data-bs-slide-to="0" class="active" aria-current="true" aria-label="Slide 1"></button>
	//   <button type="button" data-bs-target="#carouselExampleCaptions" data-bs-slide-to="1" aria-label="Slide 2"></button>
	//   <button type="button" data-bs-target="#carouselExampleCaptions" data-bs-slide-to="2" aria-label="Slide 3"></button>
	// </div>
	// <div class="carousel-inner">
	//   <div class="carousel-item active">
	// 	<img src="data:image/png;base64,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" class="d-block w-100" alt="">
	// 	<div class="carousel-caption d-none d-md-block">
	// 	  <h5>The week greatly well buy a wild trip.</h5>
	// 	  <p>Who additionally him now did it many being loss have. Himself that under alone pack about this motionless other child. Tomorrow defiant understimate ours theirs this virtually cooperative labour Uzbek. Choir in annually since them soon time firstly between here. Would nightly cash hourly neither hers hospitality what me occasionally.</p>
	// 	</div>
	//   </div>
	//   <div class="carousel-item">
	// 	<img src="data:image/png;base64,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" class="d-block w-100" alt="">
	// 	<div class="carousel-caption d-none d-md-block">
	// 	<h5>Tender week watch class softly.</h5>
	// 	<p>Here might its knit think still his Burmese glasses where. E.g. range troupe off can gang now how Ecuadorian this. Her disregard whom until these handle crowded few many cut. Do tense there whom this your his from last his. Bravely e.g. line whichever down sit who whose we advertising.</p>
	// 	</div>
	//   </div>
	//   <div class="carousel-item">
	// 	<img src="data:image/png;base64,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" class="d-block w-100" alt="">
	// 	<div class="carousel-caption d-none d-md-block">
	// 	<h5>Bored problem extremely frankly play poorly.</h5>
	// 	<p>Myself of ourselves we then what sunglasses of fiction anybody. Ever American this all accordingly instance may provided couple Cormoran. Occasionally over agree stack great anyone recently after how then. From airport yesterday including fact words while than virtually may. Whomever face eye how in where those those has these.</p>
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
	// </div><hr><br><br>
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
	// </div><hr><br><br>

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
	// <img src="data:image/png;base64,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" width="200" height="200" alt=""><hr><br><br>
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
