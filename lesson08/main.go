package  main

import (
	"fmt"
	"html/template"
	"net/http"
)


func index(w http.ResponseWriter,r *http.Request){
	//定义模板
	//解析模板,
	t,err:=template.New("index.tmpl").
		Delims("{[","]}").
		ParseFiles("./lesson08/index.tmpl")
	if err !=nil{
		fmt.Printf("parse template failed,err:%v\n",err)
		return
	}
	//渲染模板
	//渲染模板
	name:="宇文涵"
	t.Execute(w,name)
}
func xss(w http.ResponseWriter,r *http.Request){
	//定义模板
	//解析模板之前定义一个自定义的函数safe
	t,err:=template.New("xss.tmpl").Funcs(template.FuncMap{
		"safe":func(str string)template.HTML{
			return template.HTML(str)
		},
	}).ParseFiles("./lesson08/xss.tmpl")
	if err !=nil{
		fmt.Printf("parse template failed,err:%v\n",err)
		return
	}
	//渲染模板
	//渲染模板
	//str:="<script></script>"
	str1:="<a href='liwenzhou.com>李文周博客</a>"
	str2:="<a href='http://liwenzhou.com'>李文周博客</a>"
	t.Execute(w,map[string]string{
		"str1":str1,
		"str2":str2,
	})
}
func main(){

	http.HandleFunc("/index",index)
	http.HandleFunc("/xss",xss)
	err:=http.ListenAndServe(":9000",nil)
	if err !=nil{
		fmt.Printf("HTTP server start failed,err:%v",err)
		return
	}
}