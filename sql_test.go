package gofakeit

import (
	"fmt"
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
