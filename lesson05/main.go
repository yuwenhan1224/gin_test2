package main

import (
	"fmt"
	"html/template"
	"net/http"
)
type User struct {
	Name string
	Gender string
	Age int
}
//遇事不决 写注释
func sayHello(w http.ResponseWriter,r *http.Request){
	//2.解析模板
	t,err:=template.ParseFiles("./lesson05/05hello.tmpl")
	if err!=nil{
		fmt.Println("parse template failed,err:%v",err)

	}
	//3.渲染模板,数据填充
	//t.Execute(w,"我喜欢go")
	//if err!=nil{
	//	fmt.Println("render template failed,err:%v",err)
	//	return
	//}
	u1:=User{
		Name:"汪圆圆",
		Gender: "女",//注意这里gender的首字母是小写，小写不可以被外部访问，Gender可被外部访问
		Age: 18,
	}
	m1:=map[string]interface{}{
		"name":"汪圆圆",
		"gender": "女",//注意这里gender的首字母是小写，小写不可以被外部访问，Gender可被外部访问
		"age": 18,
	}
	//使用map可以传递多个变量
	hobbyList:=[]string{
		"篮球",
		"足球",
		"双色球",
	}
	err=t.Execute(w,map[string]interface{}{
		"u1":u1,
		"m1":m1,
		"hobby":hobbyList,
	})//05lesson.tmpl里面的点.就表示u1,.Name就可以访问Name字段

	if err !=nil{
		fmt.Println("render template failed,err:%v",err)
		return
	}
}
func main(){
	http.HandleFunc("/",sayHello)
	err:=http.ListenAndServe(":9000",nil)
	if err !=nil{
		fmt.Println("http server start failed ,err:%v",err)
		return
	}
}