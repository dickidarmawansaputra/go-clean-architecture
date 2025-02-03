package avatar

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/dickidarmawansaputra/go-clean-architecture/internal/lib/storage"
	"github.com/lafriks/go-avatars"
	"golang.org/x/exp/rand"
)

func Generate(name string) (string, error) {
	randomizer := rand.New(rand.NewSource(10))
	fileName := fmt.Sprintf("%s-%d", strings.ToLower(strings.Replace(name, " ", "-", -1)), randomizer.Int())
	avatar, err := avatars.Generate(fileName)
	if err != nil {
		return "", errors.New("Failed to generate user avatar")
	}

	avatarByte, _ := avatar.PNG()

	path, err := storage.CreateStorageDirectory(storage.Public, "user")
	if err != nil {
		return "", errors.New("Failed to create storage directory")
	}

	filePath := path + fileName + ".png"
	if err := os.WriteFile(filePath, avatarByte, 0644); err != nil {
		return "", errors.New("Failed to save avatar")
	}

	return filePath, nil
}
