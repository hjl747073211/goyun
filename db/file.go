package db

import (
	"database/sql"
	"fmt"
	mydb "goyun/db/mysql"
	"time"
)

//OnFileUploadFinished 文件保存好之后将信息入库
func OnFileUploadFinished(filehash string, filename string, filesize int64, fileaddr string) bool {
	createAt := time.Now().Format("2006-01-02 15:04:05")
	updateAt := time.Now().Format("2006-01-02 15:04:05")
	stmt, err := mydb.DBConn().Prepare("insert into tbl_file (`file_sha1`,`file_name`,`file_size`,`file_addr`,`status`,`create_at`,`update_at`) values (?,?,?,?,1,?,?)")
	if err != nil {
		fmt.Println("failed to prepare statement,err:" + err.Error())
		return false
	}
	defer stmt.Close()

	ret, err := stmt.Exec(filehash, filename, filesize, fileaddr, createAt, updateAt)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	//filehash是唯一键，已存在不会插入，下面返回受影响条数
	if rf, err := ret.RowsAffected(); err == nil {
		if rf <= 0 {
			fmt.Printf("file with hash:%s has been uploaded before", filehash)
		}
		return true
	}
	return false

}

//TableFile 。。
type TableFile struct {
	FileHash string
	FileName sql.NullString
	FileSize sql.NullInt64
	FileAddr sql.NullString
}

//GetFileMeta mysql获取信息
func GetFileMeta(filehash string) (*TableFile, error) {
	stmt, err := mydb.DBConn().Prepare("select file_sha1,file_name,file_size,file_addr from tbl_file where file_sha1=? and status=1 limit 1")
	if err != nil {
		fmt.Println("failed to prepare statement,err:" + err.Error())
		return nil, err
	}
	defer stmt.Close()

	tfile := TableFile{}
	err = stmt.QueryRow(filehash).Scan(&tfile.FileHash, &tfile.FileName, &tfile.FileSize, &tfile.FileAddr)
	if err != nil {
		fmt.Println("failed to Query,err:" + err.Error())
		return nil, err
	}
	return &tfile, nil

}
