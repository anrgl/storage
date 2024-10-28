package file

import (
	"fmt"

	"github.com/google/uuid"
)

type File struct {
	FileID uuid.UUID
	Name   string
	Data   []byte
}

func NewFile(filename string, blob []byte) (*File, error) {
	fileID, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	file := &File{
		FileID: fileID,
		Name:   filename,
		Data:   blob,
	}

	return file, nil
}

func (f *File) String() string {
	return fmt.Sprintf("File(%s, %v)", f.Name, f.FileID)
}
