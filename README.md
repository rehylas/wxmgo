# wxmgo: wx message for go
发送微信信息到自己手机

# install
go get github.com/rehylas/wxmgo

#   github
github.com/rehylas/wxmgo

#   usage

    ```
        package main

        import (
            "fmt"

            wxm "github.com/rehylas/wxmgo"
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

    ```    

效果
![Image text](http://www.bangnikanzhe.com/img/wxm.jpg)

#  如何获得 name 和 pwd  , 扫码并关注, 会自动获得 name 和 pwd：
![Image text](http://www.bangnikanzhe.com/img/bnkz.jpg)




