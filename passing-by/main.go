package main

type BigStruct struct {
	buf [1 << 16]byte
}

var obj BigStruct = BigStruct{}

func passByValue(obj BigStruct) {}

func passByReference(obj *BigStruct) {}

func main() {
	passByValue(obj)
	passByReference(&obj)
}
