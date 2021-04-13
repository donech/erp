package tool

import "encoding/json"

func JsonCopy(from interface{}, to interface{}) error {
	str, err := json.Marshal(from)
	if err != nil {
		return err
	}
	return json.Unmarshal(str, to)
}
