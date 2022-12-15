package gofakeit

import (
	"bytes"
	"errors"
	img "image"
	imgCol "image/color"
	"image/jpeg"
	"image/png"
	"math/rand"
	"strconv"
)

// ImageURL will generate a random Image Based Upon Height And Width. https://picsum.photos/
func ImageURL(width int, height int) string { return imageURL(globalFaker.Rand, width, height) }

// ImageURL will generate a random Image Based Upon Height And Width. https://picsum.photos/
func (f *Faker) ImageURL(width int, height int) string { return imageURL(f.Rand, width, height) }

func imageURL(r *rand.Rand, width int, height int) string {
	return "https://picsum.photos/" + strconv.Itoa(width) + "/" + strconv.Itoa(height)
}

// Image generates a random rgba image
func Image(width int, height int) *img.RGBA { return image(globalFaker.Rand, width, height) }

// Image generates a random rgba image
func (f *Faker) Image(width int, height int) *img.RGBA { return image(f.Rand, width, height) }

func image(r *rand.Rand, width int, height int) *img.RGBA {
	upLeft := img.Point{0, 0}
	lowRight := img.Point{width, height}

	img := img.NewRGBA(img.Rectangle{upLeft, lowRight})

	// Set color for each pixel
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			img.Set(x, y, imgCol.RGBA{uint8(number(r, 0, 255)), uint8(number(r, 0, 255)), uint8(number(r, 0, 255)), 0xff})
		}
	}

	return img
}

// ImageJpeg generates a random rgba jpeg image
func ImageJpeg(width int, height int) []byte { return imageJpeg(globalFaker.Rand, width, height) }

// ImageJpeg generates a random rgba jpeg image
func (f *Faker) ImageJpeg(width int, height int) []byte { return imageJpeg(f.Rand, width, height) }

func imageJpeg(r *rand.Rand, width int, height int) []byte {
	buf := new(bytes.Buffer)
	jpeg.Encode(buf, image(r, width, height), nil)
	return buf.Bytes()
}

// ImagePng generates a random rgba png image
func ImagePng(width int, height int) []byte { return imagePng(globalFaker.Rand, width, height) }

// ImagePng generates a random rgba png image
func (f *Faker) ImagePng(width int, height int) []byte { return imagePng(f.Rand, width, height) }

func imagePng(r *rand.Rand, width int, height int) []byte {
	buf := new(bytes.Buffer)
	png.Encode(buf, image(r, width, height))
	return buf.Bytes()
}

type ImageSVGOptions struct {
	Height int
	Width  int
	Type   string
	Colors []string
}

// List of available svg types
var SvgTypes = []string{"rect", "circle", "ellipse", "line", "polyline", "polygon"}

// Generate a random svg generator
func ImageSvg(options *ImageSVGOptions) string { return imageSvg(globalFaker.Rand, options) }

// Generate a random svg generator
func (f *Faker) ImageSvg(options *ImageSVGOptions) string { return imageSvg(f.Rand, options) }

func imageSvg(r *rand.Rand, options *ImageSVGOptions) string {
	// If options is nil, set it to empty struct
	if options == nil {
		options = &ImageSVGOptions{}
	}

	// If options height and weight is not set, set it to random number between 100 and 500
	if options.Width == 0 {
		options.Width = number(r, 100, 500)
	}
	widthStr := strconv.Itoa(options.Width)
	if options.Height == 0 {
		options.Height = number(r, 100, 500)
	}
	heightStr := strconv.Itoa(options.Height)

	// Check if type is set, if not set to random type
	if options.Type == "" {
		options.Type = randomString(r, SvgTypes)
	}

	// If the colors are not set, set it to a set of nice colors
	if len(options.Colors) == 0 {
		options.Colors = niceColors(r)
	}

	// Start svg string
	svgStr := `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 ` + widthStr + ` ` + heightStr + `" width="` + widthStr + `" height="` + heightStr + `">`

	// Add a rect for the background
	svgStr += `<rect x="0" y="0" width="100%" height="100%" fill="` + randomString(r, options.Colors) + `" />`

	// Add a random number of shapes
	for i := 0; i < number(r, 10, 20); i++ {
		// Add a random shape
		switch options.Type {
		case "rect":
			svgStr += `<rect x="` + strconv.Itoa(number(r, 0, options.Width)) + `" y="` + strconv.Itoa(number(r, 0, options.Height)) + `" width="` + strconv.Itoa(number(r, 0, options.Width)) + `" height="` + strconv.Itoa(number(r, 0, options.Height)) + `" fill="` + randomString(r, options.Colors) + `" />`
		case "circle":
			svgStr += `<circle cx="` + strconv.Itoa(number(r, 0, options.Width)) + `" cy="` + strconv.Itoa(number(r, 0, options.Height)) + `" r="` + strconv.Itoa(number(r, 0, options.Width)) + `" fill="` + randomString(r, options.Colors) + `" />`
		case "ellipse":
			svgStr += `<ellipse cx="` + strconv.Itoa(number(r, 0, options.Width)) + `" cy="` + strconv.Itoa(number(r, 0, options.Height)) + `" rx="` + strconv.Itoa(number(r, 0, options.Width)) + `" ry="` + strconv.Itoa(number(r, 0, options.Height)) + `" fill="` + randomString(r, options.Colors) + `" />`
		case "line":
			svgStr += `<line x1="` + strconv.Itoa(number(r, 0, options.Width)) + `" y1="` + strconv.Itoa(number(r, 0, options.Height)) + `" x2="` + strconv.Itoa(number(r, 0, options.Width)) + `" y2="` + strconv.Itoa(number(r, 0, options.Height)) + `" stroke="` + randomString(r, options.Colors) + `" />`
		case "polyline":
			svgStr += `<polyline points="` + strconv.Itoa(number(r, 0, options.Width)) + `,` + strconv.Itoa(number(r, 0, options.Height)) + ` ` + strconv.Itoa(number(r, 0, options.Width)) + `,` + strconv.Itoa(number(r, 0, options.Height)) + ` ` + strconv.Itoa(number(r, 0, options.Width)) + `,` + strconv.Itoa(number(r, 0, options.Height)) + `" fill="` + randomString(r, options.Colors) + `" />`
		case "polygon":
			svgStr += `<polygon points="` + strconv.Itoa(number(r, 0, options.Width)) + `,` + strconv.Itoa(number(r, 0, options.Height)) + ` ` + strconv.Itoa(number(r, 0, options.Width)) + `,` + strconv.Itoa(number(r, 0, options.Height)) + ` ` + strconv.Itoa(number(r, 0, options.Width)) + `,` + strconv.Itoa(number(r, 0, options.Height)) + `" fill="` + randomString(r, options.Colors) + `" />`
		}
	}

	// End svg string
	svgStr += `</svg>`

	return svgStr
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
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (interface{}, error) {
			width, err := info.GetInt(m, "width")
			if err != nil {
				return nil, err
			}
			if width < 10 || width >= 1000 {
				return nil, errors.New("invalid image width, must be greater than 10, less than 1000")
			}

			height, err := info.GetInt(m, "height")
			if err != nil {
				return nil, err
			}
			if height < 10 || height >= 1000 {
				return nil, errors.New("invalid image height, must be greater than 10, less than 1000")
			}

			return imageURL(r, width, height), nil
		},
	})

	AddFuncLookup("imagejpeg", Info{
		Display:     "Image JPEG",
		Category:    "image",
		Description: "Random jpeg image",
		Example:     "file.jpeg - bytes",
		Output:      "[]byte",
		ContentType: "image/jpeg",
		Params: []Param{
			{Field: "width", Display: "Width", Type: "int", Default: "500", Description: "Image width in px"},
			{Field: "height", Display: "Height", Type: "int", Default: "500", Description: "Image height in px"},
		},
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (interface{}, error) {
			width, err := info.GetInt(m, "width")
			if err != nil {
				return nil, err
			}
			if width < 10 || width >= 1000 {
				return nil, errors.New("invalid image width, must be greater than 10, less than 1000")
			}

			height, err := info.GetInt(m, "height")
			if err != nil {
				return nil, err
			}
			if height < 10 || height >= 1000 {
				return nil, errors.New("invalid image height, must be greater than 10, less than 1000")
			}

			return imageJpeg(r, width, height), nil
		},
	})

	AddFuncLookup("imagepng", Info{
		Display:     "Image PNG",
		Category:    "image",
		Description: "Random png image",
		Example:     "file.png - bytes",
		Output:      "[]byte",
		ContentType: "image/png",
		Params: []Param{
			{Field: "width", Display: "Width", Type: "int", Default: "500", Description: "Image width in px"},
			{Field: "height", Display: "Height", Type: "int", Default: "500", Description: "Image height in px"},
		},
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (interface{}, error) {
			width, err := info.GetInt(m, "width")
			if err != nil {
				return nil, err
			}
			if width < 10 || width >= 1000 {
				return nil, errors.New("invalid image width, must be greater than 10, less than 1000")
			}

			height, err := info.GetInt(m, "height")
			if err != nil {
				return nil, err
			}
			if height < 10 || height >= 1000 {
				return nil, errors.New("invalid image height, must be greater than 10, less than 1000")
			}

			return imagePng(r, width, height), nil
		},
	})
}
