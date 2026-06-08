package main

import "clms_website/bootstrap"

func main() {
	app := bootstrap.NewApplication()
	app.Serve()
}
