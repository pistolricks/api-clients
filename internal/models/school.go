package models

import (
	"database/sql"
	"time"
)

// School represents a school from the GeoJSON data
type School struct {
	ID         int64          `json:"id"`
	ObjectID   int            `json:"objectid"`
	Name       string         `json:"name"`
	Address    sql.NullString `json:"address"`
	City       sql.NullString `json:"city"`
	State      sql.NullString `json:"state"`
	Zip        sql.NullString `json:"zip"`
	Country    sql.NullString `json:"country"`
	County     sql.NullString `json:"county"`
	CountyFIPS sql.NullString `json:"countyfips"`
	Latitude   float64        `json:"latitude"`
	Longitude  float64        `json:"longitude"`
	Level      sql.NullString `json:"level"`
	StartGrade sql.NullString `json:"st_grade"`
	EndGrade   sql.NullString `json:"end_grade"`
	Enrollment sql.NullInt64  `json:"enrollment"`
	FTTeacher  sql.NullInt64  `json:"ft_teacher"`
	Type       sql.NullInt64  `json:"type"`
	Status     sql.NullInt64  `json:"status"`
	Population sql.NullInt64  `json:"population"`
	NCESID     sql.NullString `json:"ncesid"`
	DistrictID sql.NullString `json:"districtid"`
	NAICSCode  sql.NullString `json:"naics_code"`
	NAICSDesc  sql.NullString `json:"naics_desc"`
	Website    sql.NullString `json:"website"`
	Telephone  sql.NullString `json:"telephone"`
	SourceDate sql.NullTime   `json:"sourcedate"`
	ValDate    sql.NullTime   `json:"val_date"`
	ValMethod  sql.NullString `json:"val_method"`
	Source     sql.NullString `json:"source"`
	ShelterID  sql.NullString `json:"shelter_id"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
}

// SchoolResponse is used for API responses
type SchoolResponse struct {
	ID         int64     `json:"id"`
	ObjectID   int       `json:"objectid"`
	Name       string    `json:"name"`
	Address    string    `json:"address,omitempty"`
	City       string    `json:"city,omitempty"`
	State      string    `json:"state,omitempty"`
	Zip        string    `json:"zip,omitempty"`
	Country    string    `json:"country,omitempty"`
	County     string    `json:"county,omitempty"`
	CountyFIPS string    `json:"countyfips,omitempty"`
	Latitude   float64   `json:"latitude"`
	Longitude  float64   `json:"longitude"`
	Level      string    `json:"level,omitempty"`
	StartGrade string    `json:"st_grade,omitempty"`
	EndGrade   string    `json:"end_grade,omitempty"`
	Enrollment int64     `json:"enrollment,omitempty"`
	FTTeacher  int64     `json:"ft_teacher,omitempty"`
	Type       int64     `json:"type,omitempty"`
	Status     int64     `json:"status,omitempty"`
	Population int64     `json:"population,omitempty"`
	NCESID     string    `json:"ncesid,omitempty"`
	DistrictID string    `json:"districtid,omitempty"`
	NAICSCode  string    `json:"naics_code,omitempty"`
	NAICSDesc  string    `json:"naics_desc,omitempty"`
	Website    string    `json:"website,omitempty"`
	Telephone  string    `json:"telephone,omitempty"`
	SourceDate time.Time `json:"sourcedate,omitempty"`
	ValDate    time.Time `json:"val_date,omitempty"`
	ValMethod  string    `json:"val_method,omitempty"`
	Source     string    `json:"source,omitempty"`
	ShelterID  string    `json:"shelter_id,omitempty"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// ToResponse converts a School to a SchoolResponse
func (s *School) ToResponse() SchoolResponse {
	response := SchoolResponse{
		ID:        s.ID,
		ObjectID:  s.ObjectID,
		Name:      s.Name,
		Latitude:  s.Latitude,
		Longitude: s.Longitude,
		CreatedAt: s.CreatedAt,
		UpdatedAt: s.UpdatedAt,
	}

	if s.Address.Valid {
		response.Address = s.Address.String
	}
	if s.City.Valid {
		response.City = s.City.String
	}
	if s.State.Valid {
		response.State = s.State.String
	}
	if s.Zip.Valid {
		response.Zip = s.Zip.String
	}
	if s.Country.Valid {
		response.Country = s.Country.String
	}
	if s.County.Valid {
		response.County = s.County.String
	}
	if s.CountyFIPS.Valid {
		response.CountyFIPS = s.CountyFIPS.String
	}
	if s.Level.Valid {
		response.Level = s.Level.String
	}
	if s.StartGrade.Valid {
		response.StartGrade = s.StartGrade.String
	}
	if s.EndGrade.Valid {
		response.EndGrade = s.EndGrade.String
	}
	if s.Enrollment.Valid {
		response.Enrollment = s.Enrollment.Int64
	}
	if s.FTTeacher.Valid {
		response.FTTeacher = s.FTTeacher.Int64
	}
	if s.Type.Valid {
		response.Type = s.Type.Int64
	}
	if s.Status.Valid {
		response.Status = s.Status.Int64
	}
	if s.Population.Valid {
		response.Population = s.Population.Int64
	}
	if s.NCESID.Valid {
		response.NCESID = s.NCESID.String
	}
	if s.DistrictID.Valid {
		response.DistrictID = s.DistrictID.String
	}
	if s.NAICSCode.Valid {
		response.NAICSCode = s.NAICSCode.String
	}
	if s.NAICSDesc.Valid {
		response.NAICSDesc = s.NAICSDesc.String
	}
	if s.Website.Valid {
		response.Website = s.Website.String
	}
	if s.Telephone.Valid {
		response.Telephone = s.Telephone.String
	}
	if s.SourceDate.Valid {
		response.SourceDate = s.SourceDate.Time
	}
	if s.ValDate.Valid {
		response.ValDate = s.ValDate.Time
	}
	if s.ValMethod.Valid {
		response.ValMethod = s.ValMethod.String
	}
	if s.Source.Valid {
		response.Source = s.Source.String
	}
	if s.ShelterID.Valid {
		response.ShelterID = s.ShelterID.String
	}

	return response
}