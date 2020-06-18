package gofakeit

import (
	"fmt"
	"strings"
)

// AppName will generate a random app name
func AppName() string {
	name := ""
	switch Number(1, 3) {
	case 1:
		name = Noun() + Verb()
	case 2:
		name = Color() + Noun()
	case 3:
		name = Animal() + Verb()
	}

	return strings.Title(name)
}

// AppVersion will generate a random app version
func AppVersion() string {
	return fmt.Sprintf("%d", Number(1, 5)) + "." + fmt.Sprintf("%d", Number(1, 20)) + "." + fmt.Sprintf("%d", Number(1, 20))
}

// AppAuthor will generate a random company or person name
func AppAuthor() string {
	if Bool() {
		return Name()
	}

	return Company()
}

func addAppLookup() {
	AddFuncLookup("appname", Info{
		Display:     "App Name",
		Category:    "app",
		Description: "Random app name",
		Example:     "Parkrespond",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return AppName(), nil
		},
	})

	AddFuncLookup("appversion", Info{
		Display:     "App Version",
		Category:    "app",
		Description: "Random app version",
		Example:     "1.12.14",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return AppVersion(), nil
		},
	})

	AddFuncLookup("appauthor", Info{
		Display:     "App Author",
		Category:    "app",
		Description: "Random app author",
		Example:     "Qado Energy, Inc.",
		Output:      "string",
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			return AppAuthor(), nil
		},
	})
}
