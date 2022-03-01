# Secure Go Fiber REST API with JWT

This repository is a simple code example for implementing [JSON Web Tokens or JWT](https://jwt.io/) in Fiber framework.

To test this API locally, you need to have [Go](https://go.dev) installed on your local machine. Then, kindly download the zip file of the source code. After extract the folder, open the terminal in that extracted folder and type

```Bash
go get -u github.com/gofiber/fiber/v2
go get -u github.com/gofiber/jwt/v3
go get -u github.com/golang-jwt/jwt/v4
```

This will install all the packages for Fiber framework and JWT

```Bash
go get -u gorm.io/gorm
go get -u gorm.io/driver/sqlite
```

This will install all the packages for GORM

To run this API, just type

```Bash
go run main.go
```

Then, the server will be listening on `http://localhost:3440`.

or, if you want to use Air packages for developing/modifying this API, then type

```Bash
go install github.com/cosmtrek/air@latest
```

Then, run this API by typing

```Bash
air
```

Then, the server will be listening on `http://localhost:3440`.

## Available API routes

`GET` : `http://localhost:3440/home`  
`POST` : `http://localhost:3440/signup`  
`POST` : `http://localhost:3440/login`  
`GET` : `http://localhost:3440/profile`  
`PUT` : `http://localhost:3440/profile`

## References

-   https://go.dev
-   https://gofiber.io
-   https://github.com/gofiber/fiber
-   https://github.com/gofiber/jwt
-   https://gorm.io
-   https://github.com/go-gorm/gorm
-   https://github.com/cosmtrek/air
