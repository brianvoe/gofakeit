package gofakeit

import (
	"errors"
	"strconv"
)

// ImageURL will generate a random Image Based Upon Height And Width. https://picsum.photos/
func ImageURL(width int, height int) string {
	return "https://picsum.photos/" + strconv.Itoa(width) + "/" + strconv.Itoa(height)
}

func addImageLookup() {
	AddLookupData("imageurl", Info{
		Category:    "image",
		Description: "Random image url",
		Example:     "https://picsum.photos/500/500",
		Params: []Param{
			{Field: "width", Required: false, Type: "int", Default: "500", Description: "Image width"},
			{Field: "height", Required: false, Type: "int", Default: "500", Description: "Image height"},
		},
		Call: func(m *map[string][]string, info *Info) (interface{}, error) {
			width, err := info.GetInt(m, "width")
			if err != nil {
				return nil, err
			}
			if width < 10 || width >= 1000 {
				return nil, errors.New("Invalid image width, must be greater than 10, less than 1000")
			}

			height, err := info.GetInt(m, "height")
			if err != nil {
				return nil, err
			}
			if height < 10 || height >= 1000 {
				return nil, errors.New("Invalid image height, must be greater than 10, less than 1000")
			}

			return ImageURL(width, height), nil
		},
	})
}
