package classfile

// LineNumberTableAttribute 结构体
type LineNumberTableAttribute struct {
	lineNumberTable []*LineNumberTableEntry
}

// LineNumberTableEntry 结构体
type LineNumberTableEntry struct {
	startPc    uint16
	lineNumber uint16
}

// 读取信息
func (self *LineNumberTableAttribute) readInfo(reader *ClassReader) {
	lineNumberTableLength := reader.readUint16()
	self.lineNumberTable = make([]*LineNumberTableEntry, lineNumberTableLength)
	for i := range self.lineNumberTable {
		self.lineNumberTable[i] = &LineNumberTableEntry{
			startPc:    reader.readUint16(),
			lineNumber: reader.readUint16(),
		}
	}
}
