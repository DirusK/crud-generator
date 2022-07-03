package main

import "crud-generator/internal/app"

const appName = "CRUD generator"

func main() {
	app.New(appName).Run()
}
