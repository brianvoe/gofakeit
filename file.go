package gofakeit

// FileExtension will generate a random file extension
func FileExtension() string { return fileExtension(GlobalFaker) }

// FileExtension will generate a random file extension
func (f *Faker) FileExtension() string { return fileExtension(f) }

func fileExtension(f *Faker) string { return getRandValue(f, []string{"file", "extension"}) }

// FileMimeType will generate a random mime file type
func FileMimeType() string { return fileMimeType(GlobalFaker) }

// FileMimeType will generate a random mime file type
func (f *Faker) FileMimeType() string { return fileMimeType(f) }

func fileMimeType(f *Faker) string { return getRandValue(f, []string{"file", "mime_type"}) }

func addFileLookup() {
	AddFuncLookup("fileextension", Info{
		Display:     "File Extension",
		Category:    "file",
		Description: "Suffix appended to a filename indicating its format or type",
		Example:     "nes",
		Output:      "string",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return fileExtension(f), nil
		},
	})

	AddFuncLookup("filemimetype", Info{
		Display:     "File Mime Type",
		Category:    "file",
		Description: "Defines file format and nature for browsers and email clients using standardized identifiers",
		Example:     "application/json",
		Output:      "string",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return fileMimeType(f), nil
		},
	})
}
