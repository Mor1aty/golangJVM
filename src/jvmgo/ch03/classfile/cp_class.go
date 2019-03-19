package classfile

// COnstantClassInfo 结构体
type ConstantClassInfo struct {
	cp        ConstantPool
	nameIndex uint16
}

// 读取常量池索引
func (self *ConstantClassInfo) readInfo(reader *ClassReader) {
	self.nameIndex = reader.readUint16()
}

// 按索引从常量池中查找字符串
func (self *ConstantClassInfo) Name() string {
	return self.cp.getUtf8(self.nameIndex)
}
