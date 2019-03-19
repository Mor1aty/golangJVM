package classfile

import "math"

// ConstantIntegerInfo 结构体
type ConstantIntegerInfo struct {
	val int32
}

// 读取一个 uint32 数据，转化成一个 int32
func (self *ConstantIntegerInfo) readInfo(reader *ClassReader){
	bytes := reader.readUint32()
	self.val = int32(bytes)
}

// ConstantFloatInfo 结构体
type ConstantFloatInfo struct {
	val float32
}

// 读取一个 uint32 数据，调用 math 包的 Float32frombits() 函数把他转化成 float32 类型
func (self *ConstantFloatInfo) readInfo(reader *ClassReader){
	bytes := reader.readUint32()
	self.val = math.Float32frombits(bytes)
}

// ConstantLongInfo 结构体
type ConstantLongInfo struct {
	val int64
}

// 读取一个 uint64 数据，把他转化成 int64 类型
func (self *ConstantLongInfo) readInfo(reader *ClassReader){
	bytes := reader.readUint64()
	self.val = int64(bytes)
}

// ConstantDoubleInfo 结构体
type ConstantDoubleInfo struct {
	val float64
}

// 读取一个 uint64 数据，把他转化成 int64 类型
func (self *ConstantDoubleInfo) readInfo(reader *ClassReader){
	bytes := reader.readUint64()
	self.val = math.Float64frombits(bytes)
}

