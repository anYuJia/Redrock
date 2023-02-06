package control

import (
	sql2 "database/sql"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"sql/response"
	"sql/sql"
	"sql/struc"
	"sql/t"
	"strconv"
	"strings"
)

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	log.Println("username:" + username + "   " + "password:" + password)
	row, err := sql.GetSelectUser().Query(username)
	for row.Next() {
		var user _struc.User
		err = row.Scan(&user.Uid, &user.Username, &user.Password)
		if err != nil {
			log.Println(err)
			return
		}
		if user.Username == username {
			response.Fail(c, gin.H{}, "该用户名被注册过了")
			return
		}
	}
	defer func(row *sql2.Rows) {
		err = row.Close()
		if err != nil {

		}
	}(row)
	if len(username) > 20 {
		response.Fail(c, gin.H{}, "用户名必须要小于20位噢")
		return
	}
	if len(password) <= 6 || len(password) >= 16 {
		response.Fail(c, gin.H{}, "密码必须要大于6位噢而且小于16位噢")
		return
	}
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	newUser := _struc.User{
		Uid:      0,
		Username: username,
		Password: string(hashPassword),
	}
	_, err = sql.GetInsertUser().Exec(newUser.Uid, newUser.Username, newUser.Password)
	if err != nil {
		response.Fail(c, gin.H{}, "服务器错误，注册失败!")
		return
	}
	avatar := "http://ananqiexiafan.icu/d/tou.png"
	_, err = sql.GetInsertDetail().Exec(newUser.Uid, newUser.Username, avatar, t.GetDate(), "n", "n", 2, "n", t.GetDate(), 0, 0, 0)
	if err != nil {
		log.Println(err)
		return
	}
	response.Success(c, gin.H{}, "注册成功")
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	log.Println("username:" + username + "   " + "password:" + password)
	if len(username) > 20 {
		response.Fail(c, gin.H{}, "用户名必须要小于20位噢")
		return
	}
	if len(password) <= 6 || len(password) >= 16 {
		response.Fail(c, gin.H{}, "密码必须要大于6位噢而且小于16位噢")
		return
	}
	row, err := sql.GetSelectUser().Query(username)
	e := 0
	for row.Next() {
		e = 1
		var user _struc.User
		err = row.Scan(&user.Uid, &user.Username, &user.Password)
		if err != nil {
			log.Println(err)
			return
		}
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
			response.Fail(c, gin.H{}, "密码错误")
			return
		}
	}
	defer func(row *sql2.Rows) {
		err = row.Close()
		if err != nil {

		}
	}(row)
	if e == 0 {
		response.Fail(c, gin.H{}, "该用户不存在")
		return
	}
	c.SetCookie("user_cookie", username, 3600, "/", "", false, false)
	response.Success(c, gin.H{}, "登陆成功")
}
func UpdatePassword(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	haPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	_, err := sql.GetChangePassword().Exec(username, string(haPassword))
	if err != nil {
		log.Println(err)
		return
	}
	response.Success(c, gin.H{}, "修改密码成功")
}
func RemoveCookie(c *gin.Context) {
	c.SetCookie("user_cookie", "username", -1, "/", "", false, true)
	response.Success(c, gin.H{}, "退出登录成功!")
}
func GetDetail(c *gin.Context) {
	username := c.Query("username")
	var userDetail _struc.UserDetail
	row, err := sql.GetSelectDetail().Query(username)
	if err != nil {
		log.Println(err)
		return
	}
	for row.Next() {
		err = row.Scan(&userDetail.Uid, &userDetail.Username, &userDetail.Qq, &userDetail.Gender, &userDetail.Introduction, &userDetail.Email, &userDetail.Birthday, &userDetail.CreateTime, &userDetail.Avatar, &userDetail.Love, &userDetail.Loved, &userDetail.Stare)
		if err != nil {
			log.Println(err)
			return
		}
	}
	defer func(row *sql2.Rows) {
		err = row.Close()
		if err != nil {

		}
	}(row)
	response.Success(c, gin.H{
		"data": userDetail,
	}, "ok")

}
func UpdateDetail(c *gin.Context) {
	username, err := c.Cookie("user_cookie")
	if err != nil {
		log.Println(err)
		return
	}
	qq := c.Query("qq")
	_, err = strconv.Atoi(qq)
	if err != nil || len(qq) > 10 {
		response.Fail(c, gin.H{}, "qq必须小于10位的数字")
		return
	}
	gender, err := strconv.Atoi(c.Query("gender"))
	if gender > 1 || gender < 0 || err != nil {
		response.Fail(c, gin.H{}, "性别只能是0或者1,1为男,0为女")
		return
	}
	introduction := c.Query("introduction")
	if len(introduction) > 200 {
		response.Fail(c, gin.H{}, "介绍小于200字")
		return
	}
	email := c.Query("email")
	if !t.VerifyEmailFormat(email) {
		response.Fail(c, gin.H{}, "邮箱不合法")
		return
	}
	birthday := c.Query("birthday")
	if !t.VerifyDate(birthday) {
		response.Fail(c, gin.H{}, "生日不合法,月和日为1位的时候，请在前面加个0")
		return
	}
	_, err = sql.GetUpdateDetail().Exec(qq, gender, introduction, email, birthday, username)
	if err != nil {
		log.Println(err)
	}
	response.Success(c, gin.H{}, "修改成功")
}
func Love(c *gin.Context) {
	love, err := c.Cookie("user_cookie")
	if err != nil {
		log.Println(err)
		return
	}
	loved := c.Query("lovedUser")
	row, err := sql.GetSelectLoveList().Query(love)
	for row.Next() {
		var L _struc.Love
		var id int
		err = row.Scan(&id, &L.UsernameLove, &L.UserNameLoved)
		if err != nil {
			log.Println(err)
			return
		}
		if L.UserNameLoved == loved {
			response.Fail(c, gin.H{}, "已经关注了该用户"+loved)
			return
		}
	}
	defer func(row *sql2.Rows) {
		err = row.Close()
		if err != nil {

		}
	}(row)
	_, err = sql.GetInsertLoveList().Exec(0, love, loved)
	if err != nil {
		log.Println(err)
		return
	}

	_, err = sql.GetAddLove().Exec(love)
	if err != nil {
		log.Println(err)
		return
	}
	_, err = sql.GetAddLoved().Exec(loved)
	if err != nil {
		log.Println(err)
		return
	}
	response.Success(c, gin.H{}, "关注"+loved+"成功")
}
func GetLoveList(c *gin.Context) {
	love, err := c.Cookie("user_cookie")
	if err != nil {
		log.Println(err)
		return
	}
	row, err := sql.GetSelectLoveList().Query(love)
	if err != nil {
		log.Println(err)
		return
	}
	var LL []_struc.Love
	for row.Next() {
		var L _struc.Love
		var id int
		err = row.Scan(&id, &L.UsernameLove, &L.UserNameLoved)
		if err != nil {
			log.Println(err)
			return
		}
		LL = append(LL, L)
	}
	defer func(row *sql2.Rows) {
		err = row.Close()
		if err != nil {

		}
	}(row)
	response.Success(c, gin.H{
		"data": LL,
	}, "ok")
}
func GetLovedList(c *gin.Context) {
	loved, err := c.Cookie("user_cookie")
	if err != nil {
		log.Println(err)
		return
	}
	row, err := sql.GetSelectLovedList().Query(loved)
	if err != nil {
		log.Println(err)
		return
	}
	var LL []_struc.Love
	for row.Next() {
		var L _struc.Love
		var id int
		err = row.Scan(&id, &L.UsernameLove, &L.UserNameLoved)
		if err != nil {
			log.Println(err)
			return
		}
		LL = append(LL, L)
	}
	defer func(row *sql2.Rows) {
		err = row.Close()
		if err != nil {

		}
	}(row)
	response.Success(c, gin.H{
		"data": LL,
	}, "ok")
}
func UserStar(c *gin.Context) {
	starUser, _ := c.Cookie("user_cookie")
	staredUser := c.Query("staredUser")
	if staredUser == "" {
		response.Fail(c, gin.H{}, "您还没有输入被点赞用户的用户名呢!")
		return
	}
	row, err := sql.GetSelectUser().Query(staredUser)
	if err != nil {
		log.Println(err)
		return
	}
	defer func(row *sql2.Rows) {
		err = row.Close()
		if err != nil {

		}
	}(row)
	var flag = false
	for row.Next() {
		flag = true
		row, err = sql.GetSelectUserStarList().Query(starUser, staredUser)
		for row.Next() {
			var starList _struc.StarList
			err = row.Scan(&starList.Id, &starList.StarUser, &starList.StaredUser, &starList.Time)
			if err == nil || starList.Time == t.GetDate() {
				response.Fail(c, gin.H{}, "今天该用户已经点过赞了，不能重复点赞")
				return
			}
		}
		star1List := _struc.StarList{
			Id:         0,
			StarUser:   starUser,
			StaredUser: staredUser,
			Time:       t.GetDate(),
		}
		_, err = sql.GetInsertUserStarList().Exec(star1List.Id, star1List.StarUser, star1List.StaredUser, star1List.Time)
		if err != nil {
			log.Println(err)
			return
		}
		_, err = sql.GetAddStar().Exec(staredUser)
		if err != nil {
			log.Println(err)
			return
		}
		response.Success(c, gin.H{}, "您成功给用户"+staredUser+"点赞")
	}
	if !flag {
		response.Fail(c, gin.H{}, "该用户不存在")
		return
	}
}
func InsertArticle(c *gin.Context) {
	author, err := c.Cookie("user_cookie")
	if err != nil {
		log.Println(err)
		return
	}
	title := c.Query("title")
	content := c.Query("content")
	Tags := c.Query("tags")
	tags := strings.Split(Tags, " ")
	_, err = sql.GetInsertArticle().Exec(0, title, content, author, t.GetDate()+" "+t.GetTime())
	if err != nil {
		log.Println(err)
		return
	}
	row, err := sql.Getaid().Query(title)
	if err != nil {
		log.Println(err)
		return
	}
	var aid int
	for row.Next() {
		err = row.Scan(&aid)
		if err != nil {
			log.Println(err)
			return
		}
	}
	for i := range tags {
		_, err = sql.GetInsertTag().Exec(0, aid, tags[i])
		if err != nil {
			log.Println(err)
			return
		}
	}
	response.Success(c, gin.H{
		"articleId": t.GetAid(title),
	}, "发表文章成功！")
}
func InsertArticleComment(c *gin.Context) {
	title := c.Query("title")
	username, err := c.Cookie("user_cookie")
	content := c.Query("content")
	if err != nil {
		log.Println(err)
		return
	}
	aid := t.GetAid(title)
	_, err = sql.GetInsertArticleComment().Exec(0, aid, username, content, t.GetDate()+" "+t.GetTime())
	if err != nil {
		log.Println(err)
		return
	}
	_, err = sql.GetInsertArticleViewc().Exec(aid)
	if err != nil {
		log.Println(err)
		return
	}
	response.Success(c, gin.H{}, "评论成功!")
}
func RemoveArticleComment(c *gin.Context) {
	var flag = false
	user, err := c.Cookie("user_cookie")
	if err != nil {
		log.Println(err)
		return
	}
	acid := c.Query("acid")
	row, err := sql.GetArticleForId().Query(acid)
	var aid int
	for row.Next() {
		var username string
		err = row.Scan(&username, &aid)
		if err != nil {
			return
		}
		if user != username {
			response.Fail(c, gin.H{}, "这不是你的评论，你无法删除别人的评论!")
		}
		_, err = sql.GetArticleCommentCountD1().Exec(aid)
		if err != nil {
			log.Println(err)
			return
		}
		_, err = sql.GetRemoveArticleComment().Exec(acid)
		if err != nil {
			log.Println(err)
			return
		}
		response.Success(c, gin.H{}, "删除评论成功!")
		flag = true
	}
	if !flag {
		response.Fail(c, gin.H{}, "该评论不存在!")
	}
}
func InsertArticleStar(c *gin.Context) {
	user, err := c.Cookie("user_cookie")
	title := c.Query("title")
	if err != nil {
		log.Println(err)
		return
	}
	_, err = sql.GetInsertArticleStar().Exec(0, t.GetAid(title), user)
	if err != nil {
		log.Println(err)
		return
	}
	_, err = sql.GetInsertArticleViews().Exec(t.GetAid(title))
	response.Success(c, gin.H{}, "点赞成功")
}
func InsertArticleFavorite(c *gin.Context) {
	user, err := c.Cookie("user_cookie")
	title := c.Query("title")
	if err != nil {
		log.Println(err)
		return
	}
	_, err = sql.GetInsertArticleFavorite().Exec(0, t.GetAid(title), user)
	if err != nil {
		log.Println(err)
		return
	}
	_, err = sql.GetInsertArticleViewf().Exec(t.GetAid(title))
	response.Success(c, gin.H{}, "收藏成功")
}
func InsertArticleView(c *gin.Context) {
	user, err := c.Cookie("user_cookie")
	title := c.Query("title")
	if err != nil {
		log.Println(err)
		return
	}
	_, err = sql.GetInsertArticleViewList().Exec(0, t.GetAid(title), user, t.GetDate()+" "+t.GetTime())
	if err != nil {
		log.Println(err)
		return
	}
	_, err = sql.GetInsertArticleView().Exec(t.GetAid(title))
	response.Success(c, gin.H{}, "浏览成功")
}
func GetArticle(c *gin.Context) {
	aid := c.Query("articleId")
	var article _struc.Article
	row, err := sql.GetArticleDetail().Query(aid)
	if err != nil {
		log.Println(err)
		return
	}
	for row.Next() {
		err = row.Scan(&article.Aid, &article.Title, &article.Content, &article.Author, &article.CreateTime, &article.StarCount, &article.FavoriteCount, &article.ViewCount, &article.CommentCount)
		if err != nil {
			log.Println(err)
			return
		}
	}
	r, err := sql.GetArticleCommentList().Query(aid)
	if err != nil {
		log.Println(err)
		return
	}
	var articleCommentList []_struc.ArticleComment
	for r.Next() {
		var articleComment _struc.ArticleComment
		err = r.Scan(&articleComment.Acid, &articleComment.Aid, &articleComment.User, &articleComment.Comment, &articleComment.Time)
		if err != nil {
			log.Println(err)
			return
		}
		articleCommentList = append(articleCommentList, articleComment)
	}
	var tags []string
	R, err := sql.GetArticleTagList().Query(aid)
	if err != nil {
		log.Println(err)
		return
	}
	for R.Next() {
		var tag string
		err = R.Scan(&tag)
		if err != nil {
			log.Println(err)
			return
		}
		tags = append(tags, tag)
	}
	response.Success(c, gin.H{
		"commentList": articleCommentList,
		"detail":      article,
		"tags":        tags,
	}, "获取成功!")
}
func UserFavoriteList(c *gin.Context) {
	user := c.Query("username")
	row, err := sql.GetUserFavoriteList().Query(user)
	if err != nil {
		log.Println(err)
		return
	}
	var favoriteList []_struc.XList
	for row.Next() {
		var favorite _struc.XList
		err = row.Scan(&favorite.Axid, &favorite.Aid, &favorite.User)
		favoriteList = append(favoriteList, favorite)
	}
	response.Success(c, gin.H{
		"data": favoriteList,
	}, "获取用户收藏列表成功")
}
func UserStarList(c *gin.Context) {
	user := c.Query("username")
	row, err := sql.GetUserStarList().Query(user)
	if err != nil {
		log.Println(err)
		return
	}
	var starList []_struc.XList
	for row.Next() {
		var star _struc.XList
		err = row.Scan(&star.Axid, &star.Aid, &star.User)
		starList = append(starList, star)
	}
	response.Success(c, gin.H{
		"data": starList,
	}, "获取用户点赞列表成功")
}
func UserViewList(c *gin.Context) {
	user := c.Query("username")
	row, err := sql.GetUserViewList().Query(user)
	if err != nil {
		log.Println(err)
		return
	}
	var viewList []_struc.ViewList
	for row.Next() {
		var view _struc.ViewList
		err = row.Scan(&view.Avid, &view.Aid, &view.User, &view.Time)
		viewList = append(viewList, view)
	}
	response.Success(c, gin.H{
		"data": viewList,
	}, "获取用户浏览列表成功")
}
func GetHot(c *gin.Context) {
	var articles []_struc.Bd
	row, err := sql.GetArticleHot().Query()
	if err != nil {
		log.Println(err)
		return
	}
	for row.Next() {
		var article _struc.Article
		err = row.Scan(&article.Aid, &article.Title, &article.Content, &article.Author, &article.CreateTime, &article.StarCount, &article.FavoriteCount, &article.ViewCount, &article.CommentCount)
		if err != nil {
			log.Println(err)
			return
		}
		var tags []string
		R, err := sql.GetArticleTagList().Query(article.Aid)
		if err != nil {
			log.Println(err)
			return
		}
		for R.Next() {
			var tag string
			err = R.Scan(&tag)
			if err != nil {
				log.Println(err)
				return
			}
			tags = append(tags, tag)
		}
		var A = _struc.Bd{
			Art:  article,
			Tags: tags,
		}
		articles = append(articles, A)
	}
	response.Success(c, gin.H{
		"data": articles,
	}, "获取热榜成功")
}
func GetCommand(c *gin.Context) {
	var articles []_struc.Bd
	row, err := sql.GetArticleCommand().Query()
	if err != nil {
		log.Println(err)
		return
	}
	for row.Next() {
		var article _struc.Article
		err = row.Scan(&article.Aid, &article.Title, &article.Content, &article.Author, &article.CreateTime, &article.StarCount, &article.FavoriteCount, &article.ViewCount, &article.CommentCount)
		if err != nil {
			log.Println(err)
			return
		}
		var tags []string
		R, err := sql.GetArticleTagList().Query(article.Aid)
		if err != nil {
			log.Println(err)
			return
		}
		for R.Next() {
			var tag string
			err = R.Scan(&tag)
			if err != nil {
				log.Println(err)
				return
			}
			tags = append(tags, tag)
		}
		var A = _struc.Bd{
			Art:  article,
			Tags: tags,
		}
		articles = append(articles, A)
	}
	response.Success(c, gin.H{
		"data": articles,
	}, "获取推荐榜成功")
}
func GetArticleTimeA(c *gin.Context) {
	var articles []_struc.Bd
	row, err := sql.GetArticleTime().Query()
	if err != nil {
		log.Println(err)
		return
	}
	for row.Next() {
		var article _struc.Article
		err = row.Scan(&article.Aid, &article.Title, &article.Content, &article.Author, &article.CreateTime, &article.StarCount, &article.FavoriteCount, &article.ViewCount, &article.CommentCount)
		if err != nil {
			log.Println(err)
			return
		}
		var tags []string
		R, err := sql.GetArticleTagList().Query(article.Aid)
		if err != nil {
			log.Println(err)
			return
		}
		for R.Next() {
			var tag string
			err = R.Scan(&tag)
			if err != nil {
				log.Println(err)
				return
			}
			tags = append(tags, tag)
		}
		var A = _struc.Bd{
			Art:  article,
			Tags: tags,
		}
		articles = append(articles, A)
	}
	response.Success(c, gin.H{
		"data": articles,
	}, "获取最新榜成功")
}
func CancerArticleStar(c *gin.Context) {
	var flag = false
	user, err := c.Cookie("user_cookie")
	if err != nil {
		log.Println(err)
		return
	}
	aid := c.Query("aid")
	row, err := sql.GetSelectStarRecord().Query(aid, user)
	for row.Next() {
		var asid int
		err = row.Scan(&asid)
		_, err = sql.GetDeleteStarRecord().Exec(asid)
		if err != nil {
			response.Fail(c, gin.H{}, "输入的aid不存在!")
			log.Println(err)
			return
		}
		_, err = sql.GetArticleStarCountD1().Exec(aid)
		if err != nil {
			log.Println(err)
			return
		}
		response.Success(c, gin.H{}, "取消点赞成功!")
		flag = true
	}
	if !flag {
		response.Fail(c, gin.H{}, "只能删除自己的点赞!")
	}
}
func CancerArticleFavorite(c *gin.Context) {
	var flag = false
	user, err := c.Cookie("user_cookie")
	if err != nil {
		log.Println(err)
		return
	}
	aid := c.Query("aid")
	row, err := sql.GetSelectFavoriteRecord().Query(aid, user)
	if err != nil {
		log.Println(err)
		return
	}
	for row.Next() {
		var asid int
		err = row.Scan(&asid)
		if err != nil {
			log.Println(err)
			return
		}
		_, err = sql.GetDeleteFavoriteRecord().Exec(asid)
		if err != nil {
			response.Fail(c, gin.H{}, "输入的asid不存在!")
			log.Println(err)
			return
		}
		_, err = sql.GetArticleFavoriteCountD1().Exec(aid)
		if err != nil {
			log.Println(err)
			return
		}
		response.Success(c, gin.H{}, "取消收藏成功!")
		flag = true
	}
	if !flag {
		response.Fail(c, gin.H{}, "只能取消自己的收藏!")
	}
}
func CancerStarUser(c *gin.Context) {
	var flag = false
	starUser, err := c.Cookie("user_cookie")
	if err != nil {
		log.Println(err)
		return
	}
	staredUser := c.Query("staredUser")
	row, err := sql.GetUserLoveList().Query(starUser, staredUser)
	if err != nil {
		log.Println(err)
		return
	}
	for row.Next() {
		_, err = sql.GetDeleteUserStarRecord().Exec(starUser, staredUser)
		if err != nil {
			log.Println(err)
			return
		}
		_, err = sql.GetUserStarCountD1().Exec(starUser)
		if err != nil {
			log.Println(err)
		}
		_, err = sql.GetUserStaredCountD1().Exec(staredUser)
		if err != nil {
			log.Println(err)
			return
		}
		response.Success(c, gin.H{}, "你成功取消关注"+staredUser)
		flag = true
	}
	if !flag {
		response.Fail(c, gin.H{}, "你还未关注用户:"+staredUser)
	}

}
func GetUserFavoriteRecord(c *gin.Context) {
	var flag = false
	user, err := c.Cookie("user_cookie")
	if err != nil {
		log.Println(err)
		return
	}
	var articles []_struc.Bd
	R, err := sql.GetSelectUserFavorite().Query(user)
	if err != nil {
		log.Println(err)
		return
	}
	for R.Next() {
		var aid int
		err = R.Scan(&aid)
		row, err := sql.GetArticleById().Query(aid)
		if err != nil {
			log.Println(err)
			return
		}
		for row.Next() {
			var article _struc.Article
			err = row.Scan(&article.Aid, &article.Title, &article.Content, &article.Author, &article.CreateTime, &article.StarCount, &article.FavoriteCount, &article.ViewCount, &article.CommentCount)
			if err != nil {
				log.Println(err)
				return
			}
			var tags []string
			R, err := sql.GetArticleTagList().Query(article.Aid)
			if err != nil {
				log.Println(err)
				return
			}
			for R.Next() {
				var tag string
				err = R.Scan(&tag)
				if err != nil {
					log.Println(err)
					return
				}
				tags = append(tags, tag)
			}
			var A = _struc.Bd{
				Art:  article,
				Tags: tags,
			}
			articles = append(articles, A)
			flag = true
		}
	}
	if flag {
		response.Success(c, gin.H{
			"data": articles,
		}, "获取用户收藏列表成功!")
	} else {
		response.Fail(c, gin.H{}, "用户收藏列表为空!")
	}
}
func GetUserStarRecord(c *gin.Context) {
	var flag = false
	user, err := c.Cookie("user_cookie")
	if err != nil {
		log.Println(err)
		return
	}
	var articles []_struc.Bd
	R, err := sql.GetSelectUserStar().Query(user)
	if err != nil {
		log.Println(err)
		return
	}
	for R.Next() {
		var aid int
		err = R.Scan(&aid)
		row, err := sql.GetArticleById().Query(aid)
		if err != nil {
			log.Println(err)
			return
		}
		for row.Next() {
			var article _struc.Article
			err = row.Scan(&article.Aid, &article.Title, &article.Content, &article.Author, &article.CreateTime, &article.StarCount, &article.FavoriteCount, &article.ViewCount, &article.CommentCount)
			if err != nil {
				log.Println(err)
				return
			}
			var tags []string
			R, err := sql.GetArticleTagList().Query(article.Aid)
			if err != nil {
				log.Println(err)
				return
			}
			for R.Next() {
				var tag string
				err = R.Scan(&tag)
				if err != nil {
					log.Println(err)
					return
				}
				tags = append(tags, tag)
			}
			var A = _struc.Bd{
				Art:  article,
				Tags: tags,
			}
			articles = append(articles, A)
			flag = true
		}
	}
	if flag {
		response.Success(c, gin.H{
			"data": articles,
		}, "获取用户点赞列表成功!")
	} else {
		response.Fail(c, gin.H{}, "用户点赞列表为空!")
	}
}
func GetUserViewRecord(c *gin.Context) {
	var flag = false
	user, err := c.Cookie("user_cookie")
	if err != nil {
		log.Println(err)
		return
	}
	var articles []_struc.Bdt
	R, err := sql.GetSelectViewRecord().Query(user)
	if err != nil {
		log.Println(err)
		return
	}
	for R.Next() {
		var aid int
		var time string
		err = R.Scan(&aid, &time)
		row, err := sql.GetArticleById().Query(aid)
		if err != nil {
			log.Println(err)
			return
		}
		for row.Next() {
			var article _struc.Article
			err = row.Scan(&article.Aid, &article.Title, &article.Content, &article.Author, &article.CreateTime, &article.StarCount, &article.FavoriteCount, &article.ViewCount, &article.CommentCount)
			if err != nil {
				log.Println(err)
				return
			}
			var tags []string
			R, err := sql.GetArticleTagList().Query(article.Aid)
			if err != nil {
				log.Println(err)
				return
			}
			for R.Next() {
				var tag string
				err = R.Scan(&tag)
				if err != nil {
					log.Println(err)
					return
				}
				tags = append(tags, tag)
			}
			var A = _struc.Bdt{
				Time: time,
				Art:  article,
				Tags: tags,
			}
			articles = append(articles, A)
			flag = true
		}
	}
	if flag {
		response.Success(c, gin.H{
			"data": articles,
		}, "获取用户浏览列表成功!")
	} else {
		response.Fail(c, gin.H{}, "用户浏览列表为空!")
	}
}
func GetTagList(c *gin.Context) {
	tag := c.Query("tag")
	var articles []_struc.Bd
	var flag = false
	R, err := sql.GetAidForTag().Query(tag)
	if err != nil {
		log.Println(err)
		return
	}
	for R.Next() {
		var aid int
		err = R.Scan(&aid)
		row, err := sql.GetArticleById().Query(aid)
		if err != nil {
			log.Println(err)
			return
		}
		for row.Next() {
			var article _struc.Article
			err = row.Scan(&article.Aid, &article.Title, &article.Content, &article.Author, &article.CreateTime, &article.StarCount, &article.FavoriteCount, &article.ViewCount, &article.CommentCount)
			if err != nil {
				log.Println(err)
				return
			}
			var tags []string
			R, err := sql.GetArticleTagList().Query(article.Aid)
			if err != nil {
				log.Println(err)
				return
			}
			for R.Next() {
				var tag string
				err = R.Scan(&tag)
				if err != nil {
					log.Println(err)
					return
				}
				tags = append(tags, tag)
			}
			var A = _struc.Bd{
				Art:  article,
				Tags: tags,
			}
			articles = append(articles, A)
			flag = true
		}

	}
	if flag {
		response.Success(c, gin.H{
			"data": articles,
		}, "获取成功标签页")
	} else {
		response.Fail(c, gin.H{}, "获取标签为空!")
	}
}
func GetUserArticle(c *gin.Context) {
	user, err := c.Cookie("user_cookie")
	if err != nil {
		log.Println(err)
		return
	}
	var articles []_struc.Bd
	var flag = false
	r, err := sql.GetSelectLoveList().Query(user)
	if err != nil {
		log.Println(err)
		return
	}
	for r.Next() {
		var userLoved string
		var userLove string
		var id int
		err = r.Scan(&id, &userLove, &userLoved)
		R, err := sql.GetUserArticle().Query(userLoved)
		if err != nil {
			log.Println(err)
			return
		}
		for R.Next() {
			var article _struc.Article
			err = R.Scan(&article.Aid, &article.Title, &article.Content, &article.Author, &article.CreateTime, &article.StarCount, &article.FavoriteCount, &article.ViewCount, &article.CommentCount)
			if err != nil {
				log.Println(err)
				return
			}
			var tags []string
			R, err := sql.GetArticleTagList().Query(article.Aid)
			if err != nil {
				log.Println(err)
				return
			}
			for R.Next() {
				var tag string
				err = R.Scan(&tag)
				if err != nil {
					log.Println(err)
					return
				}
				tags = append(tags, tag)
			}
			var A = _struc.Bd{
				Art:  article,
				Tags: tags,
			}
			articles = append(articles, A)
			flag = true
		}

	}
	if flag {
		response.Success(c, gin.H{
			"data": articles,
		}, "获取成功关注页")
	} else {
		response.Success(c, gin.H{}, "获取关注为空!")
	}
}
func GetUserArticleList(c *gin.Context) {
	user := c.Query("username")
	var articles []_struc.Bd
	var flag = false
	R, err := sql.GetUserArticle().Query(user)
	if err != nil {
		log.Println(err)
		return
	}
	for R.Next() {
		var article _struc.Article
		err = R.Scan(&article.Aid, &article.Title, &article.Content, &article.Author, &article.CreateTime, &article.StarCount, &article.FavoriteCount, &article.ViewCount, &article.CommentCount)
		if err != nil {
			log.Println(err)
			return
		}
		var tags []string
		R, err := sql.GetArticleTagList().Query(article.Aid)
		if err != nil {
			log.Println(err)
			return
		}
		for R.Next() {
			var tag string
			err = R.Scan(&tag)
			if err != nil {
				log.Println(err)
				return
			}
			tags = append(tags, tag)
		}
		var A = _struc.Bd{
			Art:  article,
			Tags: tags,
		}
		articles = append(articles, A)
		flag = true
	}
	if flag {
		response.Success(c, gin.H{
			"data": articles,
		}, "获取用户文章列表成功")
	} else {
		response.Success(c, gin.H{}, "用户文章为空!")
	}
}
