# Glofox API

This is a Golang implementation of the Glofox API, designed to help boutique studios and gyms manage their classes and bookings. The application provides two main API endpoints: one for creating classes and another for booking classes. The application uses in-memory storage and adheres to RESTful standards.

---

## Table of Contents

1. [Features](#features)
2. [Setup](#setup)
   - [Prerequisites](#prerequisites)
   - [Running Locally](#running-locally)
   - [Running with Docker](#running-with-docker)
3. [API Endpoints](#api-endpoints)
   - [Create a Class](#create-a-class)
   - [Book a Class](#book-a-class)
4. [Testing](#testing)
   - [Unit Tests](#unit-tests)
   - [Integration Tests](#integration-tests)
5. [Assumptions](#assumptions)
6. [Future Improvements](#future-improvements)

---

## Features

- **Create Classes**: Studio owners can create classes with details such as class name, start date, end date, and capacity.
- **Book Classes**: Members can book classes by providing their name and the date they wish to attend.
- **In-Memory Storage**: Data is stored in memory, making it easy to test and develop without needing a database.
- **RESTful API**: The API follows REST standards and provides appropriate success and error responses.

---

## Setup

### Prerequisites

- Go 1.20 or higher
- Docker (optional, for containerization)

---

### Running Locally

1. Clone the repository:
   ```bash
   git clone https://github.com/GHATAK123/glofox-api.git
   cd glofox-api

2. Clone the repository:
   ```bash
   go run main.go

3. The server will start on `http://localhost:8080`

### Running with Docker

1. Build the Docker image:
   ```bash
   docker build -t glofox-api .

2. Run the Docker container:
   ```bash
   docker run -p 8080:8080 glofox-api

3. The server will be accessible at `http://localhost:8080`

## API Endpoints

### Create a Class

- **Endpoint** : `POST /classes`
- **Request Body** :
   ```json
   {
     "name": "Yoga",
     "start_date": "2025-02-11T00:00:00Z",
     "end_date": "2025-02-21T00:00:00Z",
     "capacity": 10
   }
- **Request Body** :
   ```json
   {
     "name": "Yoga",
     "start_date": "2025-02-11T00:00:00Z",
     "end_date": "2025-02-21T00:00:00Z",
     "capacity": 10
   }

   
 ### Book a Class

- **Endpoint** : `POST /bookings`
- **Request Body** :
   ```json
   {
     "name": "Prakash Anand",
     "date": "2025-02-16T00:00:00Z"
   }
- **Request Body** :
   ```json
   {
     "class_name": "Yoga",
     "member": "Prakash Anand",
     "date": "2025-02-16"
   }

## Testing

### Unit Tests

- Run the unit test cases:
   ```bash
   go test ./test/class_handler_test.go ./test/booking_handler_test.go
   
### Integration Tests

- Run the integration test cases:
   ```bash
   go test ./test/integration_test.go

## Assumptions
- Authentication: No authentication is required for the APIs.
- Class Uniqueness: Only one class can be created per day.
- Overbooking: Overbooking is not handled. The number of bookings can exceed the class capacity.
- In-Memory Storage: Data is stored in memory and will be lost when the application restarts.

## Future Improvements
- Database Integration: Replace in-memory storage with a database (e.g., PostgreSQL, MySQL) for persistent data storage.
- Authentication: Add authentication and authorization to secure the API endpoints.
- Overbooking Handling: Implement logic to prevent overbooking for a given class date.
- Class Scheduling: Allow multiple classes to be scheduled on the same day.
- API Documentation: Use tools like Swagger to generate API documentation.

## Contact
For questions or feedback, please reach out to ppa58514@gmail.com.

