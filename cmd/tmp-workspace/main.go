package main

import (
	"fmt"

	"github.com/IridiumNan/tmp-workspace/internal/config"
	"github.com/IridiumNan/tmp-workspace/internal/utils"
)

func main() {
	err := config.LoadConfig()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(config.Cfg)

	for i := 1; i < 10; i++ {
		fmt.Println(utils.GetWorkSpaceID())
	}
}
