# Gin User Management API

This is a simple user management API built with Golang and the Gin framework. It supports user registration, login, and basic user management functionalities. The application uses PostgreSQL as its database and JWT for authentication.

## Features

- User Registration
- User Login
- JWT Authentication
- Get All Users
- Get User by ID
- Update User
- Delete User

## Table of Contents

- [Requirements](#requirements)
- [Installation](#installation)
- [Database Schema](#database-schema)
- [API Endpoints](#api-endpoints)
- [Running the Application](#running-the-application)
- [License](#license)

## Requirements

- Golang 1.16 or higher
- PostgreSQL

## Installation

1. Clone the repository:

    `git clone https://github.com/yourusername/gin-user-management.git`

    `cd gin-user-management`

2. Install the dependencies:

    `go mod tidy`

3. Configure the database connection:

    Update the `config/database.go` file with your PostgreSQL database details.

4. Create the necessary tables in your PostgreSQL database using the schema provided in the next section.

## Database Schema

Create the following tables in your PostgreSQL database:

- **Users Table**

  - `user_id`: SERIAL, PRIMARY KEY
  - `username`: VARCHAR(100), NOT NULL, UNIQUE
  - `password`: VARCHAR(255), NOT NULL
  - `login_attempts`: INT, DEFAULT 0
  - `last_login_attempt`: TIMESTAMP, DEFAULT NULL
  - `remember_me_token`: VARCHAR(255), DEFAULT NULL
  - `created_at`: TIMESTAMP, DEFAULT CURRENT_TIMESTAMP
  - `updated_at`: TIMESTAMP, DEFAULT CURRENT_TIMESTAMP, ON UPDATE CURRENT_TIMESTAMP

- **Roles Table**

  - `role_id`: SERIAL, PRIMARY KEY
  - `role_name`: VARCHAR(100), NOT NULL, UNIQUE
  - `created_at`: TIMESTAMP, DEFAULT CURRENT_TIMESTAMP
  - `updated_at`: TIMESTAMP, DEFAULT CURRENT_TIMESTAMP, ON UPDATE CURRENT_TIMESTAMP

- **Groups Table**

  - `id`: SERIAL, PRIMARY KEY
  - `role_id`: INT
  - `user_id`: INT
  - `created_at`: TIMESTAMP, DEFAULT CURRENT_TIMESTAMP
  - `updated_at`: TIMESTAMP, DEFAULT CURRENT_TIMESTAMP, ON UPDATE CURRENT_TIMESTAMP
  - Foreign key relationships:
    - `role_id` REFERENCES `roles(role_id)`
    - `user_id` REFERENCES `users(user_id)`
  - UNIQUE constraint on `(role_id, user_id)`

## API Endpoints

- **Register**

  - `POST /register`
  - Request body: JSON with `username` and `password`
  - Response: JSON with user details

- **Login**

  - `POST /login`
  - Request body: JSON with `username` and `password`
  - Response: JSON with JWT token

- **Get All Users**

  - `GET /users`
  - Response: JSON array of user details

- **Get User by ID**

  - `GET /users/:id`
  - Response: JSON with user details

- **Update User**

  - `PUT /users/:id`
  - Request body: JSON with updated `username` and/or `password`
  - Response: JSON with updated user details

- **Delete User**

  - `DELETE /users/:id`
  - Response: JSON indicating success

## Running the Application

1. Run the application:

    `go run main.go`

2. The API will be available at `http://localhost:8080`.

## License

This project is licensed under the MIT License - see the LICENSE file for details.

---

By following this guide, you should be able to set up and run the Gin User Management API on your local machine. Adjust the database connection settings as necessary to match your environment.
