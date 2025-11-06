package common

func SomeMapKey[K int | string | Ordinal, V any](m map[K]V) K {
	for k := range m {
		return k
	}
	panic("Empty")
}
