package bitcask_go

import "errors"

var (
	ErrKeyIsEmpty            = errors.New("the key is empty")
	ErrIndexUpdateFoiled     = errors.New("failed to update index")
	ErrKeyNotFound           = errors.New("key not found in database")
	ErrDataFileNotFound      = errors.New("data file not found")
	ErrDatabaseIsUsing       = errors.New("the database directory is used by another process")
	ErrDataDiretoryCorrupted = errors.New("the data directory maybe corrupted")
)
