package classpath

import "os"
import "path/filepath"
import "strings"

func newWildcardEntry(path string) CompositeEntry {
	baseDir := path[:len(path) - 1] // 删除 * 号，将 java.lang.utils.* 变成 java.lang.utils.
	compositeEntry := []Entry{}

	walkFn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// 如果是一个目录并且 path 与 baseDir 不相等（就是给定的根目录），则用 SkipDir 跳过子目录
		if info.IsDir() && path != baseDir {
			return filepath.SkipDir
		}
		// 根据后缀名选出 jar 文件，并添加到 compositeEntry
		if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") {
			jarEntry := newZipEntry(path)
			compositeEntry = append(compositeEntry, jarEntry)
		}
		return nil
	}
	// Walk 函数遍历 baseDir，使用 walkFn 创建 ZipEntry。
	// walkFn 函数作为第二个参数
	filepath.Walk(baseDir, walkFn)

	return compositeEntry
}