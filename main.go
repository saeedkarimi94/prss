package main

import "github.com/einkaaf/prss/api"

func main() {
	engine := api.SetupRouters()
	engine.Run()
}
