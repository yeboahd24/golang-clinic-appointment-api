package repository

import (
    "clinic-appointment-system/pkg/model"
    "gorm.io/gorm"
)

// AppointmentRepository defines the interface for appointment data operations
type AppointmentRepository interface {
    Create(appointment model.Appointment) (model.Appointment, error)
    FindByID(appointmentID uint) (model.Appointment, error)
    Update(appointment model.Appointment) (model.Appointment, error)
    Delete(appointmentID uint) error
    FindAll(filterParams model.AppointmentFilterParams) ([]model.Appointment, error)
}

type appointmentRepository struct {
    db *gorm.DB
}

// NewAppointmentRepository creates a new instance of AppointmentRepository
func NewAppointmentRepository(db *gorm.DB) AppointmentRepository {
    return &appointmentRepository{
        db: db,
    }
}

// Create adds a new appointment to the database
func (r *appointmentRepository) Create(appointment model.Appointment) (model.Appointment, error) {
    err := r.db.Create(&appointment).Error
    return appointment, err
}

// FindByID finds an appointment by its ID
func (r *appointmentRepository) FindByID(appointmentID uint) (model.Appointment, error) {
    var appointment model.Appointment
    err := r.db.Where("id = ?", appointmentID).First(&appointment).Error
    return appointment, err
}

// Update modifies an existing appointment in the database
func (r *appointmentRepository) Update(appointment model.Appointment) (model.Appointment, error) {
    err := r.db.Save(&appointment).Error
    return appointment, err
}

// Delete removes an appointment from the database
func (r *appointmentRepository) Delete(appointmentID uint) error {
    err := r.db.Delete(&model.Appointment{}, appointmentID).Error
    return err
}


func (r *appointmentRepository) FindAll(filterParams model.AppointmentFilterParams) ([]model.Appointment, error) {
    var appointments []model.Appointment
    query := r.db

    // Apply filters based on filterParams
    if filterParams.DoctorID != "" {
        query = query.Where("doctor_id = ?", filterParams.DoctorID)
    }
    if filterParams.PatientID != "" {
        query = query.Where("patient_id = ?", filterParams.PatientID)
    }
    if filterParams.Date != "" {
        query = query.Where("date = ?", filterParams.Date)
    }

    err := query.Find(&appointments).Error
    return appointments, err
}

