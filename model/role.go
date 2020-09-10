package model

type Role int

const (
	// MEMBER normal permission
	MEMBER Role = iota

	// ADMIN access all resources
	ADMIN
)

func (r Role) String() string {
	return []string{"MEMBER", "ADMIN"}[r]
}
