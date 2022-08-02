package main

import (
	"fmt"
	"go-5/member"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	dsn := "root:@tcp(127.0.0.1)/bwa?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
	} else {
		newRepository := member.NewRepository(db)
		newServices := member.NewServices(newRepository)

		// newCheckEmailInput := member.CheckEmailInput{}
		// newCheckEmailInput.Nama = "rahmah@gmail.com"
		// newIsEmailAvailable, err := newServices.IsEmailAvailable(newCheckEmailInput)
		// if err != nil {
		// 	fmt.Println(err.Error())
		// } else {
		// 	fmt.Println(newIsEmailAvailable)
		// }

		// member := member.InputMember{}
		// member.Nama = "rahmah"
		// member.Password = "12345"
		// newServices.SaveServices(member)

		newHandler := member.NewHandler(newServices)

		// newLoginMember := member.LoginMember{}
		// newLoginMember.Nama = "rahmah@gmail.com"
		// newLoginMember.Password = "12345s"
		// newFindByEmailService, err := newServices.FindByEmailService(newLoginMember)
		// if err != nil {
		// 	fmt.Println(err)
		// } else {
		// 	fmt.Println(newFindByEmailService.Nama)
		// }

		// newCheckEmailInput := member.CheckEmailInput{}
		// newCheckEmailInput.Nama = "srahmah@gmail.com"

		// newIsEmailAvailable, err := newServices.IsEmailAvailable(newCheckEmailInput)
		// if err != nil {
		// 	fmt.Println(err.Error())
		// } else {
		// 	fmt.Println(newIsEmailAvailable)
		// }

		routes := gin.Default()
		api := routes.Group("v1")
		api.POST("member", newHandler.SaveHandler)
		api.POST("login", newHandler.LoginHandler)
		api.POST("email_checkers", newHandler.CheckEmailAvailability)
		routes.Run()
	}

}

func handler(h *gin.Context) {
	dsn := "root:@tcp(127.0.0.1)/bwa?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
	} else {
		var keymember []member.Member
		db.Find(&keymember)

		h.JSON(http.StatusOK, keymember)
	}
}
