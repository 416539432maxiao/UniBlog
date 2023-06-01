package user

import (
	"UniBlog/internal/pkg/core"
	"github.com/gin-gonic/gin"
)

/*

通过在 Controller 层实现有限的功能（参数解析、校验、逻辑分发、请求聚合和返回），
并将负责的业务逻辑放在 Biz 层实现，
可以使 Controller 层代码逻辑结构清晰，利于后期的代码维护。




*/

// Create 创建一个新的用户.
func (ctrl *UserController) Create(c *gin.Context) {
	log.C(c).Infow("Create user function called")

	var r v1.CreateUserRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, errno.ErrBind, nil)

		return
	}

	if _, err := govalidator.ValidateStruct(r); err != nil {
		core.WriteResponse(c, errno.ErrInvalidParameter.SetMessage(err.Error()), nil)

		return
	}

	if err := ctrl.b.Users().Create(c, &r); err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	core.WriteResponse(c, nil, nil)
}
