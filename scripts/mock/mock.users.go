package mock

import (
	"go-fiber-crud/app/model"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var users = []*model.Register{
	{
		FirstName: "test",
		LastName:  "ltest",
		Username:  "test01@test.com",
		Password:  "test1234",
	},
	{
		FirstName: "test2",
		LastName:  "ltest2",
		Username:  "test02@test.com",
		Password:  "test1234",
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
