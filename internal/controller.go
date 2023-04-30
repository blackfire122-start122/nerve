package internal

import (
	. "Nerve/pkg"
	. "Nerve/pkg/logging"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
)

func Home(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func Register(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", nil)
}

func LogIn(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func Profile(c *gin.Context) {
	c.HTML(http.StatusOK, "profile.html", nil)
}

func RegisterUser(c *gin.Context) {
	resp := make(map[string]string)

	var user UserRegister
	bodyBytes, _ := io.ReadAll(c.Request.Body)

	if err := json.Unmarshal(bodyBytes, &user); err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	if user.Password == "" || user.Username == "" {
		resp["Register"] = "Not all field"

		c.JSON(http.StatusBadRequest, resp)
		return
	}

	if err := Sign(&user); err != nil {
		resp["Register"] = "Error create user"

		c.JSON(http.StatusBadRequest, resp)
		return
	}

	resp["Register"] = "OK"
	c.JSON(http.StatusOK, resp)
}

func LoginUser(c *gin.Context) {
	resp := make(map[string]string)

	var user UserLogin
	bodyBytes, _ := io.ReadAll(c.Request.Body)

	if err := json.Unmarshal(bodyBytes, &user); err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	if Login(c.Writer, c.Request, &user) {
		resp["Login"] = "OK"
		c.JSON(http.StatusOK, resp)
	} else {
		resp["Login"] = "error login user"
		c.JSON(http.StatusForbidden, resp)
	}
}

func GetUser(c *gin.Context) {
	loginUser, user := CheckSessionUser(c.Request)

	if !loginUser {
		c.Writer.WriteHeader(http.StatusUnauthorized)
		return
	}

	var admin string

	if CheckAdmin(user) {
		admin = "true"
	} else {
		admin = "false"
	}

	resp := make(map[string]string)

	resp["Id"] = strconv.FormatUint(user.Id, 10)
	resp["Username"] = user.Username
	resp["Email"] = user.Email
	resp["Phone"] = user.Phone
	resp["Image"] = user.Image
	resp["IsAdmin"] = admin

	c.JSON(http.StatusOK, resp)
}

type FormChangeUser struct {
	Username string                `form:"Username"`
	Email    string                `form:"Email"`
	Phone    string                `form:"Phone"`
	Image    *multipart.FileHeader `form:"Image"`
}

func ChangeUser(c *gin.Context) {
	loginUser, user := CheckSessionUser(c.Request)

	if !loginUser {
		c.Writer.WriteHeader(http.StatusUnauthorized)
		return
	}

	resp := make(map[string]string)
	var form FormChangeUser
	if err := c.ShouldBind(&form); err != nil {
		c.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	var ImageName string

	if form.Image.Filename == "" && form.Image.Size == 0 {
		ImageName = user.Image
	} else {
		if err := c.SaveUploadedFile(form.Image, "./media/UserImages/"+user.Username+form.Image.Filename); err != nil {
			ErrorLogger.Println(err.Error())
		}
		if err := os.Remove("./" + user.Image); err != nil {
			ErrorLogger.Println(err.Error())
		}
		ImageName = "media/UserImages/" + user.Username + form.Image.Filename
	}

	if err := DB.Save(&User{Id: user.Id, Username: form.Username, Image: ImageName, Email: form.Email, Phone: form.Phone, Password: user.Password}).Error; err != nil {
		ErrorLogger.Println(err.Error())
	}

	if err := DB.First(&user, "id = ?", user.Id).Error; err != nil {
		ErrorLogger.Println(err.Error())
	}

	var admin string

	if CheckAdmin(user) {
		admin = "true"
	} else {
		admin = "false"
	}

	resp["Id"] = strconv.FormatUint(user.Id, 10)
	resp["Username"] = user.Username
	resp["Email"] = user.Email
	resp["Phone"] = user.Phone
	resp["Image"] = user.Image
	resp["IsAdmin"] = admin

	c.JSON(http.StatusOK, resp)
}
func UserTasks(c *gin.Context) {
	loginUser, user := CheckSessionUser(c.Request)

	if !loginUser {
		c.Writer.WriteHeader(http.StatusUnauthorized)
		return
	}
	var tasks []Task
	if err := DB.Find(&tasks, "user_id = ?", user.Id).Error; err != nil {
		ErrorLogger.Println(err.Error())
	}

	resp := make([]map[string]string, len(tasks))

	for i, typeItem := range tasks {
		item := make(map[string]string)

		item["Id"] = strconv.FormatUint(typeItem.Id, 10)
		item["Name"] = typeItem.Name
		item["Description"] = typeItem.Description

		resp[i] = item
	}

	c.JSON(http.StatusOK, resp)
}
