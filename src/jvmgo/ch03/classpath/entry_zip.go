package classpath

import "archive/zip"
import "errors"
import "io/ioutil"
import "path/filepath"

type ZipEntry struct {
	// 存放目录的绝对路径
	absPath string
}

func newZipEntry(path string) *ZipEntry {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &ZipEntry{absPath}
}

func (self *ZipEntry) String() string {
	return self.absPath
}

// 打开 zip 文件，读取其中的 .class
func (self *ZipEntry) readClass(className string) ([]byte, Entry, error){
	// 打开 zip 文件
	r, err := zip.OpenReader(self.absPath)
	if err != nil {
		return nil, nil, err
	}

	defer r.Close()
	// for range 循环遍历，遍历 zip 中的所有文件，查找与 className 相同的
	// for _, f := range r.File  --> _ 代表返回值不理会，将 r 中的每一个元素的 File 赋值给 f
	for _, f := range r.File {
		if f.Name == className {
			rc, err := f.Open()
			if err != nil {
				return nil, nil, err
			}
			defer rc.Close()
			data, err := ioutil.ReadAll(rc)
			if err != nil {
				return nil, nil, err
			}
			return data, self, nil
		}
	}
	return nil, nil, errors.New("class not found：" + className)
}