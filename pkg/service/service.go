package service

import (
    "clinic-appointment-system/pkg/model"
    "clinic-appointment-system/pkg/repository"
    "errors"
    "strings"
)

// ServiceService defines the interface for service operations
type ServiceService interface {
    CreateService(service model.Service) (model.Service, error)
    GetServiceByID(serviceID uint) (model.Service, error)
    UpdateService(service model.Service) (model.Service, error)
    DeleteService(serviceID uint) error
    GetAllServices() ([]model.Service, error)
}

type serviceService struct {
    serviceRepository repository.ServiceRepository
}

// NewServiceService creates a new instance of ServiceService
func NewServiceService(repo repository.ServiceRepository) ServiceService {
    return &serviceService{
        serviceRepository: repo,
    }
}

func (s *serviceService) CreateService(service model.Service) (model.Service, error) {
    if err := validateService(service); err != nil {
        return model.Service{}, err
    }

    createdService, err := s.serviceRepository.Create(service)
    if err != nil {
        return model.Service{}, errors.New("failed to create service: " + err.Error())
    }

    return createdService, nil
}

func (s *serviceService) GetServiceByID(serviceID uint) (model.Service, error) {
    service, err := s.serviceRepository.FindByID(serviceID)
    if err != nil {
        return model.Service{}, errors.New("service not found")
    }

    return service, nil
}

func (s *serviceService) UpdateService(service model.Service) (model.Service, error) {
    if err := validateService(service); err != nil {
        return model.Service{}, err
    }

    updatedService, err := s.serviceRepository.Update(service)
    if err != nil {
        return model.Service{}, errors.New("failed to update service: " + err.Error())
    }

    return updatedService, nil
}

func (s *serviceService) DeleteService(serviceID uint) error {
    err := s.serviceRepository.Delete(serviceID)
    if err != nil {
        return errors.New("failed to delete service")
    }
    return nil
}

// validateService checks if the service's information meets certain criteria
func validateService(service model.Service) error {
    if strings.TrimSpace(service.Name) == "" {
        return errors.New("service name is required")
    }
    if service.Duration <= 0 {
        return errors.New("service duration must be positive")
    }


    return nil
}


func (s *serviceService) GetAllServices() ([]model.Service, error) {
    // Implementation to retrieve all services from the repository
    services, err := s.serviceRepository.FindAll()
    if err != nil {
        return nil, err
    }
    return services, nil
}

