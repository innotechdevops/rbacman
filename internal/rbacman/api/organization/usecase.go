package organization

type UseCase interface {
	Count(params Params) int64
	FindList(params Params) []Organization
	FindById(id int64) Organization
	Create(data *CreateOrganization) (Organization, error)
	Update(data *UpdateOrganization) (Organization, error)
	Delete(id int64) error
}

type useCase struct {
	Repo Repository
}

func (u *useCase) Count(params Params) int64 {
	return u.Repo.Count(params)
}

func (u *useCase) FindList(params Params) []Organization {
	return u.Repo.FindList(params)
}

func (u *useCase) FindById(id int64) Organization {
	return u.Repo.FindById(id)
}

func (u *useCase) Create(data *CreateOrganization) (Organization, error) {
	err := u.Repo.Create(data)
	if err != nil {
		return Organization{}, err
	}
	return u.Repo.FindById(data.Id), nil
}

func (u *useCase) Update(data *UpdateOrganization) (Organization, error) {
	err := u.Repo.Update(data)
	if err != nil {
		return Organization{}, err
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
