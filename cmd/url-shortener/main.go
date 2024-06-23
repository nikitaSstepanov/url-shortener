package main

import "github.com/nikitaSstepanov/url-shortener/internal/app"

func main() {
	a := app.New()

	a.Run()
}