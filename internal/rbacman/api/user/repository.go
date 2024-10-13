package user

type Repository interface {
	Count(params Params) int64
	FindList(params Params) []User
	FindById(id string) User
	Create(data *CreateUser) error
	Update(data *UpdateUser) error
	Delete(id string) error
}

type repository struct {
	Ds DataSource
}

func (r *repository) Count(params Params) int64 {
	return r.Ds.Count(params)
}

func (r *repository) FindList(params Params) []User {
	return r.Ds.FindList(params)
}

func (r *repository) FindById(id string) User {
	return r.Ds.FindById(id)
}

func (r *repository) Create(data *CreateUser) error {
	return r.Ds.Create(data)
}

func (r *repository) Update(data *UpdateUser) error {
	return r.Ds.Update(data)
}

func (r *repository) Delete(id string) error {
	return r.Ds.Delete(id)
}

func NewRepository(ds DataSource) Repository {
	return &repository{
		Ds: ds,
	}
}
