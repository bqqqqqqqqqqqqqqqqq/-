package test

import (
	"dogking_shop/models"
	"fmt"
	"testing"
)

var db = models.DB

func TestHomework(t *testing.T) {

	users := make([]*models.User, 0)
	err := db.Find(&users).Error
	if err != nil {
		fmt.Print(err)
	}
	for _, v := range users {
		fmt.Printf("user ===> %v\n", v)
	}

}

func TestGetAllUser(t *testing.T) {
	users := make([]*models.UserAPI, 0)
	tx := db.Model(&models.User{}).Limit(10).Find(&users)
	tx.Order("User.id DESC")
	for _, v := range users {
		fmt.Printf("user ===> %v\n", v)
	}
}
