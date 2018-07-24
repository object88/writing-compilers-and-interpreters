package cmd

import (
	"os"

	wcai "github.com/object88/writing-compilers-and-interpreters"
	"github.com/object88/writing-compilers-and-interpreters/frontend"
	"github.com/object88/writing-compilers-and-interpreters/message"
	"github.com/spf13/cobra"
)

func createCompileCommand() *cobra.Command {
	var intermedate bool
	var xref bool

	cmd := &cobra.Command{
		Use:   "compile SOURCE",
		Short: "compiles Pascal source code",
		RunE: func(cmd *cobra.Command, args []string) error {
			path := args[0]

			f, err := os.Open(path)
			if err != nil {
				return err
			}
			defer f.Close()

			mh := message.NewMessageHandler()

			s := frontend.NewSource(f, mh)
			sl := NewSourceListener()
			s.GetMessageHandler().AddListener(sl)

			p, err := wcai.NewParser("pascal", "top-down", s, mh)
			if err != nil {
				return err
			}
			pl := NewParserListener()
			p.GetMessageHandler().AddListener(pl)

			be, err := wcai.NewBackend("compile", mh)
			if err != nil {
				return err
			}

			bl := NewBackendListener()
			be.GetMessageHandler().AddListener(bl)

			err = p.Parse()
			if err != nil {
				return err
			}

			iCode := p.GetICode()
			symTab := p.GetSymTab()

			err = be.Process(iCode, symTab)
			if err != nil {
				return err
			}

			return nil
		},
	}

	cmd.Args = cobra.ExactArgs(1)

	cmd.Flags().BoolVarP(&intermedate, "intermediate", "i", intermedate, "not implemented")
	cmd.Flags().BoolVarP(&xref, "cross-reference", "x", xref, "not implemented")

	return cmd
}
