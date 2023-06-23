// Code generated by go-bindata.
// sources:
// 001_init.down.sql
// 001_init.up.sql
// DO NOT EDIT!

package sqlite

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var __001_initDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\x09\xf2\x0f\x50\x08\x71\x74\xf2\x71\x55\xc8\xcb\x2f\xc9\x4c\xcb\x4c\x4e\x2c\xc9\xcc\xcf\x2b\x8e\xcf\xc9\x4f\xb7\x06\x04\x00\x00\xff\xff\x5d\x68\xf4\x86\x1d\x00\x00\x00")

func _001_initDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__001_initDownSql,
		"001_init.down.sql",
	)
}

func _001_initDownSql() (*asset, error) {
	bytes, err := _001_initDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "001_init.down.sql", size: 29, mode: os.FileMode(436), modTime: time.Unix(1687521753, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __001_initUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x54\x4b\x6f\xdb\x4c\x0c\xbc\xeb\x57\xf0\xcb\x25\x36\x20\x7d\x28\x02\xa4\x97\xa0\x07\xd5\x56\x11\x23\x7e\xa4\x8a\xdc\x24\x27\x67\xad\xa5\xad\x6d\xa4\x5d\x95\x4b\x25\x71\x7f\x7d\xb1\x6b\xf9\x19\xa3\x0f\xdf\xcc\x25\x67\x38\x43\x8a\xbd\x34\x89\xb3\x04\xb2\xf8\xf3\x30\x01\x6d\x58\x2d\x54\x2e\x58\x19\x6d\x67\xa5\x59\x42\x27\x00\x80\x83\xf8\x4c\x49\xe8\x5d\xc7\x69\xe7\xe2\x63\x17\xc6\x93\x0c\xc6\xd3\xe1\x10\x7a\xd7\x49\xef\x06\x3a\xc7\x89\xff\x7d\x82\xf3\xf3\x6e\x08\x51\x74\x8c\x11\x82\x6e\xaa\x39\x12\x98\x05\x70\x81\xe0\xc8\x50\x33\xad\x42\x68\x4a\x25\x3d\x2d\xab\x0a\x2d\x8b\xaa\x86\x7e\x9c\x25\xd9\x60\x94\xec\x08\xfb\xc9\x97\x78\x3a\xcc\xa0\x37\x4d\xd3\x64\x9c\xcd\xdc\xeb\x5d\x16\x8f\x6e\x43\x57\x19\x45\xbb\xe2\x10\xa4\x60\x04\xa1\xa5\x8f\x39\xc2\x0a\xad\x15\x4b\x04\xa9\x6c\x2d\x38\x2f\x94\x5e\x7a\x42\x43\x6a\xa9\x34\x7c\x8b\xd3\xb5\xc2\x0f\xef\x15\xb6\x29\x1b\x61\x9b\x5f\x14\xb5\xc5\x21\x68\xb1\x26\x71\xaa\x6c\x33\xb7\x2b\xcb\x58\x01\x17\x82\xa1\x10\x16\x72\x42\xc1\x28\xfd\xf3\xbe\x27\xbe\x81\x45\x53\x96\x93\xa3\x26\x2e\x2f\xdf\x77\xb1\x97\xd7\x76\xe2\xcb\x73\xa3\x19\x35\x67\xab\x1a\xb7\xf5\x97\x27\x44\xec\xe7\xed\xd7\x13\x2e\x90\x50\xe7\xe8\x66\xb7\x6d\xe0\xc4\x9c\x0f\x12\x77\x5e\x44\x91\x83\xd8\x3c\x84\x50\x93\x99\x97\x58\x81\x92\x50\x08\x2d\x51\x82\x79\x41\x82\xf9\xca\xab\x17\x25\x23\x29\xbd\x04\x8b\xf4\xa2\x72\xfc\x7f\x3d\x74\x12\xda\xd6\x86\x18\xb2\xe4\x21\xdb\x39\xfc\xe7\x9f\x1b\xfa\xa6\x38\x04\x54\x5c\x20\xc1\x99\xad\xb8\x3e\x03\x43\x87\x73\xc9\x49\xd5\xdc\x6a\xce\x55\xad\x50\xb3\xfd\x47\x42\x2f\x76\x53\x1b\x3a\x11\x4a\x94\xea\x27\x4a\x10\x44\x62\x15\x78\x74\xcb\x6e\xf7\x7e\xb7\x50\xeb\x8c\xbd\x0f\xc5\xfe\x28\x15\x23\x48\x83\x56\x9f\x33\x94\xea\x19\x01\x75\x53\xd9\x16\xb1\x99\x7f\xc7\xbc\x75\xc7\x47\xe6\x46\xae\xda\xbf\xeb\x2d\x6e\x36\xde\x45\x11\xbc\x16\x82\xd1\x99\x8e\x44\x86\x40\x59\x20\xe4\x86\xb4\x9b\x85\x2e\x57\xf0\x5a\xa0\x7e\xff\xb6\x06\xba\x4d\x07\xa3\x38\x7d\x84\x9b\xe4\x11\x3a\x4f\x47\x1f\xf0\x53\xb8\xfb\xc4\xba\x41\xd0\x85\xfb\x41\x76\x3d\x99\x66\x90\x4e\xee\x07\xfd\xab\x20\x8a\x82\xa0\x3d\x2e\xd3\xf1\xe0\xeb\x34\x81\xc1\xb8\x9f\x3c\x80\x92\x6f\xb3\xc3\x3b\xa3\x64\x30\x19\x9f\xba\x3d\x47\x8c\xdd\xab\xbf\x45\xdc\x36\xe6\x65\x9c\x06\xdf\x35\x7f\x15\x04\xce\x76\x16\xdc\xd8\x59\x6e\x24\x86\xe0\x96\xa6\x8d\x80\x8b\xec\x4c\x6b\x77\xf7\x6e\x94\xdd\xfa\xbd\x75\x17\x8c\x7c\x08\xdf\x14\xaf\x93\x0f\x97\xcc\x61\x57\x5c\xcf\xbc\xc9\x36\x6c\xcd\xde\x9c\xa0\x2d\xf0\x82\x4c\x75\x0c\xed\x6b\x3d\xca\x8c\xd0\xd6\x46\x5b\x0c\x61\xa1\xc8\x32\x5c\x3c\x3b\x1a\xcb\xd2\x4d\xdb\xdd\x36\xcb\x12\x89\x0e\xb9\x7f\x05\x00\x00\xff\xff\x19\x69\xa1\x6e\xdd\x05\x00\x00")

func _001_initUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__001_initUpSql,
		"001_init.up.sql",
	)
}

func _001_initUpSql() (*asset, error) {
	bytes, err := _001_initUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "001_init.up.sql", size: 1501, mode: os.FileMode(436), modTime: time.Unix(1687530028, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"001_init.down.sql": _001_initDownSql,
	"001_init.up.sql":   _001_initUpSql,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//
//	data/
//	  foo.txt
//	  img/
//	    a.png
//	    b.png
//
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"001_init.down.sql": &bintree{_001_initDownSql, map[string]*bintree{}},
	"001_init.up.sql":   &bintree{_001_initUpSql, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
