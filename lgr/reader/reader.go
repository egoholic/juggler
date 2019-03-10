package reader

import (
	"fmt"
	"time"

	"github.com/egoholic/juggler/lgr/lmsg"

	"github.com/egoholic/juggler/lgr/reader/source"
)

// `Reader` is responsible for reading collected data.
type Reader struct {
	source source.Source
}

// Returns new Reader.
func New(source source.Source) *Reader {
	return &Reader{source}
}

func (r *Reader) InBetween(begin, end time.Time) (result []*lmsg.Message, err error) {
	if begin.After(end) {
		return nil, fmt.Errorf("`begin` should be before `end`.\n\tBegin: %#v | End: %#v", begin, end)
	}

	return []*lmsg.Message{}, nil
}
