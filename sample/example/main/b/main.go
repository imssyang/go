package main

import (
	"example.com/dep"
	"github.com/pkg/profile"
)

func main() {
	defer profile.Start(
		profile.MemProfile,
		profile.MemProfileRate(1),
		profile.ProfilePath(".")).Stop()
	dep.ConcatRandomString(100)
	dep.ConcatRandomStringByBuilder(100)
}
