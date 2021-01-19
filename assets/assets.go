// Copyright 2020 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package assets

import (
	"embed"
	"net/http"

	"github.com/alimy/embedx"
)

//go:embed schema.tl
var content embed.FS

// NewFileSystem get an assets http.FileSystem instance
func NewFileSystem() http.FileSystem {
	embedFS := embedx.ChangeRoot(content, "schema.tl")
	return http.FS(embedFS)
}
