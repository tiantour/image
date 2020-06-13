package image

import (
	"fmt"
	"path"
	"time"
)

// Format format
type Format struct{}

// NewFormat new format
func NewFormat() *Format {
	return &Format{}
}

// Name format name
func (f *Format) Name(args *File) string {
	name := path.Base(args.Name)

	suffix := path.Ext(name)
	prefix := args.Prefix
	if prefix == 0 {
		prefix = time.Now().UnixNano()
	}

	return fmt.Sprintf("%d%s", prefix, suffix)
}

// Path format path
func (f *Format) Path(args *File) string {
	path := args.Path
	if path == "" {
		path = "file"
	}

	return fmt.Sprintf("%s/%s", path, args.Name)
}
