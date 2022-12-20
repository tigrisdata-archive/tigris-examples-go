# Todo Project

## Prerequisites

This project requires [Docker](https://docs.docker.com/get-docker/) and [Task](https://taskfile.dev/installation/) to be installed.

## Starting Project Locally

```sh
task run
```

This will start up the project at http://localhost:8080

Executing `task run:docker` will start the project in the detached docker container.

Run `task` without arguments to see all available commands.

## Project Structure

```
├── docker-compose.yml
├── Dockerfile
├── go.mod
├── go.sum
├── main.go
├── model
│   └── tasks.go
├── README.md
├── route
│   └── route.go
└── Taskfile.yaml
```

### Data Models
The `model` directory contains collections models, which is basically the structure of the document persisted
in the particular collection.

For example:

```golang
// Task Collection of documents with tasks details
type Task struct {
	// Completed Indicate task completion state
	Completed bool `json:"completed"`
	// CompletedAt Task completion date
	CompletedAt time.Time `json:"completed_at"`
	// Details Detail explanation of the task
	Details string `json:"details"`
	// DueAt Task due date
	DueAt time.Time `json:"due_at"`
	// Id A unique identifier for the task
	Id int64 `json:"id" tigris:"primaryKey:1,autoGenerate"`
	// Name Name of the task
	Name string `json:"name"`
	// Tags The list of task categories
	Tags []string `json:"tags"`
}
```

This model types can be modified to add new fields to the document.

### Routes

The `route/routes.go` defines SetupCRUD function which is used in the `main.go` to set up [Gin](https://github.com/gin-gonic/gin)
Web framework CRUD routes for every collection model.
Once project is started, they can be tested using curl commands.

For example:

#### Create document in the `tasks` collection:
```
curl -X POST "localhost:8080/tasks" -H 'Content-Type: application/json' 
    -d "{ JSON document body corresponding to the model.Task }"
```

#### Read document from the `tasks` collection:
```
curl -X GET "localhost:8080/tasks/{document id}"
```

#### Delete document from the `tasks` collection:
```
curl -X DELETE "localhost:8080/tasks/{document id}"
```

Full Tigris documentation [here](https://docs.tigrisdata.com).

Be brave. Have fun!
