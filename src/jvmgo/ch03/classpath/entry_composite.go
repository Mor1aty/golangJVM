package classpath

import "errors"
import "strings"

type CompositeEntry []Entry

// 将 pathList 按照分隔符分成一个单独的 Entry，存放到 compositeEntry 的 Entry 数组中
func newCompositeEntry(pathList string) CompositeEntry {
	compositeEntry := []Entry{}
	for _, path := range strings.Split(pathList, pathListSeparator) {
		entry := newEntry(path)
		compositeEntry = append(compositeEntry, entry)
	}
	return compositeEntry
}

// 依次调用每一个 Entry 的 readClass 方法
func (self CompositeEntry) readClass(className string) ([]byte, Entry, error) {
	for _, entry := range self{
		data, from, err := entry.readClass(className)
		if err == nil {
			return data, from, nil
		} 
	}
	return nil, nil, errors.New("class not found：" + className)
}

// 依次调用每一个 Entry 的 String 方法，并将其按照 pathListSeparator 分割符返回一个总的字符串
func (self CompositeEntry) String() string {
	strs := make([]string, len(self))
	for i, entry := range self {
		strs[i] = entry.String()
	}
	return strings.Join(strs, pathListSeparator)
}