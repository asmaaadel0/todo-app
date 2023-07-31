# 📝 Table of Contents

- [📝 Table of Contents](#-table-of-contents)
  - [About ](#about-)
  - [Website Features ](#website-features-)
  - [Technologies Used ](#technologies-used-)
  - [Prerequisites ](#prerequisites-)
  - [How to Use ](#how-to-use-)
  - [API Documentation ](#api-documentation-)
  - [End-to-End Tests ](#end-to-end-tests-)

## About <a name = "about"></a>

This is a Todo application implemented as a full-stack application using Go (Golang) for the backend, Vue.js for the frontend, Cypress for end-to-end testing, and Swagger for API documentation.

## Website Features <a name = "website-Features"></a>

- Add new task for tasks list.
- Update existing task if it's completed or not
- Update text for existing task
- Delete existing task from tasks list
- Filter tasks based on their completion status (All, Active, Completed)

## Technologies Used <a name = "Technologies-Used"></a>

- Backend

  - Go (Golang): A fast and efficient programming language for the backend.
  - Gin: A lightweight web framework for Go, used to build the API.
  - Swagger: Used for API documentation and testing.

- Frontend:
  - Vue.js: A progressive JavaScript framework for building user interfaces.
  - Cypress: An end-to-end testing framework for web applications.

## Prerequisites <a name = "Prerequisites"></a>

- Before running the Todo app, ensure you have the following installed:
  - Go: https://golang.org/doc/install
  - Node.js and npm: https://nodejs.org/en/download/

## How to Use <a name = "How-to-Use"></a>

- Clone the repository:

```sh
https://github.com/codescalersinternships/todoapp-Asmaa.git

cd todoapp-Asmaa
```

- Backend setup:

```sh
cd backend

go mod download

go run main.go

```

- Frontend setup:

```sh
cd frontend

npm install

npm run serve

```

- Access the application:
  - Open your web browser and go to http://localhost:8080 to access the Todo app.

## API Documentation <a name = " API-Documentation"></a>

- To view the API documentation and test the backend API using Swagger, go to http://localhost:3000/swagger/index.html after running the backend server.

## End-to-End Tests <a name = "End-to-End-Tests"></a>

- he end-to-end tests are implemented using Cypress. To run the tests, follow these steps:
  - Make sure the backend and frontend are running.
  - from frontend directory run:
  ```sh
  npx cypress open
  ```
