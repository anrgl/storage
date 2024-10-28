package storage

import (
	"fmt"

	"github.com/anrgl/storage/internal/file"
	"github.com/google/uuid"
)

type Storage struct {
	files map[uuid.UUID]*file.File
}

func NewStorage() *Storage {
	return &Storage{
		files: make(map[uuid.UUID]*file.File),
	}
}

func (s *Storage) Upload(filename string, blob []byte) (*file.File, error) {
	f, err := file.NewFile(filename, blob)
	if err != nil {
		return nil, err
	}
	s.files[f.FileID] = f
	return f, nil
}

func (s *Storage) GetByID(fileID uuid.UUID) (*file.File, error) {
	f, ok := s.files[fileID]
	if !ok {
		return nil, fmt.Errorf("file %v not found", fileID)
	}
	return f, nil
}
