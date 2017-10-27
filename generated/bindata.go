// Code generated by go-bindata.
// sources:
// files/cf-mgmt.sh
// files/cf-mgmt.yml
// files/pipeline.yml
// files/security-group.json
// files/vars.yml
// DO NOT EDIT!

package generated

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

var _filesCfMgmtSh = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x52\x56\xd4\x2f\x2d\x2e\xd2\x4f\xca\xcc\xd3\x4f\xcd\x2b\x53\x48\x4a\x2c\xce\xe0\x2a\x4e\x2d\x51\xd0\x4d\xe5\x4a\x4e\x51\x48\xce\xcf\x4b\xcb\x4c\xd7\x2d\x4a\x2d\xc8\xe7\x4a\x4e\xd3\xcd\x4d\xcf\x2d\x51\x50\x71\x76\x8b\xf7\x75\xf7\x0d\x89\x77\xf6\xf7\xf5\x75\xf4\x73\xe1\x02\x04\x00\x00\xff\xff\x4b\xdb\x02\xbb\x43\x00\x00\x00")

func filesCfMgmtShBytes() ([]byte, error) {
	return bindataRead(
		_filesCfMgmtSh,
		"files/cf-mgmt.sh",
	)
}

func filesCfMgmtSh() (*asset, error) {
	bytes, err := filesCfMgmtShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "files/cf-mgmt.sh", size: 67, mode: os.FileMode(493), modTime: time.Unix(1495821632, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _filesCfMgmtYml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x54\xcd\x41\x8a\xc3\x30\x0c\x85\xe1\xbd\x4f\x21\xb2\x1e\x4f\xf6\xbe\xcc\x60\x3c\x8a\x23\x12\x5b\x46\x92\x43\x43\xe9\xdd\xeb\x94\x52\xe8\xf6\xfd\xf0\x3d\xef\xbd\x6b\x7b\xb4\x85\xa5\x04\xd8\xa9\xf6\x9b\x73\x54\x62\xc6\x3f\x41\xe5\x2e\x09\x83\x03\xb0\xb3\x61\x80\x7f\x4e\x1b\x8a\x7f\xe5\x31\xbe\x33\xdc\x05\x1b\x2b\x19\xcb\x19\xa0\xd1\xc1\x16\x77\x45\x39\x28\xa1\xce\x69\xf1\x25\x17\xfb\x01\x8b\x39\xc0\x34\xae\x50\x6d\x7a\x8c\x93\xda\xba\xe9\x85\x7b\xa8\xb1\x0c\x27\x71\x5d\x28\xfb\x4b\x73\x4e\x7a\xbd\x5a\x8b\xb6\x7e\x95\x39\xd1\x6c\x51\xb7\x8f\xfc\xab\xab\x7b\x06\x00\x00\xff\xff\x0a\x5a\x24\x71\xc6\x00\x00\x00")

func filesCfMgmtYmlBytes() ([]byte, error) {
	return bindataRead(
		_filesCfMgmtYml,
		"files/cf-mgmt.yml",
	)
}

func filesCfMgmtYml() (*asset, error) {
	bytes, err := filesCfMgmtYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "files/cf-mgmt.yml", size: 198, mode: os.FileMode(420), modTime: time.Unix(1495821632, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _filesPipelineYml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x98\xdd\x6e\xda\x30\x14\xc7\xef\x79\x0a\x3f\x00\x29\xda\xc5\x6e\x72\x87\x80\x56\x48\x0d\x54\x84\x69\x9a\xa6\xca\x72\x13\x93\x79\x8b\x63\xcf\x76\x3a\x21\xc4\xbb\xcf\x4e\x3a\x20\xc6\x64\x40\x86\x26\x25\xbd\x4b\xeb\xe3\xf3\xf1\xff\x71\x0e\x27\x08\x2c\x59\x2e\x22\x2c\xfd\x9e\x07\x32\x44\xb1\x0f\x22\x96\xad\x48\xe2\x09\xcc\x59\x0f\x00\xb5\xe6\xfa\x7f\x09\x51\xfa\xb9\x34\xf5\xf5\x13\x00\xb9\x20\x3e\xd8\x6c\xf4\x01\x34\x96\x50\xff\xbd\xdd\x16\x27\x2f\x02\x65\xd1\x37\x1f\x50\x24\x15\x16\x3b\xb7\x1f\x3e\xd2\x9d\x3b\x45\x28\xde\xfb\x03\x1b\x92\x69\xcb\x57\x94\x16\x56\xdb\x5e\xef\x3b\x7b\x39\x4c\x48\x60\xa4\xb0\xc7\x44\x22\xf5\x25\x9e\xa2\xcc\xa4\xe0\x81\x04\x2b\x3b\x5b\x1d\x40\x90\x24\xc1\x42\xc7\x10\x39\x2e\xcc\x14\x92\x3f\x6c\x27\x00\xac\x48\x5a\xad\x75\x10\x91\x81\x31\x95\x83\x68\xe5\xd1\x84\xaa\xbb\x35\x4d\x0b\x53\x8e\x04\xa2\xb2\x2c\x1b\x80\xf0\x4b\xb8\x9c\x04\x70\x3c\x0f\x86\xd3\x99\x91\x40\xae\x75\x9d\x14\xc6\x8c\x22\x92\xbd\x69\x00\xc0\xa7\x70\xb2\x80\xd3\xb1\x31\xc8\x25\x16\x90\xc4\xbb\xa3\xa7\x61\x18\x7e\x9e\x2f\x8a\x33\x8e\xa4\xfc\xc5\xc4\xfe\x70\x34\x9f\xdd\x4f\x1f\xe0\x78\xba\xf0\xc1\xdd\xa0\xcc\xef\xcf\xd1\xe3\x74\x32\x5b\xc2\x70\x32\x5a\x4c\x96\xe6\x72\x94\x12\x9c\x29\x28\xb1\x2e\x4e\xed\x3d\xdc\xc3\xe0\x21\x58\xc2\xd1\x3c\x08\x86\xb3\x71\xb5\x74\x4b\x53\x7d\x55\x83\x53\x6b\x2f\x11\x2c\xe7\xcd\xe5\x3d\xf6\xd7\x45\xa5\x6d\x15\x76\xaa\xc7\x38\xc5\x0d\x3e\xc9\x85\x59\xd9\x47\xa7\x49\x54\x63\x74\x4a\xfe\xc3\xd2\x8f\x87\x87\xc7\x05\x79\x35\xcf\x65\x01\x67\x10\x30\x39\xe3\xd8\x07\x5f\x0f\x3a\xe8\xf9\xcc\x29\xe3\x88\xd6\x66\x16\x8f\xe3\xe1\x13\x3c\x8c\x91\xc6\x88\xc3\xe3\x40\xa7\x67\xd3\x91\x60\xf6\xa8\xe2\x48\x7f\x4d\xdd\x10\xda\x2e\xc0\x3b\xa7\xd3\x93\xad\xd4\xc8\x9a\x67\xb7\x40\x73\xc9\xac\xeb\x04\xb9\x53\xd3\xce\x42\x92\xf3\xb8\x49\xb7\x94\xf7\xea\xfb\xc5\x0e\xd1\x29\xd5\xab\xc5\xbb\x54\xf7\x4c\x8a\xff\x5c\xfa\xb3\xfa\xc1\x99\x46\x9b\xf1\x5c\x3b\xce\x1c\x42\x39\x51\xfe\xcc\x99\xd6\xe9\xf6\x6d\xb4\x0f\xd4\x66\x5a\xf5\x18\xde\x24\x70\x72\xb8\xf8\x3d\xc5\x0d\xa4\xef\x5e\xd0\x2f\xe0\xd4\xf1\x17\x9c\x5a\x2d\x2c\x72\x66\xa7\xbb\x72\x14\x36\x5e\x0c\x1c\x29\xb4\x99\x53\xc3\x31\xb8\x97\xc9\x81\xf0\xda\x11\xf8\xf7\xbd\xdb\x15\xa4\xcd\x94\xea\xe4\x77\x0f\x3f\x22\x99\xee\xb2\x84\x6a\x67\x57\xfc\x60\xe0\x66\xd2\x3f\x7c\x4d\xee\x57\x17\xfb\x7e\x75\xa9\xec\x57\xb7\x9d\xe7\x63\x76\x56\x82\xff\x95\xde\xb9\x14\x6a\x38\x9e\x00\x54\xa9\xf2\x77\x00\x00\x00\xff\xff\xeb\x9c\x6d\x1e\x32\x15\x00\x00")

func filesPipelineYmlBytes() ([]byte, error) {
	return bindataRead(
		_filesPipelineYml,
		"files/pipeline.yml",
	)
}

func filesPipelineYml() (*asset, error) {
	bytes, err := filesPipelineYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "files/pipeline.yml", size: 5426, mode: os.FileMode(420), modTime: time.Unix(1509128591, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _filesSecurityGroupJson = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8a\x8e\xe5\x02\x04\x00\x00\xff\xff\x44\xd2\x68\x70\x03\x00\x00\x00")

func filesSecurityGroupJsonBytes() ([]byte, error) {
	return bindataRead(
		_filesSecurityGroupJson,
		"files/security-group.json",
	)
}

func filesSecurityGroupJson() (*asset, error) {
	bytes, err := filesSecurityGroupJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "files/security-group.json", size: 3, mode: os.FileMode(420), modTime: time.Unix(1495821632, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _filesVarsYml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x9c\x8e\x41\x6e\x43\x21\x0c\x44\xf7\x9c\x62\x6e\xd0\x7d\x55\x71\x15\xe4\xf2\xfd\x53\x4b\xf9\x18\xd9\x46\x51\x6e\x5f\x28\xca\xae\xab\xac\x98\x79\x0c\xc3\xdc\x24\x8a\x71\xd7\x32\x4c\x3e\xf1\xf5\xd4\x61\x98\x0c\x8b\x61\xb2\x9c\xfc\xe9\xc1\x57\x39\xf4\x22\x69\xaf\x48\x3d\xb1\x39\x36\xcf\x69\x38\x5b\x91\x63\x06\x96\x02\xd5\xaa\xa3\x05\x1e\x12\x3f\xe8\x6c\x97\xb8\x8b\x36\x84\xa2\x1a\x53\x30\xd4\x6e\xfe\xe1\x9d\x2a\x7b\x4e\x9d\xdc\x1f\x6a\xeb\xf9\x4b\x42\x4f\xbc\x53\x55\xef\xc2\x2d\x8a\xf3\xbc\x8c\xd9\xb7\x3d\xb6\xc7\xa9\x86\x41\x54\xb7\xd8\x9b\x73\xba\x1f\xd4\xcb\x7f\x1b\xe6\x1f\xdf\xd2\xfe\xce\x95\xc9\xe9\x37\x00\x00\xff\xff\x27\x4c\x31\xa2\x30\x01\x00\x00")

func filesVarsYmlBytes() ([]byte, error) {
	return bindataRead(
		_filesVarsYml,
		"files/vars.yml",
	)
}

func filesVarsYml() (*asset, error) {
	bytes, err := filesVarsYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "files/vars.yml", size: 304, mode: os.FileMode(420), modTime: time.Unix(1495821632, 0)}
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
	"files/cf-mgmt.sh": filesCfMgmtSh,
	"files/cf-mgmt.yml": filesCfMgmtYml,
	"files/pipeline.yml": filesPipelineYml,
	"files/security-group.json": filesSecurityGroupJson,
	"files/vars.yml": filesVarsYml,
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
	"files": &bintree{nil, map[string]*bintree{
		"cf-mgmt.sh": &bintree{filesCfMgmtSh, map[string]*bintree{}},
		"cf-mgmt.yml": &bintree{filesCfMgmtYml, map[string]*bintree{}},
		"pipeline.yml": &bintree{filesPipelineYml, map[string]*bintree{}},
		"security-group.json": &bintree{filesSecurityGroupJson, map[string]*bintree{}},
		"vars.yml": &bintree{filesVarsYml, map[string]*bintree{}},
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

