package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// FileSystemKeyValueStore stores players in the filesystem.
type FileSystemKeyValueStore struct {
	database *json.Encoder
	keyvalue KeyValue
}

// NewFileSystemKeyValueStore creates a FileSystemPlayerStore initialising the store if needed.
func NewFileSystemKeyValueStore(file *os.File) (*FileSystemKeyValueStore, error) {

	err := initialiseKeyValueDBFile(file)

	if err != nil {
		return nil, fmt.Errorf("problem initialising player db file, %v", err)
	}

	league, err := NewKeyValue(file)

	if err != nil {
		return nil, fmt.Errorf("problem loading player store from file %s, %v", file.Name(), err)
	}

	return &FileSystemKeyValueStore{
		database: json.NewEncoder(&tape{file}),
		keyvalue: league,
	}, nil
}

func initialiseKeyValueDBFile(file *os.File) error {
	file.Seek(0, 0)

	info, err := file.Stat()

	if err != nil {
		return fmt.Errorf("problem getting file info from file %s, %v", file.Name(), err)
	}

	if info.Size() == 0 {
		file.Write([]byte("{}"))
		file.Seek(0, 0)
	}

	return nil
}

// FetchValue returns value from store based on given key
func (f *FileSystemKeyValueStore) FetchValue(key string) string {
	return f.keyvalue.Find(key)
}

// SetValue sets value for given key in store
func (f *FileSystemKeyValueStore) SetValue(key string, value string) {
	f.keyvalue.Set(key, value)
	f.database.Encode(f.keyvalue)
}

// DeleteKey deletes given key in store
func (f *FileSystemKeyValueStore) DeleteKey(key string) {
	delete(f.keyvalue, key)
	f.database.Encode(f.keyvalue)
}
