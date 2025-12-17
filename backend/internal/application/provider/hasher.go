package provider

type HasherProvider interface {
	Hash(password string) string
	CompareHash(password, hash string) bool
}
