package data

import "bitcask-go/fio"

const DataFileNameSuffix = ".data"

// DataFile 数据文件
type DataFile struct {
	FileId    uint32        //文件id
	WriteOff  int64         // 写入位置
	IoManager fio.IOManager // IO操作管理器
}

// OpenDataFile 打开新的数据文件
func OpenDataFile(dirPath string, fileId uint32, ioType fio.FileIOType) (*DataFile, error) {
	return nil, nil
}

func (df *DataFile) ReadLogRecord(offset int64) (*LogRecord, int64, error) {
	return nil, 0, nil
}

func (df *DataFile) Write(buf []byte) error {
	return nil
}

func (df *DataFile) Sync() error {
	return nil
}
