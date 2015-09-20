// Code generated by go-bindata.
// sources:
// scripts/compilation/fake-compile.sh
// scripts/compilation/fake-prerequisites.sh
// scripts/compilation/ubuntu-compile.sh
// scripts/compilation/ubuntu-prerequisites.sh
// DO NOT EDIT!

package compilation

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
	"os"
	"time"
	"io/ioutil"
	"path/filepath"
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
	name string
	size int64
	mode os.FileMode
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

var _scriptsCompilationFakeCompileSh = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x4a\xad\xc8\x2c\x51\x30\x00\x04\x00\x00\xff\xff\x4f\xf8\xd2\x86\x06\x00\x00\x00")

func scriptsCompilationFakeCompileShBytes() ([]byte, error) {
	return bindataRead(
		_scriptsCompilationFakeCompileSh,
		"scripts/compilation/fake-compile.sh",
	)
}

func scriptsCompilationFakeCompileSh() (*asset, error) {
	bytes, err := scriptsCompilationFakeCompileShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "scripts/compilation/fake-compile.sh", size: 6, mode: os.FileMode(438), modTime: time.Unix(1442783009, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _scriptsCompilationFakePrerequisitesSh = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x4a\xad\xc8\x2c\x51\x30\x00\x04\x00\x00\xff\xff\x4f\xf8\xd2\x86\x06\x00\x00\x00")

func scriptsCompilationFakePrerequisitesShBytes() ([]byte, error) {
	return bindataRead(
		_scriptsCompilationFakePrerequisitesSh,
		"scripts/compilation/fake-prerequisites.sh",
	)
}

func scriptsCompilationFakePrerequisitesSh() (*asset, error) {
	bytes, err := scriptsCompilationFakePrerequisitesShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "scripts/compilation/fake-prerequisites.sh", size: 6, mode: os.FileMode(438), modTime: time.Unix(1442783014, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _scriptsCompilationUbuntuCompileSh = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x7c\x90\x41\x6b\xf2\x40\x10\x86\xef\x81\xfc\x87\x97\x7c\xe1\x3b\x14\xd2\xd0\x5e\x4b\x0a\xa9\x04\x2b\xd5\x28\x2a\x5e\x4a\x91\x6d\x32\x31\x83\x66\x37\xec\xae\xd6\xfa\xeb\x9b\x68\x6b\xb5\x88\xc7\xcc\xfb\xcc\xf3\x66\xc7\x90\x45\x40\xf8\x07\xda\xb2\x05\x57\x15\xe5\x2c\x2c\xad\x3e\xc1\x05\x04\x0c\x57\xf5\x8a\x90\xa9\xaa\x12\x32\xdf\x43\x06\x1f\x6c\xcb\x26\x93\x4a\x06\x3b\xd2\x0a\xc6\x0a\xbb\x36\xae\xe3\x3a\xb5\xc8\x96\x62\x41\xa9\xa8\x28\xf2\xef\x8e\xdf\x33\xd2\x86\x95\x8c\xfc\xfb\x16\x6a\xc4\xaf\x08\x76\xf0\xfc\x13\xdc\xc3\xdb\x03\x6c\x49\xd2\x75\x00\xca\x4a\x05\x6f\x74\x48\x21\x9b\xb8\x29\xb3\x30\x35\x65\x5c\x30\xe5\x1e\x1e\xff\x37\xaa\x82\x2f\xea\xbe\xdb\xae\x18\x37\x07\xe2\x8a\xf4\x54\x3b\x8b\xc7\xad\xab\x1d\x56\xcb\x9c\x35\x82\x1a\xe1\x46\xe8\x70\x93\x89\xba\x9d\x66\x35\x02\x8d\xb0\x60\x63\x78\x45\x01\xcb\x63\x1a\xde\x9c\x93\xb4\xad\x95\xb6\x78\x1a\x4e\x9e\xe7\x9d\xe1\x60\xd4\xeb\x27\xf3\x69\x3c\xee\x26\xd3\xe8\x77\xc7\xa8\xb5\xce\xe8\x1c\xee\xa5\x93\x69\xdc\xef\x1f\xe1\x9f\x2e\xb5\xb6\xe7\xe0\x28\xee\xbc\xc4\xdd\x64\x9e\xc6\x83\x24\x3a\x3d\xf0\x65\x6c\x96\x8c\x27\xbd\x61\x1a\xfd\xb9\xdd\xfe\x55\x39\xfc\x0b\x3f\xea\x3a\xef\xc2\x94\x08\xb6\xb8\x0d\x0f\x4b\x2c\x17\xae\xf3\x15\x00\x00\xff\xff\xa1\x06\x62\x12\x48\x02\x00\x00")

func scriptsCompilationUbuntuCompileShBytes() ([]byte, error) {
	return bindataRead(
		_scriptsCompilationUbuntuCompileSh,
		"scripts/compilation/ubuntu-compile.sh",
	)
}

func scriptsCompilationUbuntuCompileSh() (*asset, error) {
	bytes, err := scriptsCompilationUbuntuCompileShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "scripts/compilation/ubuntu-compile.sh", size: 584, mode: os.FileMode(438), modTime: time.Unix(1442769051, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _scriptsCompilationUbuntuPrerequisitesSh = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x4c\x53\x4b\x6b\x74\x37\x0c\xdd\x07\xf2\x1f\x44\x52\xe8\xa6\x37\x43\x52\x5a\x68\x20\x8b\xb6\x49\xa1\x9b\x14\x4a\x4b\x37\x85\xa0\x6b\x6b\xee\x15\xf1\xab\x96\x3c\xaf\x5f\x5f\xd9\x33\xc9\xf7\x2d\xc6\x63\xeb\xea\x71\x74\x74\x24\xa4\x30\x11\xdc\x02\x1d\x58\x81\x63\x24\xcf\xa8\x14\x8e\xc0\x5b\x40\x10\x8e\x25\x10\xb8\x1c\x23\x26\x3f\x9c\x04\xf6\xac\xab\x7d\x4b\x39\x4d\x27\xaa\x19\x44\x51\x9b\x5c\x5f\x8d\x5c\xcd\x72\x55\x2a\xb9\x2a\xe8\x4a\xd0\x04\x17\x82\xbc\x85\x96\x38\xb1\x32\x06\x3e\x91\x87\x1d\x56\xc6\x39\x90\x45\x5d\x5f\xdd\xc2\xdf\x42\xc3\x7b\x9b\x43\xc8\x7b\x4e\x8b\xdd\x2a\x70\x92\xc2\x15\x95\x73\x7a\xec\x5e\xab\x6a\x91\xc7\xcd\x66\xb1\xfa\x6d\xbe\x33\x50\x1b\x17\x72\xf3\xdb\xdc\x92\xaf\xc7\xcd\x9c\x65\xdd\x68\x25\xda\x44\x14\xa5\xba\xb1\x23\x3a\x0a\xe1\x6d\x6e\x1c\xfc\x30\x18\x1a\xe9\xb9\xfa\xef\x1f\x02\x9f\xd3\xb7\x0a\x7b\x4c\x86\x36\x1b\xd8\x33\x0c\xc1\x48\x10\xc9\xad\x98\x58\xa2\x00\x4a\xa7\xe2\x92\xec\xae\x87\xfe\x9a\x93\x22\x27\xaa\x02\xb2\xe6\x16\x3c\xcc\x16\x91\x2b\x41\xe0\x65\xd5\x3d\xf5\xf3\x3b\xe8\x9c\xed\xe9\x2b\x97\xde\xf3\xa8\x6d\xd5\xdc\x4a\xb5\x1e\xa1\xb0\x7b\x1f\x8e\xae\x89\xe6\x68\xf4\x40\x6e\x15\x3c\x15\x4a\x9e\x92\xe3\x33\x49\x9e\x66\x79\xba\x09\x3c\x8b\x84\xc9\xd3\x0e\x82\x18\xab\xa2\x15\x1d\xc1\xcc\xc9\xff\x34\xad\x59\x14\x7c\x92\xa6\x1c\x04\xd4\x15\xdf\x62\x01\x2e\xe3\x3d\x61\x2d\x9d\xd8\x7f\xaf\xaf\x5c\xab\x01\xf6\x8b\x4d\xcb\xd2\xf5\xc7\xf7\x9f\x97\x91\x79\x66\xc9\xa9\x9b\x2a\xa1\x0f\xd6\xe5\x8f\xc3\x6c\x91\x66\x3b\xc4\xf0\x00\x97\xff\x33\x0e\xbb\x4b\xd0\xfb\xbb\xfb\xcf\xeb\xb0\x9f\xb8\xd8\xd0\xfb\x69\x81\x69\x2b\x53\x17\x91\xe5\xdd\x06\x3a\x40\x91\xc8\xe2\x00\x4b\xc1\x6a\xb4\x4d\x67\xc8\x5c\x74\x88\x02\xe4\x28\x5d\x54\x3d\xb2\xca\x31\x39\xc8\x46\x86\xc8\x3a\x09\xd5\x1d\x55\x18\x5d\xd7\xdc\xb4\xf3\x3d\x27\x83\x2e\x24\x3f\x8c\xb2\xff\xb5\xac\x78\xc1\x8a\x9c\xef\x61\xf1\xf3\x68\x0f\xcb\xc3\x64\x3c\x7d\xde\x2f\xd8\xe7\xd3\xc3\x47\x77\x2e\xe2\xbb\x29\xb6\xb1\xff\xf8\xb8\xb8\x7a\x2c\x3a\x5e\x0e\x27\x2b\xa9\xbc\x65\x67\xfb\x21\xdd\x5f\x9c\xb0\xe6\x6c\xc0\xe3\x62\xe2\xcc\xc5\x14\xe0\x5b\xa0\x09\x45\xd8\xf0\x9b\xa8\x6c\x6a\x2b\x85\x62\x90\xab\x2d\x80\x82\xb5\xab\x26\x7f\x0b\xc6\x84\xae\x1a\x1d\x36\x45\xdd\x63\xa5\xa9\x54\x6b\xd2\x0a\xd0\x07\x53\x37\x7d\xec\x74\x18\xab\xf4\xfc\xf2\xcb\xef\x3f\xbf\xbe\xfd\xf6\xe7\x1f\xaf\x7f\xbd\xbc\x3e\x3f\xd9\xee\x71\x32\x89\xa3\x53\xde\x51\x77\x44\xc3\xd9\x47\xda\x8a\x37\x7c\x5f\xde\xb6\x44\x8a\x21\xc0\x74\x84\x6f\xba\x86\xba\xef\xff\x01\x00\x00\xff\xff\x25\x2c\xe0\xcc\xf5\x03\x00\x00")

func scriptsCompilationUbuntuPrerequisitesShBytes() ([]byte, error) {
	return bindataRead(
		_scriptsCompilationUbuntuPrerequisitesSh,
		"scripts/compilation/ubuntu-prerequisites.sh",
	)
}

func scriptsCompilationUbuntuPrerequisitesSh() (*asset, error) {
	bytes, err := scriptsCompilationUbuntuPrerequisitesShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "scripts/compilation/ubuntu-prerequisites.sh", size: 1013, mode: os.FileMode(438), modTime: time.Unix(1442767899, 0)}
	a := &asset{bytes: bytes, info:  info}
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
	if (err != nil) {
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
	"scripts/compilation/fake-compile.sh": scriptsCompilationFakeCompileSh,
	"scripts/compilation/fake-prerequisites.sh": scriptsCompilationFakePrerequisitesSh,
	"scripts/compilation/ubuntu-compile.sh": scriptsCompilationUbuntuCompileSh,
	"scripts/compilation/ubuntu-prerequisites.sh": scriptsCompilationUbuntuPrerequisitesSh,
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
	Func func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"scripts": &bintree{nil, map[string]*bintree{
		"compilation": &bintree{nil, map[string]*bintree{
			"fake-compile.sh": &bintree{scriptsCompilationFakeCompileSh, map[string]*bintree{
			}},
			"fake-prerequisites.sh": &bintree{scriptsCompilationFakePrerequisitesSh, map[string]*bintree{
			}},
			"ubuntu-compile.sh": &bintree{scriptsCompilationUbuntuCompileSh, map[string]*bintree{
			}},
			"ubuntu-prerequisites.sh": &bintree{scriptsCompilationUbuntuPrerequisitesSh, map[string]*bintree{
			}},
		}},
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

