package relational

import (
	"fmt"

	"github.com/offen/offen/server/persistence"
)

func (r *relationalDAL) CreateAccountUser(u *persistence.AccountUser) error {
	local := importAccountUser(u)
	if err := r.db.Create(&local).Error; err != nil {
		return fmt.Errorf("relational: error creating account user: %w", err)
	}
	return nil
}

func (r *relationalDAL) FindAccountUser(q interface{}) (persistence.AccountUser, error) {
	var accountUser AccountUser
	switch query := q.(type) {
	case persistence.FindAccountUserQueryByHashedEmail:
		if err := r.db.Where("hashed_email = ?", string(query)).First(&accountUser).Error; err != nil {
			return accountUser.export(), fmt.Errorf("relational: error looking up account user by hashed email: %w", err)
		}
		return accountUser.export(), nil
	case persistence.FindAccountUserQueryByHashedEmailIncludeRelationships:
		if err := r.db.Preload("Relationships").Where("hashed_email = ?", string(query)).First(&accountUser).Error; err != nil {
			return accountUser.export(), fmt.Errorf("relational: error looking up account user by hashed email: %w", err)
		}
		return accountUser.export(), nil
	case persistence.FindAccountUserQueryByUserIDIncludeRelationships:
		if err := r.db.Preload("Relationships").Where("user_id = ?", string(query)).First(&accountUser).Error; err != nil {
			return accountUser.export(), fmt.Errorf("relational: error looking up account user by user id: %w", err)
		}
		return accountUser.export(), nil
	default:
		return accountUser.export(), persistence.ErrBadQuery
	}
}

func (r *relationalDAL) UpdateAccountUser(u *persistence.AccountUser) error {
	local := importAccountUser(u)
	if err := r.db.Save(&local).Error; err != nil {
		return fmt.Errorf("relational: error updating account user: %w", err)
	}
	return nil
}
