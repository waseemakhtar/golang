package main

import (
	"fmt"
)

type Vims []string

type VimsIntf interface {
	Vims() []string
	SetVims(vims []string)
}

//Vims fetches vims
func (v *Vims) Vims() Vims {
	vims := *(v)
	fmt.Printf("In Vims: %#v\n", vims)
	return vims
}

//SetVims sets vims
func (v *Vims) SetVims(vims Vims) {
	fmt.Printf("In SetVims: %#v\n", vims)
	*v = vims
}

func main() {
	vims := &Vims{}
	fmt.Printf("%#v\n", vims)

	abcVim := Vims{"abc"}

	vims.SetVims(abcVim)
	fmt.Printf("After-1: %#v\n", vims.Vims())
	//vnfDB.SetVnfRequest(instantiateVnfReq)
	//fmt.Printf("After-2 %#v\n", vnfDB)
}
