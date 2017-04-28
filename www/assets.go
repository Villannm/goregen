// Code generated by go-bindata.
// sources:
// html/home.html
// DO NOT EDIT!

package www

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

var _htmlHomeHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x92\xc1\x6e\xe2\x30\x10\x86\xcf\xce\x53\x18\x9f\x82\x94\x8d\xc5\x6d\xb5\xeb\xe4\xb0\x2c\xab\xed\x09\x04\x69\xa5\x1e\x8d\x33\x4d\xac\x9a\x31\x72\x26\x15\x08\xf1\xee\x95\x93\xa6\x70\xe0\x92\x19\xcf\x7c\xff\x68\xfe\x51\xd4\xec\xef\x7a\x59\xbd\x6e\x56\xbc\xa5\x83\x2b\x13\x15\x03\x77\x1a\x9b\x42\x00\x8a\x58\x00\x5d\x97\x09\x53\x07\x20\xcd\x4d\xab\x43\x07\x54\x88\xe7\xea\xdf\x8f\x9f\x22\xd6\xc9\x92\x83\x72\x0b\x0d\xe0\xde\x9f\x94\x1c\xdf\x09\x53\x9d\x09\xf6\x48\x65\xc2\x58\xed\x4d\x7f\x00\xa4\xbc\x01\x5a\x39\x88\xe9\x9f\xf3\x53\x9d\x8a\x00\xc6\x23\x82\x21\x31\xcf\x3d\xa6\xc2\x38\x6b\xde\x45\xc6\xdf\x7a\x34\x64\x3d\xa6\x73\x7e\x49\x18\x63\xa7\x36\xe4\xfe\x08\x98\x8a\xcd\x7a\x57\x89\x8c\xdf\x29\x33\x4e\xa1\x87\xf9\xef\xc8\x7d\xe8\xc0\x21\x04\x5e\xf0\xa8\xe8\x00\xeb\xf4\x7b\x94\xfe\x9a\xc5\x8c\xc7\xce\x3b\xc8\x9d\x6f\x52\x6a\x6d\x97\x71\x3d\xaa\xaf\x43\x18\xbe\x4a\x4e\xdb\x2b\x39\x1e\x40\xed\x7d\x7d\x8e\xb6\xda\x45\xf9\x1f\x9c\xf3\x7c\xb2\x3c\x53\xb2\x5d\xc4\x4e\xef\xa2\x59\xe5\x6c\x0c\x6c\x47\x9a\xe0\x17\xbf\x5c\xf2\x21\xbb\x5e\x63\x4f\x8e\xcd\x89\x59\xb6\x3a\x34\x70\x23\xef\xde\x0f\xf9\x0a\xb0\xb3\x1e\x07\xf6\xc5\x3b\xd2\xcd\x63\x4e\xed\x7b\x22\x8f\xdc\xd6\xc5\xdd\xa5\xca\xed\x94\x2a\x39\x02\xe5\x4d\xab\x64\x5c\x5f\xc9\xd1\xa6\x92\xe3\xef\xf0\x19\x00\x00\xff\xff\x37\x5f\x0e\xbd\x1f\x02\x00\x00")

func htmlHomeHtmlBytes() ([]byte, error) {
	return bindataRead(
		_htmlHomeHtml,
		"html/home.html",
	)
}

func htmlHomeHtml() (*asset, error) {
	bytes, err := htmlHomeHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "html/home.html", size: 543, mode: os.FileMode(420), modTime: time.Unix(1493357275, 0)}
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
	"html/home.html": htmlHomeHtml,
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
	"html": &bintree{nil, map[string]*bintree{
		"home.html": &bintree{htmlHomeHtml, map[string]*bintree{}},
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
