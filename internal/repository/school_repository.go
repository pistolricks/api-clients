package repository

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/pistolricks/api-clients/internal/models"
)

// SchoolRepository handles database operations for schools
type SchoolRepository struct {
	DB *sql.DB
}

// NewSchoolRepository creates a new SchoolRepository
func NewSchoolRepository(db *sql.DB) *SchoolRepository {
	return &SchoolRepository{DB: db}
}

// Create inserts a new school into the database
func (r *SchoolRepository) Create(school *models.School) error {
	query := `
	INSERT INTO schools (
		objectid, name, address, city, state, zip, country, county, countyfips,
		latitude, longitude, level, st_grade, end_grade, enrollment, ft_teacher,
		type, status, population, ncesid, districtid, naics_code, naics_desc,
		website, telephone, sourcedate, val_date, val_method, source, shelter_id
	) VALUES (
		$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16,
		$17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30
	) RETURNING id, created_at, updated_at;
	`

	// Execute the query
	err := r.DB.QueryRow(
		query,
		school.ObjectID, school.Name, school.Address, school.City, school.State,
		school.Zip, school.Country, school.County, school.CountyFIPS, school.Latitude,
		school.Longitude, school.Level, school.StartGrade, school.EndGrade, school.Enrollment,
		school.FTTeacher, school.Type, school.Status, school.Population, school.NCESID,
		school.DistrictID, school.NAICSCode, school.NAICSDesc, school.Website, school.Telephone,
		school.SourceDate, school.ValDate, school.ValMethod, school.Source, school.ShelterID,
	).Scan(&school.ID, &school.CreatedAt, &school.UpdatedAt)

	if err != nil {
		return fmt.Errorf("failed to create school: %w", err)
	}

	// Update the geometry column
	_, err = r.DB.Exec(
		"UPDATE schools SET location = ST_SetSRID(ST_MakePoint($1, $2), 4326) WHERE id = $3",
		school.Longitude, school.Latitude, school.ID,
	)
	if err != nil {
		return fmt.Errorf("failed to update school location: %w", err)
	}

	return nil
}

// GetByID retrieves a school by its ID
func (r *SchoolRepository) GetByID(id int64) (*models.School, error) {
	query := `
	SELECT id, objectid, name, address, city, state, zip, country, county, countyfips,
		latitude, longitude, level, st_grade, end_grade, enrollment, ft_teacher,
		type, status, population, ncesid, districtid, naics_code, naics_desc,
		website, telephone, sourcedate, val_date, val_method, source, shelter_id,
		created_at, updated_at
	FROM schools
	WHERE id = $1
	`

	var school models.School
	err := r.DB.QueryRow(query, id).Scan(
		&school.ID, &school.ObjectID, &school.Name, &school.Address, &school.City,
		&school.State, &school.Zip, &school.Country, &school.County, &school.CountyFIPS,
		&school.Latitude, &school.Longitude, &school.Level, &school.StartGrade, &school.EndGrade,
		&school.Enrollment, &school.FTTeacher, &school.Type, &school.Status, &school.Population,
		&school.NCESID, &school.DistrictID, &school.NAICSCode, &school.NAICSDesc, &school.Website,
		&school.Telephone, &school.SourceDate, &school.ValDate, &school.ValMethod, &school.Source,
		&school.ShelterID, &school.CreatedAt, &school.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get school: %w", err)
	}

	return &school, nil
}

// GetByObjectID retrieves a school by its ObjectID
func (r *SchoolRepository) GetByObjectID(objectID int) (*models.School, error) {
	query := `
	SELECT id, objectid, name, address, city, state, zip, country, county, countyfips,
		latitude, longitude, level, st_grade, end_grade, enrollment, ft_teacher,
		type, status, population, ncesid, districtid, naics_code, naics_desc,
		website, telephone, sourcedate, val_date, val_method, source, shelter_id,
		created_at, updated_at
	FROM schools
	WHERE objectid = $1
	`

	var school models.School
	err := r.DB.QueryRow(query, objectID).Scan(
		&school.ID, &school.ObjectID, &school.Name, &school.Address, &school.City,
		&school.State, &school.Zip, &school.Country, &school.County, &school.CountyFIPS,
		&school.Latitude, &school.Longitude, &school.Level, &school.StartGrade, &school.EndGrade,
		&school.Enrollment, &school.FTTeacher, &school.Type, &school.Status, &school.Population,
		&school.NCESID, &school.DistrictID, &school.NAICSCode, &school.NAICSDesc, &school.Website,
		&school.Telephone, &school.SourceDate, &school.ValDate, &school.ValMethod, &school.Source,
		&school.ShelterID, &school.CreatedAt, &school.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to get school by objectid: %w", err)
	}

	return &school, nil
}

// List retrieves all schools with pagination
func (r *SchoolRepository) List(page, pageSize int) ([]*models.School, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}

	offset := (page - 1) * pageSize

	query := `
	SELECT id, objectid, name, address, city, state, zip, country, county, countyfips,
		latitude, longitude, level, st_grade, end_grade, enrollment, ft_teacher,
		type, status, population, ncesid, districtid, naics_code, naics_desc,
		website, telephone, sourcedate, val_date, val_method, source, shelter_id,
		created_at, updated_at
	FROM schools
	ORDER BY id
	LIMIT $1 OFFSET $2
	`

	rows, err := r.DB.Query(query, pageSize, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to list schools: %w", err)
	}
	defer rows.Close()

	var schools []*models.School
	for rows.Next() {
		var school models.School
		err := rows.Scan(
			&school.ID, &school.ObjectID, &school.Name, &school.Address, &school.City,
			&school.State, &school.Zip, &school.Country, &school.County, &school.CountyFIPS,
			&school.Latitude, &school.Longitude, &school.Level, &school.StartGrade, &school.EndGrade,
			&school.Enrollment, &school.FTTeacher, &school.Type, &school.Status, &school.Population,
			&school.NCESID, &school.DistrictID, &school.NAICSCode, &school.NAICSDesc, &school.Website,
			&school.Telephone, &school.SourceDate, &school.ValDate, &school.ValMethod, &school.Source,
			&school.ShelterID, &school.CreatedAt, &school.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan school: %w", err)
		}
		schools = append(schools, &school)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating schools: %w", err)
	}

	return schools, nil
}

// Update updates a school in the database
func (r *SchoolRepository) Update(school *models.School) error {
	query := `
	UPDATE schools
	SET objectid = $1, name = $2, address = $3, city = $4, state = $5,
		zip = $6, country = $7, county = $8, countyfips = $9, latitude = $10,
		longitude = $11, level = $12, st_grade = $13, end_grade = $14, enrollment = $15,
		ft_teacher = $16, type = $17, status = $18, population = $19, ncesid = $20,
		districtid = $21, naics_code = $22, naics_desc = $23, website = $24, telephone = $25,
		sourcedate = $26, val_date = $27, val_method = $28, source = $29, shelter_id = $30,
		updated_at = $31
	WHERE id = $32
	RETURNING updated_at
	`

	// Execute the query
	err := r.DB.QueryRow(
		query,
		school.ObjectID, school.Name, school.Address, school.City, school.State,
		school.Zip, school.Country, school.County, school.CountyFIPS, school.Latitude,
		school.Longitude, school.Level, school.StartGrade, school.EndGrade, school.Enrollment,
		school.FTTeacher, school.Type, school.Status, school.Population, school.NCESID,
		school.DistrictID, school.NAICSCode, school.NAICSDesc, school.Website, school.Telephone,
		school.SourceDate, school.ValDate, school.ValMethod, school.Source, school.ShelterID,
		time.Now(), school.ID,
	).Scan(&school.UpdatedAt)

	if err != nil {
		return fmt.Errorf("failed to update school: %w", err)
	}

	// Update the geometry column
	_, err = r.DB.Exec(
		"UPDATE schools SET location = ST_SetSRID(ST_MakePoint($1, $2), 4326) WHERE id = $3",
		school.Longitude, school.Latitude, school.ID,
	)
	if err != nil {
		return fmt.Errorf("failed to update school location: %w", err)
	}

	return nil
}

// Delete deletes a school from the database
func (r *SchoolRepository) Delete(id int64) error {
	_, err := r.DB.Exec("DELETE FROM schools WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("failed to delete school: %w", err)
	}
	return nil
}

// Count returns the total number of schools
func (r *SchoolRepository) Count() (int, error) {
	var count int
	err := r.DB.QueryRow("SELECT COUNT(*) FROM schools").Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("failed to count schools: %w", err)
	}
	return count, nil
}
