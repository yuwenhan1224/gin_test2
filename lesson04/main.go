package main

import (
	"fmt"
	"html/template"
	"net/http"
)
//遇事不决 写注释
func sayHello(w http.ResponseWriter,r *http.Request){
     //2.解析模板
    t,err:=template.ParseFiles("./04hello.tmpl")
	if err!=nil{
		fmt.Println("parse template failed,err:%v",err)

	}
	//3.渲染模板
   t.Execute(w,"我喜欢go")
	if err!=nil{
		fmt.Println("render template failed,err:%v",err)
		return
	}
}
func main(){
  http.HandleFunc("/",sayHello)
  err:=http.ListenAndServe(":8088",nil)
  if err !=nil{
	  fmt.Println("http server start failed ,err:%v",err)
	  return
  }
}