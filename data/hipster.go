package data

// Hipster consists of random hipster words
var Hipster = map[string][]string{
	"word": {"Wes Anderson", "chicharrones", "narwhal", "food truck", "marfa", "aesthetic", "keytar", "art party", "sustainable", "forage", "mlkshk", "gentrify", "locavore", "swag", "hoodie", "microdosing", "VHS", "before they sold out", "pabst", "plaid", "Thundercats", "freegan", "scenester", "hella", "occupy", "truffaut", "raw denim", "beard", "post-ironic", "photo booth", "twee", "90's", "pitchfork", "cray", "cornhole", "kale chips", "pour-over", "yr", "five dollar toast", "kombucha", "you probably haven't heard of them", "mustache", "fixie", "try-hard", "franzen", "kitsch", "austin", "stumptown", "keffiyeh", "whatever", "tumblr", "DIY", "shoreditch", "biodiesel", "vegan", "pop-up", "banjo", "kogi", "cold-pressed", "letterpress", "chambray", "butcher", "synth", "trust fund", "hammock", "farm-to-table", "intelligentsia", "loko", "ugh", "offal", "poutine", "gastropub", "Godard", "jean shorts", "sriracha", "dreamcatcher", "leggings", "fashion axe", "church-key", "meggings", "tote bag", "disrupt", "readymade", "helvetica", "flannel", "meh", "roof", "hashtag", "knausgaard", "cronut", "schlitz", "green juice", "waistcoat", "normcore", "viral", "ethical", "actually", "fingerstache", "humblebrag", "deep v", "wayfarers", "tacos", "taxidermy", "selvage", "put a bird on it", "ramps", "portland", "retro", "kickstarter", "bushwick", "brunch", "distillery", "migas", "flexitarian", "XOXO", "small batch", "messenger bag", "heirloom", "tofu", "bicycle rights", "bespoke", "salvia", "wolf", "selfies", "echo", "park", "listicle", "craft beer", "chartreuse", "sartorial", "pinterest", "mumblecore", "kinfolk", "vinyl", "etsy", "umami", "8-bit", "polaroid", "banh mi", "crucifix", "bitters", "brooklyn", "PBR&B", "drinking", "vinegar", "squid", "tattooed", "skateboard", "vice", "authentic", "literally", "lomo", "celiac", "health", "goth", "artisan", "chillwave", "blue bottle", "pickled", "next level", "neutra", "organic", "Yuccie", "paleo", "blog", "single-origin coffee", "seitan", "street", "gluten-free", "mixtape", "venmo", "irony", "everyday", "carry", "slow-carb", "3 wolf moon", "direct trade", "lo-fi", "tousled", "tilde", "semiotics", "cred", "chia", "master", "cleanse", "ennui", "quinoa", "pug", "iPhone", "fanny pack", "cliche", "cardigan", "asymmetrical", "meditation", "YOLO", "typewriter", "pork belly", "shabby chic", "+1", "lumbersexual", "williamsburg"},
	"sentence": {
		// ultra-short / taglines
		"{hipsterword} vibes only.",
		"deeply {hipsterword}.",
		"casually {hipsterword}.",
		"endlessly {hipsterword}.",
		"pure {hipsterword} energy.",
		"very {hipsterword}, very {hipsterword}.",
		"{hipsterword} meets {hipsterword}.",
		"low-fi {hipsterword}.",
		"post-{hipsterword} mood.",
		"tastefully {hipsterword}.",

		// one-liners
		"living that {hipsterword} life.",
		"another day, another {hipsterword}.",
		"strictly {hipsterword} palettes.",
		"hand-picked {hipsterword}.",
		"soft {hipsterword} glow in {city}.",
		"from {city} with {hipsterword}.",
		"born in {city}, raised on {hipsterword}.",
		"powered by {hipsterword} and {hipsterword}.",
		"mildly obsessed with {hipsterword}.",
		"just add {hipsterword}.",

		// comparisons / mashups
		"{hipsterword} meets {hipsterword} in {city}.",
		"{hipsterword}, but make it {hipsterword}.",
		"{hipsterword} × {hipsterword}, minimal edition.",
		"{hipsterword} on the outside, {hipsterword} at heart.",
		"somewhere between {hipsterword} and {hipsterword}.",
		"{hipsterword} layered over {hipsterword}.",
		"{hipsterword} with a hint of {hipsterword}.",
		"heavy on {hipsterword}, light on {hipsterword}.",
		"{hipsterword} > {hipsterword}, discuss.",

		// scene setters
		"weekends are for {hipsterword} and {beerstyle}.",
		"late nights, {hipsterword} playlists, {songgenre} loops.",
		"mornings in {city}, afternoons in {hipsterword}.",
		"slow afternoons, {hipsterword} everything.",
		"back alley {hipsterword}, side-door {hipsterword}.",
		"under neon {color}, pure {hipsterword}.",
		"rooftop in {city}, whispering {hipsterword}.",
		"{hipsterword} at golden hour in {city}.",
		"between {hipsterword} stalls and {hipsterword} pop-ups.",

		// lifestyle / verbs
		"{verb} through {hipsterword} alleys.",
		"{verb} {noun} with {hipsterword} flair.",
		"{verb} the {noun}, keep it {hipsterword}.",
		"{verb} slowly; trust the {hipsterword}.",
		"{verb} toward {hipsterword} minimalism.",
		"{verb} and {verb}, always {hipsterword}.",
		"let it be {hipsterword}, let it be {adjective}.",

		// craft / food / drink
		"small-batch {hipsterword} in {city}.",
		"single-origin {hipsterword}, double {adjective}.",
		"farm-to-table {hipsterword} and {noun}.",
		"house {hipsterword} with a {adjective} finish.",
		"cold {hipsterword}, warm {hipsterword}.",
		"locally sourced {hipsterword}, globally {adjective}.",
		"{hipsterword} tasting notes: {adjective}, {adjective}, {adjective}.",
		"pairing {hipsterword} with {beerstyle}.",

		// fashion / objects
		"{color} threads, {hipsterword} cuts.",
		"{hipsterword} layers over {adjective} basics.",
		"{hipsterword} silhouettes, {adjective} textures.",
		"{hipsterword} frames and {color} accents.",
		"{hipsterword} tote with {noun} inside.",
		"imperfect by design, perfectly {hipsterword}.",

		// travel / place
		"lost in {city}, found in {hipsterword}.",
		"passport full of {hipsterword}.",
		"{hipsterword} maps and {language} menus.",
		"backstreets of {city}, front row {hipsterword}.",
		"from {country} to {city}, chasing {hipsterword}.",
		"tiny studio in {city}, big {hipsterword} dreams.",

		// meta / attitude
		"ironically {hipsterword}, sincerely {adjective}.",
		"intentionally {hipsterword}.",
		"unapologetically {hipsterword}.",
		"quietly {hipsterword} since {number:2001,2020}.",
		"probably too {hipsterword} for this.",
		"you probably haven’t tried this {hipsterword}.",
		"subtly {hipsterword}, never loud.",

		// maker / work
		"built with {hipsterword} and {hobby}.",
		"shipping {hipsterword} from {city}.",
		"{programminglanguage} by day, {hipsterword} by night.",
		"{company} meets {hipsterword} ethos.",
		"{productcategory} with {hipsterword} edges.",
		"designing around {hipsterword} constraints.",

		// music / film / culture
		"{songgenre} loops with {hipsterword} undertones.",
		"cinema nights, strictly {hipsterword}.",
		"{hipsterword} soundtrack on repeat.",
		"scored in {hipsterword}, mixed {adverb}.",
		"director’s cut: more {hipsterword}.",

		// social / shareable
		"tag it {hipsterword}, keep it {adjective}.",
		"overheard in {city}: very {hipsterword}.",
		"sent from a {hipsterword} corner of {city}.",
		"friends don’t let friends skip {hipsterword}.",
		"if you know, it’s {hipsterword}.",
	},
}
