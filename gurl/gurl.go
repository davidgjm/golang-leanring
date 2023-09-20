package main

import (
	"log/slog"
	"os"
)

func main() {

	slog.Info("hello,world", "user", os.Getenv("USER"))
}
