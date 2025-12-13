package provider

type MigrationProvider interface {
	Upgrade() error
}
