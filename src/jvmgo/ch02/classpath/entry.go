package classpath

import "os"
import "strings"

// 路径分割符，string() 相当于 Java 中的 toString() 方法 
const pathListSeparator = string(os.PathListSeparator)

// Entry 为一个接口
type Entry interface{
	// 寻找加载 class 文件
	// readClass 方法参数为 string 类型的 className，指的是 class 文件的相对路径，路径间使用 / 分割，文件名后缀为 .class
	// 返回值有三个，分别是读取到的字节数据：[]byte，最终定位到 class 文件的 Entry，错误信息 error
	readClass(className string) ([]byte,Entry,error)
	String() string
}

// 生成新的 Entry
func newEntry(path string) Entry {
	// 如果路径包含路径分割符，生成 CompositeEntry
	if strings.Contains(path,pathListSeparator){
		return newCompositeEntry(path)
	}
	// 如果路径以 * 结尾，生成 WildcardEntry
	if strings.HasSuffix(path,"*"){
		return newWildcardEntry(path)
	}
	// 如果路径以 .jar，.JAR，.zip，.ZIP 结尾，生成 WildcardEntry
	if strings.HasSuffix(path,".jar") || strings.HasSuffix(path,".JAR") || 
	   strings.HasSuffix(path,".zip") || strings.HasSuffix(path,".ZIP"){
		
		return newZipEntry(path)
	}
	// 如果不属于上面所有情况，生成 DirEntry
	return newDirEntry(path) 
}