package classfile

// ConstantMemberrefInfo 结构体
type ConstantMemberrefInfo struct {
	cp               ConstantPool
	classIndex       uint16
	nameAndTypeIndex uint16
}

// 读取信息
func (self *ConstantMemberrefInfo) readInfo(reader *ClassReader){
	self.classIndex = reader.readUint16()
	self.nameAndTypeIndex = reader.readUint16()
}

func (self *ConstantMemberrefInfo) ClassName() string {
	return self.cp.getClassName(self.classIndex)
}

func (self *ConstantMemberrefInfo) NameAndDescriptor() (string, string) {
	return self.cp.getNameAndType(self.nameAndTypeIndex)
}

// 通过嵌套结构体来实现 ConstantFieldrefInfo ConstantMethodrefInfo ConstantInterfaceMethodrefInfo 对 ConstantMemberrefInfo 的继承
type ConstantFieldrefInfo struct {
	ConstantMemberrefInfo
}
type ConstantMethodrefInfo struct {
	ConstantMemberrefInfo
}
type ConstantInterfaceMethodrefInfo struct {
	ConstantMemberrefInfo
}