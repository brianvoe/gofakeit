package gofakeit

func MusicName() string { return musicName(GlobalFaker) }

func (f *Faker) MusicName() string { return musicName(f) }

func musicName(f *Faker) string { return getRandValue(f, []string{"music", "name"}) }

func MusicArtist() string { return musicArtist(GlobalFaker) }

func (f *Faker) MusicArtist() string { return musicArtist(f) }

func musicArtist(f *Faker) string { return getRandValue(f, []string{"music", "artist"}) }

func MusicGenre() string { return musicGenre(GlobalFaker) }

func (f *Faker) MusicGenre() string { return musicGenre(f) }

func musicGenre(f *Faker) string { return getRandValue(f, []string{"music", "genre"}) }

type MusicInfo struct {
	Name   string `json:"name" xml:"name"`
	Artist string `json:"artist" xml:"artist"`
	Genre  string `json:"genre" xml:"genre"`
}

func Music() *MusicInfo { return music(GlobalFaker) }

func (f *Faker) Music() *MusicInfo { return music(f) }

func music(f *Faker) *MusicInfo {
	return &MusicInfo{
		Name:   musicName(f),
		Artist: musicArtist(f),
		Genre:  musicGenre(f),
	}
}

func addMusicLookup() {
	AddFuncLookup("music", Info{
		Display:     "Music",
		Category:    "music",
		Description: "Song with a drum and horn instrumentation",
		Example: `{
	"name": "New Rules",
	"genre": "Tropical house"
}`,
		Output:      "map[string]string",
		ContentType: "application/json",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return music(f), nil
		},
	})

	AddFuncLookup("musicname", Info{
		Display:     "Music Name",
		Category:    "music",
		Description: "Title or name of a specific music used for identification and reference",
		Example:     "New Rules",
		Output:      "string",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return musicName(f), nil
		},
	})

	AddFuncLookup("musicartist", Info{
		Display:     "Music Artist",
		Category:    "music",
		Description: "The artist of maker of mucis",
		Example:     "Dua Lipa",
		Output:      "string",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return musicArtist(f), nil
		},
	})

	AddFuncLookup("musicgenre", Info{
		Display:     "Genre",
		Category:    "music",
		Description: "Category that classifies music based on common themes, styles, and storytelling approaches",
		Example:     "Action",
		Output:      "string",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return musicGenre(f), nil
		},
	})
}
