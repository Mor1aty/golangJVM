package classfile

// ExceptionsAttribute 结构体
type ExceptionsAttribute struct {
	exceptionIndexTable []uint16
}

// 读取信息
func (self *ExceptionsAttribute) readInfo(reader *ClassReader) {
	self.exceptionIndexTable = reader.readUint16s()
}

// 返回 ExceptionIndexTable
func (self *ExceptionsAttribute) ExceptionIndexTable() []uint16 {
	return self.exceptionIndexTable
}
