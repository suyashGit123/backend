package main

import (
	"github.com/suyashGit123/backend/server"
)

func main() {
	s := server.NewServer()

	if err := s.Init(3000); err != nil {
		panic(err)
	}

	s.Start()

}
