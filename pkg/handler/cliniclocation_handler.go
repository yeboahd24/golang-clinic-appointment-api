package handler

import (
    "clinic-appointment-system/pkg/model"
    "clinic-appointment-system/pkg/service"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

type ClinicLocationHandler struct {
    clinicLocationService service.ClinicLocationService
}

func NewClinicLocationHandler(clinicLocationService service.ClinicLocationService) *ClinicLocationHandler {
    return &ClinicLocationHandler{clinicLocationService: clinicLocationService}
}

func (h *ClinicLocationHandler) AddClinicLocation(c *gin.Context) {
    var location model.ClinicLocation
    if err := c.ShouldBindJSON(&location); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    newLocation, err := h.clinicLocationService.CreateClinicLocation(location)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, newLocation)
}

func (h *ClinicLocationHandler) GetAllClinicLocations(c *gin.Context) {
    locations, err := h.clinicLocationService.GetAllClinicLocations()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, locations)
}

func (h *ClinicLocationHandler) GetClinicLocation(c *gin.Context) {
    locationID, err := strconv.ParseUint(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid location ID"})
        return
    }

    location, err := h.clinicLocationService.GetClinicLocationByID(uint(locationID))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "location not found"})
        return
    }

    c.JSON(http.StatusOK, location)
}

func (h *ClinicLocationHandler) UpdateClinicLocation(c *gin.Context) {
    locationID, err := strconv.ParseUint(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid location ID"})
        return
    }

    var location model.ClinicLocation
    if err := c.ShouldBindJSON(&location); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    location.ID = uint(locationID)
    updatedLocation, err := h.clinicLocationService.UpdateClinicLocation(location)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, updatedLocation)
}

func (h *ClinicLocationHandler) DeleteClinicLocation(c *gin.Context) {
    locationID, err := strconv.ParseUint(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid location ID"})
        return
    }

    err = h.clinicLocationService.DeleteClinicLocation(uint(locationID))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "location deleted"})
}

