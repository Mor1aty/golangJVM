package classpath

import "os"
import "path/filepath"

type Classpath struct {
	// 分别存放的是三种类路径
	bootClasspath Entry	// 启动类路径
	extClasspath Entry	// 扩展类路径
	userClasspath Entry	// 用户类路径
}

// 使用 -Xjre 选项解析启动类路径和扩展类路径，使用 -classpath/-cp 选项解析用户类路径
func Parse(jreOption, cpOption string) *Classpath {
	cp := &Classpath{}
	cp.parseBootAndExtClasspath(jreOption)
	cp.parseUserClasspath(cpOption)
	return cp
}

// 解析 Boot 和 Ext 类路径
func (self *Classpath) parseBootAndExtClasspath(jreOption string) {
	jreDir := getJreDir(jreOption)

	// jre/lib/*
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	self.bootClasspath = newWildcardEntry(jreLibPath)

	// jre.lib.ext/*
	jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
	self.extClasspath = newWildcardEntry(jreExtPath)
}

// 寻找 jre 目录
func getJreDir(jreOption string) string {
	// 先使用用户使用的 -Xjre 选项作为 jre 目录
	if jreOption != "" && exists(jreOption) {
		return jreOption
	}
	// 没有输入则在当前目录下寻找 jre 目录
	if exists("./jre") {
		return "./jre"
	}
	// 在 JAVA_HOME 中寻找 jre 目录
	if jh := os.Getenv("JAVA_HOME"); jh != "" {
		return filepath.Join(jh, "jre")
	}
	panic("Can not find jre folder!")
}

// 判断目录是否存在
func exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// 解析 User 类路径，如果用户没有提供 -classpath/-cp 选项，则使用当前目录作为用户类路径
func (self *Classpath) parseUserClasspath(cpOption string) {
	if cpOption == "" {
		cpOption = "."
	}
	self.userClasspath = newEntry(cpOption)
}

// ReadClass 方法依次从 bootClasspath，extClasspath，userClasspath 搜索 class 文件
func (self *Classpath) ReadClass(className string) ([]byte, Entry, error) {
	// 参数 className 不包含 .class 后缀
	className = className + ".class"
	if data, entry, err := self.bootClasspath.readClass(className); err == nil {
		return data, entry, err
	}
	if data, entry, err := self.extClasspath.readClass(className); err == nil {
		return data, entry, err
	}
	return self.userClasspath.readClass(className)
}

// String 方法返回用户类路径的字符串
func (self *Classpath) String() string {
	return self.userClasspath.String()
}