package user_permission

type UserPermission struct {
	GroupName          string `json:"groupName"`
	ResourceName       string `json:"resourceName"`
	PermissionName     string `json:"permissionName"`
	PermissionValue    string `json:"permissionValue"`
	ResourcePermission string `json:"resourcePermission"`
}

type UserResourcePermission struct {
	Permission string `json:"permission"`
}
