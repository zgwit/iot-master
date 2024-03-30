package tunnel

import "github.com/zgwit/iot-master/v4/lib"

var links lib.Map[Link]

func GetLink(id string) *Link {
	return links.Load(id)
}
