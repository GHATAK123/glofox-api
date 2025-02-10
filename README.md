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
   - [End-to-End Tests](#end-to-end-tests)
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
