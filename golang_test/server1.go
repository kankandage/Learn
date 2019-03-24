//server2 迷你的回声和计数器服务器
package main

import(
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex//？？？
var count int

func  main1()  {
	http.HandleFunc("/",handler)
	http.HandleFunc("/count",counter)
	err:=http.ListenAndServe(":9090",nil)//设置监听的端口
       if err!=nil{
		   log.Fatal("ListenAndServe:",err)
	   }
}

//处理程序回显请求的URL的路径部分
func handler1(w http.ResponseWriter, r *http.Request)  {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w,"URL.Path=%q\n",r.URL.Path)
}
//counter回显目前为止的调用次数
func counter(w http.ResponseWriter, r *http.Request)  {
	mu.Lock()
	fmt.Fprintf(w,"Count %d\n",count)
	mu.Unlock()
}