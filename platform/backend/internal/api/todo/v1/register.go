package api_todo_v1

import (
	"api/go-sdk/todo/v1/todoV1connect"
	"net/http"
	todogendb "platform/backend/internal/domain/todo/gendb"

	"github.com/gin-gonic/gin"
)

type Dependencies struct {
	Store todogendb.Querier
}

func Register(apis *gin.RouterGroup, deps Dependencies) {
	service := newTodoService(deps.Store)
	servicePath, serviceHandler := todoV1connect.NewTodoServiceHandler(service)

	apis.POST(servicePath+"*rpc", gin.WrapH(http.StripPrefix("/api", serviceHandler)))
}
