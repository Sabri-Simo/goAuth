# GoAuth - Authentication Service

## Overview

GoAuth is a high-performance backend authentication service built with Go. The service leverages the Gin framework for efficient HTTP routing, GORM for ORM with PostgreSQL, and bcrypt for secure password hashing. It provides a robust solution for user authentication, including secure signup and login functionalities.

## Features

- **User Signup**: Register new users with hashed passwords.
- **User Login**: Authenticate users by validating credentials.
- **Secure Password Handling**: Utilize bcrypt for hashing and verifying passwords.
- **Containerized Deployment**: Easy deployment with Docker.

## Technologies

### Gin
- **Description**: A lightweight and fast web framework for Go.
- **Role**: Manages HTTP routing and middleware.

### GORM
- **Description**: An ORM library for Go that simplifies database interactions.
- **Role**: Handles database operations with PostgreSQL.

### PostgreSQL
- **Description**: A powerful, open-source relational database.
- **Role**: Stores user data securely.

### bcrypt
- **Description**: A cryptographic algorithm for secure password hashing.
- **Role**: Hashes passwords before storage.

### Docker
- **Description**: A platform for containerizing applications.
- **Role**: Packages the Go application and PostgreSQL database for consistent deployment.

## Installation

### Prerequisites

- Docker and Docker Compose installed on your machine.
- Git for version control.

### Steps

1. **Clone the Repository**

    ```bash
    git clone https://github.com/yourusername/goauth.git
    cd goauth
    ```

2. **Configure Environment Variables**

   Create a `.env` file in the root directory with the following content:

    ```env
    PORT=3030
    DATABASE="postgresql://username:password@hostname:port/database?sslmode=require"
    ```

3. **Build and Run**

    Use Docker Compose to build and run the application:

    ```bash
    docker-compose up --build
    ```

4. **Access the API**

   - **Signup**: `POST /sing-up`
   - **Login**: `GET /login`

