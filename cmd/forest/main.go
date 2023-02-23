package main

import (
	"fmt"
	"os"

	"github.com/greenchainearth/go-forest/cmd/forest/launcher"
)

func main() {
	if err := launcher.Launch(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
