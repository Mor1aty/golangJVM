package classfile

import "fmt"

type ClassFile struct {
	// magic uint32
	minorVersion uint16
	majorVersion uint16
	constantPool ConstantPool
	accessFlags  uint16
	thisClass    uint16
	superClass   uint16
	interfaces   []uint16
	fields       []*MemberInfo
	methods      []*MemberInfo
	attributes   []AttributeInfo
}

func Parse(classData []byte) (cf *ClassFile, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()

	cr := &ClassReader{classData}
	cf = &ClassFile{}
	cf.read(cr)
	return
}

func (self *ClassFile) read(reader *ClassReader) {
	self.readAndCheckMagic(reader)
	self.readAndCheckVersion(reader)
	self.constantPool = readConstantPool(reader)
	self.accessFlags = reader.readUint16()
	self.thisClass = reader.readUint16()
	self.superClass = reader.readUint16()
	self.interfaces = reader.readUint16s()
	self.fields = readMembers(reader, self.constantPool)
	self.methods = readMembers(reader, self.constantPool)
	self.attributes = readAttributes(reader, self.constantPool)
}

/**
	很多文件格式都会规定满足该格式的文件必须以某几个固定字节开头，这几个字节起标志作用，叫做魔数（magic number）
	例如：PDF 文件以 4 字节“%PDF”（0x25，0x50，0x44，0x46）开头，ZIP 文件以 2 字节“PK”（0x50，0x4B）
	class 文件的魔数是“0xCAFEBABE”
 */
// 读取和查验魔数
func (self *ClassFile) readAndCheckMagic(reader *ClassReader) {
	magic := reader.readUint32()
	if magic != 0xCAFEBABE {
		panic("java.lang.ClassFormatError:magic!")
	}
}

/**
	魔数之后是 class 文件的次版本号和主版本号，都是 u2 类型。若某 class 文件的主版本号是 M，次版本号是 m，那么完整的版本号就是 M.m
	次版本号只在 J2SE 1.2 之前用过，从 1.2 开始基本就没什么用了（都是 0）。主版本号在 J2SE 1.2 之前是 45，从 1.2 开始，每次有大的版本号发布，都会 +1
	我们这里参考 Java8，支持 45.0 到 52.0 的 class 文件，如果遇到其他版本号，先调用 panic() 方法终止程序执行
 */
// 读取和查验版本
func (self *ClassFile) readAndCheckVersion(reader *ClassReader) {
	self.minorVersion = reader.readUint16()
	self.majorVersion = reader.readUint16()
	switch self.majorVersion {
	case 45:
		return
	case 46,47,48,49,50,51,52:
		if self.minorVersion == 0 {
			return
		}
	}
	panic("java.lang.UnsupportedClassVersionError!")
}



// 从常量池查找类名
func (self *ClassFile) ClassName() string {
	return self.constantPool.getClassName(self.thisClass)
}

// 从常量池中查找超类类名
func (self *ClassFile) SuperClassName() string {
	if self.superClass > 0 {
		return self.constantPool.getClassName(self.superClass)
	}
	return "" //只有 java.lang.Object 只有没有超类
}

// 从常量池查询接口名
func (self *ClassFile) InterfaceNames() []string {
	interfaceNames := make([]string, len(self.interfaces))
	for i, cpIndex := range self.interfaces {
		interfaceNames[i] = self.constantPool.getClassName(cpIndex)
	}
	return interfaceNames
}

/*
	下面 6 个方法是结构体字段的 Getter 方法
 */
func (self *ClassFile) MinorVersion() uint16 {
	return self.minorVersion
}

func (self *ClassFile) MajorVersion() uint16 {
	return self.majorVersion
}

func (self *ClassFile) ConstantPool() ConstantPool {
	return self.constantPool
}

func (self *ClassFile) AccessFlags() uint16 {
	return self.accessFlags
}

func (self *ClassFile) Fields() []*MemberInfo {
	return self.fields
}

func (self *ClassFile) Methods() []*MemberInfo {
	return self.methods
}
