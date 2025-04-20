package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// GeoJSON structure
type GeoJSONCollection struct {
	Type     string        `json:"type"`
	Features []GeoJSONFeature `json:"features"`
}

type GeoJSONFeature struct {
	Type       string                 `json:"type"`
	Properties map[string]interface{} `json:"properties"`
	Geometry   GeoJSONGeometry        `json:"geometry"`
}

type GeoJSONGeometry struct {
	Type        string          `json:"type"`
	Coordinates json.RawMessage `json:"coordinates"`
}

func main() {
	// Open the GeoJSON file
	file, err := os.Open("../../us-public-schools-part1.geojson")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	// Decode the JSON
	var collection GeoJSONCollection
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&collection); err != nil {
		fmt.Printf("Error decoding JSON: %v\n", err)
		return
	}

	// Print the type and number of features
	fmt.Printf("Type: %s\n", collection.Type)
	fmt.Printf("Number of features: %d\n", len(collection.Features))

	// Print the objectid of the first 5 features
	fmt.Println("\nObjectIDs of the first 5 features:")
	for i := 0; i < 5 && i < len(collection.Features); i++ {
		objectid := collection.Features[i].Properties["objectid"]
		fmt.Printf("Feature %d - objectid: %v (type: %T)\n", i, objectid, objectid)
	}
}
