# Fiber MongoDB HRMS

## Project Overview

This project is an API server built with Go (Golang) using the Fiber web framework. It provides endpoints for managing employees in a Human Resource Management System (HRMS). The server interacts with a MongoDB database to perform CRUD operations on employee data.

## File Structure

The project has the following file structure:
```
- main.go
- go.mod
- go.sum
- app/
  - server.go
  - server-2.go
- handlers/
  - employee_handler.go
  - employee_handler-2.go
- models/
  - employee.go
- database/
  - mongo.go
```

- `main.go`: The entry point of the application that starts the server and initializes configurations.
- `go.mod` and `go.sum`: Files for managing dependencies using Go modules.
- `app/`: Directory containing server implementations.
  - `server.go`: The main server file that sets up routes and handles HTTP requests.
  - `server-2.go`: An alternative or additional server implementation.
- `handlers/`: Directory containing handler functions for API endpoints.
  - `employee_handler.go`: Handler functions for employee-related API endpoints.
  - `employee_handler-2.go`: Additional or alternative handler functions.
- `models/`: Directory containing model definitions.
  - `employee.go`: Struct definition for the employee data structure.
- `database/`: Directory containing database-related operations.
  - `mongo.go`: Functionality for connecting to a MongoDB database and interacting with it.

## Getting Started

To start the project, follow these steps:

1. Make sure you have Go installed on your machine.

2. Clone the repository:

   ```shell
    git clone github.com/Mdromi/Crash-Course/golang-11-projects/9-fiber-hrms/
    ```
3. Install the project dependencies by running the following command in the project root directory:
    ```
    go mod download
    ```
4. Set up a MongoDB database and update the connection details in database/mongo.go.
   
5. Start the server by running the following command in the project root directory:
   ```
   go run main.go
   ```
6. The server will start running on the specified port, and you can now access the API endpoints.
   
# API Endpoints

- `/employee`: GET - Retrieve all employees.
- `/employee`: POST - Create a new employee.
- `/employee/:id`: PUT - Update an employee by ID.
- `/employee/:id`: DELETE - Delete an employee by ID.

Please refer to the source code and API handler files for more details on the endpoint implementations.
