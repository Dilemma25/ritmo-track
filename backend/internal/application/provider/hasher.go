package provider

type HasherProvider interface {
	Hash(password string) string
	CompareHashAndPassword(hash string, password string) bool
}
