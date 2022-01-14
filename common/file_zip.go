package common

import (
	"archive/zip"
	"errors"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

//压缩为zip格式
//source为要压缩的文件或文件夹, 绝对路径和相对路径都可以
//target是目标文件
//filter是过滤正则(Golang 的 包 path.Match)
func ZipArchive(source, target, filter string) error {
	var err error
	if isAbs := filepath.IsAbs(source); !isAbs {
		source, err = filepath.Abs(source) // 将传入路径直接转化为绝对路径
		if err != nil {
			return err
		}
	}
	//创建zip包文件
	zipfile, err := os.Create(target)
	if err != nil {
		return err
	}

	defer func() {
		if err := zipfile.Close(); err != nil {
			log.Printf("*File close error: %s, file: %s", err.Error(), zipfile.Name())
		}
	}()

	//创建zip.Writer
	zw := zip.NewWriter(zipfile)

	defer func() {
		if err := zw.Close(); err != nil {
			log.Printf("zipwriter close error: %s", err.Error())
		}
	}()

	info, err := os.Stat(source)
	if err != nil {
		return err
	}

	var baseDir string
	if info.IsDir() {
		baseDir = filepath.Base(source)
	}
	err = filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		//将遍历到的路径与pattern进行匹配
		ism, err := filepath.Match(filter, info.Name())

		if err != nil {
			return err
		}
		//如果匹配就忽略
		if ism {
			return nil
		}
		//创建文件头
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}
		if baseDir != "" {
			//header.Name = filepath.Join(baseDir, strings.TrimPrefix(path, source))
			header.Name = filepath.Join(".", strings.TrimPrefix(path, source))
		}
		if info.IsDir() {
			header.Name += "/"
		} else {
			header.Method = zip.Deflate
		}
		//写入文件头信息
		writer, err := zw.CreateHeader(header)
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		//写入文件内容
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer func() {
			if err := file.Close(); err != nil {
				log.Printf("*File close error: %s, file: %s", err.Error(), file.Name())
			}
		}()
		_, err = io.Copy(writer, file)
		return err
	})

	if err != nil {
		return err
	}

	return nil
}

// 打包指定的文件集合
func ZipArchiveFiles(files map[string]string, target string) error {
	if len(files) == 0 {
		return errors.New("文件集合为空")
	}
	zipfile, err := os.Create(target)
	if err != nil {
		return err
	}
	//创建zip.Writer
	zw := zip.NewWriter(zipfile)
	for filePath, fileName := range files {
		if info, err := os.Stat(filePath); err != nil {
			return err
		} else if header, err := zip.FileInfoHeader(info); err != nil {
			return err
		} else {
			header.Name = fileName
			if writer, err := zw.CreateHeader(header); err != nil {
				return err
			} else if fileData, err := os.Open(filePath); err != nil {
				return err
			} else {
				defer fileData.Close()
				if _, err = io.Copy(writer, fileData); err != nil {
					return err
				}
			}
		}
	}
	zw.Close()
	return nil
}
