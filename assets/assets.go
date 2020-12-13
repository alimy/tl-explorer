// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package assets

import (
	"embed"
	"io/fs"
	"net/http"
	"strings"
)

//go:embed schema.tl
var schemaTL embed.FS

type fsFunc func(name string) (fs.File, error)

func (f fsFunc) Open(name string) (fs.File, error) {
	return f(name)
}

// chroot change dir as root directory
func chroot(dir string, emFS embed.FS) fs.FS {
	if dir == "" {
		return emFS
	}

	root := strings.TrimSuffix(dir, "/")
	rootSlash := root + "/"

	return fsFunc(func(name string) (fs.File, error) {
		if name == "." {
			return emFS.Open(root)
		}
		return emFS.Open(rootSlash + name)
	})
}

// NewFS get an assets http.FileSystem instance
func NewFS() http.FileSystem {
	return http.FS(chroot("schema.tl", schemaTL))
}
