package classpath

import (
	"path/filepath"
	"io/ioutil"
)

type DirEntry struct {
	absDir string
}

func newDirEntry(dir string) *DirEntry {
	absDir, err := filepath.Abs(dir)
	if err != nil{
		panic(err)
	}
	return  &DirEntry{absDir}
}

func (self *DirEntry) ReadClass(className string) ([]byte, Entry, error) {
	fileName := filepath.Join(self.absDir,className)
	data, err := ioutil.ReadFile(fileName)
	return data, self, err
}

func (self *DirEntry) String() string  {
	return  self.absDir
}