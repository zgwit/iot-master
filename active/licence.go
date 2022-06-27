package active

import (
	"encoding/hex"
	"github.com/zgwit/go-license"
	"os"
	"path/filepath"
	"strings"
)

var _filename = "iot-master.lic"
var _license *license.Licence
var _product = "iot-master-ce"
var _key = "e32edc3719d919ef1b6da8bf8def2baa1e7a197866c09e3fbce8d45db6012c61"

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
	key, _ := hex.DecodeString(_key)
	err := lic.Verify(key)
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
