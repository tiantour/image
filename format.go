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
func (f *Format) Name(args *Image) string {
	name := path.Base(args.Name)

	suffix := path.Ext(name)
	prefix := args.Prefix
	if prefix == "" {
		prefix = fmt.Sprintf("%d", time.Now().UnixNano())
	}

	return fmt.Sprintf("%s%s", prefix, suffix)
}

// Path format path
func (f *Format) Path(args *Image) string {
	path := args.Path
	if path == "" {
		path = "file"
	}

	return fmt.Sprintf("%s/%s", path, args.Name)
}
