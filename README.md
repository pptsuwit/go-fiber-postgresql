# Basic golang + auth + crud

##

[![pptsuwit](https://avatars.githubusercontent.com/u/90542847?v=4)](https://github.com/pptsuwit)

## Installation

requires [Golang](https://go.dev/) v1.21.3+ to run.
requires [Air live reload](https://github.com/cosmtrek/air)

## Config

Copy example environment.

```sh
cp example.env.yaml env.yaml
```

Change environments .env file

```sh
app:
  port: 8000 # debug port
  jwtSecret: "jwt-secret" # jwt secret
  tokenLiftHour: 48 # token lifetime
db:
  driver: "postgres" # database connection
  ## if use docker in localhost change host to --> host: "host.docker.internal"
  host: "localhost" # host name
  port: 5432 # database connection port
  username: "database-username" # database username
  password: "database-password" # database password
  database: "database-name" # database name
```

Seeding database (Optional)

```sh
go run .\scripts\seed.go
```

Run

```sh
air
```

## Go pkg

Instructions on how to use them in your own application are linked below.

| Name      | link                                                    |
| --------- | ------------------------------------------------------- |
| validator | [github.com/go-playground/validator/v10 v10.16.0][pkg1] |
| fiber     | [githubgithub.com/gofiber/fiber v1.14.6][pkg2]          |
| fiber v2  | [github.com/gofiber/fiber/v2 v2.51.0][pkg3]             |
| jwt       | [github.com/gofiber/jwt/v2 v2.2.7][pkg4]                |
| jwt v3    | [github.com/golang-jwt/jwt v3.2.2+incompatible][pkg5]   |
| viper     | [github.com/spf13/viper v1.17.0][pkg6]                  |
| zap       | [go.uber.org/zap v1.26.0][pkg7]                         |
| crypto    | [golang.org/x/crypto v0.15.0][pkg8]                     |
| postgres  | [gorm.io/driver/postgres v1.5.4][pkg9]                  |
| gorm      | [gorm.io/gorm v1.25.5][pkg10]                           |

## Features

- Login
- Register
- CRUD

## Docker

to install and deploy in a Docker container.

By default, the Docker will expose port 8000, so change this within the
Dockerfile if necessary. When ready, simply use the Dockerfile to
build the image.

```sh
docker build -t <image-name> .
```

This will create the image and pull in the necessary dependencies.

Once done, run the Docker image and map the port to whatever you wish on
your host. In this example, we simply map port 8000 of the host to
port 8000 of the Docker (or whatever port was exposed in the Dockerfile):

```sh
docker run --rm -p 8000:8000 -d --name [container-name] [image-name]
```

# example

```sh
http://localhost:8000/api/abouts
```

[pkg1]: https://pkg.go.dev/github.com/go-playground/validator/v10
[pkg2]: https://pkg.go.dev/github.com/gofiber/fiber@v1.14.6
[pkg3]: https://pkg.go.dev/github.com/gofiber/fiber/v2@v2.51.0
[pkg4]: https://pkg.go.dev/github.com/gofiber/jwt/v2@v2.2.7
[pkg5]: https://pkg.go.dev/github.com/golang-jwt/jwt@v3.2.2+incompatible
[pkg6]: https://pkg.go.dev/github.com/spf13/viper@v1.17.0
[pkg7]: https://pkg.go.dev/go.uber.org/zap@v1.26.0
[pkg8]: https://pkg.go.dev/golang.org/x/crypto@v0.15.0
[pkg9]: https://pkg.go.dev/gorm.io/driver/postgres@v1.5.4
[pkg10]: https://pkg.go.dev/gorm.io/gorm@v1.25.5
