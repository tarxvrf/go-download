package main

import (
	"go-download/databases"
	"go-download/routers"
)

func main() {
	databases.Konekdb()
	routers.Routerx()
}
