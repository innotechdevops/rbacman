package resource

type UseCase interface {
	Count(params Params) int64
	FindList(params Params) []Resource
	FindById(id int64) Resource
	Create(data *CreateResource) (Resource, error)
	Update(data *UpdateResource) (Resource, error)
	Delete(id int64) error
}

type useCase struct {
	Repo Repository
}

func (u *useCase) Count(params Params) int64 {
	return u.Repo.Count(params)
}

func (u *useCase) FindList(params Params) []Resource {
	return u.Repo.FindList(params)
}

func (u *useCase) FindById(id int64) Resource {
	return u.Repo.FindById(id)
}

func (u *useCase) Create(data *CreateResource) (Resource, error) {
	err := u.Repo.Create(data)
	if err != nil {
		return Resource{}, err
	}
	return u.Repo.FindById(data.Id), nil
}

func (u *useCase) Update(data *UpdateResource) (Resource, error) {
	err := u.Repo.Update(data)
	if err != nil {
		return Resource{}, err
	}
	return u.Repo.FindById(data.Id), nil
}

func (u *useCase) Delete(id int64) error {
	return u.Repo.Delete(id)
}

func NewUseCase(repo Repository) UseCase {
	return &useCase{
		Repo: repo,
	}
}
