package d

import "github.com/binary-soup/go-command/d/fs"

// Dependency object for a file system. Default implementation is Disk.
var FileSystem = fileSystemDependency{
	FileSystem: fs.Disk{},
}

type fileSystemDependency struct {
	fs.FileSystem
	cache fs.FileSystem
}

// Override the current implementation with the new one.
// Primary used for testing and mocking.
//
// Note: Overriding is NOT thread-safe.
func (d *fileSystemDependency) Override(new fs.FileSystem) {
	d.cache = d.FileSystem
	d.FileSystem = new
}

// Restore the dependency to its previous value.
func (d *fileSystemDependency) Restore() {
	d.FileSystem = d.cache
}
