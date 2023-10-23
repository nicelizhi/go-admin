package actions

import (
	"net/http"

	ginI18n "github.com/gin-contrib/i18n"
	"github.com/gin-gonic/gin"
	log "github.com/nicelizhi/go-admin-core/logger"
	"github.com/nicelizhi/go-admin-core/sdk/pkg"
	"github.com/nicelizhi/go-admin-core/sdk/pkg/jwtauth/user"
	"github.com/nicelizhi/go-admin-core/sdk/pkg/response"

	"go-admin/common/dto"
	"go-admin/common/models"
)

// UpdateAction 通用更新动作
func UpdateAction(control dto.Control) gin.HandlerFunc {
	return func(c *gin.Context) {
		db, err := pkg.GetOrm(c)
		if err != nil {
			log.Error(err)
			return
		}

		msgID := pkg.GenerateMsgIDFromContext(c)
		req := control.Generate()
		//更新操作
		err = req.Bind(c)
		if err != nil {
			response.Error(c, http.StatusUnprocessableEntity, err, ginI18n.MustGetMessage(c, "Parameter validation failed"))
			return
		}
		var object models.ActiveRecord
		object, err = req.GenerateM()
		if err != nil {
			response.Error(c, 500, err, ginI18n.MustGetMessage(c, "Model generation failed"))
			return
		}
		object.SetUpdateBy(user.GetUserId(c))

		//数据权限检查
		p := GetPermissionFromContext(c)

		db = db.WithContext(c).Scopes(
			Permission(object.TableName(), p),
		).Where(req.GetId()).Updates(object)
		if err = db.Error; err != nil {
			log.Errorf("MsgID[%s] Update error: %s", msgID, err)
			response.Error(c, 500, err, ginI18n.MustGetMessage(c, "Update failed"))
			return
		}
		if db.RowsAffected == 0 {
			response.Error(c, http.StatusForbidden, nil, ginI18n.MustGetMessage(c, "Don not have permission to update this data"))
			return
		}
		response.OK(c, object.GetId(), ginI18n.MustGetMessage(c, "Update completed"))
		c.Next()
	}
}
