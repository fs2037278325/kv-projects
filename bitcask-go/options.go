package bitcask_go

type Options struct {
	DirPath string // 数据库数据目录

	// 数据文件的大小
	DataFileSize int64
}
