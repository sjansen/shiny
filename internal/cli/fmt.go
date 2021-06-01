package cli

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/hashicorp/hcl/v2/hclwrite"
)

type fmtCmd struct {
	Directory string `kong:"arg,name='DIR',default='.',type='existingdir'"`
}

func (cmd *fmtCmd) Run() error {
	return filepath.WalkDir(cmd.Directory, func(path string, d fs.DirEntry, err error) error {
		t := d.Type()
		switch {
		case t.IsDir() && path != cmd.Directory:
			return fs.SkipDir
		case err != nil:
			return err
		case t.IsRegular():
			n := d.Name()
			if n[0] != '.' && strings.HasSuffix(n, ".hcl") {
				data, err := os.ReadFile(path)
				if err != nil {
					return err
				}
				data = hclwrite.Format(data)
				if err := os.WriteFile(path, data, 0666); err != nil {
					return err
				}
			}
		}
		return nil
	})
}
