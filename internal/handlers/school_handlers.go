package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/pistolricks/api-clients/internal/models"
	"github.com/pistolricks/api-clients/internal/repository"
	"net/http"
	"strconv"
)

// SchoolHandler handles HTTP requests for schools
type SchoolHandler struct {
	Repo *repository.SchoolRepository
}

// NewSchoolHandler creates a new SchoolHandler
func NewSchoolHandler(repo *repository.SchoolRepository) *SchoolHandler {
	return &SchoolHandler{Repo: repo}
}

// GetSchools handles GET requests to list schools
func (h *SchoolHandler) GetSchools(w http.ResponseWriter, r *http.Request) {
	// Parse query parameters for pagination
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil || page < 1 {
		page = 1
	}

	pageSize, err := strconv.Atoi(r.URL.Query().Get("pageSize"))
	if err != nil || pageSize < 1 || pageSize > 200 {
		pageSize = 200
	}

	city := r.URL.Query().Get("city")
	state := r.URL.Query().Get("state")

	// Get schools from repository
	schools, err := h.Repo.List(page, pageSize, city, state)
	if err != nil {
		http.Error(w, "Error retrieving schools: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Get total count for pagination
	count, err := h.Repo.Count()
	if err != nil {
		http.Error(w, "Error counting schools: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Convert to response objects
	var response struct {
		Schools  []models.SchoolResponse `json:"schools"`
		Total    int                     `json:"total"`
		Page     int                     `json:"page"`
		PageSize int                     `json:"pageSize"`
	}
	response.Total = count
	response.Page = page
	response.PageSize = pageSize
	response.Schools = make([]models.SchoolResponse, len(schools))

	for i, school := range schools {
		response.Schools[i] = school.ToResponse()
	}

	// Write response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// GetSchool handles GET requests to retrieve a single school
func (h *SchoolHandler) GetSchool(w http.ResponseWriter, r *http.Request) {
	// Get ID from URL
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid school ID", http.StatusBadRequest)
		return
	}

	// Get school from repository
	school, err := h.Repo.GetByID(id)
	if err != nil {
		http.Error(w, "Error retrieving school: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if school == nil {
		http.Error(w, "School not found", http.StatusNotFound)
		return
	}

	// Convert to response object
	response := school.ToResponse()

	// Write response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// CreateSchool handles POST requests to create a new school
func (h *SchoolHandler) CreateSchool(w http.ResponseWriter, r *http.Request) {
	// Parse request body
	var schoolData struct {
		ObjectID   int     `json:"objectid"`
		Name       string  `json:"name"`
		Address    string  `json:"address"`
		City       string  `json:"city"`
		State      string  `json:"state"`
		Zip        string  `json:"zip"`
		Country    string  `json:"country"`
		County     string  `json:"county"`
		CountyFIPS string  `json:"countyfips"`
		Latitude   float64 `json:"latitude"`
		Longitude  float64 `json:"longitude"`
		Level      string  `json:"level"`
		StartGrade string  `json:"st_grade"`
		EndGrade   string  `json:"end_grade"`
		Enrollment int64   `json:"enrollment"`
		FTTeacher  int64   `json:"ft_teacher"`
		Type       int64   `json:"type"`
		Status     int64   `json:"status"`
		Population int64   `json:"population"`
		NCESID     string  `json:"ncesid"`
		DistrictID string  `json:"districtid"`
		NAICSCode  string  `json:"naics_code"`
		NAICSDesc  string  `json:"naics_desc"`
		Website    string  `json:"website"`
		Telephone  string  `json:"telephone"`
		SourceDate string  `json:"sourcedate"`
		ValDate    string  `json:"val_date"`
		ValMethod  string  `json:"val_method"`
		Source     string  `json:"source"`
		ShelterID  string  `json:"shelter_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&schoolData); err != nil {
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Validate required fields
	if schoolData.Name == "" {
		http.Error(w, "Name is required", http.StatusBadRequest)
		return
	}

	// Create school object
	school := models.School{
		ObjectID:  schoolData.ObjectID,
		Name:      schoolData.Name,
		Latitude:  schoolData.Latitude,
		Longitude: schoolData.Longitude,
	}

	// Set optional fields
	if schoolData.Address != "" {
		school.Address.String = schoolData.Address
		school.Address.Valid = true
	}
	if schoolData.City != "" {
		school.City.String = schoolData.City
		school.City.Valid = true
	}
	if schoolData.State != "" {
		school.State.String = schoolData.State
		school.State.Valid = true
	}
	if schoolData.Zip != "" {
		school.Zip.String = schoolData.Zip
		school.Zip.Valid = true
	}
	if schoolData.Country != "" {
		school.Country.String = schoolData.Country
		school.Country.Valid = true
	}
	if schoolData.County != "" {
		school.County.String = schoolData.County
		school.County.Valid = true
	}
	if schoolData.CountyFIPS != "" {
		school.CountyFIPS.String = schoolData.CountyFIPS
		school.CountyFIPS.Valid = true
	}
	if schoolData.Level != "" {
		school.Level.String = schoolData.Level
		school.Level.Valid = true
	}
	if schoolData.StartGrade != "" {
		school.StartGrade.String = schoolData.StartGrade
		school.StartGrade.Valid = true
	}
	if schoolData.EndGrade != "" {
		school.EndGrade.String = schoolData.EndGrade
		school.EndGrade.Valid = true
	}
	if schoolData.Enrollment != 0 {
		school.Enrollment.Int64 = schoolData.Enrollment
		school.Enrollment.Valid = true
	}
	if schoolData.FTTeacher != 0 {
		school.FTTeacher.Int64 = schoolData.FTTeacher
		school.FTTeacher.Valid = true
	}
	if schoolData.Type != 0 {
		school.Type.Int64 = schoolData.Type
		school.Type.Valid = true
	}
	if schoolData.Status != 0 {
		school.Status.Int64 = schoolData.Status
		school.Status.Valid = true
	}
	if schoolData.Population != 0 {
		school.Population.Int64 = schoolData.Population
		school.Population.Valid = true
	}
	if schoolData.NCESID != "" {
		school.NCESID.String = schoolData.NCESID
		school.NCESID.Valid = true
	}
	if schoolData.DistrictID != "" {
		school.DistrictID.String = schoolData.DistrictID
		school.DistrictID.Valid = true
	}
	if schoolData.NAICSCode != "" {
		school.NAICSCode.String = schoolData.NAICSCode
		school.NAICSCode.Valid = true
	}
	if schoolData.NAICSDesc != "" {
		school.NAICSDesc.String = schoolData.NAICSDesc
		school.NAICSDesc.Valid = true
	}
	if schoolData.Website != "" {
		school.Website.String = schoolData.Website
		school.Website.Valid = true
	}
	if schoolData.Telephone != "" {
		school.Telephone.String = schoolData.Telephone
		school.Telephone.Valid = true
	}
	if schoolData.SourceDate != "" {
		// Parse date string
	}
	if schoolData.ValDate != "" {
		// Parse date string
	}
	if schoolData.ValMethod != "" {
		school.ValMethod.String = schoolData.ValMethod
		school.ValMethod.Valid = true
	}
	if schoolData.Source != "" {
		school.Source.String = schoolData.Source
		school.Source.Valid = true
	}
	if schoolData.ShelterID != "" {
		school.ShelterID.String = schoolData.ShelterID
		school.ShelterID.Valid = true
	}

	// Save to repository
	if err := h.Repo.Create(&school); err != nil {
		http.Error(w, "Error creating school: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Return the created school
	response := school.ToResponse()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

// UpdateSchool handles PUT requests to update a school
func (h *SchoolHandler) UpdateSchool(w http.ResponseWriter, r *http.Request) {
	// Get ID from URL
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid school ID", http.StatusBadRequest)
		return
	}

	// Get existing school
	school, err := h.Repo.GetByID(id)
	if err != nil {
		http.Error(w, "Error retrieving school: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if school == nil {
		http.Error(w, "School not found", http.StatusNotFound)
		return
	}

	// Parse request body
	var schoolData struct {
		ObjectID   int     `json:"objectid"`
		Name       string  `json:"name"`
		Address    string  `json:"address"`
		City       string  `json:"city"`
		State      string  `json:"state"`
		Zip        string  `json:"zip"`
		Country    string  `json:"country"`
		County     string  `json:"county"`
		CountyFIPS string  `json:"countyfips"`
		Latitude   float64 `json:"latitude"`
		Longitude  float64 `json:"longitude"`
		Level      string  `json:"level"`
		StartGrade string  `json:"st_grade"`
		EndGrade   string  `json:"end_grade"`
		Enrollment int64   `json:"enrollment"`
		FTTeacher  int64   `json:"ft_teacher"`
		Type       int64   `json:"type"`
		Status     int64   `json:"status"`
		Population int64   `json:"population"`
		NCESID     string  `json:"ncesid"`
		DistrictID string  `json:"districtid"`
		NAICSCode  string  `json:"naics_code"`
		NAICSDesc  string  `json:"naics_desc"`
		Website    string  `json:"website"`
		Telephone  string  `json:"telephone"`
		SourceDate string  `json:"sourcedate"`
		ValDate    string  `json:"val_date"`
		ValMethod  string  `json:"val_method"`
		Source     string  `json:"source"`
		ShelterID  string  `json:"shelter_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&schoolData); err != nil {
		http.Error(w, "Invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Update fields
	if schoolData.ObjectID != 0 {
		school.ObjectID = schoolData.ObjectID
	}
	if schoolData.Name != "" {
		school.Name = schoolData.Name
	}
	if schoolData.Latitude != 0 {
		school.Latitude = schoolData.Latitude
	}
	if schoolData.Longitude != 0 {
		school.Longitude = schoolData.Longitude
	}

	// Update optional fields
	if schoolData.Address != "" {
		school.Address.String = schoolData.Address
		school.Address.Valid = true
	}
	if schoolData.City != "" {
		school.City.String = schoolData.City
		school.City.Valid = true
	}
	if schoolData.State != "" {
		school.State.String = schoolData.State
		school.State.Valid = true
	}
	if schoolData.Zip != "" {
		school.Zip.String = schoolData.Zip
		school.Zip.Valid = true
	}
	if schoolData.Country != "" {
		school.Country.String = schoolData.Country
		school.Country.Valid = true
	}
	if schoolData.County != "" {
		school.County.String = schoolData.County
		school.County.Valid = true
	}
	if schoolData.CountyFIPS != "" {
		school.CountyFIPS.String = schoolData.CountyFIPS
		school.CountyFIPS.Valid = true
	}
	if schoolData.Level != "" {
		school.Level.String = schoolData.Level
		school.Level.Valid = true
	}
	if schoolData.StartGrade != "" {
		school.StartGrade.String = schoolData.StartGrade
		school.StartGrade.Valid = true
	}
	if schoolData.EndGrade != "" {
		school.EndGrade.String = schoolData.EndGrade
		school.EndGrade.Valid = true
	}
	if schoolData.Enrollment != 0 {
		school.Enrollment.Int64 = schoolData.Enrollment
		school.Enrollment.Valid = true
	}
	if schoolData.FTTeacher != 0 {
		school.FTTeacher.Int64 = schoolData.FTTeacher
		school.FTTeacher.Valid = true
	}
	if schoolData.Type != 0 {
		school.Type.Int64 = schoolData.Type
		school.Type.Valid = true
	}
	if schoolData.Status != 0 {
		school.Status.Int64 = schoolData.Status
		school.Status.Valid = true
	}
	if schoolData.Population != 0 {
		school.Population.Int64 = schoolData.Population
		school.Population.Valid = true
	}
	if schoolData.NCESID != "" {
		school.NCESID.String = schoolData.NCESID
		school.NCESID.Valid = true
	}
	if schoolData.DistrictID != "" {
		school.DistrictID.String = schoolData.DistrictID
		school.DistrictID.Valid = true
	}
	if schoolData.NAICSCode != "" {
		school.NAICSCode.String = schoolData.NAICSCode
		school.NAICSCode.Valid = true
	}
	if schoolData.NAICSDesc != "" {
		school.NAICSDesc.String = schoolData.NAICSDesc
		school.NAICSDesc.Valid = true
	}
	if schoolData.Website != "" {
		school.Website.String = schoolData.Website
		school.Website.Valid = true
	}
	if schoolData.Telephone != "" {
		school.Telephone.String = schoolData.Telephone
		school.Telephone.Valid = true
	}
	if schoolData.SourceDate != "" {
		// Parse date string
	}
	if schoolData.ValDate != "" {
		// Parse date string
	}
	if schoolData.ValMethod != "" {
		school.ValMethod.String = schoolData.ValMethod
		school.ValMethod.Valid = true
	}
	if schoolData.Source != "" {
		school.Source.String = schoolData.Source
		school.Source.Valid = true
	}
	if schoolData.ShelterID != "" {
		school.ShelterID.String = schoolData.ShelterID
		school.ShelterID.Valid = true
	}

	// Save to repository
	if err := h.Repo.Update(school); err != nil {
		http.Error(w, "Error updating school: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Return the updated school
	response := school.ToResponse()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// DeleteSchool handles DELETE requests to delete a school
func (h *SchoolHandler) DeleteSchool(w http.ResponseWriter, r *http.Request) {
	// Get ID from URL
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid school ID", http.StatusBadRequest)
		return
	}

	// Check if school exists
	school, err := h.Repo.GetByID(id)
	if err != nil {
		http.Error(w, "Error retrieving school: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if school == nil {
		http.Error(w, "School not found", http.StatusNotFound)
		return
	}

	// Delete from repository
	if err := h.Repo.Delete(id); err != nil {
		http.Error(w, "Error deleting school: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Return success
	w.WriteHeader(http.StatusNoContent)
}

// ImportGeoJSON handles POST requests to import schools from a GeoJSON file
func (h *SchoolHandler) ImportGeoJSON(w http.ResponseWriter, r *http.Request) {
	// Import from GeoJSON file
	// Use absolute path to the file in the project root
	count, err := h.Repo.ImportFromGeoJSON("./us-public-schools-part1.geojson")
	if err != nil {
		http.Error(w, "Error importing schools: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Return success
	response := struct {
		Message string `json:"message"`
		Count   int    `json:"count"`
	}{
		Message: "Schools imported successfully",
		Count:   count,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

type GeoJSONFeatureCollection struct {
	Type     string                      `json:"type"`
	Features []repository.GeoJSONFeature `json:"features"`
}
type GeoJSONFeature struct {
	Type       string          `json:"type"`
	Properties models.School   `json:"properties"`
	Geometry   GeoJSONGeometry `json:"geometry"`
}

type GeoJSONGeometry struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}
