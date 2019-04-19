package mock

import (
	"os"
	"time"

	"github.com/stretchr/testify/mock"
)

// FileInfo is a mock FileInfo
type FileInfo struct{ mock.Mock }

// Name ...
func (mfi *FileInfo) Name() string { return "mock-machine.log" }

// Size ...
func (mfi *FileInfo) Size() int64 { return 0 }

// Mode ...
func (mfi *FileInfo) Mode() os.FileMode { return os.FileMode(0777) }

// ModTime ...
func (mfi *FileInfo) ModTime() time.Time { return time.Now() }

// IsDir ...
func (mfi *FileInfo) IsDir() bool { return false }

// Sys ...
func (mfi *FileInfo) Sys() interface{} { return nil }
