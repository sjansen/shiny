package cli

import "fmt"

type generateCmd struct {
	Directory string `kong:"arg,name='DIR',default='.',type='existingdir'"`
}

func (cmd *generateCmd) Run() error {
	fmt.Println(cmd.Directory)
	return nil
}
