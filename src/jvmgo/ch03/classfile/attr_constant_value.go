package classfile

// ConstantValueAttribute 结构体
type ConstantValueAttribute struct {
	constantValueIndex uint16
}

// 读取信息
func (self *ConstantValueAttribute) readInfo(reader *ClassReader) {
	self.constantValueIndex = reader.readUint16()
}

// 返回 constantValueIndex
func (self *ConstantValueAttribute) ConstantValueIndex() uint16 {
	return self.constantValueIndex
}
