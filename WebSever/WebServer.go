package main

import (
    "io"
    "net/http"
)

//当使用字符串常量表示html文本时，包含<html><body></body></html>对于让浏览器识别它收到了一个html非常重要。
const form = `
    <html><body>
        <form action="#" method="post" name="bar">
            <input type="text" name="in" />
            <input type="submit" value="submit"/>
        </form>
    </body></html>
`

/* handle a simple get request */
func SimpleServer(w http.ResponseWriter, request *http.Request) {
    io.WriteString(w, "<h1>hello, world</h1>")
}

func FormServer(w http.ResponseWriter, request *http.Request) {
	//在处理器中使用w.Header().Set("Content-Type", "text/html")在写入返回之前将header的content-type设置为text/html
	//content-type会让浏览器认为它可以使用函数http.DetectContentType([]byte(form))来处理收到的数据
	w.Header().Set("Content-Type", "text/html")
    //1.url最初由浏览器请求，那么它是一个GET请求，并且返回form常量，包含简单的input表单，这个表单里有一个文本框和一个提交按钮.
	//2.当文本框输入一些东西并点击提交按钮时候，会发起post请求。request.FormValue("in")获取内容，并写回浏览器页面.
    switch request.Method {
    case "GET":
        /* display the form to the user */
        io.WriteString(w, form)
    case "POST":
        /* handle the form data, note that ParseForm must
           be called before we can extract form data */
        //request.ParseForm();
        //io.WriteString(w, request.Form["in"][0])
        io.WriteString(w, request.FormValue("in"))
    }
}

func main() {
    http.HandleFunc("/test1", SimpleServer)//simpleServer处理test1上的请求
	http.HandleFunc("/test2", FormServer)//FormServer 处理test2上的请求
    if err := http.ListenAndServe(":8088", nil); err != nil {
        panic(err)
    }
}