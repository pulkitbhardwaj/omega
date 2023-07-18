/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/pulkitbhardwaj/omega/pkgs/ats"
	"go.uber.org/fx"
)

func main() {
	fx.New(ats.Module).Run()
}
