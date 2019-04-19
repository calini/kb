package mock

import (
	"os"
	"time"

	"github.com/stretchr/testify/mock"
)

type FileInfo struct{ mock.Mock }

func (mfi *FileInfo) Name() string       { return "mock-machine.log" }
func (mfi *FileInfo) Size() int64        { return 0 }
func (mfi *FileInfo) Mode() os.FileMode  { return os.FileMode(0777) }
func (mfi *FileInfo) ModTime() time.Time { return time.Now() }
func (mfi *FileInfo) IsDir() bool        { return false }
func (mfi *FileInfo) Sys() interface{}   { return nil }
