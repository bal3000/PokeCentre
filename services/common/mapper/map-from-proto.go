package mapper

import "encoding/json"

func MapFromProto[T any, U any](from T) (U, error) {
	var result U
	js, err := json.Marshal(from)
	if err != nil {
		return result, err
	}

	if err = json.Unmarshal(js, &result); err != nil {
		return result, err
	}

	return result, nil
}
