package repository

import (
    "clinic-appointment-system/pkg/model"
    "gorm.io/gorm"
)

// PatientRepository defines the interface for patient data operations
type PatientRepository interface {
    Create(patient model.Patient) (model.Patient, error)
    FindByID(patientID uint) (model.Patient, error)
    Update(patient model.Patient) (model.Patient, error)
    Delete(patientID uint) error
    FindAll() ([]model.Patient, error)
}

type patientRepository struct {
    db *gorm.DB
}

// NewPatientRepository creates a new instance of PatientRepository
func NewPatientRepository(db *gorm.DB) PatientRepository {
    return &patientRepository{
        db: db,
    }
}

// Create adds a new patient to the database
func (r *patientRepository) Create(patient model.Patient) (model.Patient, error) {
    err := r.db.Create(&patient).Error
    return patient, err
}

// FindByID finds a patient by their ID
func (r *patientRepository) FindByID(patientID uint) (model.Patient, error) {
    var patient model.Patient
    err := r.db.Where("id = ?", patientID).First(&patient).Error
    return patient, err
}

// Update modifies an existing patient in the database
func (r *patientRepository) Update(patient model.Patient) (model.Patient, error) {
    err := r.db.Save(&patient).Error
    return patient, err
}

// Delete removes a patient from the database
func (r *patientRepository) Delete(patientID uint) error {
    err := r.db.Delete(&model.Patient{}, patientID).Error
    return err
}


func (r *patientRepository) FindAll() ([]model.Patient, error) {
    var patients []model.Patient
    err := r.db.Find(&patients).Error
    return patients, err
}

