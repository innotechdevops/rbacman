package user_group

type Repository interface {
	Count(params Params) int64
	FindList(params Params) []UserGroup
	FindById(id int64) UserGroup
	Create(data *CreateUserGroup) error
	Update(data *UpdateUserGroup) error
	Delete(id int64) error
}

type repository struct {
	Ds DataSource
}

func (r *repository) Count(params Params) int64 {
	return r.Ds.Count(params)
}

func (r *repository) FindList(params Params) []UserGroup {
	return r.Ds.FindList(params)
}

func (r *repository) FindById(id int64) UserGroup {
	return r.Ds.FindById(id)
}

func (r *repository) Create(data *CreateUserGroup) error {
	return r.Ds.Create(data)
}

func (r *repository) Update(data *UpdateUserGroup) error {
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
