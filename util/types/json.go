package types

import (
	"fmt"
	"github.com/pingcap/tidb/util/hack"
)

type Json struct {
	origin string
}

func JsonFromString(s string) Json {
	return Json{origin: s}
}

func (j *Json) JsonSerialize() []byte {
	s := fmt.Sprintf("serialized: %s", j.origin)
	sink(s)
	return hack.Slice(s)
}

func JsonDeserialize(data []byte) Json {
	return Json{origin: hack.String(data[12:])}
}

func (j *Json) JsonDumpToText() string {
	return j.origin
}
