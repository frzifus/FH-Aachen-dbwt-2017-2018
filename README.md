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
# amd64
make

# win64
make win64

# arm
make arm
```

# Setup Docker environment
```sh
make docker
```
creates a container with mariadb under the name dbwt-db. It is located in the dbwtnet network. The default mysql root password is "1234".

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
