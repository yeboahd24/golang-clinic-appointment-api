package service

import (
    "clinic-appointment-system/pkg/model"
    "clinic-appointment-system/pkg/repository"
    "errors"
    "strings"
)

// ClinicLocationService defines the interface for clinic location service operations
type ClinicLocationService interface {
    CreateClinicLocation(location model.ClinicLocation) (model.ClinicLocation, error)
    GetClinicLocationByID(locationID uint) (model.ClinicLocation, error)
    UpdateClinicLocation(location model.ClinicLocation) (model.ClinicLocation, error)
    DeleteClinicLocation(locationID uint) error
    GetAllClinicLocations() ([]model.ClinicLocation, error)
}

type clinicLocationService struct {
    clinicLocationRepository repository.ClinicLocationRepository
}

// NewClinicLocationService creates a new instance of ClinicLocationService
func NewClinicLocationService(repo repository.ClinicLocationRepository) ClinicLocationService {
    return &clinicLocationService{
        clinicLocationRepository: repo,
    }
}

func (s *clinicLocationService) CreateClinicLocation(location model.ClinicLocation) (model.ClinicLocation, error) {
    if err := validateClinicLocation(location); err != nil {
        return model.ClinicLocation{}, err
    }

    createdLocation, err := s.clinicLocationRepository.Create(location)
    if err != nil {
        return model.ClinicLocation{}, errors.New("failed to create clinic location: " + err.Error())
    }

    return createdLocation, nil
}

func (s *clinicLocationService) GetClinicLocationByID(locationID uint) (model.ClinicLocation, error) {
    location, err := s.clinicLocationRepository.FindByID(locationID)
    if err != nil {
        return model.ClinicLocation{}, errors.New("clinic location not found")
    }

    return location, nil
}

func (s *clinicLocationService) UpdateClinicLocation(location model.ClinicLocation) (model.ClinicLocation, error) {
    if err := validateClinicLocation(location); err != nil {
        return model.ClinicLocation{}, err
    }

    updatedLocation, err := s.clinicLocationRepository.Update(location)
    if err != nil {
        return model.ClinicLocation{}, errors.New("failed to update clinic location: " + err.Error())
    }

    return updatedLocation, nil
}

func (s *clinicLocationService) DeleteClinicLocation(locationID uint) error {
    err := s.clinicLocationRepository.Delete(locationID)
    if err != nil {
        return errors.New("failed to delete clinic location")
    }
    return nil
}

// validateClinicLocation checks if the clinic location's information meets certain criteria
func validateClinicLocation(location model.ClinicLocation) error {
    if strings.TrimSpace(location.Name) == "" {
        return errors.New("clinic location name is required")
    }
    if strings.TrimSpace(location.Address) == "" {
        return errors.New("clinic location address is required")
    }
    // Add additional validation as needed
    // ...

    return nil
}

func (s *clinicLocationService) GetAllClinicLocations() ([]model.ClinicLocation, error) {
    return s.clinicLocationRepository.FindAll()
}
