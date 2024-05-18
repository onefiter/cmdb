package main

import (
	"fmt"

	"github.com/onefier/cmdb/cli"
)

func main() {
	if err := cli.RootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
