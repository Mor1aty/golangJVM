package classpath

import "io/ioutil"
import "path/filepath"

type DirEntry struct{
	// 存放目录的绝对路径
	absDir string
}

/*
 Go 语言没有专门的构造函数，这里使用 new 函数代替
*/
// 调用参数将 path 转换为绝对路径 absDir，如果转换过程出现错误则调用 panic() 函数终止程序执行，否则建立 DirEntry实例并返回
func newDirEntry(path string) *DirEntry{
	absDir, err := filepath.Abs(path)
	// nil 零值，初始化时赋的值
	if err != nil {
		panic(err)
	}
	return &DirEntry{absDir}
}

func (self *DirEntry) readClass(className string) ([]byte,Entry,error){
	fileName := filepath.Join(self.absDir,className)
	data, err := ioutil.ReadFile(fileName)
	return data, self, err
}

func (self *DirEntry) String() string{
	return self.absDir
}