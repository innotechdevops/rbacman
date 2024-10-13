package permission

type Repository interface {
	Count(params Params) int64
	FindList(params Params) []Permission
	FindById(id int64) Permission
	Create(data *CreatePermission) error
	Update(data *UpdatePermission) error
	Delete(id int64) error
}

type repository struct {
	Ds DataSource
}

func (r *repository) Count(params Params) int64 {
	return r.Ds.Count(params)
}

func (r *repository) FindList(params Params) []Permission {
	return r.Ds.FindList(params)
}

func (r *repository) FindById(id int64) Permission {
	return r.Ds.FindById(id)
}

func (r *repository) Create(data *CreatePermission) error {
	return r.Ds.Create(data)
}

func (r *repository) Update(data *UpdatePermission) error {
	return r.Ds.Update(data)
}

func (r *repository) Delete(id int64) error {
	return r.Ds.Delete(id)
}

func NewRepository(ds DataSource) Repository {
	return &repository{
		Ds: ds,
	}
}
