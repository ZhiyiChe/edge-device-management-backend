package routers

import (
	"edge-device-management-backend/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/v1/testapi", &controllers.TestController{})

	beego.Router("/v1/joinCluster", &controllers.JoinClusterController{})
	beego.Router("/v1/downloadK3s", &controllers.DownloadK3sController{})
	beego.Router("/v1/downloadK3sImages", &controllers.DownloadK3sImagesController{})
	beego.Router("/v1/downloadInstallScript", &controllers.DownloadInstallScriptController{})

	beego.Router("/v1/approveNode", &controllers.ApproveNodeController{})
	beego.Router("/v1/queryNode", &controllers.QueryNodeController{})
	beego.Router("/v1/deleteNode", &controllers.DeleteNodeController{})
	beego.Router("/v1/updateNode", &controllers.UpdateNodeController{})

	beego.Router("/v1/queryToken", &controllers.QueryTokenController{})

	beego.Router("/v1/queryNamespace", &controllers.QueryNamespaceController{})
	beego.Router("/v1/addNamespace", &controllers.AddNamespaceController{})
	beego.Router("/v1/deleteNamespace", &controllers.DeleteNamespaceController{})

	beego.Router("/v1/queryPod", &controllers.QueryPodController{})
	beego.Router("/v1/addPod", &controllers.AddPodController{})
	beego.Router("/v1/updatePod", &controllers.UpdatePodController{})
	beego.Router("/v1/deletePod", &controllers.DeletePodController{})

	beego.Router("/v1/queryDeployment", &controllers.QueryDeploymentController{})

	beego.Router("/v1/queryService", &controllers.QueryServiceController{})

	beego.Router("/v1/execKubectl", &controllers.ExecKubectlController{})

	beego.Router("/v1/login", &controllers.LoginController{})
	beego.Router("/v1/queryUser", &controllers.QueryUserController{})
	beego.Router("/v1/addUser", &controllers.AddUserController{})
	beego.Router("/v1/updateUser", &controllers.UpdateUserController{})
	beego.Router("/v1/deleteUser", &controllers.DeleteUserController{})
}
