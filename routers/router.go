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

	beego.Router("/v1/queryNode", &controllers.QueryNodeController{})
	beego.Router("/v1/approveNode", &controllers.ApproveNodeController{})
	beego.Router("/v1/deleteNode", &controllers.DeleteNodeController{})

	beego.Router("/v1/queryToken", &controllers.QueryTokenController{})

	beego.Router("/v1/queryNamespace", &controllers.QueryNamespaceController{})
	beego.Router("/v1/addNamespace", &controllers.AddNamespaceController{})
	beego.Router("/v1/deleteNamespace", &controllers.DeleteNamespaceController{})

	beego.Router("/v1/queryPod", &controllers.QueryPodController{})
	beego.Router("/v1/addPod", &controllers.AddPodController{})
	beego.Router("/v1/deletePod", &controllers.DeletePodController{})
}
