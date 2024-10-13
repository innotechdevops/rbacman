package user_organization

type UseCase interface {
	Count(params Params) int64
	FindList(params Params) []UserOrganization
	FindById(id int64) UserOrganization
	Create(data *CreateUserOrganization) (UserOrganization, error)
	Update(data *UpdateUserOrganization) (UserOrganization, error)
	Delete(id int64) error
}

type useCase struct {
	Repo Repository
}

func (u *useCase) Count(params Params) int64 {
	return u.Repo.Count(params)
}

func (u *useCase) FindList(params Params) []UserOrganization {
	return u.Repo.FindList(params)
}

func (u *useCase) FindById(id int64) UserOrganization {
	return u.Repo.FindById(id)
}

func (u *useCase) Create(data *CreateUserOrganization) (UserOrganization, error) {
	err := u.Repo.Create(data)
	if err != nil {
		return UserOrganization{}, err
	}
	return u.Repo.FindById(data.Id), nil
}

func (u *useCase) Update(data *UpdateUserOrganization) (UserOrganization, error) {
	err := u.Repo.Update(data)
	if err != nil {
		return UserOrganization{}, err
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
