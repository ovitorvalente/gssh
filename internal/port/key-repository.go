package port

type KeyRepository interface {
	Exists() bool
	Generate() error
	ReadPublicKey() (string, error)
}
