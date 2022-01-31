package interval

import "github.com/asdine/storm/v3"

var Storm *storm.DB

func init()  {
	Storm, _ = storm.Open("my.db")
}
