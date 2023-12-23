package service

import (
    "clinic-appointment-system/pkg/model"
    "clinic-appointment-system/pkg/repository"
    "errors"
    "time"
)

// PatientService defines the interface for patient service operations
type PatientService interface {
    CreatePatient(patient model.Patient) (model.Patient, error)
    GetPatientByID(patientID uint) (model.Patient, error)
    UpdatePatient(patient model.Patient) (model.Patient, error)
    DeletePatient(patientID uint) error
    GetAllPatients() ([]model.Patient, error)
}

type patientService struct {
    patientRepository repository.PatientRepository
}

// NewPatientService creates a new instance of PatientService
func NewPatientService(repo repository.PatientRepository) PatientService {
    return &patientService{
        patientRepository: repo,
    }
}

func (s *patientService) CreatePatient(patient model.Patient) (model.Patient, error) {
    if err := validatePatient(patient); err != nil {
        return model.Patient{}, err
    }

    createdPatient, err := s.patientRepository.Create(patient)
    if err != nil {
        return model.Patient{}, errors.New("failed to create patient: " + err.Error())
    }

    return createdPatient, nil
}

func (s *patientService) GetPatientByID(patientID uint) (model.Patient, error) {
    patient, err := s.patientRepository.FindByID(patientID)
    if err != nil {
        return model.Patient{}, errors.New("patient not found")
    }

    return patient, nil
}

func (s *patientService) UpdatePatient(patient model.Patient) (model.Patient, error) {
    if err := validatePatient(patient); err != nil {
        return model.Patient{}, err
    }

    updatedPatient, err := s.patientRepository.Update(patient)
    if err != nil {
        return model.Patient{}, errors.New("failed to update patient: " + err.Error())
    }

    return updatedPatient, nil
}

func (s *patientService) DeletePatient(patientID uint) error {
    err := s.patientRepository.Delete(patientID)
    if err != nil {
        return errors.New("failed to delete patient")
    }
    return nil
}

// validatePatient checks if the patient's information meets certain criteria
func validatePatient(patient model.Patient) error {
    if time.Since(patient.DateOfBirth).Hours() < 0 {
        return errors.New("patient date of birth cannot be in the future")
    }


    return nil
}



func (s *patientService) GetAllPatients() ([]model.Patient, error) {
    // Implementation to retrieve all patients from the repository
    patients, err := s.patientRepository.FindAll()
    if err != nil {
        return nil, err
    }
    return patients, nil
}

