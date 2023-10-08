package memory

import (
	"bufio"
	"encoding/json"
	"os"

	"github.com/Orendev/gokeeper/pkg/logger"
	"github.com/pkg/errors"
)

var ErrNotFoundFile = errors.New("file not found")

// FileStorage - structure describing the File.
type FileStorage struct {
	filePath string
}

type Memory struct {
	ID       string `json:"id"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}

// NewFileStorage - constructor for the File.
func NewFileStorage(filePath string) *FileStorage {
	return &FileStorage{
		filePath: filePath,
	}
}

// Save saves data in a file.
func (f *FileStorage) Save(memo Memory) error {

	file, err := os.OpenFile(f.filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer func() {
		if err := file.Close(); err != nil {
			logger.Log.Sugar().Errorf("error while closing file: %s", err)
		}
	}()

	// serializing the structure in JSON format
	writeData, err := json.Marshal(memo)

	if err != nil {
		return err
	}

	_, err = file.Write(append(writeData, '\n'))
	if err != nil {
		return err
	}

	return nil
}

// Remove delete the file.
func (f *FileStorage) Remove() error {
	return os.Remove(f.filePath)
}

// Memory read the data from the file.
func (f *FileStorage) Memory() (*Memory, error) {
	file, err := os.OpenFile(f.filePath, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err = file.Close(); err != nil {
			logger.Log.Sugar().Errorf("error when closing a file while reading: %s", err)
		}
	}()

	scan := bufio.NewScanner(file)
	model := Memory{}

	if !scan.Scan() {
		return nil, ErrNotFoundFile
	}

	err = json.Unmarshal(scan.Bytes(), &model)

	if err != nil {
		return nil, err
	}

	return &model, nil
}
