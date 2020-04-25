package gofakeit

import (
	"bytes"
	"errors"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"strconv"
)

// ImageURL will generate a random Image Based Upon Height And Width. https://picsum.photos/
func ImageURL(width int, height int) string {
	return "https://picsum.photos/" + strconv.Itoa(width) + "/" + strconv.Itoa(height)
}

// Image generates a random rgba image
func Image(width int, height int) *image.RGBA {
	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	// Set color for each pixel
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			img.Set(x, y, color.RGBA{uint8(Number(0, 255)), uint8(Number(0, 255)), uint8(Number(0, 255)), 0xff})
		}
	}

	return img
}

// ImageJpeg generates a random rgba jpeg image
func ImageJpeg(width int, height int) []byte {
	buf := new(bytes.Buffer)
	jpeg.Encode(buf, Image(width, height), nil)
	return buf.Bytes()
}

// ImagePng generates a random rgba png image
func ImagePng(width int, height int) []byte {
	buf := new(bytes.Buffer)
	png.Encode(buf, Image(width, height))
	return buf.Bytes()
}

func addImageLookup() {
	AddFuncLookup("imageurl", Info{
		Display:     "Image URL",
		Category:    "image",
		Description: "Random image url",
		Example:     "https://picsum.photos/500/500",
		Output:      "string",
		Params: []Param{
			{Field: "width", Display: "Width", Type: "int", Default: "500", Description: "Image width in px"},
			{Field: "height", Display: "Height", Type: "int", Default: "500", Description: "Image height in px"},
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

	AddFuncLookup("imagejpeg", Info{
		Display:     "Image JPEG",
		Category:    "image",
		Description: "Random jpeg image",
		Example:     "file.jpeg - bytes",
		Output:      "[]byte",
		Params: []Param{
			{Field: "width", Display: "Width", Type: "int", Default: "500", Description: "Image width in px"},
			{Field: "height", Display: "Height", Type: "int", Default: "500", Description: "Image height in px"},
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

			return ImageJpeg(width, height), nil
		},
	})

	AddFuncLookup("imagepng", Info{
		Display:     "Image PNG",
		Category:    "image",
		Description: "Random png image",
		Example:     "file.png - bytes",
		Output:      "[]byte",
		Params: []Param{
			{Field: "width", Display: "Width", Type: "int", Default: "500", Description: "Image width in px"},
			{Field: "height", Display: "Height", Type: "int", Default: "500", Description: "Image height in px"},
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

			return ImagePng(width, height), nil
		},
	})
}
