package main

import (
	"Assignment_Vivasoft/package/containers"
	"github.com/labstack/echo/v4"
)

func main(){
	e:= echo.New()
	containers.Serve(e)
}