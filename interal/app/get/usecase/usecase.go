package usecase

type Usecase struct {
	Repo Repo
}

type Repo interface {
	Get(uid string) string
}

func (u Usecase) Get(uid string) string {
	return u.Repo.Get(uid)
}
