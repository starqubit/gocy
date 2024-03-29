package common

import (
	"archive/zip"
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
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
func Md5FileHex(filePath string) string {
	file, err := os.Open(filePath)
	if err != nil {
		return "00000000000000000000000000000000"
	}
	defer file.Close()

	w := md5.New()
	if _, err = io.Copy(w, file); err != nil {
		return "00000000000000000000000000000000"
	}
	strMd5 := fmt.Sprintf("%x", w.Sum(nil))
	return strMd5
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

// 异或加密
func Xor(data, key []byte) []byte {
	var result []byte
	for i := 0; i < len(data); i++ {
		result = append(result, data[i]^key[0])
	}
	return result
}

// 遍历文件
func GetAllFile(pathname string, whiteExt ...string) ([]string, error) {
	filePaths := make([]string, 0)
	rd, err := ioutil.ReadDir(pathname)
	if err != nil {
		return filePaths, err
	}
	for _, fi := range rd {
		// 不是文件夹
		if !fi.IsDir() {
			// 过滤后缀
			_ext := strings.ToLower(filepath.Ext(fi.Name()))

			// 扫描白名单
			for _, ext := range whiteExt {
				if _ext == strings.ToLower(ext) {
					// 白名单内的加入队列
					filePaths = append(filePaths, filepath.Join(pathname, fi.Name()))
					break
				}
			}
			continue
		}
		// 文件夹
		resp, err := GetAllFile(filepath.Join(pathname, fi.Name()), whiteExt...)
		if err != nil {
			return filePaths, err
		}
		filePaths = append(filePaths, resp...)

	}
	return filePaths, nil
}
