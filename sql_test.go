package gofakeit

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
)

func ExampleSQL() {
	Seed(11)

	res, _ := SQL(&SQLOptions{
		Table: "people",
		Count: 2,
		Fields: []Field{
			{Name: "id", Function: "autoincrement"},
			{Name: "first_name", Function: "firstname"},
			{Name: "price", Function: "price"},
			{Name: "age", Function: "number", Params: MapParams{"min": {"1"}, "max": {"99"}}},
			{Name: "created_at", Function: "date", Params: MapParams{"format": {"2006-01-02 15:04:05"}}},
		},
	})

	fmt.Println(string(res))

	// Output:
	// INSERT INTO people (id, first_name, price, age, created_at) VALUES (1, 'Markus', 804.92, 21, '1989-01-30 07:58:01'),(2, 'Santino', 235.13, 40, '1926-07-07 22:25:40');
}

func ExampleFaker_SQL() {
	f := New(11)

	res, _ := f.SQL(&SQLOptions{
		Table: "people",
		Count: 2,
		Fields: []Field{
			{Name: "id", Function: "autoincrement"},
			{Name: "first_name", Function: "firstname"},
			{Name: "price", Function: "price"},
			{Name: "age", Function: "number", Params: MapParams{"min": {"1"}, "max": {"99"}}},
			{Name: "created_at", Function: "date", Params: MapParams{"format": {"2006-01-02 15:04:05"}}},
		},
	})

	fmt.Println(string(res))

	// Output:
	// INSERT INTO people (id, first_name, price, age, created_at) VALUES (1, 'Markus', 804.92, 21, '1996-11-22 07:34:00'),(2, 'Anibal', 674.87, 60, '1973-01-03 11:07:53');
}

func TestSQLJSON(t *testing.T) {
	Seed(11)

	AddFuncLookup("jsonperson", Info{
		Category:    "custom",
		Description: "random JSON of a person",
		Example:     `{"first_name":"Bob", "last_name":"Jones"}`,
		Output:      "[]byte",
		Generate: func(r *rand.Rand, m *MapParams, info *Info) (any, error) {

			v, _ := JSON(&JSONOptions{
				Type: "object",
				Fields: []Field{
					{Name: "first_name", Function: "firstname"},
					{Name: "last_name", Function: "lastname"},
				},
			})

			return v, nil
		},
	})
	defer RemoveFuncLookup("jsonperson")

	res, err := SQL(&SQLOptions{
		Table: "people",
		Count: 2,
		Fields: []Field{
			{Name: "data", Function: "jsonperson"},
		},
	})

	if err != nil {
		t.Fatal(err)
	}

	if res != `INSERT INTO people (data) VALUES ('{"first_name":"Markus","last_name":"Moen"}'),('{"first_name":"Alayna","last_name":"Wuckert"}');` {
		t.Error("SQL query does not match")
	}
}

func TestSQLAll(t *testing.T) {
	Seed(11)

	res, err := SQL(&SQLOptions{
		Table: "people",
		Count: 3,
		Fields: []Field{
			{Name: "id", Function: "autoincrement"},
			{Name: "first_name", Function: "firstname"},
			{Name: "balance", Function: "float32"},
		},
	})

	if err != nil {
		t.Fatal(err)
	}

	// Split by VALUES
	values := strings.Split(res, "VALUES")

	// Check to make sure there are 3 ( and ) symbols
	if strings.Count(values[1], "(") != 3 || strings.Count(values[1], ")") != 3 {
		t.Error("SQL query does not have 3 values")
	}
}

func TestSingleSQL(t *testing.T) {
	Seed(11)

	res, err := SQL(&SQLOptions{
		Table: "People",
		Count: 1,
		Fields: []Field{
			{Name: "first_name", Function: "firstname"},
			{Name: "last_name", Function: "lastname"},
		},
	})

	if err != nil {
		t.Fatal(err)
	}

	// Split by VALUES
	values := strings.Split(res, "VALUES")

	// Check to make sure there are 3 ( and ) symbols
	if strings.Count(values[1], "(") != 1 || strings.Count(values[1], ")") != 1 {
		t.Error("SQL query should have 1 value")
	}
}

func TestSQLNoCount(t *testing.T) {
	Seed(11)

	_, err := SQL(&SQLOptions{
		Table: "People",
		Fields: []Field{
			{Name: "first_name", Function: "firstname"},
			{Name: "last_name", Function: "lastname"},
		},
	})

	if err == nil {
		t.Fatal("should have failed for no count")
	}
}

func TestSQLNoFields(t *testing.T) {
	Seed(11)

	_, err := SQL(&SQLOptions{
		Table:  "People",
		Count:  1,
		Fields: []Field{},
	})

	if err == nil {
		t.Fatal("should have failed for no fields")
	}
}

func TestSQLNilFields(t *testing.T) {
	Seed(11)

	_, err := SQL(&SQLOptions{
		Table: "People",
		Count: 1,
	})

	if err == nil {
		t.Fatal("should have failed for nil fields")
	}
}

func TestSQLInvalidFunction(t *testing.T) {
	Seed(11)

	_, err := SQL(&SQLOptions{
		Table: "People",
		Fields: []Field{
			{Name: "thing", Function: "stuff"},
		},
		Count: 1,
	})

	if err == nil {
		t.Fatal("should have failed for invalid function")
	}
}

func TestSQLNilTable(t *testing.T) {
	Seed(11)

	_, err := SQL(&SQLOptions{
		Count: 3,
		Fields: []Field{
			{Name: "first_name", Function: "firstname"},
			{Name: "last_name", Function: "lastname"},
		},
	})

	if err == nil {
		t.Fatal("should have failed for no table")
	}
}
