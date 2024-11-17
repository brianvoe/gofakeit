package data

// Ref:
// 	- https://www.isbn-international.org/range_file_generation
// 	- https://github.com/joke2k/faker/blob/master/faker/providers/isbn/en_US/__init__.py

const PrefixGroup1 = "978"

type RegistrantElements struct {
	Min    string
	Max    string
	Length int
}
type Rule map[string][]RegistrantElements

var ISBNRules = map[string]Rule{
	PrefixGroup1: {
		// Registration Group
		"0": {
			// Registrant Rule (min, max, registrant element length)
			{"0000000", "1999999", 2},
			{"2000000", "2279999", 3},
			{"2280000", "2289999", 4},
			{"2290000", "6479999", 3},
			{"6480000", "6489999", 7},
			{"6490000", "6999999", 3},
			{"7000000", "8499999", 4},
			{"8500000", "8999999", 5},
			{"9000000", "9499999", 6},
			{"9500000", "9999999", 7},
		},
		"1": {
			{"0000000", "0999999", 2},
			{"1000000", "3999999", 3},
			{"4000000", "5499999", 4},
			{"5500000", "7319999", 5},
			{"7320000", "7399999", 7},
			{"7400000", "8697999", 5},
			{"8698000", "9729999", 6},
			{"9730000", "9877999", 4},
			{"9878000", "9989999", 6},
			{"9990000", "9999999", 7},
		},
	},
}
