package main

import (
	"fmt"
	"reflect"
)
type abc struct {
	a string
}

func (s *abc) seta(input string) {
	s.a = input
}

type abcIntf interface {
	seta(input string)
}

type FlvDetailsMap map[string]*abc

type FlvDetailsMapIntf interface {
	GetByKey(k string) *abc
	SetByKey(string, *abc)
	PrintMap()
}

/*
func (f *FlvDetailsMap) GetByKey(k string) interface{} {
	var v interface{}

	fMap := *(f)
	if val, present := fMap[k]; present {
		return val
	}
	return v
}

func (f *FlvDetailsMap) SetByKey(k string, v interface{}) {
	fMap := *(f)
	fMap[k] = v.(*abc)
}

func (f *FlvDetailsMap) PrintMap() {
	for _, item := range *(f) {
		fmt.Printf("PrintMap: %#v\n", item)
	}
}

 */


func (f FlvDetailsMap) GetByKey(k string) (v interface{}) {

	if val, present := f[k]; present {
		return val
	}
	return v
}

func (f FlvDetailsMap) SetByKey(k string, v *abc) {
	f[k] = v
}

func (f FlvDetailsMap) PrintMap() {
	fmt.Printf("Print Whole Map:%s\n", f)
	for _, item := range f {
		fmt.Printf("PrintMap: %#v\n", item)
	}
}

func main() {
	flv := &FlvDetailsMap{}
	fmt.Printf("%#v\n", flv)

	flv.SetByKey("first", &abc{a:"waseem"})
	flv.SetByKey("second", &abc{a:"akhtar"})
	abc2 := flv.GetByKey("first")
	fmt.Println(reflect.ValueOf(abc2))
	abc2.(abcIntf).seta("husna")
	fmt.Println(reflect.ValueOf(abc2))
	abc3:= flv.GetByKey("second")
	fmt.Println(reflect.ValueOf(abc3))
	fmt.Printf("abc2: %#v\n", abc2)
	fmt.Printf("abc2: %#v\n", abc3)
	flv.PrintMap()

}
