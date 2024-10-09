package common

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type Storage interface {
	// Base
	Save(key string, bucket string, data []byte) error
	Read(key string, bucket string) ([]byte, error)
	Remove(key string, bucket string) error
	Download(url string, key string, bucket string) error

	// Avatars
	GetAvatar(userId int) ([]byte, error)
	SaveAvatar(userId int, data []byte) error
	DefaultAvatar() ([]byte, error)
}

type FileStorage struct {
	dataPath string
}

func NewFileStorage(dataPath string) Storage {
	return &FileStorage{dataPath: dataPath}
}

func (storage *FileStorage) Read(key string, folder string) ([]byte, error) {
	path := fmt.Sprintf("%s/%s/%s", storage.dataPath, folder, key)
	return os.ReadFile(path)
}

func (storage *FileStorage) Save(key string, folder string, data []byte) error {
	path := fmt.Sprintf("%s/%s", storage.dataPath, folder)
	err := os.MkdirAll(path, 0755)

	if err != nil {
		return err
	}

	os.WriteFile(fmt.Sprintf("%s/%s", path, key), data, os.ModePerm)
	return nil
}

func (storage *FileStorage) Remove(key string, folder string) error {
	path := fmt.Sprintf("%s/%s/%s", storage.dataPath, folder, key)
	return os.Remove(path)
}

func (storage *FileStorage) Download(url string, key string, folder string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return storage.Save(key, folder, data)
}

func (storage *FileStorage) GetAvatar(userId int) ([]byte, error) {
	avatar, err := storage.Read(string(userId), "avatars")
	if err != nil {
		return storage.DefaultAvatar()
	}
	return avatar, nil
}

func (storage *FileStorage) DefaultAvatar() ([]byte, error) {
	return storage.Read("unknown", "avatars")
}

func (storage *FileStorage) SaveAvatar(userId int, data []byte) error {
	return storage.Save(string(userId), "avatars", data)
}
