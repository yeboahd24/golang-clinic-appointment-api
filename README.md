# Clinic Appointment System API Documentation

## Overview

This document provides sample POST requests for the Clinic Appointment System API.

### Base URL

The base URL for the API (when running locally) is `http://localhost:8080`.

## Endpoints

### User Registration

- **URL**: `/api/users/register`
- **Method**: `POST`
- **Curl Example**:
  ```bash
  curl -X POST http://localhost:8080/api/users/register \\
      -H "Content-Type: application/json" \\
      -d '{
          "Name": "John Doe",
          "Email": "johndoe@example.com",
          "Password": "yourpassword",
          "Role": "user",
          "ContactNumber": "1234567890"
      }'  

  ```

- URL: `/api/users/login`
- Method: `POST`
- Curl Example:
  ```bash
  curl -X POST http://localhost:8080/api/users/login \\
    -H "Content-Type: application/json" \\
    -d '{
        "Email": "johndoe@example.com",
        "Password": "yourpassword"
    }'

  ```
- URL: /api/appointments
- Method: `POST`
- Curl Example:

  ```bash
  curl -X POST http://localhost:8080/api/appointments \
    -H "Content-Type: application/json" \
    -d '{
        "patient_id": 1,
        "doctor_id": 2,
        "date": "2023-04-15",
        "time": "14:00",
        "status": "scheduled"
    }'
  ```

- URL: /api/services
- Method: `POST`
- Curl Example:

  ```bash
  curl -X POST http://localhost:8080/api/services \
    -H "Content-Type: application/json" \
    -d '{
        "name": "General Consultation",
        "description": "A general medical consultation for routine health check-ups and non-emergency issues.",
        "duration": 30
    }'
  ```


- URL: /api/locations
- Method: `POST`
- Curl Example:

  ```bash
  curl -X POST http://localhost:8080/api/locations \
    -H "Content-Type: application/json" \
    -d '{
        "name": "Downtown Clinic",
        "address": "123 Main St, Downtown City, DC 12345",
        "contact_info": "555-1234"
    }'

  ```
# golang-clinic-appointment-api
