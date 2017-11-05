package main

import (
    "github.com/kataras/iris"
    "os"
)

func main() {
	port := os.Getenv("PORT")

	app := iris.New()
    api := app.Party("/api")
    {
    	api.Get("/commingsoon/{index:int min(0)}", specificCommingsoonHandler)
    	api.Get("/commingsoon", commingsoonHandler)
    	api.Get("/theaters", theatersHandler)
    }

    app.Run(iris.Addr(":" + port))
}
