package main

import "github.com/gopherjs/gopherjs/js"

func main() {
	js.Global.Get("window").Set("onload", func() {
		go func() {
			body := js.Global.Get("document").Get("body")
			body.Get("style").Set("background-color", "red")

			img := js.Global.Get("document").Call("createElement", "img")
			img.Set("src", "gopher.png")

			body.Call("appendChild", img)
		}()
	})
}
