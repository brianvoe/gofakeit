package gofakeit

import (
	"fmt"
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

	// Output: INSERT INTO people (id, first_name, price, age, created_at) VALUES (1, 'Priscilla', 985.96, 20, '1925-07-08 17:32:57'),(2, 'Ramon', 639.32, 74, '2018-12-29 07:25:24');
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

	// Output: INSERT INTO people (id, first_name, price, age, created_at) VALUES (1, 'Priscilla', 985.96, 20, '1925-07-08 17:32:57'),(2, 'Ramon', 639.32, 74, '2018-12-29 07:25:24');
}

func TestSQLJSON(t *testing.T) {
	Seed(11)

	AddFuncLookup("jsonperson", Info{
		Category:    "custom",
		Description: "random JSON of a person",
		Example:     `{"first_name":"Bob", "last_name":"Jones"}`,
		Output:      "[]byte",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {

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

	if res != `INSERT INTO people (data) VALUES ('{"first_name":"Priscilla","last_name":"Thornton"}'),('{"first_name":"Brooke","last_name":"Duncan"}');` {
		t.Errorf("SQL query did not match expected output: %s", res)
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

func TestSQLQuotes(t *testing.T) {
	AddFuncLookup("withquote", Info{
		Category:    "custom",
		Description: "a string with a single quote",
		Example:     "O'Brien",
		Output:      "string",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return "O'Brien", nil
		},
	})
	defer RemoveFuncLookup("withquote")

	AddFuncLookup("withquotebytes", Info{
		Category:    "custom",
		Description: "bytes with a single quote",
		Example:     "it's",
		Output:      "[]byte",
		Generate: func(f *Faker, m *MapParams, info *Info) (any, error) {
			return []byte("it's"), nil
		},
	})
	defer RemoveFuncLookup("withquotebytes")

	res, err := SQL(&SQLOptions{
		Table: "people",
		Count: 1,
		Fields: []Field{
			{Name: "last_name", Function: "withquote"},
			{Name: "note", Function: "withquotebytes"},
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	// A single quote inside a string is written twice
	if res != `INSERT INTO people (last_name, note) VALUES ('O''Brien', 'it''s');` {
		t.Errorf("SQL query did not match expected output: %s", res)
	}
}

// sqlStringsClosed reports whether every string in a statement is closed, a doubled quote inside a
// string is part of it
func sqlStringsClosed(sql string) bool {
	inString := false
	for i := 0; i < len(sql); i++ {
		if sql[i] != '\'' {
			continue
		}
		if inString && i+1 < len(sql) && sql[i+1] == '\'' {
			i++
			continue
		}
		inString = !inString
	}
	return !inString
}

// Some names, companies and hobbies have a single quote in them, like Rubik's Cube
func TestSQLGeneratedStringsAreClosed(t *testing.T) {
	f := New(11)
	for range 2000 {
		res, err := f.SQL(&SQLOptions{
			Table: "things",
			Count: 1,
			Fields: []Field{
				{Name: "company", Function: "company"},
				{Name: "hobby", Function: "hobby"},
				{Name: "quote", Function: "quote"},
			},
		})
		if err != nil {
			t.Fatal(err)
		}
		if !sqlStringsClosed(res) {
			t.Fatalf("A string in the SQL query is not closed: %s", res)
		}
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
