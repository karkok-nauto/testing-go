package main

import "github.com/hoisie/web"

func main() {
	web.Get("/(.*)", hello)
	web.Run("0.0.0.0:9999")
}
