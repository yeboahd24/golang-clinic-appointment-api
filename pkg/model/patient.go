package model

import (
    "gorm.io/gorm"
    "time"
)

type Patient struct {
    gorm.Model
    PatientID     uint      `gorm:"primaryKey"`
    UserID        uint      `gorm:"unique;index"`
    DateOfBirth   time.Time
    Gender        string
    MedicalHistory string   `gorm:"type:text"`
}

