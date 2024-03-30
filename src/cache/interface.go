package cache

type WebCache interface {
	Inc(string) (int64, error)
}
