package main

import (
    "clinic-appointment-system/pkg/handler"
    "clinic-appointment-system/pkg/repository"
    "clinic-appointment-system/pkg/service"
    "clinic-appointment-system/pkg/config"
    "log"

    "github.com/gin-gonic/gin"
)

func main() {
    // Initialize Gin Router
    router := gin.Default()

    // Setup Database Connection (replace with your database details)
    db := config.SetupDatabaseConnection()
    defer config.CloseDatabaseConnection(db)

    // Initialize Repositories
    userRepository := repository.NewUserRepository(db)
    appointmentRepository := repository.NewAppointmentRepository(db)
    doctorRepository := repository.NewDoctorRepository(db)
    patientRepository := repository.NewPatientRepository(db)
    serviceRepository := repository.NewServiceRepository(db)
    clinicLocationRepository := repository.NewClinicLocationRepository(db)

    // Initialize Services
    userService := service.NewUserService(userRepository)
    appointmentService := service.NewAppointmentService(appointmentRepository)
    doctorService := service.NewDoctorService(doctorRepository)
    patientService := service.NewPatientService(patientRepository)
    serviceService := service.NewServiceService(serviceRepository)
    clinicLocationService := service.NewClinicLocationService(clinicLocationRepository)

    // Initialize Handlers
    userHandler := handler.NewUserHandler(userService)
    appointmentHandler := handler.NewAppointmentHandler(appointmentService)
    doctorHandler := handler.NewDoctorHandler(doctorService)
    patientHandler := handler.NewPatientHandler(patientService)
    serviceHandler := handler.NewServiceHandler(serviceService)
    clinicLocationHandler := handler.NewClinicLocationHandler(clinicLocationService)

    // Register Routes
    // User Routes
    router.POST("/api/users/register", userHandler.RegisterUser)
    router.POST("/api/users/login", userHandler.LoginUser)
    router.GET("/api/users/:id", userHandler.GetUser)
    router.PUT("/api/users/:id", userHandler.UpdateUser)
    router.DELETE("/api/users/:id", userHandler.DeleteUser)

    // Appointment Routes
    router.POST("/api/appointments", appointmentHandler.CreateAppointment)
    router.GET("/api/appointments/:id", appointmentHandler.GetAppointment)
    router.GET("/api/appointments", appointmentHandler.GetAllAppointments)
    router.PUT("/api/appointments/:id", appointmentHandler.UpdateAppointment)
    router.DELETE("/api/appointments/:id", appointmentHandler.DeleteAppointment)

    // Doctor Routes
    router.GET("/api/doctors", doctorHandler.GetAllDoctors)
    router.GET("/api/doctors/:id", doctorHandler.GetDoctor)
    router.PUT("/api/doctors/:id", doctorHandler.UpdateDoctor)
    router.DELETE("/api/doctors/:id", doctorHandler.DeleteDoctor)

    // Patient Routes
    router.GET("/api/patients", patientHandler.GetAllPatients)
    router.GET("/api/patients/:id", patientHandler.GetPatient)
    router.PUT("/api/patients/:id", patientHandler.UpdatePatient)
    router.DELETE("/api/patients/:id", patientHandler.DeletePatient)

    // Service Routes
    router.POST("/api/services", serviceHandler.AddService)
    router.GET("/api/services", serviceHandler.GetAllServices)
    router.GET("/api/services/:id", serviceHandler.GetService)
    router.PUT("/api/services/:id", serviceHandler.UpdateService)
    router.DELETE("/api/services/:id", serviceHandler.DeleteService)

    // Clinic Location Routes
    router.POST("/api/locations", clinicLocationHandler.AddClinicLocation)
    router.GET("/api/locations", clinicLocationHandler.GetAllClinicLocations)
    router.GET("/api/locations/:id", clinicLocationHandler.GetClinicLocation)
    router.PUT("/api/locations/:id", clinicLocationHandler.UpdateClinicLocation)
    router.DELETE("/api/locations/:id", clinicLocationHandler.DeleteClinicLocation)

    // Start server
    if err := router.Run(":8080"); err != nil {
        log.Fatal("Failed to run server:", err)
    }
}


