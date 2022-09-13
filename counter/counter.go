package counter

type Counter interface {
	Add(uint64) uint64
	Read() uint64
}
