package common

import (
	"archive/zip"
	"crypto/md5"
	"encoding/hex"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
)

// 剪切文件
func CutFile(dstName, srcName string) error {
	// 先删除再cut
	os.Remove(dstName)
	src, err := os.Open(srcName)
	if err != nil {
		return err
	}
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer dst.Close()
	_, err = io.Copy(dst, src)
	src.Close()
	if err == nil {
		err = os.Remove(srcName)
	}
	return err
}

// 复制文件
func CopyFile(dstName, srcName string) error {
	// 先删除再copy
	os.Remove(dstName)
	src, err := os.Open(srcName)
	if err != nil {
		return err
	}
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer dst.Close()
	_, err = io.Copy(dst, src)
	src.Close()
	return err
}

// 文件是否存在
func Isfile(filename string) bool {
	_, err := os.Stat(filename)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}

// 获取文件md5的hex编码
func Md5FileHex(path string) string {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "00000000000000000000000000000000"
	}
	hash := md5.New()
	hash.Write(data)
	return hex.EncodeToString(hash.Sum(nil))
}

// 文件文件或文件夹是否存在
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

// 创建文件夹，文件夹如果存在 直接返回
func MakePath(filePath string) error {
	ok, err := PathExists(filePath)
	if err != nil {
		return err
	}
	if !ok {
		err = os.MkdirAll(filePath, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

// 解压zip文件
func Unzip(zipFile string, destDir string) error {
	zipReader, err := zip.OpenReader(zipFile)
	if err != nil {
		return err
	}
	defer zipReader.Close()

	for _, f := range zipReader.File {
		fpath := filepath.Join(destDir, f.Name)
		if f.FileInfo().IsDir() {
			os.MkdirAll(fpath, os.ModePerm)
		} else {
			if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
				return err
			}

			inFile, err := f.Open()
			if err != nil {
				return err
			}
			defer inFile.Close()

			outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer outFile.Close()

			_, err = io.Copy(outFile, inFile)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// 获取纯净文件名 用于保存到磁盘本地
func PureFilename(name string) string {
	reg := regexp.MustCompile(`[/|?|*|:|\||\\|<|>|&|#|@|$|(|)|;|'|"|%]+`)
	return reg.ReplaceAllString(name, "")
}
