package handler

import (
    "clinic-appointment-system/pkg/model"
    "clinic-appointment-system/pkg/service"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

type ServiceHandler struct {
    serviceService service.ServiceService
}

func NewServiceHandler(serviceService service.ServiceService) *ServiceHandler {
    return &ServiceHandler{serviceService: serviceService}
}

func (h *ServiceHandler) AddService(c *gin.Context) {
    var service model.Service
    if err := c.ShouldBindJSON(&service); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    newService, err := h.serviceService.CreateService(service)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, newService)
}

func (h *ServiceHandler) GetAllServices(c *gin.Context) {
    services, err := h.serviceService.GetAllServices()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, services)
}

func (h *ServiceHandler) GetService(c *gin.Context) {
    serviceID, err := strconv.ParseUint(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid service ID"})
        return
    }

    service, err := h.serviceService.GetServiceByID(uint(serviceID))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "service not found"})
        return
    }

    c.JSON(http.StatusOK, service)
}

func (h *ServiceHandler) UpdateService(c *gin.Context) {
    serviceID, err := strconv.ParseUint(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid service ID"})
        return
    }

    var service model.Service
    if err := c.ShouldBindJSON(&service); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    service.ID = uint(serviceID)
    updatedService, err := h.serviceService.UpdateService(service)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, updatedService)
}

func (h *ServiceHandler) DeleteService(c *gin.Context) {
    serviceID, err := strconv.ParseUint(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid service ID"})
        return
    }

    err = h.serviceService.DeleteService(uint(serviceID))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "service deleted"})
}

