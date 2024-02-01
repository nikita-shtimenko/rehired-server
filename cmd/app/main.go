package main

import "github.com/nikita-shtimenko/rehired-server/internal/app"

func main() {
	app := app.New()
	app.Run()
}
