package json

import (
	"encoding/json"
	"fmt"
	"io"
)

func Encoder(w io.Writer) *json.Encoder {
	return json.NewEncoder(w)
}

func Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

const maxDepth = 10000

func Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

func ConvertInterfaceNumber(v *interface{}, depth int) error {
	var err error
	switch v2 := (*v).(type) {
	case json.Number:
		*v, err = ConvertNumber(v2)
	case map[string]interface{}:
		err = ConvertMapNumber(v2, depth+1)
	case []interface{}:
		err = ConvertSliceNumbers(v2, depth+1)
	}
	return err
}

func ConvertMapNumber(m map[string]interface{}, depth int) error {
	if depth > maxDepth {
		return fmt.Errorf("max depth exceeded")
	}
	var err error
	for k, v := range m {
		switch v := v.(type) {
		case json.Number:
			m[k], err = ConvertNumber(v)
		case map[string]interface{}:
			err = ConvertMapNumber(v, depth+1)
		case []interface{}:
			err = ConvertSliceNumbers(v, depth+1)
		}
	}
	return err
}

func ConvertSliceNumbers(s []interface{}, depth int) error {
	if depth > maxDepth {
		return fmt.Errorf("max depth exceeded")
	}
	var err error
	for i, v := range s {
		switch v := v.(type) {
		case json.Number:
			s[i], err = ConvertNumber(v)
		case map[string]interface{}:
			err = ConvertMapNumber(v, depth+1)
		case []interface{}:
			err = ConvertSliceNumbers(v, depth+1)
		}
	}
	return err
}

func ConvertNumber(n json.Number) (interface{}, error) {
	if i, err := n.Int64(); err == nil {
		return i, nil
	}
	return n.Float64()
}
