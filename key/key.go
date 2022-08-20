package key

type Key interface {
	GenerateKey(name string, size int)
}
