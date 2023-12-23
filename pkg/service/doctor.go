package service

import (
    "clinic-appointment-system/pkg/model"
    "clinic-appointment-system/pkg/repository"
    "errors"
    "strings"
)

type DoctorService interface {
    CreateDoctor(doctor model.Doctor) (model.Doctor, error)
    GetDoctorByID(doctorID uint) (model.Doctor, error)
    UpdateDoctor(doctor model.Doctor) (model.Doctor, error)
    DeleteDoctor(doctorID uint) error
    GetAllDoctors() ([]model.Doctor, error)
}

type doctorService struct {
    doctorRepository repository.DoctorRepository
}

func NewDoctorService(repo repository.DoctorRepository) DoctorService {
    return &doctorService{
        doctorRepository: repo,
    }
}

func (s *doctorService) CreateDoctor(doctor model.Doctor) (model.Doctor, error) {
    if err := validateDoctor(doctor); err != nil {
        return model.Doctor{}, err
    }

    createdDoctor, err := s.doctorRepository.Create(doctor)
    if err != nil {
        return model.Doctor{}, errors.New("failed to create doctor: " + err.Error())
    }

    return createdDoctor, nil
}

func (s *doctorService) GetDoctorByID(doctorID uint) (model.Doctor, error) {
    doctor, err := s.doctorRepository.FindByID(doctorID)
    if err != nil {
        return model.Doctor{}, errors.New("doctor not found")
    }

    return doctor, nil
}

func (s *doctorService) UpdateDoctor(doctor model.Doctor) (model.Doctor, error) {
    if err := validateDoctor(doctor); err != nil {
        return model.Doctor{}, err
    }

    updatedDoctor, err := s.doctorRepository.Update(doctor)
    if err != nil {
        return model.Doctor{}, errors.New("failed to update doctor: " + err.Error())
    }

    return updatedDoctor, nil
}

func (s *doctorService) DeleteDoctor(doctorID uint) error {
    err := s.doctorRepository.Delete(doctorID)
    if err != nil {
        return errors.New("failed to delete doctor")
    }
    return nil
}

// validateDoctor checks if the doctor's information meets certain criteria
func validateDoctor(doctor model.Doctor) error {
    if strings.TrimSpace(doctor.Specialization) == "" {
        return errors.New("doctor specialization is required")
    }


    return nil
}


func (s *doctorService) GetAllDoctors() ([]model.Doctor, error) {
    // Implementation to retrieve all doctors from the repository
    doctors, err := s.doctorRepository.FindAll()
    if err != nil {
        return nil, err
    }
    return doctors, nil
}
