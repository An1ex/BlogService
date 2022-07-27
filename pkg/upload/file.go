package upload

import (
	"io"
	"mime/multipart"
	"os"
	"path"
	"strings"

	"BlogService/global"
	"BlogService/pkg/util"
)

type FileType int

const (
	TypeImage FileType = iota + 1
	TypeExcel
	TypeTxt
)

// EncryptFileName MD5加密文件名称
func EncryptFileName(name string) string {
	ext := GetFileExt(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = util.EncodeMD5(fileName)
	return fileName + ext
}

// GetFileExt 获取文件格式
func GetFileExt(name string) string {
	return path.Ext(name)
}

// GetSavePath 获取保存路径
func GetSavePath() string {
	return global.Config.App.UploadSavePath
}

// CheckFileType 检查文件类型是否允许
func CheckFileType(t FileType) bool {
	switch t {
	case TypeImage:
		return true
	case TypeExcel:
		return true
	case TypeTxt:
		return true
	}
	return false
}

// CheckSavePath 检查保存目录是否存在
func CheckSavePath(dst string) bool {
	_, err := os.Stat(dst)    //若文件夹存在,err=nil
	return os.IsNotExist(err) //若文件夹不存在,err!=nil,IsNotExist(err)=true
}

// CheckContainExt 检查文件格式是否允许
func CheckContainExt(t FileType, name string) bool {
	ext := GetFileExt(name)
	ext = strings.ToUpper(ext)
	switch t {
	case TypeImage:
		for _, allowExt := range global.Config.App.UploadImageAllowExts {
			if strings.ToUpper(allowExt) == ext {
				return true
			}
		}

	}
	return false
}

// CheckMaxSize 检查文件大小是否超出限制
func CheckMaxSize(t FileType, header *multipart.FileHeader) bool {
	size := int(header.Size)
	switch t {
	case TypeImage:
		if size <= global.Config.App.UploadImageMaxSize*1024*1024 {
			return true
		}
	}
	return false
}

// CheckPermission 检查文件权限是否足够
func CheckPermission(dst string) bool {
	_, err := os.Stat(dst)
	return os.IsPermission(err)
}

// CreateSavePath 创建保存文件的目录
func CreateSavePath(dst string, perm os.FileMode) error {
	err := os.MkdirAll(dst, perm)
	if err != nil {
		return err
	}
	return nil
}

// SaveFile 保存上传的文件
func SaveFile(dst string, header *multipart.FileHeader) error {
	src, err := header.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, src)
	return err
}
