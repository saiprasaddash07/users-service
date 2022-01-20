# User Service

> User Service built with Golang ang gin package

## Features

- CRUD operation with the user

### Env Variables

Create a .env file in then root and add the following

```
APP_ENV=service name
DB_USERNAME=user name of db
DB_HOST_READER=reader db url
DB_HOST_WRITER=writer db url
DB_PORT=port on which db is running
DB_PASSWORD=password of the db
DB_NAME=name of the database
DB_MAX_IDLE_CONNECTION=maximum no of idle connections
DB_MAX_OPEN_CONNECTIONS=maximum no of open connections
SERVICE_NAME=name of the service
SERVER_PORT=port on ehich server will run
REDIS_ADDRESS=redis url
CACHE_ENABLED=boolean value to control redis cache
REDIS_PASSWORD=password of the redis
```

### Install Dependencies

```
go mod tidy
```

### Run

```
# Run backend (:4000)
go run .
```
