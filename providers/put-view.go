package providers

import "encoding/json"

func PutView(entity interface{}, view interface{}) error {
	jason, err := json.Marshal(entity)

	if err != nil {
		return err
	}

	err = json.Unmarshal(jason, view)

	if err != nil {
		return err
	}

	return nil
}
