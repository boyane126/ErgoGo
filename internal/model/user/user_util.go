package user

import (
	"ErgoGo/internal/pkg/app"
	"ErgoGo/pkg/database"
	"ErgoGo/pkg/http/paginator"

	"github.com/gin-gonic/gin"
)

// Paginate 分页内容
func Paginate(c *gin.Context, perPage int) (users []User, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.DB.Model(User{}),
		&users,
		app.V1URL(database.TableName(&User{})),
		perPage,
	)
	return
}
