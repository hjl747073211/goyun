package meta

//文件元信息结构体
type FileMeta struct {
	FileSha1 string
	FileName string
	FileSize int64
	Location string //保存位置
	UploadAt string //上传时间

}

var fileMetas map[string]FileMeta

func init() {
	fileMetas = make(map[string]FileMeta)
}

//UpdateFileMeta 新增或更新文件元信息
func UpdateFileMeta(fmeta FileMeta) {
	fileMetas[fmeta.FileSha1] = fmeta
}

//GetFileMeta 通过sha1获取文件元信息对象
func GetFileMeta(fileSha1 string) FileMeta {
	return fileMetas[fileSha1]
}
