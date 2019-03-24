//fetchall 并发获取URL并报告它们时间和大小
package main
import (
	 "fmt"
	 "io"
	 "io/ioutil"
	 "net/http"
	 "os"
	 "time"
)

func mainFetchall()  {
	start:=time.Now()
	//创建字符串通道
	ch:=make(chan string)
	//用for循环来启动一个新的goroutine 它异步调用fetch来使用http.get获取URL内容
	for _, url:=range os.Args[1:]{
		go fetch(url,ch)
	}
	for range os.Args[1:]{
        fmt.Println(<-ch)//
	}
    fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string,ch chan<- string){
	start:=time.Now()
	resp,err :=http.Get(url)
	if err != nil {
		ch<-fmt.Sprint(err)//发送到通道ch  ch????
		return
	}

	nbytes,err :=io.Copy(ioutil.Discard,resp.Body)
	resp.Body.Close()//不要泄漏资源
	if err != nil {
		ch<-fmt.Sprintf("while reading %s:%v",url,err)
		return 
	}

	secs:=time.Since(start).Seconds()
	ch<- fmt.Sprintf("%.2fs %7d %s",secs,nbytes,url)
}