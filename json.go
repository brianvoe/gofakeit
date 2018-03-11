package gofakeit

import (
	"encoding/json"
	"errors"
)

// JSONString creates a random JSON using a JSON template
func JSONString(template string) (interface{}, error) {
	var templateJSON map[string]interface{}
	t := []byte(template)
	err := json.Unmarshal(t, &templateJSON)
	if err != nil {
		return nil, err
	}
	return randJSON(templateJSON)
}

// JSON creates a random JSON from a map or slice
func JSON(m interface{}) (interface{}, error) {
	return randJSON(m)
}

func randJSON(m interface{}) (interface{}, error) {
	switch v := m.(type) {
	case map[string]interface{}:
		return randMap(v)
	case []interface{}:
		return randJSONArray(v)
	default:
		return nil, errors.New("unknown type")
	}
}

func randMap(m map[string]interface{}) (map[string]interface{}, error) {
	result := make(map[string]interface{})
	for k, v := range m {
		randKey, randValue, err := randJSONPair(k, v)
		if err != nil {
			return nil, err
		}
		result[randKey] = randValue
	}
	return result, nil
}

func randJSONArray(p []interface{}) ([]interface{}, error) {
	var result []interface{}
	for _, elem := range p {
		randElem, err := randJSONValue(elem)
		if err != nil {
			return nil, err
		}
		result = append(result, randElem)
	}
	return result, nil
}

func randJSONPair(key string, value interface{}) (string, interface{}, error) {
	randKey := Generate(key)
	randValue, err := randJSONValue(value)
	return randKey, randValue, err
}

func randJSONValue(e interface{}) (interface{}, error) {
	if e == nil {
		return nil, nil
	}
	var err error
	var randValue interface{}
	switch v := e.(type) {
	// ints, floats and booleans remaing as is
	case int:
		randValue = v
	case float64:
		randValue = v
	case bool:
		randValue = v
	case nil:
		randValue = nil
	// randomize strings
	case string:
		randValue = Generate(v)
	// randomize nested objects
	case map[string]interface{}:
		randValue, err = randMap(v)
		if err != nil {
			return nil, err
		}
	// randomize arrays
	case []interface{}:
		randValue, err = randJSONArray(v)
		if err != nil {
			return nil, err
		}
	default:
		return nil, errors.New("unknown type")
	}
	return randValue, nil
}
