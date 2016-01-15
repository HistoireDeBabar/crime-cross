// Processes the Data Collection for each force.

package main

import (
	"encoding/json"
)

type ForceDataTransformer struct {
}

// Transforms data byte array to a []Forces array
func (fdc *ForceDataTransformer) Transform(data []byte) (forces []Force, err error) {
	if len(data) == 0 {
		return forces, nil
	}

	err = json.Unmarshal(data, &forces)
	if err != nil {
		return nil, err
	}

	return forces, nil
}
