//迷你回声服务器
package main

import (
	"fmt"
	"log"
	"net/http"
)

func  main0()  {
	  http.HandleFunc("/",handler)//设置访问路由
	  err:=http.ListenAndServe(":8000",nil)//设置监听的端口
       if err!=nil{
		   log.Fatal("ListenAndServe:",err)
	   }	
}

//处理程序回显请求 URL r路径部分
func handler(w http.ResponseWriter, r *http.Request){
  //http 响应片段  
      fmt.Fprintf(w,"URL.Path=%q\n",r.URL.Path)
}