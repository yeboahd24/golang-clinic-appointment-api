package handler

import (
    "clinic-appointment-system/pkg/model"
    "clinic-appointment-system/pkg/service"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

type DoctorHandler struct {
    doctorService service.DoctorService
}

func NewDoctorHandler(doctorService service.DoctorService) *DoctorHandler {
    return &DoctorHandler{doctorService: doctorService}
}
//
// func (h *DoctorHandler) GetAllDoctors(c *gin.Context) {
//     doctors, err := h.doctorService.GetAllDoctors()
//     if err != nil {
//         c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//         return
//     }
//
//     c.JSON(http.StatusOK, doctors)
// }

func (h *DoctorHandler) GetDoctor(c *gin.Context) {
    doctorID, err := strconv.ParseUint(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid doctor ID"})
        return
    }

    doctor, err := h.doctorService.GetDoctorByID(uint(doctorID))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "doctor not found"})
        return
    }

    c.JSON(http.StatusOK, doctor)
}

func (h *DoctorHandler) UpdateDoctor(c *gin.Context) {
    doctorID, err := strconv.ParseUint(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid doctor ID"})
        return
    }

    var doctor model.Doctor
    if err := c.ShouldBindJSON(&doctor); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    doctor.ID = uint(doctorID)
    updatedDoctor, err := h.doctorService.UpdateDoctor(doctor)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, updatedDoctor)
}

func (h *DoctorHandler) DeleteDoctor(c *gin.Context) {
    doctorID, err := strconv.ParseUint(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid doctor ID"})
        return
    }

    err = h.doctorService.DeleteDoctor(uint(doctorID))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "doctor deleted"})
}



func (h *DoctorHandler) GetAllDoctors(c *gin.Context) {
    doctors, err := h.doctorService.GetAllDoctors()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, doctors)
}

