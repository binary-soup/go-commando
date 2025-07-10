package d

import "github.com/binary-soup/go-command/d/fs"

var FileSystem = fileSystemDependency{
	FileSystem: fs.Disk{},
}

type fileSystemDependency struct {
	fs.FileSystem
	cache fs.FileSystem
}

func (d *fileSystemDependency) Override(new fs.FileSystem) {
	d.cache = d.FileSystem
	d.FileSystem = new
}

func (d *fileSystemDependency) Restore() {
	d.FileSystem = d.cache
}
