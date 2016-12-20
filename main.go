package main

import (
	"os"

	"github.com/teambition/gear"
	"github.com/teambition/gear/middleware"
)

func main() {
	app := gear.New()
	// open log file
	f, err := os.OpenFile("/data/log/k8s-demo/output.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND|os.O_SYNC, 0660)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	// logger middleware
	logger := &middleware.DefaultLogger{
		W: f,
	}
	app.Use(middleware.NewLogger(logger))

	demoRouter := gear.NewRouter()
	// demo json data
	demoData := struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}{
		Code:    0,
		Message: "hello world!",
	}
	// GET /
	demoRouter.Get("/", func(c *gear.Context) error {
		return c.JSON(200, demoData)
	})
	app.UseHandler(demoRouter)
	app.Error(app.Listen(":8080"))
}
