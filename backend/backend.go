package backend

import "github.com/object88/writing-compilers-and-interpreters/intermediate"

type BaseBackend struct {
	symTab *intermediate.SymTab
	iCode  *intermediate.ICode
}
