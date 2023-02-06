package sql

import "database/sql"

func GetSelectUser() *sql.Stmt {
	return selectUser
}
func GetInsertUser() *sql.Stmt {
	return insertUser
}
func GetChangePassword() *sql.Stmt {
	return changePassword
}
func GetInsertDetail() *sql.Stmt {
	return insertDetail
}
func GetSelectDetail() *sql.Stmt {
	return selectUserDetail
}
func GetUpdateDetail() *sql.Stmt {
	return updateDetail
}
func GetInsertLoveList() *sql.Stmt {
	return insertLoveList
}
func GetSelectLovedList() *sql.Stmt {
	return selectLovedList
}
func GetSelectLoveList() *sql.Stmt {
	return selectLoveList
}
func GetAddLove() *sql.Stmt {
	return addLove
}
func GetAddLoved() *sql.Stmt {
	return addLoved
}
func GetAddStar() *sql.Stmt {
	return addStar
}
func GetSelectUserStarList() *sql.Stmt {
	return selectUserStarList
}
func GetInsertUserStarList() *sql.Stmt {
	return insertUserStarList
}
func GetInsertArticle() *sql.Stmt {
	return insertArticle
}
func GetInsertTag() *sql.Stmt {
	return insertTag
}
func GetInsertArticleStar() *sql.Stmt {
	return insertArticleStar
}
func GetInsertArticleFavorite() *sql.Stmt {
	return insertArticleFavorite
}
func GetInsertArticleView() *sql.Stmt {
	return insertArticleView
}
func GetInsertArticleComment() *sql.Stmt {
	return insertArticleComment
}
func GetInsertArticleViews() *sql.Stmt {
	return insertArticleViews
}
func GetInsertArticleViewf() *sql.Stmt {
	return insertArticleViewf
}
func GetInsertArticleViewc() *sql.Stmt {
	return insertArticleViewc
}
func Getaid() *sql.Stmt {
	return getaid
}
func GetInsertArticleViewList() *sql.Stmt {
	return insertViewList
}
func GetArticleStarList() *sql.Stmt {
	return getArticleStarList
}
func GetArticleCommentList() *sql.Stmt {
	return getArticleCommentList
}
func GetArticleDetail() *sql.Stmt {
	return getArticleDetail
}
func GetUserStarList() *sql.Stmt {
	return getUserStarList
}
func GetUserFavoriteList() *sql.Stmt {
	return getUserFavoriteList
}
func GetUserViewList() *sql.Stmt {
	return getUserViewList
}
func GetArticleTime() *sql.Stmt {
	return getArticleTime
}
func GetArticleHot() *sql.Stmt {
	return getArticleHot
}
func GetArticleCommand() *sql.Stmt {
	return getArticleCommand
}
func GetDeleteViewRecord() *sql.Stmt {
	return deleteViewRecord
}
func GetDeleteStarRecord() *sql.Stmt {
	return deleteStarRecord
}
func GetDeleteFavoriteRecord() *sql.Stmt {
	return deleteFavoriteRecord
}

func GetSelectStarRecord() *sql.Stmt {
	return selectStarRecord
}
func GetSelectFavoriteRecord() *sql.Stmt {
	return selectFavoriteRecord
}
func GetUserStaredCountD1() *sql.Stmt {
	return userStaredCountD1
}
func GetUserStarCountD1() *sql.Stmt {
	return userStarCountD1
}
func GetArticleFavoriteCountD1() *sql.Stmt {
	return articleFavoriteCountD1
}
func GetArticleStarCountD1() *sql.Stmt {
	return articleStarCountD1
}
func GetDeleteUserStarRecord() *sql.Stmt {
	return deleteUserStarRecord
}
func GetRemoveArticleComment() *sql.Stmt {
	return removeArticleComment
}
func GetArticleCommentCountD1() *sql.Stmt {
	return articleCommentCountD1
}
func GetArticleForId() *sql.Stmt {
	return getArticleForId
}
func GetArticleTagList() *sql.Stmt {
	return getArticleTagList
}
func GetSelectUserStar() *sql.Stmt {
	return selectUserStar
}
func GetSelectUserFavorite() *sql.Stmt {
	return selectUserFavorite
}
func GetSelectViewRecord() *sql.Stmt {
	return selectViewRecord
}
func GetUserLoveList() *sql.Stmt {
	return getUserLoveList
}
func GetArticleById() *sql.Stmt {
	return getArticleById
}
func GetAidForTag() *sql.Stmt {
	return getAidForTag
}
func GetUserArticle() *sql.Stmt {
	return getUserArticle
}
func GetUserLoveUser() *sql.Stmt {
	return getUserLoveUser
}
