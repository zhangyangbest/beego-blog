package routers

import (
	"myblog/controllers"
	"github.com/astaxie/beego"
)

func init() {

	beego.Router("/", &controllers.HomeController{})

	//注册
	beego.Router("/register", &controllers.RegisterController{})
	//登录
	beego.Router("/login", &controllers.LoginController{})

	//退出
	beego.Router("/exit", &controllers.ExitController{})

	//写文章
	beego.Router("/article/add", &controllers.AddArticleController{})

	//显示文章详情
	beego.Router("/article/:id", &controllers.ShowArticleController{})

	//更新文章
	beego.Router("/article/update", &controllers.UpdateArticleController{})

	// 删除文章
	beego.Router("/article/delete", &controllers.DeleteArticleController{})

	//标签
	beego.Router("/tags", &controllers.TagsController{})

	//相册
	beego.Router("/album", &controllers.AlbumController{})

	//文件上传
	beego.Router("/upload", &controllers.UploadController{})

	//关于我页面
	beego.Router("/aboutme",&controllers.AboutMeController{})

}
