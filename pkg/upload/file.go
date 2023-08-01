package upload 

import (
	"path"
	"strings"
	"os"
	"io"
	"io/ioutil"
	"mime/multipart"
	"douyin/pkg/util"
	"douyin/global"
)

type FileType int

const TypeVideo FileType = iota + 1

//获取文件名称，先是通过获取文件后缀并筛出原始文件名进行 MD5 加密，最后返回经过加密处理后的文件名。
func GetFileName(name string) string {
	ext := GetFileExt(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = util.EncodeMD5(fileName)

	return fileName + ext
}

//GetFileExt：获取文件后缀，主要是通过调用 path.Ext 方法进行循环查找”.“符号，最后通过切片索引返回对应的文化后缀名称。
func GetFileExt(name string) string {
	return path.Ext(name)
}

//GetSavePath：获取文件保存地址，这里直接返回配置中的文件保存目录即可，也便于后续的调整。
func GetSavePath() string {
	return global.AppSetting.UploadSavePath
}

//检查保存目录是否存在，通过调用 os.Stat 方法获取文件的描述信息 FileInfo，并调用 os.IsNotExist 方法进行判断
func CheckSavePath(dst string) bool {
	_, err := os.Stat(dst)
	return os.IsNotExist(err)
}

//检查文件后缀是否包含在约定的后缀配置项中，需要的是所上传的文件的后缀有可能是大写、小写、大小写等，因此我们需要调用 strings.ToUpper 方法统一转为大写（固定的格式）来进行匹配。
func CheckContainExt(t FileType, name string) bool {
	ext := GetFileExt(name)
	ext = strings.ToUpper(ext)
	switch t {
	case TypeVideo:
		for _, allowExt := range global.AppSetting.UploadVideoAllowExts {
			if strings.ToUpper(allowExt) == ext {
				return true
			}
		}

	}

	return false
}

//检查文件大小是否超出最大大小限制。
func CheckMaxSize(t FileType, f multipart.File) bool {
	content, _ := ioutil.ReadAll(f)
	size := len(content)
	switch t {
	case TypeVideo:
		if size >= global.AppSetting.UploadVideoMaxSize*1024*1024 {
			return true
		}
	}

	return false
}

//检查文件权限是否足够
func CheckPermission(dst string) bool {
	_, err := os.Stat(dst)
	return os.IsPermission(err)
}

//创建在上传文件时所使用的保存目录
func CreateSavePath(dst string, perm os.FileMode) error {
	err := os.MkdirAll(dst, perm)
	if err != nil {
		return err
	}

	return nil
}

//保存所上传的文件
func SaveFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
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