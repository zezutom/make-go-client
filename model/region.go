package model

type Region int64

const (
	EU Region = iota
	US
)

func (v Region) String() string {
	return [...]string{"Europe", "USA"}[v]
}
