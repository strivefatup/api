# Gin Web RESTful API Example
this example is based on [gin](https://github.com/gin-gonic/gin) and [gorm](https://github.com/jinzhu/gorm). Gin is a web framework written in Go (Golang), Gorm is ORM library for Golang. This example, there are still many shortcomings, to be improved, welcome to correct.

## Installation

To install API Example, you need to install Go and set your Go workspace first.

1. Download and [install](https://learnku.com/docs/build-web-application-with-golang/011-installs-go/3153) it

2. Test whether the installation is successful:

```sh
$ go 
$ go env 
```
3. Via git install:

```sh
$ git clone https://github.com/strivefatup/api.git
```

## Example Directory Structure

The application directory structure is as follows

```git
api                                                Application root directory
├── README.md
├── config                                         Configuration file directory
│   ├── config.go                                  Config helper function file
│   ├── db.yml                                     Databases configuration file
│   └── pagination.yml                             Pagination configuration file
├── http                       
│   ├── Middleware                                 Customize middleware directory
│   └── controller                                 Controller directory
│       ├── baseController.go                      Base Controller
│       ├── order                                  
│       │   └── order.go                           
│       └── user                                     
├── main.go                                        Entry file
├── models                                         models
│   ├── BaseModel.go                               Base Model
│   └── order                                                     
│       └── order.go                               
├── resources                                      Static resource directory
│   └── log                                        Save log file directory
├── routes                                         Route directory
│   ├── order                                      
│   │   └── order.go                               
│   ├── route.go                                   
│   └── user                                       
├── serveres                                       Server layer
│   └── order                                      
│       └── order.go                               
├── tools                                          tool directory
│   ├── databases
│   │   ├── connPool.go
│   │   └── mysql
│   │       └── mysqlConnPool.go
│   └── http
│       └── response
│           └── response_list_struct.go
└── vendor
    └── vendor.json
```

## Install A Third-party Package

use govendor install third-party package

1. install govendor

```sh
$ go get -u -v github.com/kardianos/govendor
```

2. Install according to the vendor.json file

```sh
$ govendor sync
```

## API Examples
 
### Using GET, POST, PUT and DELETE

/routes/order/order.go file 

```go
package order

import (
	"github.com/gin-gonic/gin"
	"api/http/controller/order"
)

func Order(engine *gin.Engine){
    route := engine.Group("/api")
    {
        route.GET("/order", order.GetOrderList)
        route.GET("/order/:id", order.GetOrderById)
	route.POST("/order", order.CreateOrder)
        route.PUT("/order/:id", order.UpdateOrder)
        route.DELETE("/order/:id", order.DeleteOrderById)
    }
}
```
## Run This Example 

Default Port is 8080

```sh
$ go run main.go
```

## Hint

Config function helper example

```go
package main

import "api/config"

func main() {
    config.Config("db.mysql.connection.password")
}
```


