package classfile

// ConstantNameAndTypeInfo 结构体
type ConstantNameAndTypeInfo struct {
	nameIndex       uint16
	descriptorIndex uint16
}

// 读取 nameIndex 和 descriptorIndex
func (self *ConstantNameAndTypeInfo) readInfo(reader *ClassReader){
	self.nameIndex = reader.readUint16()
	self.descriptorIndex = reader.readUint16()
}
