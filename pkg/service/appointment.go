package service

import (
    "clinic-appointment-system/pkg/model"
    "clinic-appointment-system/pkg/repository"
    "errors"
    "time"
)

// AppointmentService defines the interface for appointment service operations
type AppointmentService interface {
    CreateAppointment(appointment model.Appointment) (model.Appointment, error)
    GetAppointmentByID(appointmentID uint) (model.Appointment, error)
    UpdateAppointment(appointment model.Appointment) (model.Appointment, error)
    DeleteAppointment(appointmentID uint) error
    GetAllAppointments(filterParams model.AppointmentFilterParams) ([]model.Appointment, error)
}

type appointmentService struct {
    appointmentRepository repository.AppointmentRepository
}

// NewAppointmentService creates a new instance of AppointmentService
func NewAppointmentService(repo repository.AppointmentRepository) AppointmentService {
    return &appointmentService{
        appointmentRepository: repo,
    }
}

func (s *appointmentService) CreateAppointment(appointment model.Appointment) (model.Appointment, error) {
    if err := validateAppointment(appointment); err != nil {
        return model.Appointment{}, err
    }

    createdAppointment, err := s.appointmentRepository.Create(appointment)
    if err != nil {
        return model.Appointment{}, errors.New("failed to create appointment: " + err.Error())
    }

    return createdAppointment, nil
}

func (s *appointmentService) GetAppointmentByID(appointmentID uint) (model.Appointment, error) {
    appointment, err := s.appointmentRepository.FindByID(appointmentID)
    if err != nil {
        return model.Appointment{}, errors.New("appointment not found")
    }

    return appointment, nil
}

func (s *appointmentService) UpdateAppointment(appointment model.Appointment) (model.Appointment, error) {
    if err := validateAppointment(appointment); err != nil {
        return model.Appointment{}, err
    }

    updatedAppointment, err := s.appointmentRepository.Update(appointment)
    if err != nil {
        return model.Appointment{}, errors.New("failed to update appointment: " + err.Error())
    }

    return updatedAppointment, nil
}

func (s *appointmentService) DeleteAppointment(appointmentID uint) error {
    err := s.appointmentRepository.Delete(appointmentID)
    if err != nil {
        return errors.New("failed to delete appointment")
    }
    return nil
}

// validateAppointment checks if the appointment's information meets certain criteria
func validateAppointment(appointment model.Appointment) error {
    if appointment.Date.Before(time.Now()) {
        return errors.New("appointment date cannot be in the past")
    }
    if appointment.PatientID == 0 || appointment.DoctorID == 0 {
        return errors.New("both patient and doctor must be specified")
    }

    return nil
}


// GetAllAppointments retrieves appointments, potentially filtered by criteria
func (s *appointmentService) GetAllAppointments(filterParams model.AppointmentFilterParams) ([]model.Appointment, error) {
    appointments, err := s.appointmentRepository.FindAll(filterParams)
    if err != nil {
        return nil, err
    }

    return appointments, nil
}

