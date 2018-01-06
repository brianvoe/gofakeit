package gofakeit

import (
	"encoding/json"
	"errors"
)

// JSON produces a random JSON using a JSON template
func JSON(template string) (map[string]interface{}, error) {
	var templateJSON map[string]interface{}
	t := []byte(template)
	err := json.Unmarshal(t, &templateJSON)
	if err != nil {
		return nil, err
	}
	return randJSON(templateJSON)
}

func randJSON(m map[string]interface{}) (map[string]interface{}, error) {
	result := make(map[string]interface{})
	for k, v := range m {
		var val interface{}
		switch vv := v.(type) {
		case string:
			val = Generate(vv)
		case map[string]interface{}:
			var err error
			val, err = randJSON(vv)
			if err != nil {
				return nil, err
			}
		default:
			return nil, errors.New("unknown type")
		}
		k = Generate(k)
		result[k] = val
	}
	return result, nil
}
