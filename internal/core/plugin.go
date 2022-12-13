package core

import (
	"github.com/zgwit/iot-master/v3/model"
	"os"
)

type Plugin struct {
	Id      string
	Process *os.Process
	Status  model.Status
}
