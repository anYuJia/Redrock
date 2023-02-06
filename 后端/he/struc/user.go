package _struc

type User struct {
	Uid      int
	Username string
	Password string
}
type UserDetail struct {
	Uid          int
	Username     string
	Avatar       string
	Birthday     string
	Introduction string
	Qq           string
	Gender       int
	Email        string
	CreateTime   string
	Love         int
	Loved        int
	Stare        int
}
type Love struct {
	UsernameLove  string
	UserNameLoved string
}
type StarList struct {
	Id         int
	StarUser   string
	StaredUser string
	Time       string
}
type Article struct {
	Aid           int
	Title         string
	Content       string
	Author        string
	CreateTime    string
	StarCount     int
	FavoriteCount int
	ViewCount     int
	CommentCount  int
}
type ArticleComment struct {
	Acid    int
	Aid     int
	User    string
	Comment string
	Time    string
}
type XList struct {
	Axid int
	Aid  int
	User string
}
type ViewList struct {
	Avid int
	Aid  int
	User string
	Time string
}
type Bd struct {
	Art  Article
	Tags []string
}
type Bdt struct {
	Time string
	Art  Article
	Tags []string
}
