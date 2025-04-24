# go-task-management

This is a go boilerplate project with `fiber` and `gin`.

## Table of Contents

- [go-task-management](#go-task-management)
  - [Table of Contents](#table-of-contents)
  - [Installation](#installation)
  - [Database](#database)
    - [Run MySQL](#run-mysql)
    - [Migrate tables](#migrate-tables)
    - [Run Redis](#run-redis)
  - [Run](#run)
    - [Make File](#make-file)
    - [Air](#air)
    - [Bash](#bash)
  - [TODO](#todo)
  - [Change Log](#change-log)
  - [Author](#author)

## Installation

- To run with `fiber` checkout branch `main`
- To run with `gin` checkout branch `feature/gin-server`

## Database

To using project with MySQL, set `DB_DRIVER` to `mysql`. If you want `SQLite` set it to `sqlite`

### Run MySQL

To run MySQL with docker compose run `docker compose up -d`.

for running wihtout compose run this command (Set your username and password):

```bash
docker run --name mysql-container -e MYSQL_ROOT_PASSWORD=my-secret-pw -p 3306:3306 -d mysql:latest
```

Create a new database in the container:

```bash
docker exec -it mysql-container mysql -u root -p 

# Then run:
CREATE DATABASE task_management;

SHOW DATABASES;
```

### Migrate tables

To migrate tables you can use make commands:

```bash
# Migrate up
make migrate-up:

# Migrate Down
make migrate-down:

# Get status of migrations
migrate-status:

# Create new migration file
make migrate-new
```

### Run Redis

To run Redis container with docker compose you can run `docker compose up -d` or if you want to run it separately, run this:

```bash
docker compose up redis -d
```

Check redis with:

```bash
docker exec -it redis-container redis-cli
```

## Run

Clone Project

```bash
git clone https://github.com/amirkenarang/go-task-management.git
cd go-task-management
```

### Make File

```bash
make run
```

### Air

Run with hot-reload

```bash
air
```

### Bash

Or run with Go

```bash
go run ./cmd
```

## TODO

- [x] Best Practices for project structure
- [x] Separate Models, Repository, Routes, Handlers, Middlewares
- [x] Add authentication
- [x] Managment
- [x] using environment
- [x] Using fiber and gin
- [x] Initialize SQLite database
- [x] Initialize air
- [x] Initialize MySQL
  - [x] Initialize Migration
  - [x] Docker: MySQL
- [x] Initialize Redis
  - [x] Docker: Redis
- [ ] Initialize NATS
- [ ] Loging structure
- [ ] Monitoring
  - [x] Prometheus server
  - [ ] Grafana
- [ ] Initialize Kubernetese
- [ ] Unit-test coverage
- [x] Front-End Application using NextJS [NextJS APP](https://github.com/amirkenarang/nextjs-task-management)

## Change Log

See [Changelog](CHANGELOG.md) for more information.

## Author

Amir Kenarang (<amir.kenarang@gmail.com>)
