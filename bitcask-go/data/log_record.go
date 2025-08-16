package data

type LogRecordType = byte

const (
	LogRecordNormal LogRecordType = iota
	LogRecordDeleted
)

// LogRecord 写入到数据文件到记录
// 之所以叫日志，是因为数据文件中的数据是追加写入的， 类似日志的格式
type LogRecord struct {
	Key   []byte
	Value []byte
	Type  LogRecordType
}

// LogRecordPos 数据内存索引，主要描述数据在磁盘上的位置
type LogRecordPos struct {
	Fid    uint32 //文件 id， 表示将数据存储到哪个文件当中
	Offset int64  //偏移，表示将数据存储到了数据文件中的哪个位置
}

// 对LogRecord进行编码， 返回字节数组及长度
func EncodeLogRecord(LogRecord *LogRecord) ([]byte, int64) {
	return nil, 0
}
