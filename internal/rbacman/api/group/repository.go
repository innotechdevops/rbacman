package group

type Repository interface {
	Count(params Params) int64
	FindList(params Params) []Group
	FindById(id string) Group
	Create(data *CreateGroup) error
	Update(data *UpdateGroup) error
	Delete(id string) error
}

type repository struct {
	Ds DataSource
}

func (r *repository) Count(params Params) int64 {
	return r.Ds.Count(params)
}

func (r *repository) FindList(params Params) []Group {
	return r.Ds.FindList(params)
}

func (r *repository) FindById(id string) Group {
	return r.Ds.FindById(id)
}

func (r *repository) Create(data *CreateGroup) error {
	return r.Ds.Create(data)
}

func (r *repository) Update(data *UpdateGroup) error {
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
