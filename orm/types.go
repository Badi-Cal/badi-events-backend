package orm

import (
	"fmt"
	"time"
)

type JSONTime time.Time // JS Date#toISOString

func (t JSONTime) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format(time.RFC3339))
	return []byte(stamp), nil
}
