package main

import (
	"fmt"
	"wxm"
)

func main() {
	fmt.Println("test:")
	test()
}

func test() {
	name := "80010121"
	pwd := "888889"
	txt1 := "a"
	txt2 := "b"
	txt3 := "c"

	wxm.Init(name, pwd)
	result, err := wxm.SendMsg(txt1, txt2, txt3)
	fmt.Println("wxm.SendMsg:", result, err)

	result, err = wxm.SendMsgToUser(name, pwd, txt1, txt2, txt3)
	fmt.Println("wxm.SendMsg:", result, err)

}
