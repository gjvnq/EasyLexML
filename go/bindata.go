// Code generated by go-bindata.
// sources:
// res/standalone.html
// DO NOT EDIT!

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

var _resStandaloneHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x55\x4f\x6f\xfb\x36\x0c\x3d\xbb\x9f\x82\x6b\x51\xa0\x03\xe2\xc4\x49\x9b\x6d\x50\xdc\x60\x45\xb7\x9d\x06\x6c\x87\xee\xb0\xa3\x6c\xd3\x96\x56\x59\x32\x64\x3a\x59\x1a\xe4\xbb\x0f\x92\xec\xc4\x49\xdb\xf5\x77\x89\x15\xf2\xbd\x67\xfe\x13\x9d\x7e\xf7\xcb\x1f\xcf\x2f\x7f\xff\xf9\x2b\x08\xaa\xd5\xfa\x2a\x0d\x8f\x28\x15\xc8\x8b\xf5\x55\x14\xa5\x35\x12\x87\x5c\x70\xdb\x22\x3d\x5e\xff\xf5\xf2\x5b\xfc\xd3\xb5\x77\x90\x24\x85\xeb\xfd\x7e\xfa\xe2\x0e\x87\x43\x3a\x0b\x16\xe7\x6b\x69\x17\x4e\x51\x66\x8a\x1d\xec\xdd\x29\xaa\xb9\xad\xa4\x8e\xc9\x34\x0c\x16\x16\xeb\x95\xb7\x96\x46\x53\x5c\xf2\x5a\xaa\x1d\x83\x96\xeb\x36\x6e\xd1\xca\x72\xe4\x6c\xe5\x1b\x32\x98\x2f\x1a\x1a\x19\x37\xdc\x4a\xae\x29\xd6\x5d\x8d\x56\xe6\x0c\x88\x67\x9d\xe2\xd6\x19\xda\x80\x23\xfc\x97\x62\xae\x64\xa5\x19\xfc\xd3\xb5\x24\xcb\xdd\xc8\x61\x51\x17\x68\xa5\xae\x18\x98\x86\x64\x2d\xdf\xf0\x77\xac\x64\x26\x95\xa4\x80\x3b\xb8\x9f\x9f\x6b\x2c\x24\x87\xbb\x5a\xea\x78\x2b\x0b\x12\x0c\x1e\x5c\xf0\xdf\xf7\x49\x8d\xf2\x8b\xc6\xfe\xf0\xa6\x21\x67\x85\x25\x31\xc8\xb9\xca\xef\x96\xc9\x2d\xc4\xb0\x98\x3b\x8d\x00\x3a\x1c\x5f\xd6\x62\x4e\xd2\xe8\x1b\x32\x39\x74\xaa\x97\x55\xb2\xa5\xd8\x17\x34\xa6\x5d\x83\x0c\xb4\xd1\xb8\xfa\x90\xb3\x3e\xb1\x1a\x5e\x14\x3e\xb9\xe4\x04\x15\xf3\x09\x88\xc5\x04\xc4\xfd\x04\xc4\xc3\x04\xc4\x72\x02\xe2\x87\x9e\x30\xae\x56\x8e\x9a\xd0\x8e\x88\xf7\xc0\xa7\x8a\x67\xa8\x1c\x79\x38\xb3\x8d\x6c\x25\x61\xd1\x0b\xe4\x46\x19\xcb\xe0\x26\x49\x92\x51\x99\x0b\xcc\x8d\xe5\x2e\xc2\xcb\xc0\x9b\x93\x66\x73\x29\x39\x81\x3e\x2d\x58\x9f\x50\xef\x4c\x9f\x04\xf0\xf4\xf4\xb4\x1a\x0f\x9c\x95\x95\x20\x06\xc9\x74\x79\xec\xcb\x97\xa1\x8d\x27\x76\xa8\xe1\xd9\x00\xcf\x8f\x5a\xbd\x35\x33\x44\xa6\x1e\x39\xce\xba\x93\x1b\xdb\x74\xed\x90\xc2\x34\x57\x2d\xac\x8f\x09\x9d\x5d\x8f\x30\x2a\x27\x15\xf0\x0d\xb0\x1f\xdc\xa1\x6f\x0c\x81\x11\xb7\x15\x52\x2f\xc0\xb5\xac\x7d\xce\xb1\xe6\x35\x32\x10\xb2\x12\xca\x15\x68\x75\xe1\x2e\xba\xa1\x36\x8b\x76\x74\x1d\x5e\x71\x57\x5a\x5e\x63\x7b\x62\xf6\xca\xc9\x2d\xec\x33\x9e\xbf\x56\xd6\x74\xba\x88\x87\x6e\x94\x25\x5f\x79\x6a\xf4\xe3\xf2\x2b\xc4\x3c\xf9\x58\xc4\x77\xe7\x74\x4d\xf8\x04\xf8\x45\xef\x3f\x6f\x68\xb4\x35\xb6\x88\x33\x8b\xfc\x95\x81\x7f\xc4\x5c\xa9\xe0\x32\x1b\xb4\xa5\x32\xdb\x78\x6b\x79\x33\x78\x1d\x7e\x35\x1e\xa8\x4c\x75\xe7\x63\xcb\x84\x23\xfa\xb1\x65\x3c\x27\xb9\xc1\x70\x2d\x06\xbb\x3f\x07\xc7\x67\xe1\x75\x6e\xf9\x28\x39\x1a\xba\x74\x36\x2c\xcd\x74\xd6\x2f\xdf\xd4\x2d\x17\xbf\x4f\x9d\x01\x2d\xc8\xe2\xf1\x3a\x1c\xfd\x0a\x8e\x52\x31\x3f\x5b\xc0\x62\xee\xd1\xb3\x80\x09\x9b\xb8\x9f\x31\x47\x75\xab\xbc\xe0\xc4\xc3\xfe\x9e\xf5\xae\x77\x30\x32\x79\x90\x77\xd2\x26\x3f\x1c\xfe\x17\x1d\x26\xfb\x48\x78\xf6\x7f\xdf\x71\xd2\x59\xc8\x25\x9d\xf9\x0f\xcc\x7f\x01\x00\x00\xff\xff\x65\x9b\xd2\x76\x77\x06\x00\x00")

func resStandaloneHtmlBytes() ([]byte, error) {
	return bindataRead(
		_resStandaloneHtml,
		"res/standalone.html",
	)
}

func resStandaloneHtml() (*asset, error) {
	bytes, err := resStandaloneHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "res/standalone.html", size: 1655, mode: os.FileMode(420), modTime: time.Unix(1552392077, 0)}
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
	"res/standalone.html": resStandaloneHtml,
}

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
	"res": &bintree{nil, map[string]*bintree{
		"standalone.html": &bintree{resStandaloneHtml, map[string]*bintree{}},
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
