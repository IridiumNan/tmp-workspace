package main

import (
	"fmt"

	"github.com/IridiumNan/tmp-workspace/internal/config"
)

func main() {
	err := config.LoadConfig()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(config.Cfg)
}
