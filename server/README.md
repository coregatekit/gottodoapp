# Todo Application Server with gin-gonic

## Prerequisite

- Go version 1.22+
- Air - Live reload for Go apps download from [Github](https://github.com/air-verse/air])

## Quick start

Before start the application please set environment variable to specific database path.

```bash
DATABASE_URL=todo.db
(.env file)

or

export DATABASE_URL=todo.db
```

After install all requirements use these command to start develop.
The default port is `8080` visit [http://localhost:8080/](http://localhost:8080/)

To start local

```bash
air
```

To build the application

```bash
go build -o /build/server .
```

To start the application from build file

```bash
./build/server
```

After start the application please visit ping path the verify server is starting correctly. [http://localhost:8080/api/ping](http://localhost:8080/api/ping)
