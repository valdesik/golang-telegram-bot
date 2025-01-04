package files

import (
	"os"
	"path/filepath"
	"telegram-bot-golang/lib/errors"
	"telegram-bot-golang/storage"
)

type Storage struct {
	path string
}

const defaultPerm = 0777

func New(path string) *Storage {
	return &Storage{path: path}
}
func (s Storage) Save(page *storage.Page) (err error) {
	defer func() { err = errors.WrapIfErr("cant save", err) }()
	filePath := filepath.Join(s.path, page.UserName)
	if err := os.MkdirAll(filePath, defaultPerm); err != nil {
		return err
	}
	return err
}

func (s *Storage) PickRandom(userName string) (page *storage.Page, err error) {
	return nil, nil
}
func (s *Storage) Remove(page *storage.Page) error {
	return nil
}
func (s *Storage) IsExists(page *storage.Page) (bool, error) {
	return false, nil
}
