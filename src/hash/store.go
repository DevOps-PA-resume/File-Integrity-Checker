package hash

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const hashFilePath = "data/hashes.json"

func readHashes() (map[string]string, error) {
	hashes := make(map[string]string)

	if _, err := os.Stat(hashFilePath); os.IsNotExist(err) {
		return hashes, nil
	}

	data, err := os.ReadFile(hashFilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read hash file: %w", err)
	}

	if err := json.Unmarshal(data, &hashes); err != nil {
		return nil, fmt.Errorf("failed to parse hash file: %w", err)
	}

	return hashes, nil
}

func writeHashes(hashes map[string]string) error {
	if err := os.MkdirAll(filepath.Dir(hashFilePath), 0755); err != nil {
		return err
	}

	data, err := json.MarshalIndent(hashes, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to encode hashes: %w", err)
	}

	return os.WriteFile(hashFilePath, data, 0644)
}

func GetKey(filePath string) (string, bool) {
	hashes, err := readHashes()
	if err != nil {
		return "", false
	}

	val, exists := hashes[filePath]
	return val, exists
}

func WriteKey(filePath, hashValue string) error {
	hashes, _ := readHashes()
	hashes[filePath] = hashValue
	return writeHashes(hashes)
}

func IsFile(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !info.IsDir()
}

func IsDir(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}
