package sql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var DB *sql.DB
var selectUser *sql.Stmt
var selectUserDetail *sql.Stmt
var insertUser *sql.Stmt
var changePassword *sql.Stmt
var insertDetail *sql.Stmt
var updateDetail *sql.Stmt
var insertLoveList *sql.Stmt
var selectLoveList *sql.Stmt
var selectLovedList *sql.Stmt
var addLove *sql.Stmt
var addLoved *sql.Stmt
var addStar *sql.Stmt
var selectUserStarList *sql.Stmt
var insertUserStarList *sql.Stmt
var insertArticle *sql.Stmt
var insertTag *sql.Stmt
var insertArticleComment *sql.Stmt
var insertArticleStar *sql.Stmt
var insertArticleFavorite *sql.Stmt
var insertViewList *sql.Stmt
var insertArticleView *sql.Stmt
var insertArticleViews *sql.Stmt
var insertArticleViewf *sql.Stmt
var insertArticleViewc *sql.Stmt
var getaid *sql.Stmt
var getArticleStarList *sql.Stmt
var getArticleCommentList *sql.Stmt
var getArticleDetail *sql.Stmt
var getUserStarList *sql.Stmt
var getUserFavoriteList *sql.Stmt
var getUserViewList *sql.Stmt
var getArticleTime *sql.Stmt
var getArticleHot *sql.Stmt
var getArticleCommand *sql.Stmt
var deleteViewRecord *sql.Stmt
var deleteFavoriteRecord *sql.Stmt
var deleteStarRecord *sql.Stmt
var deleteUserStarRecord *sql.Stmt
var selectViewRecord *sql.Stmt
var selectFavoriteRecord *sql.Stmt
var selectStarRecord *sql.Stmt
var articleStarCountD1 *sql.Stmt
var articleFavoriteCountD1 *sql.Stmt
var userStarCountD1 *sql.Stmt
var userStaredCountD1 *sql.Stmt
var removeArticleComment *sql.Stmt
var getArticleForId *sql.Stmt
var articleCommentCountD1 *sql.Stmt
var getArticleTagList *sql.Stmt
var getUserLoveList *sql.Stmt
var selectUserStar *sql.Stmt
var selectUserFavorite *sql.Stmt
var getArticleById *sql.Stmt
var getAidForTag *sql.Stmt
var getUserArticle *sql.Stmt
var getUserLoveUser *sql.Stmt

func InitUser() {
	// 连接数据库
	db, err := sql.Open("mysql", "root:z0114b3586@tcp(127.0.0.1:3306)/JJ")
	if err != nil {
		fmt.Println("数据库连接失败：", err.Error())
		return
	}
	Create(db)
	prepare(db)
	DB = db
}
func Create(db *sql.DB) {
	//创建用户表
	CreateTable(db, "user", "CREATE TABLE IF NOT EXISTS `user` (`uid` int(10) NOT NULL AUTO_INCREMENT,`username` varchar(20) NOT NULL DEFAULT '1', `password` varchar(256) NOT NULL DEFAULT '1' ,PRIMARY KEY (`uid`)) ;")
	//创建用户细节表
	CreateTable(db, "userDetail", "CREATE TABLE IF NOT EXISTS `userDetail` (`uid` int(10) NOT NULL AUTO_INCREMENT,`username` varchar(20) NOT NULL DEFAULT '1',`qq` varchar(10) NOT NULL DEFAULT '1', `gender` tinyint(1) DEFAULT NULL,`introduction` varchar(256) DEFAULT NULL,`email` varchar(30) DEFAULT NULL,`birthday` date DEFAULT NULL,`createtime` date DEFAULT NULL,`avatar` varchar(100) NOT NULL DEFAULT '1',`love` INT NOT NULL DEFAULT 0,`loved` INT NOT NULL DEFAULT 0,`star` INT NOT NULL DEFAULT 0,PRIMARY KEY (`uid`)) ;")
	//创建存放cookie的表
	//CreateTable(db, "userCookie", "CREATE TABLE `userCookie` (`cookie` varchar(100) NOT NULL,PRIMARY KEY (`uid`)) ;")
	//创建关注列表
	CreateTable(db, "loveList", "create table IF NOT EXISTS `loveList` (`id` int(10) NOT NULL AUTO_INCREMENT,`loveUsername` varchar(20) NOT NULL DEFAULT '1',`lovedUsername` varchar(20) NOT NULL DEFAULT '1',PRIMARY KEY (`id`)) ;")
	//创建点赞列表te
	CreateTable(db, "userStarList", "create table IF NOT EXISTS `userStarList` (`id` int(10) not null auto_increment,`starUser` varchar(20) not null default '1',`staredUser` varchar(20) not null default '1',`time` date not null,primary key (`id`)) ;")
	//创建文章表
	CreateTable(db, "articleList", "create table IF NOT EXISTS `articleList` (`aid` int(10) not null auto_increment,`title` varchar(50) not null default '1',`content` varchar(3000) not null default '1' ,`author` varchar(20) not null default '1',`publicTime` datetime not null , `starCount` int not null default 0,`favorite` int not null default 0,`view` int not null default 0 ,`commentCount` int default 0,primary key (`aid`))")
	//创建文章标签表
	CreateTable(db, "tagList", "create table IF NOT EXISTS `tagList` (`tid` int(10) not null auto_increment,`aid` int(10) not null default 0,`tag` varchar(20) not null default '1',primary key(`tid`))")
	//创建文章点赞表
	CreateTable(db, "articleStar", "create table IF NOT EXISTS `articleStar` (`asid` int(10) not null auto_increment,`aid` int(10) not null,`user` varchar(20) not null default '1',primary key(`asid`))")
	//创建文章收藏表
	CreateTable(db, "articleFavorite", "create table IF NOT EXISTS `articleFavorite` (`afid` int(10) not null auto_increment,`aid` int(10) not null,`user` varchar(20) not null default '1',primary key(`afid`))")
	//创建文章评论表
	CreateTable(db, "articleComment", "create table IF NOT EXISTS `articleComment` (`acid` int(10) not null auto_increment,`aid` int(10) not null,`user` varchar(20) not null default '1',`comment` varchar(200) not null default '1',`time` datetime not null,primary key(`acid`))")
	//创建浏览记录表
	CreateTable(db, "viewList", "create table IF NOT EXISTS `viewList` (`avid` int(10) not null auto_increment,`aid` int(10) not null,`user` varchar(20) not null default '1',`time` datetime not null,primary key(`avid`))")
}
func CreateTable(db *sql.DB, tableName string, s string) {
	rows, err := db.Query("select * from ?", tableName)
	if err != nil {
		// 如果不存在user表就创建该表
		_, err := db.Exec(s)
		if err != nil {
			log.Println("创建"+tableName+"表失败", err)
			return
		}
		log.Println("创建" + tableName + "表成功")
	} else {
		defer func(rows *sql.Rows) {
			err := rows.Close()
			if err != nil {

			}
		}(rows)
	}
}
func GetDB() *sql.DB {
	return DB
}
func CloseDB() {
	err := DB.Close()
	if err != nil {
		return
	}
}
func prepare(db *sql.DB) {
	//查询用户
	su, err := db.Prepare("select * from user where username = ?")
	if err != nil {
		log.Println(err)
		return
	}
	selectUser = su
	//查询用户信息
	sd, err := db.Prepare("select * from userDetail where username = ?")
	if err != nil {
		log.Println(err)
		return
	}
	selectUserDetail = sd

	//新增用户
	iu, err := db.Prepare("insert into user (uid,username,password) values (?,?,?)")
	if err != nil {
		log.Println(err)
		return
	}
	insertUser = iu
	//新增用户细节表
	id, err := db.Prepare("insert into userDetail (uid,username,avatar,birthday,introduction,qq,gender,email,createTime,love,loved,star) values (?,?,?,?,?,?,?,?,?,?,?,?)")
	if err != nil {
		log.Println(err)
		return
	}
	insertDetail = id
	//更新密码
	cp, err := db.Prepare("update user set password = ? where username = ?")
	if err != nil {
		log.Println(err)
		return
	}
	changePassword = cp
	//更新用户信息
	ud, err := db.Prepare("update userDetail set qq = ? , gender=?,introduction=?,email=?,birthday=? where username=?")
	if err != nil {
		log.Println(err)
		return
	}
	updateDetail = ud
	//新增关注列表
	il, err := db.Prepare("insert into loveList (id,loveUsername,lovedUsername) values (?,?,?)")
	if err != nil {
		log.Println(err)
		return
	}
	insertLoveList = il
	sll, err := db.Prepare("select * from loveList where loveUsername=?")
	if err != nil {
		log.Println(err)
		return
	}
	selectLoveList = sll
	slld, err := db.Prepare("select * from loveList where lovedUsername=?")
	if err != nil {
		log.Println(err)
		return
	}
	selectLovedList = slld
	al, err := db.Prepare("update userDetail set love = love+1 where username = ?")
	if err != nil {
		log.Println(err)
		return
	}
	addLove = al
	ald, err := db.Prepare("update userDetail set loved = loved+1 where username = ?")
	if err != nil {
		log.Println(err)
		return
	}
	addLoved = ald

	as, err := db.Prepare("update userDetail set star = star+1 where username = ?")
	if err != nil {
		log.Println(err)
		return
	}
	addStar = as
	//
	susl, err := db.Prepare("select * from userStarList where starUser=? and staredUser=?")
	if err != nil {
		log.Println(err)
		return
	}
	selectUserStarList = susl
	//
	iusl, err := db.Prepare("insert into userStarList (id,starUser,staredUser,time) values (?,?,?,?)")
	if err != nil {
		log.Println(err)
		return
	}
	insertUserStarList = iusl
	//新增文章
	ia, err := db.Prepare("insert into articleList (aid,title,content,author,publicTime) values (?,?,?,?,?)")
	if err != nil {
		log.Println(err)
		return
	}
	insertArticle = ia
	//新增文章标签列表
	it, err := db.Prepare("insert into tagList (tid,aid,tag) values (?,?,?)")
	if err != nil {
		log.Println(err)
		return
	}
	insertTag = it
	//新增文章评论列表
	iac, err := db.Prepare("insert into articleComment (acid,aid,user,comment,time) values (?,?,?,?,?)")
	if err != nil {
		log.Println(err)
		return
	}
	insertArticleComment = iac
	//新增文章点赞列表
	ias, err := db.Prepare("insert into articleStar (asid,aid,user) values (?,?,?)")
	if err != nil {
		log.Println(err)
		return
	}
	insertArticleStar = ias
	//新增文章收藏列表
	iaf, err := db.Prepare("insert into articleFavorite (afid,aid,user) values (?,?,?)")
	if err != nil {
		log.Println(err)
		return
	}
	insertArticleFavorite = iaf
	//新增文章浏览列表
	iavl, err := db.Prepare("insert into viewList (avid,aid,user,time) values (?,?,?,?)")
	if err != nil {
		log.Println(err)
		return
	}
	insertViewList = iavl
	//新增文章浏览
	iav, err := db.Prepare("update articleList set view = view+1 where aid = ?")
	if err != nil {
		log.Println(err)
		return
	}
	insertArticleView = iav
	//新增文章点赞
	iavs, err := db.Prepare("update articleList set starCount = starCount+1 where aid = ?")
	if err != nil {
		log.Println(err)
		return
	}
	insertArticleViews = iavs
	//新增文章收藏
	iavf, err := db.Prepare("update articleList set favorite = favorite+1 where aid = ?")
	if err != nil {
		log.Println(err)
		return
	}
	insertArticleViewf = iavf
	//新增文章评论
	iavc, err := db.Prepare("update articleList set commentCount = commentCount+1 where aid = ?")
	if err != nil {
		log.Println(err)
		return
	}
	insertArticleViewc = iavc
	//获取文章aid
	gaid, err := db.Prepare("select aid from articleList where title = ?")
	if err != nil {
		log.Println(err)
		return
	}
	getaid = gaid
	//获取文章的点赞列表
	gasl, err := db.Prepare("select user from articleStar where aid = ?")
	if err != nil {
		log.Println(err)
		return
	}
	getArticleStarList = gasl
	//获取文章的评论列表
	dacl, err := db.Prepare("select * from articleComment where aid = ?")
	if err != nil {
		log.Println(err)
		return
	}
	getArticleCommentList = dacl
	//获取文章详情信息
	gad, err := db.Prepare("select * from articleList where aid = ?")
	if err != nil {
		log.Println(err)
		return
	}
	getArticleDetail = gad
	//获取用户点赞列表
	gusl, err := db.Prepare("select * from articleStar where user = ? order by asid desc")
	if err != nil {
		log.Println(err)
		return
	}
	getUserStarList = gusl
	//获取用户收藏列表
	gufl, err := db.Prepare("select * from articleFavorite where user = ? order by afid desc")
	if err != nil {
		log.Println(err)
		return
	}
	getUserFavoriteList = gufl
	//获取用户浏览列表
	guvl, err := db.Prepare("select * from viewList where user = ? order by avid desc")
	if err != nil {
		log.Println(err)
		return
	}
	getUserViewList = guvl
	gaTime, err := db.Prepare("select * from articleList order by aid desc limit 1000")
	if err != nil {
		log.Println(err)
		return
	}
	getArticleTime = gaTime
	gaRe, err := db.Prepare("select * from articleList order by view desc limit 1000")
	if err != nil {
		log.Println(err)
		return
	}
	getArticleHot = gaRe
	gac, err := db.Prepare("select * from articleList order by starCount desc limit 1000")
	if err != nil {
		log.Println(err)
		return
	}
	getArticleCommand = gac
	//删除浏览记录
	dvc, err := db.Prepare("delete from viewList where avid = ?")
	if err != nil {
		log.Println(err)
		return
	}
	deleteViewRecord = dvc

	//查询收藏记录
	sfc, err := db.Prepare("select afid from articleFavorite where aid=? and user=?")
	if err != nil {
		log.Println(err)
		return
	}
	selectFavoriteRecord = sfc
	//删除收藏记录
	dfc, err := db.Prepare("delete from articleFavorite where afid = ?")
	if err != nil {
		log.Println(err)
		return
	}
	deleteFavoriteRecord = dfc
	//查询点赞记录
	ssc, err := db.Prepare("select asid from articleStar where aid=? and user=?")
	if err != nil {
		log.Println(err)
		return
	}
	selectStarRecord = ssc
	//删除点赞记录
	dsc, err := db.Prepare("delete from articleStar where asid = ?")
	if err != nil {
		log.Println(err)
		return
	}
	deleteStarRecord = dsc
	//文章点赞数-1
	asd1, err := db.Prepare("update articleList set starCount = starCount-1 where aid = ?")
	if err != nil {
		log.Println(err)
		return
	}
	articleStarCountD1 = asd1
	//文章收藏数-1
	afd1, err := db.Prepare("update articleList set favorite = favorite-1 where aid = ?")
	if err != nil {
		log.Println(err)
		return
	}
	articleFavoriteCountD1 = afd1
	//关注作者数-1
	usd1, err := db.Prepare("update userDetail set love = love-1 where username = ?")
	if err != nil {
		log.Println(err)
		return
	}
	userStarCountD1 = usd1
	//被关注作者数-1
	usdd1, err := db.Prepare("update userDetail set loved = loved-1 where username = ?")
	if err != nil {
		log.Println(err)
		return
	}
	userStaredCountD1 = usdd1
	//评论数量-1
	accd1, err := db.Prepare("update articleList set commentCount=commentCount-1  where aid=?")
	if err != nil {
		log.Println(err)
		return
	}
	articleCommentCountD1 = accd1
	//删除关注记录
	dsdc, err := db.Prepare("delete from loveList where loveUsername=? and lovedUsername=?")
	if err != nil {
		log.Println(err)
		return
	}
	deleteUserStarRecord = dsdc
	//获取是否存在关注关系
	gull, err := db.Prepare("select * from loveList where loveUsername=? and lovedUsername=?")
	if err != nil {
		log.Println(err)
		return
	}
	getUserLoveList = gull
	//删除评论
	racl, err := db.Prepare("delete from articleComment where acid = ?")
	if err != nil {
		log.Println(err)
		return
	}
	removeArticleComment = racl
	gacf, err := db.Prepare("select user,aid from articleComment where acid = ?")
	if err != nil {
		log.Println(err)
		return
	}
	getArticleForId = gacf
	gatl, err := db.Prepare("select tag from tagList where aid = ?")
	if err != nil {
		log.Println(err)
		return
	}
	getArticleTagList = gatl
	//查询用户浏览记录
	svc, err := db.Prepare("select aid,time from viewList where user = ? order by avid desc ")
	if err != nil {
		log.Println(err)
		return
	}
	selectViewRecord = svc
	//查询用户点赞列表
	sus, err := db.Prepare("select aid from articleStar where user = ? order by asid desc ")
	if err != nil {
		log.Println(err)
		return
	}
	selectUserStar = sus
	//查询用户收藏
	suf, err := db.Prepare("select aid from articleFavorite where user = ? order by afid desc ")
	if err != nil {
		log.Println(err)
		return
	}
	selectUserFavorite = suf
	//获取文章通过aid
	gabid, err := db.Prepare("select * from articleList where aid = ?;")
	if err != nil {
		log.Println(err)
		return
	}
	getArticleById = gabid
	//获取不同标签aid
	gaft, err := db.Prepare("select aid from tagList where tag = ? order by aid desc")
	if err != nil {
		log.Println(err)
		return
	}
	getAidForTag = gaft
	//获取用户发布的文章
	gua, err := db.Prepare("select * from articleList where author = ? order by aid desc")
	if err != nil {
		log.Println(err)
		return
	}
	getUserArticle = gua

}
