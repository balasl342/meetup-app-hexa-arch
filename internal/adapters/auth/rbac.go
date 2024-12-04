package auth

type Role string

const (
	Admin       Role = "admin"
	Organizer   Role = "organizer"
	Participant Role = "participant"
)

type RBACService struct{}

func (r *RBACService) CheckAccess(userRole Role, requiredRole Role) bool {
	return userRole == requiredRole || userRole == Admin
}
