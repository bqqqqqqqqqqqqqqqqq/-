package main

import "dogking_shop/router"

func main() {
	r := router.Router()
	r.Run(":8888")
}
