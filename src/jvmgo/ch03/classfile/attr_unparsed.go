package classfile

// UnparsedAttribute 结构体
type UnparsedAttribute struct {
	name   string
	length uint32
	info   []byte
}

// 读取信息
func (self *UnparsedAttribute) readInfo(reader *ClassReader) {
	self.info = reader.readBytes(self.length)
}
