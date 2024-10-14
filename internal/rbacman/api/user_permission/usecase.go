package user_permission

type UseCase interface {
	PermissionAllowed(userId string, resourcePermission string) *UserResourcePermission
	PermissionList(userId string) []UserPermission
}

type useCase struct {
	Repo Repository
}

func (u *useCase) PermissionAllowed(userId string, resourcePermission string) *UserResourcePermission {
	return u.Repo.PermissionAllowed(userId, resourcePermission)
}

func (u *useCase) PermissionList(userId string) []UserPermission {
	return u.Repo.PermissionList(userId)
}

func NewUseCase(repo Repository) UseCase {
	return &useCase{
		Repo: repo,
	}
}
