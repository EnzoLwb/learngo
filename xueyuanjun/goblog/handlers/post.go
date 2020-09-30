package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type Post struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

//获取json格式的数据
func AddPost(w http.ResponseWriter, r *http.Request) {
	len := r.ContentLength    // 获取请求实体长度
	body := make([]byte, len) // 创建存放请求实体的字节切片
	r.Body.Read(body)         // 调用 Read 方法读取请求实体并将返回内容存放到上面创建的字节切片
	//io.WriteString(w, string(body))
	post := Post{}
	json.Unmarshal(body, &post)   // 对读取的 JSON 数据进行解析
	fmt.Fprintf(w, "%#v\n", post) // 格式化输出结果
}

func GetPosts(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "All posts")
}
func EditPost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.PostForm) //只會打印post表单数据map类型 [content:[hello], title:[test]]
	//id1 := r.Form["id"] //[1]
	//id2 := r.Form.Get("id") //1

	//不需要 parseform 直接获取单个表单值的方式
	fmt.Println("post id:", r.FormValue("id"))
	fmt.Println("post title:", r.PostFormValue("title"))
	fmt.Println("post title:", r.PostFormValue("content"))
	//FormValue/PostFormValue 之所以不用显式调用 ParseForm 解析请求数据，是因为底层对其进行了封装，实际上还是要调用这个方法。
	fmt.Fprintln(w, r.Form)
}
func UploadImage(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10240 * 1000)          // 最大支持 10000KB，即 10 M
	name := r.MultipartForm.Value["name"][0]    // text
	content := r.MultipartForm.Value["content"] // text
	image := r.MultipartForm.File["image"][0]   // 文件

	fmt.Println("图片上传成功: ", name)
	fmt.Println("content: ", content[0])

	file, err := image.Open()
	if err == nil {
		data, err := ioutil.ReadAll(file) // 读取二进制文件字节流
		if err == nil {
			//fmt.Fprintln(w, string(data)) // 将读取的字节信息输出
			// 将文件存储到项目根目录下的 images 子目录
			// 从上传文件中读取文件名并获取文件后缀
			fmt.Println(image)
			names := strings.Split(image.Filename, ".")
			suffix := names[len(names)-1]
			// 将上传文件名字段值和源文件后缀拼接出新的文件名
			filename := name + "." + suffix
			fmt.Println(filename)
			// 创建这个文件
			newFile, _ := os.Create(filename)
			fmt.Println(newFile)
			defer newFile.Close()
			// 将上传文件的二进制字节信息写入新建的文件
			size, err := newFile.Write(data)
			if err == nil {
				fmt.Fprintf(w, "图片上传成功，图片大小: %d 字节\n", size/1000)
			}
		}
	}
}
