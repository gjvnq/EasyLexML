// Code generated by go-bindata. DO NOT EDIT.
// sources:
// res/standalone.html

package easyLexML

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
	info  fileInfoEx
}

type fileInfoEx interface {
	os.FileInfo
	MD5Checksum() string
}

type bindataFileInfo struct {
	name        string
	size        int64
	mode        os.FileMode
	modTime     time.Time
	md5checksum string
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
func (fi bindataFileInfo) MD5Checksum() string {
	return fi.md5checksum
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _bindataResStandalonehtml = []byte(
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x54\x4f\x6f\xdb\x3e\x0c\x3d\xbb\x9f\x82\xbf\x16\x05\xfa\x03\xe2\xc4" +
		"\x4e\x9b\x0d\x50\xbc\x60\x45\xb7\xf3\x76\xe8\x65\x47\xc6\x92\x23\xa1\xb2\x64\xc8\x74\xb2\x2c\xc8\x77\x1f\x24\xd9" +
		"\x8d\xd3\x3f\xeb\x25\x66\xc8\xf7\x9e\x68\xea\x99\xc5\x7f\xdf\x7e\x3c\x3c\xfe\xfa\xf9\x1d\x24\xd5\x7a\x75\x51\xc4" +
		"\x47\x52\x48\x81\x7c\x75\x91\x24\x05\x29\xd2\x62\x75\x38\x4c\x1f\x7d\x70\x3c\x16\xb3\x98\xf1\xb5\x96\xf6\x31\x4a" +
		"\xd6\x96\xef\xe1\xe0\xa3\xa4\x46\xb7\x51\x26\x25\xdb\x30\x98\x3b\x51\x2f\x43\xb6\xb2\x86\xd2\x0a\x6b\xa5\xf7\x0c" +
		"\x5a\x34\x6d\xda\x0a\xa7\xaa\x51\xb1\x55\x7f\x04\x83\x7c\xde\xd0\x28\xb9\x45\xa7\xd0\x50\x6a\xba\x5a\x38\x55\x32" +
		"\x20\x5c\x77\x1a\x9d\x4f\xb4\x01\x77\xf4\x3f\x5f\x6b\xc1\x15\xc2\x4d\xad\x4c\xba\x53\x9c\x24\x83\x3b\x7f\xf6\xff" +
		"\x7d\x4f\xa3\xf6\x92\x71\x3d\x9e\x34\xb4\xac\x45\x45\x0c\x4a\xd4\xe5\xcd\x22\xbb\x86\x14\xe6\xb9\xd7\x88\xa0\xe3" +
		"\xf3\x61\xad\x28\x49\x59\x73\x45\xb6\x84\x4e\xf7\xb2\x5a\xb5\x94\x86\x79\xa4\xb4\x6f\x04\x03\x63\x8d\x58\xbe\xc9" +
		"\x59\x9d\x58\x0d\x72\xae\xcc\x86\x41\x76\x82\xca\x7c\x02\x72\x3e\x01\x79\x3b\x01\x79\x37\x01\xb9\x98\x80\xfc\xd4" +
		"\x13\x48\xfc\xa6\x14\xb5\xda\x18\x06\xa5\x30\x24\xdc\x88\x78\x0b\x38\xd5\xb8\x16\xda\x93\x87\x98\x6d\x55\xab\x48" +
		"\xf0\x5e\xa0\xb4\xda\x3a\x06\x57\x59\x16\x8f\x8c\x8a\x5c\x94\xd6\xa1\xef\xf0\x65\xe3\xcd\x49\xb3\xf9\x40\xf2\xfe" +
		"\xfe\x7e\x39\x76\x80\x53\x1b\x49\x0c\xb2\xe9\xe2\x79\xd2\x1f\x1e\x36\xb6\xd0\x30\x95\x33\x47\x65\xd3\xd3\xbd\xf5" +
		"\xf9\xb5\x25\xb2\xf5\x59\x29\xce\xc3\xbd\xe1\xc8\xfc\x3d\x76\x7e\xc6\x65\x84\x6e\x23\xa8\x17\x40\xa3\xea\xd0\x70" +
		"\x6a\xb0\x16\x0c\xa4\xda\x48\xed\xdf\x6e\xf9\xa2\xcc\xbb\xe1\xc5\xe6\x63\x77\x3e\x89\x7d\xe5\xb0\x16\xed\x89\xd9" +
		"\x2b\x67\xd7\x70\x58\x63\xf9\xb4\x71\xb6\x33\x3c\x1d\x46\x59\x55\xb8\x0c\xd4\xe4\xf3\xe2\x23\x44\x9e\xbd\x2d\x12" +
		"\x46\x7b\x72\x2d\x4e\x00\x5f\x5c\xdc\xfb\xb7\x91\xec\xac\xe3\xe9\xda\x09\x7c\x62\x10\x1e\x29\x6a\x1d\x4b\x76\x2b" +
		"\x5c\xa5\xed\x2e\xdd\x39\x6c\x86\xaa\xc7\x2f\xc7\x6e\x58\xeb\xee\xdc\x45\x4c\x7a\x62\x70\x11\xc3\x92\xd4\x56\x44" +
		"\x97\x0e\xf9\x10\xc7\xc2\x7b\xed\x75\x86\x0b\xa7\xd5\xc8\x31\xc5\x6c\x58\x41\xc5\xac\xdf\x58\x85\xff\xd6\xc3\x76" +
		"\xf2\x09\xe1\x40\xf1\x2f\x97\x31\xbc\x0c\xab\xaa\x90\xf9\xd9\x3a\x93\x79\x40\xcf\x22\x26\xee\xb5\xf8\xb5\x06\x6a" +
		"\x2d\x08\x39\x12\x5e\x46\x54\x5f\x7a\x05\x23\x5b\x46\x79\x2f\x6d\xcb\xe3\xf1\x9f\xe8\xd2\xba\xa6\x6b\x9f\x09\x0f" +
		"\xe1\xef\x2b\x4e\x31\x8b\xef\x52\xcc\xc2\x56\xfe\x1b\x00\x00\xff\xff\x97\x6a\xad\x1a\xac\x05\x00\x00")

func bindataResStandalonehtmlBytes() ([]byte, error) {
	return bindataRead(
		_bindataResStandalonehtml,
		"res/standalone.html",
	)
}

func bindataResStandalonehtml() (*asset, error) {
	bytes, err := bindataResStandalonehtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name:        "res/standalone.html",
		size:        1452,
		md5checksum: "",
		mode:        os.FileMode(420),
		modTime:     time.Unix(1552348877, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

//
// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
//
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, &os.PathError{Op: "open", Path: name, Err: os.ErrNotExist}
}

//
// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
// nolint: deadcode
//
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

//
// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or could not be loaded.
//
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, &os.PathError{Op: "open", Path: name, Err: os.ErrNotExist}
}

//
// AssetNames returns the names of the assets.
// nolint: deadcode
//
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

//
// _bindata is a table, holding each asset generator, mapped to its name.
//
var _bindata = map[string]func() (*asset, error){
	"res/standalone.html": bindataResStandalonehtml,
}

//
// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
//
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, &os.PathError{
					Op:   "open",
					Path: name,
					Err:  os.ErrNotExist,
				}
			}
		}
	}
	if node.Func != nil {
		return nil, &os.PathError{
			Op:   "open",
			Path: name,
			Err:  os.ErrNotExist,
		}
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

var _bintree = &bintree{Func: nil, Children: map[string]*bintree{
	"res": {Func: nil, Children: map[string]*bintree{
		"standalone.html": {Func: bindataResStandalonehtml, Children: map[string]*bintree{}},
	}},
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
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
