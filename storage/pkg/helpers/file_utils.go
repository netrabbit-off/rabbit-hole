package helpers

import (
	"os"

	"github.com/netrabbit-off/rabbit-hole/pkg/models"
)

func CreateFile(file *models.File, content []byte) error {
	created, err := os.Create("data/" + file.ID)
	if err != nil {
		return err
	}

	_, err = created.Write(content)
	if err != nil {
		return err
	}

	return nil
}

func ReadFile(id string) ([]byte, error) {
	content, err := os.ReadFile("data/" + id)
	if err != nil {
		return nil, err
	}
	return content, nil
}
