package handler

import (
	"fmt"
	"goyun/meta"
	"goyun/util"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

//UploadHandler 上传文件
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		//返回上传的html页面
		data, err := ioutil.ReadFile("./static/view/index.html")
		if err != nil {
			io.WriteString(w, "internel server err")
			return
		}
		io.WriteString(w, string(data))
	} else if r.Method == "POST" {
		//接收文件流并保存到本地目录
		file, head, err := r.FormFile("file") //文件句柄，文件头，错误信息
		if err != nil {
			fmt.Printf("failed to get data ,err:%s", err.Error())
			return
		}
		defer file.Close()

		fileMeta := meta.FileMeta{
			FileName: head.Filename,
			Location: "/tmp/" + head.Filename,
			UploadAt: time.Now().Format("2006-01-02 15:04:05"),
		}

		newFile, err := os.Create(fileMeta.Location) //文件保存到本地的位置
		if err != nil {
			fmt.Printf("failed to create file ,err:%s", err.Error())
			return
		}
		defer newFile.Close()

		fileMeta.FileSize, err = io.Copy(newFile, file) //返回copy的字节长度和err
		if err != nil {
			fmt.Printf("failed to save data into file ,err:%s", err.Error())
			return
		}

		newFile.Seek(0, 0)
		fileMeta.FileSha1 = util.FileSha1(newFile)
		meta.UpdateFileMeta(fileMeta)

		http.Redirect(w, r, "/file/upload/suc", http.StatusFound)
	}
}

//UploadSucHandler 上传成功
func UploadSucHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Upload finished!")
}
