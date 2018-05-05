# Info
This project is developed for learning purposes. This is a site that allows you to order and view meals in the canteen. Based on the programming language golang. The mvc framework [utron](https://github.com/gernest/utron) is used.

# Entity Relationship Modell
The model described here is used.
![ER Modell](./doc/erdiagramm.jpg "Entity-Relationship-Modell")

# Source

##### Installation:
```sh
go get -u -d -v github.com/Frzifus/dbwt
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

# Documentation
##### Run godoc:
```sh
godoc -http=":6060"
```
Creates a local webserver with package documentation and can be found at [localhost:6060/pkg/github.com/frzifus/dbwt](http://localhost:6060/pkg/github.com/frzifus/dbwt/)

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
