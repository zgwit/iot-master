package core

import (
	"os"
)

type Plugin struct {
	Id      string
	Process *os.Process
}
