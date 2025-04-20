package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// GeoJSON structure
type FeatureCollection struct {
	Type     string    `json:"type"`
	Features []Feature `json:"features"`
}

type Feature struct {
	Type       string                 `json:"type"`
	Properties map[string]interface{} `json:"properties"`
	Geometry   Geometry               `json:"geometry"`
}

type Geometry struct {
	Type        string          `json:"type"`
	Coordinates json.RawMessage `json:"coordinates"`
}

func main() {
	// Open the GeoJSON file
	file, err := os.Open("../../us-public-schools.geojson")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	// Decode the JSON
	var featureCollection FeatureCollection
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&featureCollection); err != nil {
		fmt.Printf("Error decoding JSON: %v\n", err)
		return
	}

	// Print the type and number of features
	fmt.Printf("Type: %s\n", featureCollection.Type)
	fmt.Printf("Number of features: %d\n", len(featureCollection.Features))

	// Print the properties of the first feature to understand the structure
	if len(featureCollection.Features) > 0 {
		fmt.Println("\nProperties of the first feature:")
		for key, value := range featureCollection.Features[0].Properties {
			fmt.Printf("%s: %v\n", key, value)
		}

		// Print the geometry type
		fmt.Printf("\nGeometry Type: %s\n", featureCollection.Features[0].Geometry.Type)
	}
}
