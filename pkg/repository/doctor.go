package repository

import (
    "clinic-appointment-system/pkg/model"
    "gorm.io/gorm"
)

// DoctorRepository defines the interface for doctor data operations
type DoctorRepository interface {
    Create(doctor model.Doctor) (model.Doctor, error)
    FindByID(doctorID uint) (model.Doctor, error)
    Update(doctor model.Doctor) (model.Doctor, error)
    Delete(doctorID uint) error
    FindAll() ([]model.Doctor, error)
}

type doctorRepository struct {
    db *gorm.DB
}

// NewDoctorRepository creates a new instance of DoctorRepository
func NewDoctorRepository(db *gorm.DB) DoctorRepository {
    return &doctorRepository{
        db: db,
    }
}

// Create adds a new doctor to the database
func (r *doctorRepository) Create(doctor model.Doctor) (model.Doctor, error) {
    err := r.db.Create(&doctor).Error
    return doctor, err
}

// FindByID finds a doctor by their ID
func (r *doctorRepository) FindByID(doctorID uint) (model.Doctor, error) {
    var doctor model.Doctor
    err := r.db.Where("id = ?", doctorID).First(&doctor).Error
    return doctor, err
}

// Update modifies an existing doctor in the database
func (r *doctorRepository) Update(doctor model.Doctor) (model.Doctor, error) {
    err := r.db.Save(&doctor).Error
    return doctor, err
}

// Delete removes a doctor from the database
func (r *doctorRepository) Delete(doctorID uint) error {
    err := r.db.Delete(&model.Doctor{}, doctorID).Error
    return err
}


func (r *doctorRepository) FindAll() ([]model.Doctor, error) {
    var doctors []model.Doctor
    err := r.db.Find(&doctors).Error
    return doctors, err
}

