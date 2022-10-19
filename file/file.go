package file

import (
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
)

//获取文件大小
func GetSize(f multipart.File) (int, error) {
	content, err := ioutil.ReadAll(f)
	return len(content), err
}

//获取文件后缀
func GetExt(fileName string) string {
	return path.Ext(fileName)
}

//检查文件是否存在
func CheckExist(src string) bool {
	_, err := os.Stat(src)
	return os.IsNotExist(err)
}

//检查文件权限
func CheckPermission(src string) bool {
	_, err := os.Stat(src)
	return os.IsPermission(err)
}

//如果不存在则新建文件夹
func IsNotExistMkDir(src string) error {
	//fmt.Println(src)
	exist := CheckExist(src)
	//fmt.Println(exist)
	if exist {
		fmt.Println("exist:%s", exist)
		if err := MkDir(src); err != nil {
			return err
		}
	}
	return nil
}

//新建文件夹
func MkDir(src string) error {
	err := os.MkdirAll(src, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

//打开文件
func Open(name string, flag int, perm os.FileMode) (*os.File, error) {
	f, err := os.OpenFile(name, flag, perm)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func IsExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

// 判断所给路径是否为文件夹
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// 判断所给路径是否为文件
func IsFile(path string) bool {
	return !IsDir(path)
}

// 判断文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func IsDirCreate(pathDir string) bool {
	_dir := pathDir
	exist, err := PathExists(_dir)
	if err != nil {
		fmt.Printf("get dir error![%v]\n", err)
		return false
	}

	if exist {
		fmt.Printf("has dir![%v]\n", _dir)
		return true
	} else {
		fmt.Printf("no dir![%v]\n", _dir)
		// 创建文件夹
		err := os.Mkdir(_dir, os.ModePerm)
		if err != nil {
			fmt.Printf("mkdir failed![%v]\n", err)
			allErr := os.MkdirAll(_dir, os.ModePerm)
			if allErr != nil {
				fmt.Printf("all mkdir failed![%v]\n", err)
				return false
			} else {
				fmt.Printf("all mkdir success!\n")
				return true
			}

		} else {
			fmt.Printf("mkdir success!\n")
			return true
		}
	}
}
