package main

import (
	"fmt"
	"html/template"
	"net/http"
	//ファイル分割したモジュール
	//DB関係
	"wordapp/db"
)

// global変数
var user_id int

// ユーザー登録画面
func HandlerUserForm(w http.ResponseWriter, r *http.Request) {
	// テンプレートをパースする
	tpl := template.Must(template.ParseFiles("templates/user-form.html"))
	// テンプレートに出力する値をマップにセット
	values := map[string]string{}
	// マップを展開してテンプレートを出力する
	if err := tpl.ExecuteTemplate(w, "user-form.html", values); err != nil {
		fmt.Println(err)
	}
}

// ユーザー登録の確認画面
func HandlerUserConfirm(w http.ResponseWriter, req *http.Request) {
	// テンプレートをパースする
	tpl := template.Must(template.ParseFiles("templates/user-confirm.html"))
	// テンプレートに出力する値をマップにセット
	values := map[string]string{
		"account": req.FormValue("account"),
		"passwd":  req.FormValue("passwd"),
	}
	fmt.Println("登録", values["account"]) //登録した名前が出る
	//dbに登録
	db.User_db(values["account"], values["passwd"])
	// マップを展開してテンプレートを出力する
	if err := tpl.ExecuteTemplate(w, "user-confirm.html", values); err != nil {
		fmt.Println(err)
	}
}

// ログイン画面
func HandlerUserFormLogin(w http.ResponseWriter, r *http.Request) {
	// テンプレートをパースする
	tpl := template.Must(template.ParseFiles("templates/user-login.html"))
	// テンプレートに出力する値をマップにセット
	values := map[string]string{}
	// マップを展開してテンプレートを出力する
	if err := tpl.ExecuteTemplate(w, "user-login.html", values); err != nil {
		fmt.Println(err)
	}
}

// ログイン入力内容の確認画面
func HandlerUserConfirmLogin(w http.ResponseWriter, req *http.Request) {
	// テンプレートをパースする
	tpl := template.Must(template.ParseFiles("templates/user-confirm-login.html"))
	// テンプレートに出力する値をマップにセット
	values := map[string]string{
		"account": req.FormValue("account"),
		"passwd":  req.FormValue("passwd"),
	}
	fmt.Printf("ログインしたユーザ名: %s\n", values["account"])
	//ログイン
	//ユーザーIDの取得
	user_id = db.User_login(values["account"], values["passwd"])
	fmt.Printf("ログインuser_id %d\n", user_id)
	if user_id > 0 {
		// マップを展開してテンプレートを出力する
		if err := tpl.ExecuteTemplate(w, "user-confirm-login.html", values); err != nil {
			fmt.Println(err)
		}
	} else { //ユーザー名またはパスワードを間違えるとuser_idに'0'がはいる
		fmt.Fprintf(w, "入力が間違っています")
	}
}


func main() {
	//ユーザー登録
	http.HandleFunc("/user-form", HandlerUserForm)
	//ユーザー登録確認
	http.HandleFunc("/user-confirm", HandlerUserConfirm)
	//ログイン
	http.HandleFunc("/user-login", HandlerUserFormLogin)
	//ログイン確認
	http.HandleFunc("/user-confirm-login", HandlerUserConfirmLogin)


	fmt.Println("Server Start Up........")

	// サーバーを起動
	http.ListenAndServe(":8080", nil)
}
