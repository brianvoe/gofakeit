package gofakeit

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestMultiSQLInsert(t *testing.T) {
	Seed(11)

	res, _ := SQLInsert(&SQLOptions{
		Table:      "People",
		EntryCount: 3,
		Fields: []Field{
			{Name: "first_name", Function: "firstname"},
			{Name: "last_name", Function: "lastname"},
			{Name: "age", Function: "number", Params: MapParams{"min": {"1"}, "max": {"99"}}},
		},
	})

	fmt.Println(string(res))

	// Output:
	// INSERT INTO People VALUES ('Markus', 'Moen', 21), ('Anibal', 'Kozey', 60), ('Sylvan', 'Mraz', 59);
}

func TestMultiSQLInsertJSON(t *testing.T) {

	Seed(12)

	AddFuncLookup("jsonperson", Info{
		Category:    "custom",
		Description: "random JSON of a person",
		Example:     `{"first_name":"Bob", "last_name":"Jones"}`,
		Output:      "[]byte",
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (interface{}, error) {

			v, _ := JSON(&JSONOptions{
				Type:     "object",
				RowCount: 1,
				Fields: []Field{
					{Name: "first_name", Function: "firstname"},
					{Name: "last_name", Function: "lastname"},
				},
			})

			return v, nil
		},
	})

	type RandomPerson struct {
		FakePerson []byte `fake:{jsonperson}"`
	}

	var f RandomPerson
	Struct(&f)
	fmt.Printf("%s\n", (f.FakePerson))
}

func TestMultiSQLInsertFloat(t *testing.T) {
	Seed(11)

	res, _ := SQLInsert(&SQLOptions{
		Table:      "People",
		EntryCount: 3,
		Fields: []Field{
			{Name: "first_name", Function: "firstname"},
			{Name: "last_name", Function: "lastname"},
			{Name: "balance", Function: "float32"},
		},
	})

	fmt.Println(string(res))

	// Output:
	// INSERT INTO People VALUES ('Markus', 'Moen', 2.7766049e+38), ('Anibal', 'Kozey', 3.8815667e+37), ('Sylvan', 'Mraz', 1.6268478e+38);
}

func TestMultiSQLInsertAutoincrement(t *testing.T) {
	Seed(11)

	res, _ := SQLInsert(&SQLOptions{
		Table:      "People",
		EntryCount: 3,
		Fields: []Field{
			{Name: "id", Function: "autoincrement"},
			{Name: "first_name", Function: "firstname"},
			{Name: "last_name", Function: "lastname"},
			{Name: "age", Function: "number", Params: MapParams{"min": {"1"}, "max": {"99"}}},
		},
	})

	fmt.Println(string(res))

	// Output:
	// INSERT INTO People VALUES (1, 'Markus', 'Moen', 21), (2, 'Anibal', 'Kozey', 60), (3, 'Sylvan', 'Mraz', 59);
}

func TestSignleSQLInsert(t *testing.T) {
	Seed(114)

	res, _ := SQLInsert(&SQLOptions{
		Table:      "People",
		EntryCount: 1,
		Fields: []Field{
			{Name: "first_name", Function: "firstname"},
			{Name: "last_name", Function: "lastname"},
		},
	})
	fmt.Println(string(res))

	// Output:
	// INSERT INTO People VALUES ('Colt', 'Koss');
}

func TestSQLInsertNoFields(t *testing.T) {
	Seed(114)

	_, err := SQLInsert(&SQLOptions{
		Table:      "People",
		EntryCount: 1,
		Fields:     []Field{},
	})

	if err != nil {
		fmt.Println(err.Error())
	}

	// Output:
	// must pass fields in order to generate SQL queries
}

func TestSQLInsertNilFields(t *testing.T) {
	Seed(114)

	_, err := SQLInsert(&SQLOptions{
		Table:      "People",
		EntryCount: 1,
	})

	if err != nil {
		fmt.Println(err.Error())
	}

	// Output:
	// must pass fields in order to generate SQL queries
}

func TestSQLInsertNilTable(t *testing.T) {
	Seed(11)

	_, err := SQLInsert(&SQLOptions{
		EntryCount: 3,
		Fields: []Field{
			{Name: "first_name", Function: "firstname"},
			{Name: "last_name", Function: "lastname"},
		},
	})

	if err != nil {
		fmt.Println(err.Error())
	}

	// Output:
	// must provide table name to generate SQL
}
