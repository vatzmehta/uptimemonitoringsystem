package url

import "github.com/jinzhu/gorm"

type UrlRepo struct {
	db *gorm.DB
}

func RepoProvider(db *gorm.DB) UrlRepo {
	return UrlRepo{db: db}
}

func (u *UrlRepo) Save(url UrlResource) UrlResource {
	u.db.Save(&url)
	return url
}

func (u *UrlRepo) FindById(id uint) UrlResource {
	var urlObj UrlResource
	u.db.First(&urlObj, id)
	return urlObj
}

func (u *UrlRepo) FindAll() []UrlResource {
	var urlObjs []UrlResource
	u.db.Find(&urlObjs)
	return urlObjs
}

func (u *UrlRepo) Delete(urlObj UrlResource) {
	u.db.Delete(&urlObj)
}

func (u *UrlRepo) DeleteAll() {
	u.db.Delete(&UrlResource{})
}

func (u *UrlRepo) Activate(url UrlResource, b bool) {
	url.Enable = b
	u.db.Save(&url)
}

func (u *UrlRepo) FindAllActive() []UrlResource {
	var urlObjs []UrlResource
	u.db.Find(&urlObjs, UrlResource{Enable: true})
	return urlObjs
}
