# Todo App

This is a Todo application where users can interact with their tasks and manage their todo list.

## Installation

- Clone the repository:

```sh
$ git clone https://github.com/codescalersinternships/todoapp-Asmaa.git
```

- Change into the project directory:

```sh
$ cd todoapp-Asmaa/backend
```

- Install Go dependencies:

```sh
go get ./...
```

- To start the Todo App, run the following command:

```sh
go run main.go
```

## API Endpoints

The Todo App provides the following API endpoints:

- GET /tasks: Get a list of all tasks.
- POST /tasks: Add a new task to the list.
- PUT /tasks: Update a task.
- DELETE /tasks/:id: Delete a task by its ID.

## How to test

- Run the tests by running:

```sh
go test ./...
```

- If all tests pass, the output indicate that the tests have passed. if there is failure, the output will provide information about it.
