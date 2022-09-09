package terraform

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/hashicorp/terraform-config-inspect/tfconfig"
)

type Module struct {
	M *tfconfig.Module
}

// CollectModules returns all valid Terraform modules recursively found in the
// path specified.
func CollectModules(root string) ([]Module, error) {
	var results []Module

	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return fs.SkipDir
		}

		if strings.Contains(path, ".terraform") {
			return fs.SkipDir
		}

		if d.IsDir() && tfconfig.IsModuleDir(path) {
			module, diags := tfconfig.LoadModule(path)

			if diags != nil {
				fmt.Printf("Warning: cannot load terraform module '%s': %v\n", path, diags)
				return fs.SkipDir
			}

			results = append(results, Module{M: module})
		}

		return nil
	})

	if err != nil {
		return []Module{}, err
	}

	return results, nil
}
