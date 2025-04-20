package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// GeoJSON structures
type FeatureCollection struct {
	Type     string    `json:"type"`
	Features []Feature `json:"features"`
}

type Feature struct {
	Type       string     `json:"type"`
	Geometry   Geometry   `json:"geometry"`
	Properties Properties `json:"properties"`
}

type Geometry struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}

type Properties map[string]interface{}

func main() {
	// Path to the large GeoJSON file
	inputFile := "us-public-schools-part2.geojson"

	// Read the file
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	// Parse the GeoJSON
	var featureCollection FeatureCollection
	err = json.Unmarshal(data, &featureCollection)
	if err != nil {
		fmt.Printf("Error parsing GeoJSON: %v\n", err)
		return
	}

	// Get the total number of features
	totalFeatures := len(featureCollection.Features)
	fmt.Printf("Total features: %d\n", totalFeatures)

	// Calculate features per file (approximately)
	featuresPerFile := totalFeatures / 4
	fmt.Printf("Features per file: ~%d\n", featuresPerFile)

	// Split and save the features into 4 files
	for i := 0; i < 4; i++ {
		// Calculate start and end indices for this part
		start := i * featuresPerFile
		end := start + featuresPerFile

		// Adjust the end index for the last part to include any remaining features
		if i == 3 {
			end = totalFeatures
		}

		// Create a new feature collection for this part
		part := FeatureCollection{
			Type:     "FeatureCollection",
			Features: featureCollection.Features[start:end],
		}

		// Marshal the part to JSON
		partData, err := json.MarshalIndent(part, "", "  ")
		if err != nil {
			fmt.Printf("Error marshaling part %d: %v\n", i+1, err)
			continue
		}

		// Create the output file name
		outputFile := fmt.Sprintf("us-public-schools-part%d.geojson", i+1)

		// Write the part to a new file
		err = ioutil.WriteFile(outputFile, partData, 0644)
		if err != nil {
			fmt.Printf("Error writing part %d: %v\n", i+1, err)
			continue
		}

		fmt.Printf("Created %s with %d features\n", outputFile, len(part.Features))
	}

	fmt.Println("Split completed successfully!")
}
