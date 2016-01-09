// Retrieves data from the police api about specific forces

package forces

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type ForcesDataCollector struct {
	Client   *http.Client
	endpoint string
}

func NewForcesDataCollector(endpoint string) (collector *ForcesDataCollector) {
	return &ForcesDataCollector{endpoint: endpoint}

}

// Retreieves all the basic data about the forces from the police api
func (fdc *ForcesDataCollector) GetForcesIdentifier() (forces []Force, err error) {
	response, err := fdc.Client.Get(fdc.endpoint)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, errors.New("Error Response from Server")
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	if len(body) == 0 {
		return forces, nil
	}

	err = json.Unmarshal(body, &forces)
	if err != nil {
		return nil, err
	}
	return forces, nil
}
