package repository

import (
    "clinic-appointment-system/pkg/model"
    "gorm.io/gorm"
)

// ClinicLocationRepository defines the interface for clinic location data operations
type ClinicLocationRepository interface {
    Create(location model.ClinicLocation) (model.ClinicLocation, error)
    FindByID(locationID uint) (model.ClinicLocation, error)
    Update(location model.ClinicLocation) (model.ClinicLocation, error)
    Delete(locationID uint) error
    FindAll() ([]model.ClinicLocation, error)
}

type clinicLocationRepository struct {
    db *gorm.DB
}

// NewClinicLocationRepository creates a new instance of ClinicLocationRepository
func NewClinicLocationRepository(db *gorm.DB) ClinicLocationRepository {
    return &clinicLocationRepository{
        db: db,
    }
}

// Create adds a new clinic location to the database
func (r *clinicLocationRepository) Create(location model.ClinicLocation) (model.ClinicLocation, error) {
    err := r.db.Create(&location).Error
    return location, err
}

// FindByID finds a clinic location by its ID
func (r *clinicLocationRepository) FindByID(locationID uint) (model.ClinicLocation, error) {
    var location model.ClinicLocation
    err := r.db.Where("id = ?", locationID).First(&location).Error
    return location, err
}

// Update modifies an existing clinic location in the database
func (r *clinicLocationRepository) Update(location model.ClinicLocation) (model.ClinicLocation, error) {
    err := r.db.Save(&location).Error
    return location, err
}

// Delete removes a clinic location from the database
func (r *clinicLocationRepository) Delete(locationID uint) error {
    err := r.db.Delete(&model.ClinicLocation{}, locationID).Error
    return err
}


func (r *clinicLocationRepository) FindAll() ([]model.ClinicLocation, error) {
    var locations []model.ClinicLocation
    err := r.db.Find(&locations).Error
    return locations, err
}

