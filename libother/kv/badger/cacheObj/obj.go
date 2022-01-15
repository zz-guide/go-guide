package cacheObj

import (
	"github.com/dgraph-io/badger/v3"
	"kv/common"
)

type BadgerDb struct {
	DB     *badger.DB
	GetNum uint64
	SetNum uint64
	DelNum uint64
	BSync  bool
}

func NewBadgerDB() common.DBInterface {
	return &BadgerDb{}
}

func (db *BadgerDb) Open(path string, sync bool) error {

	opt := badger.DefaultOptions(path).WithSyncWrites(sync)
	//opt.SyncWrites = sync
	database, err := badger.Open(opt)
	if err != nil {
		return err
	}
	db.DB = database
	return nil
}

func (db *BadgerDb) Close() error {
	return db.DB.Close()
}

func (db *BadgerDb) Get(key []byte) ([]byte, error) {
	var ret []byte
	err := db.DB.View(func(tx *badger.Txn) error {

		db.GetNum++
		v, err := tx.Get(key)
		if err != nil {
			return err
		}
		ret, err = v.ValueCopy(nil)
		return nil
	})
	return ret, err
}

func (db *BadgerDb) Set(key, val []byte) error {
	err := db.DB.Update(func(tx *badger.Txn) error {
		db.SetNum++
		return tx.Set(key, val)
	})
	return err
}

func (db *BadgerDb) Del(key []byte) error {
	err := db.DB.Update(func(tx *badger.Txn) error {

		return tx.Delete(key)
	})
	return err
}

func (db *BadgerDb) GetAll() (int, error) {
	var cout int
	err := db.DB.View(func(tx *badger.Txn) error {
		opt := badger.DefaultIteratorOptions
		it := tx.NewIterator(opt)
		defer it.Close()

		for it.Rewind(); it.Valid(); it.Next() {
			// item := it.Item()
			// k := item.Key()
			// err := item.Value(func(v []byte) error {
			//   fmt.Printf("key=%s, value=%s\n", k, v)
			//   return nil
			// })
			// if err != nil {
			//   return err
			// }
			cout++
		}

		return nil
	})
	return cout, err
}
