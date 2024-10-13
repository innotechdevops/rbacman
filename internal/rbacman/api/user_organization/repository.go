package user_organization

type Repository interface {
	Count(params Params) int64
	FindList(params Params) []UserOrganization
	FindById(id int64) UserOrganization
	Create(data *CreateUserOrganization) error
	Update(data *UpdateUserOrganization) error
	Delete(id int64) error
}

type repository struct {
	Ds DataSource
}

func (r *repository) Count(params Params) int64 {
	return r.Ds.Count(params)
}

func (r *repository) FindList(params Params) []UserOrganization {
	return r.Ds.FindList(params)
}

func (r *repository) FindById(id int64) UserOrganization {
	return r.Ds.FindById(id)
}

func (r *repository) Create(data *CreateUserOrganization) error {
	return r.Ds.Create(data)
}

func (r *repository) Update(data *UpdateUserOrganization) error {
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
