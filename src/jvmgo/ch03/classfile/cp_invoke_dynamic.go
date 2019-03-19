package classfile

// ConstantMethodHandleInfo 结构体
type ConstantMethodHandleInfo struct {
	referenceKind  uint8
	referenceIndex uint16
}

// 读取信息
func (self *ConstantMethodHandleInfo) readInfo(reader *ClassReader) {
	self.referenceKind = reader.readUint8()
	self.referenceIndex = reader.readUint16()
}

// ConstantMethodTypeInfo
type ConstantMethodTypeInfo struct {
	descriptorIndex uint16
}

// 读取信息
func (self *ConstantMethodTypeInfo) readInfo(reader *ClassReader) {
	self.descriptorIndex = reader.readUint16()
}

// ConstantInvokeDynamicInfo 结构体
type ConstantInvokeDynamicInfo struct {
	bootstrapMethodAttrIndex uint16
	nameAndTypeIndex         uint16
}

// 读取信息
func (self *ConstantInvokeDynamicInfo) readInfo(reader *ClassReader) {
	self.bootstrapMethodAttrIndex = reader.readUint16()
	self.nameAndTypeIndex = reader.readUint16()
}
