package domain

type PackageManager interface {
	Install(packages []string) error
	Remove(packages []string) error
	Update() error
	Upgrade() error
}
