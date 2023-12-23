package model

import (
    "gorm.io/gorm"
    "time"
)

type Appointment struct {
    gorm.Model
    AppointmentID uint      `gorm:"primaryKey"`
    PatientID     uint      `gorm:"index"`
    DoctorID      uint      `gorm:"index"`
    Date          time.Time
    Time          time.Time
    Status        string
}


type AppointmentFilterParams struct {
    DoctorID  string
    PatientID string
    Date      string // Depending on how you handle dates, this might be a time.Time or similar
}

