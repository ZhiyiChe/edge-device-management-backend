package controllers

import (
	"log"
	"time"

	beego "github.com/beego/beego/v2/server/web"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	"github.com/dgrijalva/jwt-go"

	"edge-device-management-backend/models"
)

const (
	KEY                    string = "JWT-ARY-STARK"
	DEFAULT_EXPIRE_SECONDS int    = 60 * 30 // 默认过期时间（s）
)

// JWT -- json web token
// HEADER PAYLOAD SIGNATURE
// This struct is the PAYLOAD
type MyCustomClaims struct {
	models.User
	jwt.StandardClaims
}

var ClientSet *kubernetes.Clientset

func init() {
	kubeConfig := "/etc/rancher/k3s/k3s.yaml"
	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", kubeConfig)
	if err != nil {
		log.Printf("clientcmd.BuildConfigFromFlags() failed: %v \n", err)
	}
	// create the clientset
	ClientSet, err = kubernetes.NewForConfig(config)
	if err != nil {
		log.Printf("kubernetes.NewForConfig() failed: %v \n", err)
	}
}

// CommonResponse 标准响应
type CommonResponse struct {
	Code int
	Data interface{}
}

// GenerateToken 获取jwt token
func GenerateToken(info *models.User, expiredSeconds int) (tokenString string, err error) {
	if expiredSeconds == 0 {
		expiredSeconds = DEFAULT_EXPIRE_SECONDS
	}
	// Create the Claims
	mySigningKey := []byte(KEY)
	expireAt := time.Now().Add(time.Second * time.Duration(expiredSeconds)).Unix()
	// pass parameter to this func or not
	user := *info
	claims := MyCustomClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: expireAt,
			Issuer:    user.Account,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(mySigningKey)
	if err != nil {
		log.Printf("token.SignedString() failed: %v \n", err)
	} else {
		tokenString = tokenStr
	}
	return
}

// ValidateToken 验证jtw token
func ValidateToken(tokenString string) (info models.User, err error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&MyCustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(KEY), nil
		})
	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		//fmt.Printf("%v %v", claims.User, claims.StandardClaims.ExpiresAt)
		//fmt.Println("token will be expired at ", time.Unix(claims.StandardClaims.ExpiresAt, 0))
		info = claims.User
	} else {
		log.Printf("validate tokenString failed: %v \n", err)
	}
	return
}

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
