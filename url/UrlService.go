package url

type UrlService struct {
	urlRepo UrlRepo
}

func UrlServiceProvider(repo UrlRepo) UrlService {
	return UrlService{urlRepo: repo}
}

func (u *UrlService) Create(obj UrlResource) UrlResource {
	return u.urlRepo.Save(obj)
}

func (u *UrlService) Delete(obj UrlResource) {
	u.urlRepo.Delete(obj)
}

func (u *UrlService) FindById(id uint) UrlResource {
	return u.urlRepo.FindById(id)
}

func (u *UrlService) FindAll() []UrlResource {
	return u.urlRepo.FindAll()
}

func (u *UrlService) DeleteAll() {
	u.urlRepo.DeleteAll()
}

func (u *UrlService) Activate(url UrlResource, b bool) {
	u.urlRepo.Activate(url, b)
}

func (u *UrlService) FindAllActive() []UrlResource {
	return u.urlRepo.FindAllActive()
}
