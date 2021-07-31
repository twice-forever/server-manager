package v1

import (
	"net/http"
	"server-manager/models"
	"server-manager/pkg/utils"
	"server-manager/service"
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

	// 判断用户名重复
	if err := requestData.GetUser(); err == nil {
		utils.HandleErrorResponse(c, http.StatusInternalServerError, nil, "用户名重复，请重试！")
		return
	}

	// 增加数据到数据库
	requestData.PWDSalt = utils.GenerateSalt()
	requestData.Password = utils.GeneratHashPWD(requestData.Password, requestData.PWDSalt)
	err := requestData.AddUser()
	if err != nil {
		utils.HandleErrorResponse(c, http.StatusInternalServerError, nil, "增加用户失败，请重试！")
		return
	}

	// 返回成功信息
	utils.HandleSuccessResponse(c, http.StatusOK, gin.H{"id": requestData.ID}, "")
}

// 删除用户
func DeleteUser(c *gin.Context) {
	// 解析id
	userId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		utils.HandleErrorResponse(c, http.StatusInternalServerError, nil, "无用户id")
		return
	}

	// 删除用户
	deleteUser := models.User{
		ID: userId,
	}
	if err := deleteUser.DeleteUser(); err != nil {
		utils.HandleErrorResponse(c, http.StatusInternalServerError, nil, "删除用户错误")
		return
	}

	// 返回成功信息
	utils.HandleSuccessResponse(c, http.StatusOK, nil, "删除用户成功")
}

// 查询用户列表
func GetUsers(c *gin.Context) {
	var count int64

	// 获取列表
	users, err := service.GetUsers(c)
	if err != nil {
		utils.HandleErrorResponse(c, http.StatusInternalServerError, nil, "查询用户错误")
		return
	}

	// 获取总数
	if err := models.DB.Model(&models.User{}).Count(&count).Error; err != nil {
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
	// 解析id
	userId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		utils.HandleErrorResponse(c, http.StatusInternalServerError, nil, "无用户id")
		return
	}

	user := models.User{
		ID: userId,
	}
	if err := user.GetUser(); err != nil {
		utils.HandleErrorResponse(c, http.StatusInternalServerError, nil, "查询用户错误")
		return
	}

	// 返回成功信息
	utils.HandleSuccessResponse(c, http.StatusOK, &user, "")
}

// 更新用户
func UpdateUsers(c *gin.Context) {
	// 解析id
	userId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		utils.HandleErrorResponse(c, http.StatusInternalServerError, nil, "无用户id")
		return
	}

	// 解析请求体
	requestBodyData := &models.User{}
	if err := c.ShouldBind(&requestBodyData); err != nil {
		utils.HandleErrorResponse(c, http.StatusInternalServerError, nil, "请求解析失败")
		return
	}
	requestBodyData.ID = userId

	// 更新数据库
	if err := requestBodyData.UpdateUser(); err != nil {
		utils.HandleErrorResponse(c, http.StatusInternalServerError, nil, "无法找到该用户")
		return
	}

	// 返回成功信息
	utils.HandleSuccessResponse(c, http.StatusOK, nil, "更新成功")
}

// 修改密码
func ChangePassword(c *gin.Context) {
	// 解析id
	userId, err := strconv.Atoi(c.Param("userId"))
	if err != nil {
		utils.HandleErrorResponse(c, http.StatusInternalServerError, nil, "无用户id")
		return
	}

	// 解析请求体
	requestBodyData := &service.ChangePassword{}
	if err := c.ShouldBind(&requestBodyData); err != nil {
		utils.HandleErrorResponse(c, http.StatusInternalServerError, nil, "请求解析失败")
		return
	}

	// 查找用户
	databaseData := &models.User{
		ID: userId,
	}
	if err := databaseData.GetUser(); err != nil {
		utils.HandleErrorResponse(c, http.StatusInternalServerError, nil, "查询用户错误")
		return
	}

	// 密码对比
	hashPassword := utils.GeneratHashPWD(requestBodyData.Password, databaseData.PWDSalt)
	if hashPassword != databaseData.Password {
		utils.HandleErrorResponse(c, http.StatusInternalServerError, nil, "密码不正确")
		return
	}
	newHashPassword := utils.GeneratHashPWD(requestBodyData.NewPassword, databaseData.PWDSalt)
	databaseData.Password = newHashPassword
	if err := databaseData.UpdateUserPassword(); err != nil {
		utils.HandleErrorResponse(c, http.StatusInternalServerError, nil, "无法找到该用户")
		return
	}

	// 返回成功信息
	utils.HandleSuccessResponse(c, http.StatusOK, nil, "更新成功")
}
