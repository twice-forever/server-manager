package v1

import (
	"net/http"
	"server-manager/models"
	"server-manager/pkg/utils"
	"strconv"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

// 创建用户
func CreateUser(c *gin.Context) {
	// 解析请求数据
	var requestData models.User
	if err := c.ShouldBind(&requestData); err != nil {
		utils.HandleErrorResponse(c, http.StatusInternalServerError, nil, "请求失败，请重试！")
		return
	}

	// 表单验证
	valid := validation.Validation{}
	valid.Required(requestData.Username, "username").Message("username")
	valid.Required(requestData.Password, "password").Message("password")

	// 增加数据到数据库
	requestData.PWDSalt = utils.GenerateSalt()
	requestData.Password = utils.GeneratHashPWD(requestData.Password, requestData.PWDSalt)
	result := utils.ConnectDB.Create(&requestData)
	if result.Error != nil {
		utils.HandleErrorResponse(c, http.StatusInternalServerError, nil, "增加用户失败，请重试！")
		return
	}

	// 返回成功信息
	utils.HandleSuccessResponse(c, http.StatusOK, gin.H{"id": requestData.ID}, "")
}

// 删除用户
func DeleteUser(c *gin.Context) {
	userIdStr := c.Param("userId")
	if userIdStr == "" {
		utils.HandleErrorResponse(c, http.StatusInternalServerError, nil, "无用户id")
		return
	}
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		utils.HandleErrorResponse(c, http.StatusInternalServerError, nil, "id错误")
		return
	}

	deleteUser := models.User{
		ID: userId,
	}
	if err := utils.ConnectDB.Delete(&deleteUser).Error; err != nil {
		utils.HandleErrorResponse(c, http.StatusInternalServerError, nil, "删除用户错误")
		return
	}

	// 返回成功信息
	utils.HandleSuccessResponse(c, http.StatusOK, nil, "删除用户成功")
}

// 查询用户列表
func GetUsers(c *gin.Context) {
	users := make([]models.ShowUser, 0, 10)

	var count int64
	tempDB := utils.ConnectDB.Scopes(utils.Paginate(c)).Model(models.User{})
	if err := tempDB.Find(&users).Error; err != nil {
		utils.HandleErrorResponse(c, http.StatusInternalServerError, nil, "查询用户错误")
		return
	}
	if err := tempDB.Count(&count).Error; err != nil {
		utils.HandleErrorResponse(c, http.StatusInternalServerError, nil, "查询用户错误")
		return
	}

	// 返回成功信息
	utils.HandleSuccessResponse(c, http.StatusOK, gin.H{
		"list":  &users,
		"count": count,
	}, "")
}

// 查询用户
func GetUser(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("userId"))

	user := models.ShowUser{}
	if err := utils.ConnectDB.Model(models.User{}).First(&user, userId).Error; err != nil {
		utils.HandleErrorResponse(c, http.StatusInternalServerError, nil, "查询用户错误")
		return
	}
	// 返回成功信息
	utils.HandleSuccessResponse(c, http.StatusOK, &user, "")
}

// 更新用户
func UpdateUsers(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		utils.HandleErrorResponse(c, http.StatusInternalServerError, nil, "无用户id")
		return
	}

	requestBodyData := &models.User{}
	if err := c.ShouldBind(&requestBodyData); err != nil {
		utils.HandleErrorResponse(c, http.StatusInternalServerError, nil, "请求解析失败")
		return
	}
	requestBodyData.ID = userId
	if err := utils.ConnectDB.Model(&requestBodyData).Select("username", "real_name", "avatar_url").Updates(&requestBodyData).Error; err != nil {
		utils.HandleErrorResponse(c, http.StatusInternalServerError, nil, "无法找到该用户")
		return
	}

	// 返回成功信息
	utils.HandleSuccessResponse(c, http.StatusOK, nil, "更新成功")
}
