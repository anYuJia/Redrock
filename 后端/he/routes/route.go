package routes

import (
	"github.com/gin-gonic/gin"
	"sql/control"
)

func CollectRouter(r *gin.Engine) {
	//注册
	r.GET("/register", control.Register)
	//登录
	r.GET("/login", control.Login)
	//更新密码
	r.GET("/updatePassword", HandlerFunc(), control.UpdatePassword)
	//退出登录
	r.GET("/logout", HandlerFunc(), control.RemoveCookie)
	//获取用户详情信息
	r.GET("/getDetail", HandlerFunc(), control.GetDetail)
	//更新用户信息
	r.GET("/updateDetail", HandlerFunc(), control.UpdateDetail)
	//用户关注用户
	r.GET("/love", HandlerFunc(), control.Love)
	//获取用户关注列表
	r.GET("/getLoveList", HandlerFunc(), control.GetLoveList)
	//获取用户被关注列表
	r.GET("/getLovedList", HandlerFunc(), control.GetLovedList)
	//给别的用户点赞
	r.GET("/userStar", HandlerFunc(), control.UserStar)
	//发表文章
	r.GET("/insertArticle", HandlerFunc(), control.InsertArticle)
	//发布文章评论
	r.GET("/insertArticleComment", HandlerFunc(), control.InsertArticleComment)
	//给文章点赞
	r.GET("/insertArticleStar", HandlerFunc(), control.InsertArticleStar)
	//收藏文章
	r.GET("/insertArticleFavorite", HandlerFunc(), control.InsertArticleFavorite)
	//浏览文章
	r.GET("/insertArticleView", HandlerFunc(), control.InsertArticleView)
	//获取文章详情信息
	r.GET("/GetArticle", HandlerFunc(), control.GetArticle)
	//获取用户收藏的文章列表
	r.GET("/userFavoriteList", HandlerFunc(), control.UserFavoriteList)
	//获取用户的浏览列表
	r.GET("/userViewList", HandlerFunc(), control.UserViewList)
	//获取用户的点赞列表
	r.GET("/userStarList", HandlerFunc(), control.UserStarList)
	//获取热榜
	r.GET("/getHot", HandlerFunc(), control.GetHot)
	//获取最新榜
	r.GET("/GetArticleTimeA", HandlerFunc(), control.GetArticleTimeA)
	//获取推荐榜单
	r.GET("/GetCommand", HandlerFunc(), control.GetCommand)
	//删除用户对文章的评论
	r.GET("/RemoveArticleComment", HandlerFunc(), control.RemoveArticleComment)
	//取消用户对文章的收藏
	r.GET("/CancerArticleFavorite", HandlerFunc(), control.CancerArticleFavorite)
	//取消用户对文章的点赞
	r.GET("/CancerArticleStar", HandlerFunc(), control.CancerArticleStar)
	//取消关注的用户
	r.GET("/CancerStarUser", HandlerFunc(), control.CancerStarUser)
	//获取自己的喜欢列表详情
	r.GET("/GetUserFavoriteRecord", HandlerFunc(), control.GetUserFavoriteRecord)
	//获取自己的点赞列表详情
	r.GET("/GetUserStarRecord", HandlerFunc(), control.GetUserStarRecord)
	//获取自己的浏览列表详情
	r.GET("/GetUserViewRecord", HandlerFunc(), control.GetUserViewRecord)
	//获取文章标签列表
	r.GET("/GetTagList", HandlerFunc(), control.GetTagList)
	//获取专注的用户发布的文章
	r.GET("/GetUserArticle", HandlerFunc(), control.GetUserArticle)
	//获取用户文章列表
	r.GET("/GetUserArticleList", HandlerFunc(), control.GetUserArticleList)
}
