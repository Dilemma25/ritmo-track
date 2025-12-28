package provider

type MigrationProvider interface {
	Upgrade() error
	Down() error
	Reset() error
}
