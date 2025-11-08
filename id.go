package gofakeit

const (
	idLength       = 20
	idBitsPerChar  = 5
	idAlphabetMask = (1 << idBitsPerChar) - 1
	// readable 32 chars, (no 0, o, 1, i, l)
	// 0 and o are removed to avoid confusion with each other
	// 1, i, l are removed to avoid confusion with each other
	// extra g was added to fit 32 chars
	idAlphabetStr = "23456789abcdefgghjkmnpqrstuvwxyz"
	hexDigits     = "0123456789abcdef"
)

var (
	idAlphabet = []byte(idAlphabetStr)
)

// ID will return a random unique identifier
func ID() string { return id(GlobalFaker) }

// ID will return a random unique identifier
func (f *Faker) ID() string { return id(f) }

func id(f *Faker) string {
	out := make([]byte, idLength)

	var cache uint64
	var bits uint

	for i := 0; i < idLength; {
		if bits < idBitsPerChar {
			cache = f.Uint64()
			bits = 64
		}

		index := cache & idAlphabetMask
		cache >>= idBitsPerChar
		bits -= idBitsPerChar

		// optimization: remove this check to avoid bounds check
		// if index >= uint64(idAlphabetLen) {
		// 	continue
		// }

		out[i] = idAlphabet[index]
		i++
	}

	return string(out)
}

// UUID (version 4) will generate a random unique identifier based upon random numbers
func UUID() string { return uuid(GlobalFaker) }

// UUID (version 4) will generate a random unique identifier based upon random numbers
func (f *Faker) UUID() string { return uuid(f) }

func uuid(f *Faker) string {
	const version = byte(4)

	var uuid [16]byte
	var r uint64

	r = f.Uint64()
	for i := 0; i < 8; i++ {
		uuid[i] = byte(r)
		r >>= 8
	}

	r = f.Uint64()
	for i := 8; i < 16; i++ {
		uuid[i] = byte(r)
		r >>= 8
	}

	uuid[6] = (uuid[6] & 0x0f) | (version << 4)
	uuid[8] = (uuid[8] & 0x3f) | 0x80

	var buf [36]byte
	encodeHexLower(buf[0:8], uuid[0:4])
	buf[8] = dash
	encodeHexLower(buf[9:13], uuid[4:6])
	buf[13] = dash
	encodeHexLower(buf[14:18], uuid[6:8])
	buf[18] = dash
	encodeHexLower(buf[19:23], uuid[8:10])
	buf[23] = dash
	encodeHexLower(buf[24:], uuid[10:])

	return string(buf[:])
}

func encodeHexLower(dst, src []byte) {
	for i, b := range src {
		dst[i*2] = hexDigits[b>>4]
		dst[i*2+1] = hexDigits[b&0x0f]
	}
}

func addIDLookup() {
	AddFuncLookup("id", Info{
		Display:     "ID",
		Category:    "id",
		Description: "Generates a short, URL-safe base32 identifier using a custom alphabet that avoids lookalike characters",
		Example:     "pfsfktb87rcmj6bqha2fz9",
		Output:      "string",
		Aliases:     []string{"unique id", "random id", "base32 id", "url-safe id", "slug id", "short id"},
		Keywords:    []string{"id", "random", "base32", "slug", "token", "url", "identifier", "nonsequential"},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return id(f), nil
		},
	})

	AddFuncLookup("uuid", Info{
		Display:     "UUID",
		Category:    "id",
		Description: "Generates a RFC 4122 compliant version 4 UUID using the faker random source",
		Example:     "b4ddf623-4ea6-48e5-9292-541f028d1fdb",
		Output:      "string",
		Aliases:     []string{"identifier", "guid", "uuid v4", "128-bit", "uuid generator"},
		Keywords:    []string{"unique", "v4", "hex", "computer", "system", "random"},
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return uuid(f), nil
		},
	})
}
