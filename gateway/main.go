package main

import (
	"gateway/pkg/manager"

	"github.com/gookit/slog"
)

func main() {

	slog.Info("starting the kube-gateway 🐝")
	c, err := manager.Setup()
	if err != nil {
		slog.Fatal(err)
	}
	err = manager.LoadEPF(c)
	if err != nil {
		slog.Fatal(err)
	}
	err = manager.Start(c)
	if err != nil {
		slog.Fatal(err)
	}
}
