package liadl

import (
	"bytes"
	"context"

	"github.com/BeanWei/li/liadl/lihcl"
	"github.com/gogf/gf/v2/os/gfile"
)

type (
	ATM struct {
		Commands map[string]string
		Pages    map[string]map[string]interface{}
	}

	Operation struct {
		Name        string `spec:",name"`
		Description string `spec:"description"`
		Command     string `spec:"command"`
	}

	Page struct {
		Path        string                 `spec:",name"`
		Description string                 `spec:"description"`
		Properties  map[string]interface{} `spec:"properties"`
	}

	// PageProperty struct {
	// 	Name           string                 `spec:",name"`
	// 	Component      string                 `spec:"component" json:"x-component"`
	// 	ComponentProps map[string]interface{} `spec:"componentProps" json:"x-component-props"`
	// 	Properties     []*PageProperty        `spec:"properties" json:"properties"`
	// }
)

var atm = new(ATM)

func atmsetup(ctx context.Context) error {
	var (
		hclb bytes.Buffer
		res  struct {
			Operations []*Operation `spce:"operation"`
			Pages      []*Page      `spce:"page"`
		}
	)
	paths, err := gfile.ScanDirFile(gfile.MainPkgPath(), "*.hcl", true)
	if err != nil {
		return err
	}
	for _, path := range paths {
		err = gfile.ReadLinesBytes(path, func(bytes []byte) error {
			_, e := hclb.Write(bytes)
			return e
		})
		if err != nil {
			return err
		}
	}

	err = lihcl.Unmarshal(hclb.Bytes(), &res)
	if err != nil {
		return err
	}

	for _, op := range res.Operations {
		atm.Commands[op.Name] = op.Command
	}
	for _, page := range res.Pages {
		atm.Pages[page.Path] = page.Properties
	}

	return nil
}
