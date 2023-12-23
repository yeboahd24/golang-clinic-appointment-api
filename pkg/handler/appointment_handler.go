package handler

import (
    "clinic-appointment-system/pkg/model"
    "clinic-appointment-system/pkg/service"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

type AppointmentHandler struct {
    appointmentService service.AppointmentService
}

func NewAppointmentHandler(appointmentService service.AppointmentService) *AppointmentHandler {
    return &AppointmentHandler{appointmentService: appointmentService}
}

func (h *AppointmentHandler) CreateAppointment(c *gin.Context) {
    var appointment model.Appointment
    if err := c.ShouldBindJSON(&appointment); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    newAppointment, err := h.appointmentService.CreateAppointment(appointment)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, newAppointment)
}

func (h *AppointmentHandler) GetAppointment(c *gin.Context) {
    appointmentID, err := strconv.ParseUint(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid appointment ID"})
        return
    }

    appointment, err := h.appointmentService.GetAppointmentByID(uint(appointmentID))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "appointment not found"})
        return
    }

    c.JSON(http.StatusOK, appointment)
}


func (h *AppointmentHandler) GetAllAppointments(c *gin.Context) {
    // Optional: Implement query parameters for filtering (date, doctor, patient)
    filterParams := model.AppointmentFilterParams{
        DoctorID: c.Query("doctorId"),
        PatientID: c.Query("patientId"),
        Date: c.Query("date"),
    }

    appointments, err := h.appointmentService.GetAllAppointments(filterParams)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, appointments)
}


func (h *AppointmentHandler) UpdateAppointment(c *gin.Context) {
    appointmentID, err := strconv.ParseUint(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid appointment ID"})
        return
    }

    var appointment model.Appointment
    if err := c.ShouldBindJSON(&appointment); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    appointment.ID = uint(appointmentID)
    updatedAppointment, err := h.appointmentService.UpdateAppointment(appointment)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, updatedAppointment)
}

func (h *AppointmentHandler) DeleteAppointment(c *gin.Context) {
    appointmentID, err := strconv.ParseUint(c.Param("id"), 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid appointment ID"})
        return
    }

    err = h.appointmentService.DeleteAppointment(uint(appointmentID))
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "appointment deleted"})
}

