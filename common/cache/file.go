package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type FileProvider struct {
	path string
}

type FileItem struct {
	Value     string    `json:"value"`
	ExpiresAt time.Time `json:"expires_at"`
	NoExpiry  bool      `json:"no_expiry"`
}

func NewFileProvider(path string) *FileProvider {
	os.MkdirAll(path, 0755)

	return &FileProvider{
		path: path,
	}
}

func (f *FileProvider) Set(ctx context.Context, key string, value any, exp time.Duration) error {
	item := FileItem{
		Value:     value.(string),
		ExpiresAt: time.Now().Add(exp),
	}

	fmt.Printf("DEBUG: exp=%v\n", exp)

	if exp == 0 {
		item.NoExpiry = true
	} else {
		item.ExpiresAt = time.Now().Add(exp)
	}
	data, _ := json.Marshal(item)

	return os.WriteFile(
		filepath.Join(f.path, key+".json"),
		data,
		0644,
	)
}

func (f *FileProvider) Get(ctx context.Context, key string) (string, error) {
	data, err := os.ReadFile(filepath.Join(f.path, key+".json"))
	if err != nil {
		return "", err
	}

	var item FileItem

	if err := json.Unmarshal(data, &item); err != nil {
		return "", err
	}

	if !item.NoExpiry && time.Now().After(item.ExpiresAt) {
		_ = os.Remove(filepath.Join(f.path, key+".json"))
		return "", os.ErrNotExist
	}

	return item.Value, nil
}

func (f *FileProvider) Delete(ctx context.Context, key string) error {
	return os.Remove(filepath.Join(f.path, key+".json"))
}

func (f *FileProvider) Ping(ctx context.Context) error {
	return nil
}

func (f *FileProvider) Close() error {
	return nil
}
