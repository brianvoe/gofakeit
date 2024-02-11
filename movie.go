package gofakeit

func MovieName() string { return movieName(GlobalFaker) }

func (f *Faker) MovieName() string { return movieName(f) }

func movieName(f *Faker) string { return getRandValue(f, []string{"movie", "name"}) }

func MovieGenre() string { return movieGenre(GlobalFaker) }

func (f *Faker) MovieGenre() string { return movieGenre(f) }

func movieGenre(f *Faker) string { return getRandValue(f, []string{"movie", "genre"}) }

type MovieInfo struct {
	Name  string `json:"name" xml:"name"`
	Genre string `json:"genre" xml:"genre"`
}

func Movie() *MovieInfo { return movie(GlobalFaker) }

func (f *Faker) Movie() *MovieInfo { return movie(f) }

func movie(f *Faker) *MovieInfo {
	return &MovieInfo{
		Name:  movieName(f),
		Genre: movieGenre(f),
	}
}

func addMovieLookup() {
	AddFuncLookup("movie", Info{
		Display:     "Movie",
		Category:    "movie",
		Description: "A story told through moving pictures and sound",
		Example: `{
	"name": "Psycho",
	"genre": "Mystery"
}`,
		Output:      "map[string]string",
		ContentType: "application/json",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return movie(f), nil
		},
	})

	AddFuncLookup("moviename", Info{
		Display:     "Movie Name",
		Category:    "movie",
		Description: "Title or name of a specific film used for identification and reference",
		Example:     "The Matrix",
		Output:      "string",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return movieName(f), nil
		},
	})

	AddFuncLookup("moviegenre", Info{
		Display:     "Genre",
		Category:    "movie",
		Description: "Category that classifies movies based on common themes, styles, and storytelling approaches",
		Example:     "Action",
		Output:      "string",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return movieGenre(f), nil
		},
	})
}
