package user_permission

type Repository interface {
	PermissionAllowed(userId string, resourcePermission string) *UserPermission
	PermissionList(userId string) []UserPermission
}

type repository struct {
	Ds DataSource
}

func (r *repository) PermissionAllowed(userId string, resourcePermission string) *UserPermission {
	return r.Ds.PermissionAllowed(userId, resourcePermission)
}

func (r *repository) PermissionList(userId string) []UserPermission {
	return r.Ds.PermissionList(userId)
}

func NewRepository(ds DataSource) Repository {
	return &repository{
		Ds: ds,
	}
}
