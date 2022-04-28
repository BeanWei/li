//go:build ignore
// +build ignore

package main

// package ent

import (
	"log"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"github.com/BeanWei/li/li-engine/contrib/lient"
)

func main() {
	err := entc.Generate("./schema",
		&gen.Config{},
		entc.Extensions(&lient.Extension{}),
	)
	if err != nil {
		log.Fatal("running ent codegen:", err)
	}
}
