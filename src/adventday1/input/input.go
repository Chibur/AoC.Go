package input

import (
	"../../adventclient"
)

func GetFrequencies() []string {
	return adventclient.GetStringList("2018/day/1/input")
}
