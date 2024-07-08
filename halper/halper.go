package halper

import (
	"encoding/json"
	"fmt"
)

func DataParser1[T1 any, T2 any](src T1, dst T2) error{
	bytData, err := json.Marshal(src)
	if err != nil {
		return err
	}
	json.Unmarshal(bytData, dst)
	return nil
}

func DataParser2(src interface{}, dst interface{}) error {
	bytData, err := json.Marshal(src)
	if err != nil {
		return fmt.Errorf("error marshalling JSON: %v", err)
	}

	if err := json.Unmarshal(bytData, dst); err != nil {
		return fmt.Errorf("error unmarshalling JSON: %v", err)
	}

	return nil
}

func DataParser3(src interface{}, dst interface{}) error {
    bytData, err := json.Marshal(src)
    if err != nil {
        return err
    }
    return json.Unmarshal(bytData, dst)
}
