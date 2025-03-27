package main

import (
	"fmt"
	"github.com/astaxie/beeku"
	"github.com/peaceLT/build-web-application-with-golang/helper"
	"html/template"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"time"
	"unicode"
)

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("./htmlcode/login.html")
		log.Println(t.Execute(w, nil))
	} else {
		r.ParseForm()
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])

		// 判断整数
		getint, err := strconv.Atoi(r.Form.Get("age"))
		if err != nil {
			//	非整数
		}
		if getint > 100 {

		}

		// 判断全是中文-方法一
		username := r.Form.Get("username")
		for _, r := range username {
			if !unicode.Is(unicode.Han, r) {
				//	包含非中文
			}
		}
		// 全是中文判断-方法二
		if m, _ := regexp.MatchString("^\\p{Han}+$", username); !m {
			// 不满足全是中文
		}

		// 判断全是英文
		password := r.Form.Get("password")
		if m, _ := regexp.MatchString("^[a-zA-Z]+$", password); !m {
			// 不满足
		}

		// 判断电子邮件
		email := r.Form.Get("email")
		if m, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,})\.([a-z]{2,4})$`, email); !m {
			// 不是电子邮件格式
		}

		// 判断手机号码
		phone := r.Form.Get("phone")
		if m, _ := regexp.MatchString(`^(1[3|4|5|8][0-9]\d{4,8})$`, phone); !m {
			// 格式不对
		}

		// 下拉菜单中是否有被选中的项目
		slice := []string{"apple", "pear", "banana"}
		v := r.Form.Get("fruit")
		for _, item := range slice {
			if item == v {
				// 有选中的值
			}
		}

		// 判断单选按钮是否为预设值
		slice = []string{"1", "2"}
		for _, item := range slice {
			if item == r.Form.Get("gender") {
				// 是合法的值
			}
		}

		// 复选框判断,复选框接收到的数据是一个slice
		slice = []string{"football", "basketball", "tennis"}
		slice1, slice2 := helper.StringSliceToInterface(r.Form["interest"]), helper.StringSliceToInterface(slice)
		if a := beeku.Slice_diff(slice1, slice2); a != nil {
			// 非法数据
		}

		// 日期和时间
		t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
		fmt.Printf("Go launched at %s\n", t.Local())

		// 验证15位身份证，15位的是全部数字
		if m, _ := regexp.MatchString(`^(\d{15})$`, r.Form.Get("usercard")); !m {
			// 非法
		}
		//验证18位身份证，18位前17位为数字，最后一位是校验位，可能为数字或字符X。
		if m, _ := regexp.MatchString(`^(\d{17})([0-9]|X)$`, r.Form.Get("usercard")); !m {
			// 非法
		}

		// 在浏览器输出<script>alert()</script>-方法一
		temple, err := template.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
		err = temple.ExecuteTemplate(w, "T", "<script>alert('you have been pwned')</script>")
		// 在浏览器输出<script>alert()</script>-方法二
		temple, err = template.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
		err = temple.ExecuteTemplate(w, "T", template.HTML("<script>alert('you have been pwned')</script>"))
	}
}
