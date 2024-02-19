package repository

import (
	"go-fiber-crud/app/model"
	"go-fiber-crud/app/utils/errs"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type authRepository struct {
	db *gorm.DB
}
type AuthRepository interface {
	Login(*model.Login) (*model.LoginWithUser, error)
	Register(*model.Register) (*model.UserResponse, error)
}

func NewAuthRepositoryDB(db *gorm.DB) authRepository {
	return authRepository{db: db}
}
func (r authRepository) Login(login *model.Login) (*model.LoginWithUser, error) {
	entity := model.User{}
	tx := r.db.Where("username = ?", login.Username).Preload(clause.Associations).First(&entity)
	if tx.Error != nil {
		return nil, tx.Error
	}

	err := bcrypt.CompareHashAndPassword([]byte(entity.Password), []byte(login.Password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return nil, errs.New("password mismatch")
	}
	token, err := generateToken(entity.ID)
	response := model.LoginWithUser{
		Token: token,
		User:  entity,
	}
	return &response, nil
}
func (r authRepository) Register(data *model.Register) (*model.UserResponse, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	entity := model.User{
		FirstName: data.FirstName,
		LastName:  data.LastName,
		Username:  data.Username,
		Password:  string(hashedPassword),
	}
	tx := r.db.Create(&entity)
	if tx.Error != nil {
		return nil, tx.Error
	}

	response := model.UserResponse{
		Username:  entity.Username,
		FirstName: data.FirstName,
		LastName:  data.LastName,
	}
	return &response, nil
}

func generateToken(id uint) (string, error) {
	tokenLifeTimeHour, err := strconv.Atoi(viper.GetString("app.tokenLiftHour"))
	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{
		"exp":    time.Now().Add(time.Hour * time.Duration(tokenLifeTimeHour)).Unix(),
		"iat":    time.Now().Unix(),
		"userId": id,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(viper.GetString("app.jwtSecret")))

}
