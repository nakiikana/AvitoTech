package main

import (
	api "tools/api/middleware"
	"tools/internals/server"
)

func main() {
	srv := new(server.Server)
	srv.Run(api.InitRoutes())
}
