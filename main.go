package main

import (
	"packer-plugin-shell/shell"

	"github.com/hashicorp/packer/packer/plugin"
)

func main() {
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}

	p := new(shell.Provisioner)
	server.RegisterProvisioner(p)
	server.Serve()
}
