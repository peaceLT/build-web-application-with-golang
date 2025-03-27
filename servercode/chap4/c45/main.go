package main

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"time"
)

func upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		// 处理GET请求，通常用于显示上传页面

		// 生成一个唯一的token，用于防止CSRF攻击
		crutime := time.Now().Unix()                      // 获取当前时间的Unix时间戳
		h := md5.New()                                    // 创建一个新的MD5哈希对象
		io.WriteString(h, strconv.FormatInt(crutime, 10)) // 将时间戳写入哈希对象
		token := fmt.Sprintf("%x", h.Sum(nil))            // 生成MD5哈希值作为token

		// 解析并执行模板，传递token到模板中
		t, _ := template.ParseFiles("./htmlcode/upload.html") // 解析upload.html模板文件
		t.Execute(w, token)                                   // 执行模板，将token传递给模板并输出到客户端
	} else {
		// 处理POST请求，通常用于提交文件上传

		r.ParseMultipartForm(32 << 20)                 // 解析表单，设置最大内存为32MB
		file, handler, err := r.FormFile("uploadfile") // 获取上传的文件
		if err != nil {
			fmt.Println(err) // 输出错误信息
			return
		}
		defer file.Close() // 确保文件在函数结束时关闭

		// 输出文件头信息到客户端
		fmt.Fprintf(w, "%v", handler.Header)

		// 打开或创建文件，准备写入上传的内容
		f, err := os.OpenFile("./test/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666) // 假设当前目录下已存在test目录
		if err != nil {
			fmt.Println(err) // 输出错误信息
			return
		}
		defer f.Close() // 确保文件在函数结束时关闭

		// 将上传的文件内容复制到目标文件
		io.Copy(f, file)
	}
}

// 模拟客户端表单功能支持文件上传,用于将文件上传到指定的URL
func postFile(fileName, targetUrl string) error {
	// 创建一个缓冲区和一个multipart.Writer
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	// 关键的一步操作：创建一个表单文件字段
	fileWriter, err := bodyWriter.CreateFormFile("uploadfile", fileName)
	if err != nil {
		fmt.Println("error writing to buffer")
		return err
	}

	// 打开文件句柄操作
	fh, err := os.Open(fileName)
	if err != nil {
		fmt.Println("error opening file")
		return err
	}
	defer fh.Close() // 确保文件在函数结束时关闭

	// 将文件内容复制到fileWriter
	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		return err
	}

	// 获取Content-Type并关闭bodyWriter
	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	// 发送POST请求
	resp, err := http.Post(targetUrl, contentType, bodyBuf)
	if err != nil {
		return err
	}
	defer resp.Body.Close() // 确保响应体在函数结束时关闭

	// 读取响应体
	resp_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// 输出响应状态和响应体
	fmt.Println(resp.Status)
	fmt.Println(string(resp_body))
	return nil
}

func main() {
	http.HandleFunc("/upload", upload)
	if err := http.ListenAndServe(":9090", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}

	target_url := "http://localhost:9090/upload"
	filename := "./corver.jpg"
	postFile(filename, target_url)
}
