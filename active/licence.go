package active

import (
	"github.com/zgwit/go-license"
	"os"
	"path/filepath"
	"strings"
)

var _filename = "iot-master.lic"
var _license *license.Licence
var _product = "iot-master-ce"
var _key = []byte{
	0x09, 0xc1, 0x68, 0xe6, 0xb8, 0x19, 0x52, 0x68,
	0xfc, 0xee, 0x59, 0xe2, 0xc9, 0x75, 0x9a, 0xd9,
	0x6f, 0x33, 0xc0, 0x5a, 0x77, 0xca, 0xee, 0xf9,
	0xf7, 0x8d, 0x0a, 0xf2, 0x1e, 0xcf, 0x1f, 0x12,
}

func init() {
	app, _ := filepath.Abs(os.Args[0])
	ext := filepath.Ext(os.Args[0])
	_filename = strings.TrimSuffix(app, ext) + ".lic"
}

func Licence() *license.Licence {
	return _license
}

func Load() error {
	// 如果没有文件，则使用默认信息创建
	if _, err := os.Stat(_filename); os.IsNotExist(err) {
		return nil
	} else {
		content, err := os.ReadFile(_filename)
		if err != nil {
			return err
		}

		if _license == nil {
			_license = &license.Licence{}
		}

		err = _license.Decode(string(content))
		if err != nil {
			return err
		}
	}
	return nil
}

func Validate(lic *license.Licence) error {
	err := lic.Verify(_key)
	if err != nil {
		return err
	}
	return lic.Match(_product)
}

func Save(lic *license.Licence) error {
	_license = lic
	content := lic.Encode()
	return Write(content)
}

func Write(lic string) error {
	return os.WriteFile(_filename, []byte(lic), os.ModePerm)
}
