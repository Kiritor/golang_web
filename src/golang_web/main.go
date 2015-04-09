package main

import (
    "net/http"
    "log"
)

func main() {
	log.Println("服务器已启动")
	//设置静态资源
    http.Handle("/css/", http.FileServer(http.Dir("template")))
    http.Handle("/js/", http.FileServer(http.Dir("template")))
	http.Handle("/img/login/", http.FileServer(http.Dir("template")))
    //设置路由
    http.HandleFunc("/admin/", adminHandler)
    http.HandleFunc("/login/",urlHandler)
    http.HandleFunc("/ajax/",ajaxHandler)
	http.HandleFunc("/getUsers/",ajaxHandler)
    http.HandleFunc("/",NotFoundHandler)
    err:=http.ListenAndServe(":8888", nil)    //设置监听的接口
	if err !=nil {
		log.Println(err)
	}
}
