package classfile

// SourceFileAttribute 结构体
type SourceFileAttribute struct {
	cp              ConstantPool
	sourceFileIndex uint16
}

// 读取信息
func (self *SourceFileAttribute) readInfo(reader *ClassReader) {
	self.sourceFileIndex = reader.readUint16()
}

// 文件名
func (self *SourceFileAttribute) FileName() string {
	return self.cp.getUtf8(self.sourceFileIndex)
}
