package mock

import (
	"go-fiber-crud/app/model"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var users = []*model.Register{
	{
		FirstName: "admin",
		LastName:  "admin",
		Username:  "admin@admin.com",
		Password:  "password",
	},
	{
		FirstName: "test",
		LastName:  "testaccount",
		Username:  "test@test.com",
		Password:  "testpassword",
	},
}

func SeedUser(db *gorm.DB) (string, error) {
	for _, user := range users {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return "", err
		}
		entity := model.User{
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Username:  user.Username,
			Password:  string(hashedPassword),
		}
		tx := db.Create(&entity)
		if tx.Error != nil {
			return "", tx.Error
		}
	}
	return "Seeded User Success", nil
}
