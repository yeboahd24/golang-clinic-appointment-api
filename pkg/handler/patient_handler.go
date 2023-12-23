package handler

import (
    "clinic-appointment-system/pkg/model"
    "clinic-appointment-system/pkg/service"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

type PatientHandler struct {
    patientService service.PatientService
}

func NewPatientHandler(patientService service.PatientService) *PatientHandler {
    return &PatientHandler{patientService: patientService}
}

// func (h *PatientHandler) GetAllPatients(c *gin.Context) {
//     patients, err := h.patientService.GetAllPatients()
//     if err != nil {
//         c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//         return
//     }
//
//     c.JSON(http.StatusOK, patients)
// }

func (h *PatientHandler) GetPatient(c *gin.Context) {
    patientID, err := strconv.ParseUint(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid patient ID"})
        return
    }

    patient, err := h.patientService.GetPatientByID(uint(patientID))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "patient not found"})
        return
    }

    c.JSON(http.StatusOK, patient)
}

func (h *PatientHandler) UpdatePatient(c *gin.Context) {
    patientID, err := strconv.ParseUint(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid patient ID"})
        return
    }

    var patient model.Patient
    if err := c.ShouldBindJSON(&patient); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    patient.ID = uint(patientID)
    updatedPatient, err := h.patientService.UpdatePatient(patient)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, updatedPatient)
}

func (h *PatientHandler) DeletePatient(c *gin.Context) {
    patientID, err := strconv.ParseUint(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid patient ID"})
        return
    }

    err = h.patientService.DeletePatient(uint(patientID))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "patient deleted"})
}



func (h *PatientHandler) GetAllPatients(c *gin.Context) {
    patients, err := h.patientService.GetAllPatients()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, patients)
}




