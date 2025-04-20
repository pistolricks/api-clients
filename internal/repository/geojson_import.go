package repository

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/pistolricks/api-clients/internal/models"
)

// GeoJSONFeatureCollection represents the GeoJSON structure
type GeoJSONFeatureCollection struct {
	Type     string           `json:"type"`
	Features []GeoJSONFeature `json:"features"`
}

type GeoJSONFeature struct {
	Type       string                 `json:"type"`
	Properties map[string]interface{} `json:"properties"`
	Geometry   GeoJSONGeometry        `json:"geometry"`
}

type GeoJSONGeometry struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}

// ImportFromGeoJSON imports schools from a GeoJSON file
func (r *SchoolRepository) ImportFromGeoJSON(filePath string) (int, error) {
	// Open the GeoJSON file
	file, err := os.Open(filePath)
	if err != nil {
		return 0, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	// Decode the JSON
	var featureCollection GeoJSONFeatureCollection
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&featureCollection); err != nil {
		return 0, fmt.Errorf("error decoding JSON: %w", err)
	}

	// Begin transaction
	tx, err := r.DB.Begin()
	if err != nil {
		return 0, fmt.Errorf("error beginning transaction: %w", err)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	// Prepare the insert statement
	stmt, err := tx.Prepare(`
	INSERT INTO schools (
		objectid, name, address, city, state, zip, country, county, countyfips,
		latitude, longitude, level, st_grade, end_grade, enrollment, ft_teacher,
		type, status, population, ncesid, districtid, naics_code, naics_desc,
		website, telephone, sourcedate, val_date, val_method, source, shelter_id
	) VALUES (
		$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16,
		$17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30
	) ON CONFLICT (objectid) DO NOTHING
	`)
	if err != nil {
		return 0, fmt.Errorf("error preparing statement: %w", err)
	}
	defer stmt.Close()

	// Prepare the geometry update statement
	geoStmt, err := tx.Prepare(`
	UPDATE schools SET location = ST_SetSRID(ST_MakePoint($1, $2), 4326) 
	WHERE objectid = $3
	`)
	if err != nil {
		return 0, fmt.Errorf("error preparing geometry statement: %w", err)
	}
	defer geoStmt.Close()

	// Insert each feature as a school
	count := 0
	for _, feature := range featureCollection.Features {
		// Skip if not a Point geometry
		if feature.Geometry.Type != "Point" {
			continue
		}

		// Extract properties
		school := r.ExtractSchoolFromFeature(feature)

		fmt.Println("objectid:", school.ObjectID)

		// Execute the insert
		_, err = stmt.Exec(
			school.ObjectID, school.Name, school.Address, school.City, school.State,
			school.Zip, school.Country, school.County, school.CountyFIPS, school.Latitude,
			school.Longitude, school.Level, school.StartGrade, school.EndGrade, school.Enrollment,
			school.FTTeacher, school.Type, school.Status, school.Population, school.NCESID,
			school.DistrictID, school.NAICSCode, school.NAICSDesc, school.Website, school.Telephone,
			school.SourceDate, school.ValDate, school.ValMethod, school.Source, school.ShelterID,
		)
		if err != nil {
			return 0, fmt.Errorf("error inserting school: %w", err)
		}

		// Update the geometry
		_, err = geoStmt.Exec(school.Longitude, school.Latitude, school.ObjectID)
		if err != nil {
			return 0, fmt.Errorf("error updating geometry: %w", err)
		}

		count++
	}

	// Commit the transaction
	if err = tx.Commit(); err != nil {
		return 0, fmt.Errorf("error committing transaction: %w", err)
	}

	return count, nil
}

// ExtractSchoolFromFeature converts a GeoJSON feature to a School model
func (h *SchoolRepository) ExtractSchoolFromFeature(feature GeoJSONFeature) models.School {
	props := feature.Properties
	school := models.School{}

	// Extract coordinates
	if len(feature.Geometry.Coordinates) >= 2 {
		school.Longitude = feature.Geometry.Coordinates[0]
		school.Latitude = feature.Geometry.Coordinates[1]
	}

	fmt.Println("objectid:", props["objectid"])

	// Extract properties
	if objectIDStr, ok := props["objectid"].(string); ok {
		// Convert string to int
		if objectID, err := strconv.Atoi(objectIDStr); err == nil {
			school.ObjectID = objectID
		}
	} else if objectIDFloat, ok := props["objectid"].(float64); ok {
		// Handle float64 case as well for backward compatibility
		school.ObjectID = int(objectIDFloat)
	}

	if name, ok := props["name"].(string); ok {
		school.Name = name
	}

	if address, ok := props["address"].(string); ok {
		school.Address.String = address
		school.Address.Valid = true
	}

	if city, ok := props["city"].(string); ok {
		school.City.String = city
		school.City.Valid = true
	}

	if state, ok := props["state"].(string); ok {
		school.State.String = state
		school.State.Valid = true
	}

	if zip, ok := props["zip"].(string); ok {
		school.Zip.String = zip
		school.Zip.Valid = true
	}

	if country, ok := props["country"].(string); ok {
		school.Country.String = country
		school.Country.Valid = true
	}

	if county, ok := props["county"].(string); ok {
		school.County.String = county
		school.County.Valid = true
	}

	if countyfips, ok := props["countyfips"].(string); ok {
		school.CountyFIPS.String = countyfips
		school.CountyFIPS.Valid = true
	}

	if level, ok := props["level"].(string); ok {
		school.Level.String = level
		school.Level.Valid = true
	}

	if stGrade, ok := props["st_grade"].(string); ok {
		school.StartGrade.String = stGrade
		school.StartGrade.Valid = true
	}

	if endGrade, ok := props["end_grade"].(string); ok {
		school.EndGrade.String = endGrade
		school.EndGrade.Valid = true
	}

	if enrollment, ok := props["enrollment"].(float64); ok {
		school.Enrollment.Int64 = int64(enrollment)
		school.Enrollment.Valid = true
	}

	if ftTeacher, ok := props["ft_teacher"].(float64); ok {
		school.FTTeacher.Int64 = int64(ftTeacher)
		school.FTTeacher.Valid = true
	}

	if typeVal, ok := props["type"].(float64); ok {
		school.Type.Int64 = int64(typeVal)
		school.Type.Valid = true
	}

	if status, ok := props["status"].(float64); ok {
		school.Status.Int64 = int64(status)
		school.Status.Valid = true
	}

	if population, ok := props["population"].(float64); ok {
		school.Population.Int64 = int64(population)
		school.Population.Valid = true
	}

	if ncesid, ok := props["ncesid"].(string); ok {
		school.NCESID.String = ncesid
		school.NCESID.Valid = true
	}

	if districtid, ok := props["districtid"].(string); ok {
		school.DistrictID.String = districtid
		school.DistrictID.Valid = true
	}

	if naicsCode, ok := props["naics_code"].(string); ok {
		school.NAICSCode.String = naicsCode
		school.NAICSCode.Valid = true
	}

	if naicsDesc, ok := props["naics_desc"].(string); ok {
		school.NAICSDesc.String = naicsDesc
		school.NAICSDesc.Valid = true
	}

	if website, ok := props["website"].(string); ok {
		school.Website.String = website
		school.Website.Valid = true
	}

	if telephone, ok := props["telephone"].(string); ok {
		school.Telephone.String = telephone
		school.Telephone.Valid = true
	}

	if sourcedate, ok := props["sourcedate"].(string); ok {
		if t, err := time.Parse(time.RFC3339, sourcedate); err == nil {
			school.SourceDate.Time = t
			school.SourceDate.Valid = true
		}
	}

	if valDate, ok := props["val_date"].(string); ok {
		if t, err := time.Parse(time.RFC3339, valDate); err == nil {
			school.ValDate.Time = t
			school.ValDate.Valid = true
		}
	}

	if valMethod, ok := props["val_method"].(string); ok {
		school.ValMethod.String = valMethod
		school.ValMethod.Valid = true
	}

	if source, ok := props["source"].(string); ok {
		school.Source.String = source
		school.Source.Valid = true
	}

	if shelterId, ok := props["shelter_id"].(string); ok {
		school.ShelterID.String = shelterId
		school.ShelterID.Valid = true
	}

	return school
}
