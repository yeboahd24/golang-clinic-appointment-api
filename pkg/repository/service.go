package repository

import (
    "clinic-appointment-system/pkg/model"
    "gorm.io/gorm"
)

// ServiceRepository defines the interface for service data operations
type ServiceRepository interface {
    Create(service model.Service) (model.Service, error)
    FindByID(serviceID uint) (model.Service, error)
    Update(service model.Service) (model.Service, error)
    Delete(serviceID uint) error
    FindAll() ([]model.Service, error)
}

type serviceRepository struct {
    db *gorm.DB
}

// NewServiceRepository creates a new instance of ServiceRepository
func NewServiceRepository(db *gorm.DB) ServiceRepository {
    return &serviceRepository{
        db: db,
    }
}

// Create adds a new service to the database
func (r *serviceRepository) Create(service model.Service) (model.Service, error) {
    err := r.db.Create(&service).Error
    return service, err
}

// FindByID finds a service by its ID
func (r *serviceRepository) FindByID(serviceID uint) (model.Service, error) {
    var service model.Service
    err := r.db.Where("id = ?", serviceID).First(&service).Error
    return service, err
}

// Update modifies an existing service in the database
func (r *serviceRepository) Update(service model.Service) (model.Service, error) {
    err := r.db.Save(&service).Error
    return service, err
}

// Delete removes a service from the database
func (r *serviceRepository) Delete(serviceID uint) error {
    err := r.db.Delete(&model.Service{}, serviceID).Error
    return err
}

func (r *serviceRepository) FindAll() ([]model.Service, error) {
    var services []model.Service
    err := r.db.Find(&services).Error
    return services, err
}

