# Info
This project is developed for learning purposes. This is a site that allows you to order meals in the canteen. Based on the programming language golang. The mvc framework utron is used. Located at "github. com/gernest/utron".

# Entity Relationship Modell
The model described here is used.
![ER Modell](./doc/erdiagramm.jpg "Entity-Relationship-Modell")

# Source

##### Installation:
```sh
go get git@github.com:Frzifus/go-dbwt.git
```

# Building

You will need golang (1.6 or newer).
Currently supported/tested are:
 - linux amd64
 - win64
 - arm

##### Build => "./build/bin/":
```sh
make amd64 # linux-amd64
make win64 # win-64.exe
make arm   # linux-arm
```

# Setup Docker environment
Note: docker compose v2+ is required.
```sh
make docker-build
```
creates a container with mariadb under the name dbwt_database_1. The default mysql root password is "1234". Furthermore, a second container dbwt_app_1 is created, where the project will be compiled.

# Test

Write output to ./build/log/test_[date].log
##### Run tests:
```sh
make test
```

# Visualize call graph of your Go program using dot format
[More info](https://github.com/TrueFurby/go-callvis)
### Usage

```sh
make viz
```
