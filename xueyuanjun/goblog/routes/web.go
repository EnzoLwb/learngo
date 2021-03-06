package routes

import (
	"learngo/xueyuanjun/goblog/handlers"
	"net/http"
)

// 定义一个 WebRoute 结构体用于存放单个路由
type WebRoute struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// 声明 WebRoutes 切片存放所有 Web 路由
type WebRoutes []WebRoute

// 定义所有 Web 路由
var webRoutes = WebRoutes{
	WebRoute{
		"Home",
		"GET",
		"/",
		handlers.Home,
	},
	WebRoute{
		"Posts",
		"GET",
		"/posts",
		handlers.GetPosts,
	},
	WebRoute{
		"User",
		"GET",
		"/user/{id}",
		handlers.GetUser,
	},
	WebRoute{
		"UpdatePost",
		"POST",
		"/post/edit",
		handlers.EditPost,
	},
	WebRoute{
		"AddPost",
		"POST",
		"/post/add",
		handlers.AddPost,
	},
	WebRoute{
		"UploadImage",
		"POST",
		"/image/upload",
		handlers.UploadImage,
	},
	WebRoute{
		"Redirect",
		"GET",
		"/redirect",
		handlers.Redirect,
	},
	WebRoute{
		"Session",
		"GET",
		"/count",
		handlers.Counter,
	},
}
