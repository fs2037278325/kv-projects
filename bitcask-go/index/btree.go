package index

import (
	"bitcask-go/data"
	"github.com/google/btree"
	"sync"
)

// BTree 索引， 主要封装了 google 的 btree 库
// "https://github.com/google/btree"
type Btree struct {
	tree *btree.BTree
	lock *sync.RWMutex
}

// NewBTree 初始化BTree索引结构
func NewBTree() *Btree {
	return &Btree{
		tree: btree.New(32),
		lock: new(sync.RWMutex),
	}
}

func (bt *Btree) Put(key []byte, pos *data.LogRecordPos) bool {
	it := &Item{key: key, pos: pos}
	bt.lock.Lock()
	bt.tree.ReplaceOrInsert(it)
	bt.lock.Unlock()
	return true
}

func (bt *Btree) Get(key []byte) *data.LogRecordPos {
	it := &Item{key: key}
	btreeItem := bt.tree.Get(it)
	if btreeItem == nil {
		return nil
	}
	return btreeItem.(*Item).pos
}

func (bt *Btree) Delete(key []byte) bool {
	it := &Item{key: key}
	bt.lock.Lock()
	oldItem := bt.tree.Delete(it)
	bt.lock.Unlock()
	if oldItem == nil {
		return false
	}
	return true
}
