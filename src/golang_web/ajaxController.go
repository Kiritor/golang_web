package main

import (
	"net/http"
    //"github.com/ziutek/mymysql/mysql"
	_ "github.com/ziutek/mymysql/thrsafe"
	"encoding/json"
	"log"
)

type Result struct {
	Ret int            //结果状态码:0/1
	Reason string      //描述信息
	Data interface{}   //数据
}

type ajaxController struct {

}

/**
   处理登录请求
*/
func (controller *ajaxController) LoginAction(w http.ResponseWriter,r *http.Request) {
	w.Header().Set("content-type","application/json")  //设置数据格式为json
	err :=r.ParseForm()  //解析post请求
	if err!=nil {
		OutputJson(w,0,"参数错误",nil)
		return
	}
	username :=r.FormValue("username")
	password:=r.FormValue("password")
	if username == "" || password == "" {
		OutputJson(w,0,"参数错误",nil)
	}

	db:=NewDataMng()
	defer db.Close()
	rows,res,err:= db.Query("select * from user where username='%s'",username)
	if err !=nil {
		log.Println(err)
		OutputJson(w,0,"数据库操作失败",nil)
		return
	}

	name:=res.Map("password")
	log.Println(username)
	password_db :=rows[0].Str(name)
	if password !=password_db {
		OutputJson(w, 0, "密码输入错误", nil)
        return
	}

	//存入cookie

	cookie :=http.Cookie{Name:"username",Value:rows[0].Str(res.Map("username")),Path:"/"}
	http.SetCookie(w,&cookie)
	OutputJson(w, 1, "操作成功", nil)
    return
}

/**
  输出json
*/
func OutputJson(w http.ResponseWriter, ret int, reason string, i interface{}) {
    out := &Result{ret, reason, i}
    b, err := json.Marshal(out)
    if err != nil {
        return
    }
    w.Write(b)
}
/**
   得到所有用户
*/
func (controller *ajaxController) GetUsersAction(w http.ResponseWriter,r *http.Request) {
	w.Header().Set("content-type","application/json")  //设置数据格式为json
	err :=r.ParseForm()  //解析post请求
	if err!=nil {
		OutputJson(w,0,"参数错误",nil)
		return
	}

	db:=NewDataMng()
	if err :=db.Connect();err !=nil {
		log.Println(err)
		OutputJson(w,0,"数据库操作失败",nil)
		return
	}
	defer db.Close()
	rows,res,err:= db.Query("select * from user where username='%s'","LCore")
	if err !=nil {
		log.Println(err)
		OutputJson(w,0,"数据库操作失败",nil)
		return
	}
	OutputJson(w, 1, "操作成功", rows[0].Str(res.Map("username")))
    return
}
