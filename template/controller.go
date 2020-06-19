package template

import (
	"fmt"
	"html/template"
	"net/http"
)

type Content map[string]interface{}

func LayoutTpl() {
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		t, err := setAdminContentFile("template/view/users.html")
		if err != nil {
			fmt.Println("parse file err:", err.Error())
			return
		}
		p := Content{"title": "admin-users", "users": "layouts - template - users"}
		_ = t.Execute(w, p)
	})
	http.HandleFunc("/news", func(w http.ResponseWriter, r *http.Request) {
		t, err := setAdminContentFile("template/view/news.html")
		if err != nil {
			fmt.Println("parse file err:", err.Error())
			return
		}
		p := Content{"title": "admin-news", "news": "layouts - template - news"}
		_ = t.Execute(w, p)
	})
	_ = http.ListenAndServe("127.0.0.1:8080", nil)
}

func setAdminContentFile(contentFileName string) (*template.Template, error) {
	files := []string{
		"template/view/template_admin.html",
		"template/view/layouts/header.html",
		"template/view/layouts/sidebar.html",
		contentFileName,
		"template/view/layouts/footer.html",
	}
	return template.ParseFiles(files...)
}

