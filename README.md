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