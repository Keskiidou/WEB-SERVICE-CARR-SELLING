package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	// The correct API endpoint format with a carID in the URL path
	carPriceAPI = "http://localhost:3000/car/price/"
	carModelAPI = "http://localhost:3000/car/model/"
)

func FetchCarPrice(carID int32) (float64, error) {
	// Construct the correct URL by appending carID directly to the path
	url := fmt.Sprintf("%s%d", carPriceAPI, carID)

	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	// Initialize a map to hold the response data
	var result map[string]float64

	// Unmarshal the JSON response into the result map
	if err := json.Unmarshal(body, &result); err != nil {
		return 0, err
	}

	// Retrieve the price from the result map
	price, ok := result["price"]
	if !ok {
		return 0, fmt.Errorf("failed to fetch car price")
	}

	// Return the fetched price
	return price, nil
}
func FetchCarmodel(carID int32) (string, error) {
	// Construct the correct URL
	url := fmt.Sprintf("%s%d", carModelAPI, carID)

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Unmarshal the JSON response into a map
	var result map[string]string
	if err := json.Unmarshal(body, &result); err != nil {
		return "", err
	}

	carInfo := fmt.Sprintf("%s %s", result["make"], result["modelName"])

	return carInfo, nil
}
