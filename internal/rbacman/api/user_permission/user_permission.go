package user_permission

type UserPermission struct {
	GroupName          string `json:"groupName"`
	ResourceName       string `json:"resourceName"`
	PermissionName     string `json:"permissionName"`
	PermissionCode     string `json:"permissionCode"`
	ResourcePermission string `json:"resourcePermission"`
}
