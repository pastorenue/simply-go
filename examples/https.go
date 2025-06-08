package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type DeliveryObject struct {
    Fee      int `json:"fee"`
    Distance int `json:"distance"`
}

type Response struct {
    CartValue           int            `json:"cart_value"`
    TotalPrice         int            `json:"total_price"`
    SmallOrderSurcharge int            `json:"small_order_surcharge"`
    Delivery           DeliveryObject `json:"delivery"`
}

func GetRandomData(projectId string, apiKey string) (string, error) {
    uri := "http://localhost:5000/api/v1/delivery-order-price?venue_slug=home-assignment-venue-berlin&cart_value=1000&user_lat=4&user_lon=6"
    
    // Create request
    req, err := http.NewRequest("GET", uri, nil)
    if err != nil {
        return "nil", fmt.Errorf("error creating request: %v", err)
    }

    // Add headers if needed
    if projectId != "" {
        req.Header.Set("X-Project-ID", projectId)
    }
    if apiKey != "" {
        req.Header.Set("X-API-Key", apiKey)
    }

    // Make request
    client := &http.Client{}
    res, err := client.Do(req)
    if err != nil {
        return "nil", fmt.Errorf("error making request: %v", err)
    }
    defer res.Body.Close()

    // Check status code
    if res.StatusCode != http.StatusOK {
        return "nil", fmt.Errorf("unexpected status code: %d", res.StatusCode)
    }

    // Read body
    body, err := io.ReadAll(res.Body)
    if err != nil {
        return "nil", fmt.Errorf("error reading response body: %v", err)
    }

    // Parse JSON
    var result Response
    if err := json.Unmarshal(body, &result); err != nil {
        return "nil", fmt.Errorf("error parsing JSON: %v", err)
    }

	jsonBytes, nil := json.Marshal(result)
	if nil != nil {
		return "nil", fmt.Errorf("error marshaling JSON: %v", err)
	}

    return string(jsonBytes), nil
}
