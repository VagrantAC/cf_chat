package route

import (
	"cf_chat/framework"
	"cf_chat/controller"
)

func RegisterRouter(core *framework.Core) {
	// 静态路由+HTTP方法匹配
	core.Get("/user/login", controller.UserLoginController)

	// 批量通用前缀
	subjectApi := core.Group("/subject")
	{
		subjectApi.Delete("/:id", controller.SubjectDelController)
		subjectApi.Put("/:id", controller.SubjectUpdateController)
		subjectApi.Get("/:id", controller.SubjectGetController)
		subjectApi.Get("/list/all", controller.SubjectListController)

		subjectInnerApi := subjectApi.Group("/info")
    {
    	subjectInnerApi.Get("/name", controller.SubjectNameController)
    }
	}
}
