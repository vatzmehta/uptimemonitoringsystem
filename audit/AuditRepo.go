package audit

import "github.com/jinzhu/gorm"

type AuditRepo struct {
	db *gorm.DB
}

func RepoProvider(db *gorm.DB) AuditRepo {
	return AuditRepo{db: db}
}

func (u *AuditRepo) Save(audit Audit) Audit {
	u.db.Save(&audit)
	return audit
}
