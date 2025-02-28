# go-task-management

This is a go boilerplate project with `fiber` and `gin`.

## Table of Contents

- [go-task-management](#go-task-management)
  - [Table of Contents](#table-of-contents)
  - [Installation](#installation)
  - [Database](#database)
    - [Run MySQL](#run-mysql)
    - [Migrate tables](#migrate-tables)
  - [Run](#run)
    - [Make File](#make-file)
    - [Air](#air)
    - [Bash](#bash)
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

```
docker run --name mysql-container -e MYSQL_ROOT_PASSWORD=my-secret-pw -p 3306:3306 -d mysql:latest
```

Create a new database in the container:

```
docker exec -it mysql-container mysql -u root -p 

# Then run:
CREATE DATABASE taask_managment
```

### Migrate tables

To migrate tables you can use make commands:

```
# Migrate up
make migrate-up:

# Migrate Down
make migrate-down:

# Get status of migrations
migrate-status:

# Create new migration file
make migrate-new
```

## Run

Clone Project

```bash
git clone https://github.com/amirkenarang/go-task-management.git
cd go-task-management
```

### Make File

```
make run
```

### Air

Run with hot-reload

```bash
air
```

### Bash

Or run with Go

```
go run ./cmd
```

## Change Log

See [Changelog](CHANGELOG.md) for more information.

## Author

Amir Kenarang (<amir.kenarang@gmail.com>)
