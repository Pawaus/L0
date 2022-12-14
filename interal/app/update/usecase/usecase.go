package usecase

type Usecase struct {
	repo Repo
}

type Repo interface {
	Update(uid string, data string)
	GetAll()
}

func NewUseCase(r Repo) *Usecase {
	return &Usecase{repo: r}
}

func (u *Usecase) Update(uid string, data string) {
	u.repo.Update(uid, data)
}

func (u *Usecase) GetAll() {
	u.repo.GetAll()
}
