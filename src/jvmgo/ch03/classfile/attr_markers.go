package classfile

// DeprecatedAttribute 结构体
type DeprecatedAttribute struct {
	MarkerAttribute
}

// SyntheticAttribute 结构体
type SyntheticAttribute struct {
	MarkerAttribute
}

type MarkerAttribute struct{}

func (self *MarkerAttribute) readInfo(reader *ClassReader){
	// read nothing
	// 属性没有数据，方法为空
}