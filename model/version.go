package model

type Version int64

const (
	V1 Version = iota
	V2
)

func (v Version) String() string {
	return [...]string{"v1", "v2"}[v]
}
