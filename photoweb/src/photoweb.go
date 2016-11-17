package main

import (
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"runtime/debug"
)

// 定义常量
const (
	UploadDir   = "./uploads"
	TemplateDir = "./views"
	ListDir     = 0x0001
)

// 定义模板缓存容器
var (
	templates = make(map[string]*template.Template)
)

// 此函数会在 main 函数之前执行
func init() {
	// 读取 views 目录并获取里面的模板
	fileInfoArr, err := ioutil.ReadDir(TemplateDir)
	if err != nil {
		// 抛出异常
		panic(err)
	}

	var templateName, templatePath string
	// 循环读取目录下的文件
	for _, fileInfo := range fileInfoArr {
		templateName = fileInfo.Name()
		// 获取后缀为 html 的文件
		if ext := path.Ext(templateName); ext != ".html" {
			continue
		}

		// 构造模板路径
		templatePath = TemplateDir + "/" + templateName
		log.Println("Loading template: ", templatePath)
		// 读取模板
		t := template.Must(template.ParseFiles(templatePath))
		// 缓存模板到变量中
		templates[templateName] = t
	}
}

// 主函数
func main() {
	mux := http.NewServeMux()
	staticDirHandler(mux, "/assets/", "./public", 0)

	mux.HandleFunc("/", safeHandler(listHandler))
	mux.HandleFunc("/upload", safeHandler(uploadHandler))
	mux.HandleFunc("/view", safeHandler(viewHandler))

	log.Println("Server run at localhost:8080...")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}

// 上传回调函数
func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		check(renderHTML(w, "upload", nil))
	}

	if r.Method == "POST" {
		f, h, err := r.FormFile("image")
		check(err)

		defer f.Close()
		filename := h.Filename

		t, err := os.Create(UploadDir + "/" + filename)
		check(err)

		defer t.Close()

		_, err = io.Copy(t, f)
		check(err)

		http.Redirect(w, r, "/view?id="+filename, http.StatusFound)
	}
}

// 查看回调函数
func viewHandler(w http.ResponseWriter, r *http.Request) {
	imageID := r.FormValue("id")
	imagePath := UploadDir + "/" + imageID

	if _, err := os.Stat(imagePath); err != nil {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "image")
	http.ServeFile(w, r, imagePath)
}

// 列表回调函数
func listHandler(w http.ResponseWriter, r *http.Request) {
	fileInfoArr, err := ioutil.ReadDir(UploadDir)
	check(err)

	locals := make(map[string]interface{})
	images := []string{}
	for _, fileInfo := range fileInfoArr {
		images = append(images, fileInfo.Name())
	}
	locals["images"] = images

	check(renderHTML(w, "list", locals))
}

// 静态文件服务
func staticDirHandler(mux *http.ServeMux, prefix string, staticDir string, flags int) {
	mux.HandleFunc(prefix, func(w http.ResponseWriter, r *http.Request) {
		// 从请求地址中得到本地文件地址
		file := staticDir + r.URL.Path[len(prefix)-1:]
		if (flags & ListDir) == 0 {
			if _, err := os.Stat(file); err != nil {
				http.NotFound(w, r)
				return
			}
		}
		http.ServeFile(w, r, file)
	})
}

// 渲染模板函数
func renderHTML(w http.ResponseWriter, tmpl string, locals map[string]interface{}) (err error) {
	err = templates[tmpl+".html"].Execute(w, locals)
	return
}

// 错误检查函数
func check(err error) {
	if err != nil {
		panic(err)
	}
}

// 错误处理函数
func safeHandler(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err, ok := recover().(error); ok {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				log.Printf("WARN: panic in %v - %v\n", fn, err)
				log.Println(string(debug.Stack()))
			}
		}()

		fn(w, r)
	}
}
