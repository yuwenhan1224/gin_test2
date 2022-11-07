package  main

import (
	"fmt"
	"html/template"
	"net/http"
)

func f1(w http.ResponseWriter,r *http.Request){
		//解析模板第一种方法，分开写：
		//告诉模板引擎，我现在多了一个自定义的函数kua
		kua:=func(name string)(string,error){
			return  name+"好年轻",nil
		}
		t:=template.New("f.tmpl")//创建一个名字是f的模板对象,名字一定要与模板的名字能对应上
		//告诉模板引擎，我现在多了一个自定义的函数kua
		t.Funcs(template.FuncMap{
			 "kua99":kua,
		})
		_,err:=t.ParseFiles("./lesson06/f.tmpl ")
		if err !=nil{
			fmt.Printf("parse template failed,err:%v\n",err)
			return
		}
	//解析模板第二种方法，合并写：
	//t,err:=template.New("f").ParseFiles("./lesson06/f.tmpl ")
		name:="小王子"
		//渲染模板
		t.Execute(w,name)
		//2.解析模板

}
func demo1(w http.ResponseWriter,r *http.Request){
	//定义模板
	//解析模板,
	//注意，t.tmpl包含了ul.tmpl,所以先写t再写ul
   t,err:=template.ParseFiles("./lesson06/t.tmpl","./lesson06/ul.tmpl")
	if err !=nil{
		fmt.Printf("parse template failed,err:%v\n",err)
		return
	}
	//渲染模板
	name:="小王子"
	t.Execute(w,name)
}
func index(w http.ResponseWriter,r *http.Request){
	//定义模板
	//解析模板,
	//注意，t.tmpl包含了ul.tmpl,所以先写t再写ul
	t,err:=template.ParseFiles("./lesson06/index.tmpl")
	if err !=nil{
		fmt.Printf("parse template failed,err:%v\n",err)
		return
	}
	//渲染模板
	msg:="这是index页面"
	t.Execute(w,msg)
}
func home(w http.ResponseWriter,r *http.Request){
	//定义模板
	//解析模板,
	//注意，t.tmpl包含了ul.tmpl,所以先写t再写ul
	t,err:=template.ParseFiles("./lesson06/home.tmpl")
	if err !=nil{
		fmt.Printf("parse template failed,err:%v\n",err)
		return
	}
	//渲染模板
	msg:="这是小王子"
	t.Execute(w,msg)
}
func index2(w http.ResponseWriter,r *http.Request){
	//定义模板
	//解析模板,
	t,err:=template.ParseFiles("./lesson06/templates/base.tmpl","./lesson06/templates/index2.tmpl")
	if err !=nil{
		fmt.Printf("parse template failed,err:%v\n",err)
		return
	}
	//渲染模板
	//渲染模板
	name:="宇文涵"
	t.ExecuteTemplate(w,"./lesson06/templates/index2.tmpl",name)
}
func home2(w http.ResponseWriter,r *http.Request){
	//定义模板
	//解析模板,
	t,err:=template.ParseFiles("./lesson06/templates/base.tmpl","./lesson06/templates/home2.tmpl")
	if err !=nil{
		fmt.Printf("parse template failed,err:%v\n",err)
		return
	}
	//渲染模板
	//渲染模板
	name:="homeeee"
	t.ExecuteTemplate(w,"./lesson06/templates/home2.tmpl",name)

}
func main(){
	http.HandleFunc("/",f1)
	http.HandleFunc("/tmplDemo",demo1)
	http.HandleFunc("/home",home)
	http.HandleFunc("/index",index)
	http.HandleFunc("/home2",home)
	http.HandleFunc("/index2",index)
	err:=http.ListenAndServe(":9000",nil)
	if err !=nil{
		fmt.Printf("HTTP server start failed,err:%v",err)
		return
	}
}