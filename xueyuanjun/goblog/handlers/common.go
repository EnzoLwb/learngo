package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"text/template"

	"github.com/gorilla/sessions"
)

type Greeting struct {
	Message string `json:"message"`
}

func Home(w http.ResponseWriter, r *http.Request) {
	//io.WriteString(w, "Welcome to my blog site")
	w.Write([]byte("欢迎访问学院君个人网站👏"))
	// 返回 JSON 格式数据
	greeting := Greeting{
		"欢迎访问学院君个人网站👏",
	}
	w.Header().Set("Content-Type", "application/json")
	message, _ := json.Marshal(greeting)
	w.Write(message)
}
func Redirect(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "www.baidu.com")
	w.WriteHeader(301) //writeheader 要写在header头之后
}

// 初始化存储器（基于 Cookie）
var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY)")))

func Counter(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "GOSESSID")
	count := session.Values["count"] //session.Values 是一个字典结构（map[interface{}]interface{}），
	if count == nil {
		session.Values["count"] = 1
	} else {
		session.Values["count"] = count.(int) + 1
	}
	err := session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t, _ := template.ParseFiles("counter.gtpl")
	w.Header().Set("Content-Type", "text/html")
	t.Execute(w, session.Values["count"])
}
