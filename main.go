// Copyright 2017 Johan Brandhorst. All Rights Reserved.
// See LICENSE for licensing terms.

package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
	usersv1 "github.com/johanbrandhorst/grpc-web-go-react-example/gen/users/v1"
	"github.com/johanbrandhorst/grpc-web-go-react-example/users"
	"google.golang.org/grpc"
)

//go:embed dist
var frontend embed.FS

func main() {
	gs := grpc.NewServer()
	usersv1.RegisterUserServiceServer(gs, &users.UserService{})
	wrappedServer := grpcweb.WrapServer(gs)

	http.Handle("/api/", http.StripPrefix("/api/", wrappedServer))
	distFS, err := fs.Sub(frontend, "dist")
	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/", http.FileServer(http.FS(distFS)))

	log.Println("Serving on http://0.0.0.0:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
