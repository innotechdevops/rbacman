package group

type UseCase interface {
	Count(params Params) int64
	FindList(params Params) []Group
	FindById(id string) Group
	Create(data *CreateGroup) (Group, error)
	Update(data *UpdateGroup) (Group, error)
	Delete(id string) error
}

type useCase struct {
	Repo Repository
}

func (u *useCase) Count(params Params) int64 {
	return u.Repo.Count(params)
}

func (u *useCase) FindList(params Params) []Group {
	return u.Repo.FindList(params)
}

func (u *useCase) FindById(id string) Group {
	return u.Repo.FindById(id)
}

func (u *useCase) Create(data *CreateGroup) (Group, error) {
	err := u.Repo.Create(data)
	if err != nil {
		return Group{}, err
	}
	return u.Repo.FindById(data.Id), nil
}

func (u *useCase) Update(data *UpdateGroup) (Group, error) {
	err := u.Repo.Update(data)
	if err != nil {
		return Group{}, err
	}
	return u.Repo.FindById(data.Id), nil
}

func (u *useCase) Delete(id string) error {
	return u.Repo.Delete(id)
}

func NewUseCase(repo Repository) UseCase {
	return &useCase{
		Repo: repo,
	}
}
