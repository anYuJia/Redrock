# 接口文档--py

### 注册(get)

```url
http://ananqiexiafan.icu:9999/register
```

#### 请求体(参数1：username，参数2：password)

```json
username : username
password : password
```

请求示例

```纯文本
http://ananqiexiafan.icu:9999/register?username=11111111&password=2222222
```

#### 返回示例1：



```json
{
    "code": 200,
    "data": {},
    "msg": "注册成功"
}
```

### 登录(get)

```url
http://ananqiexiafan.icu:9999/login
```

#### 请求体(参数1：username，参数2：password)

```json
username : username
password : password
```

请求示例

```纯文本
http://ananqiexiafan.icu:9999/login?username=11111111&password=2222222
```

返回示例1：

```json
{
    "code": 200,
    "data": {
        "cookie": "NotSet"
    },
    "msg": "登陆成功"
}
```

### 修改密码(get)

```url
http://ananqiexiafan.icu:9999/updatePassword
```

注：用户在登录之后才能进行改密码，另外注销登录之后也不能改密码，每次登录时间持续1小时，一小时后需要重新登陆

#### 请求体(参数1：username，参数2：password)

```json
username : username
password : password
```

请求示例

```纯文本
http://ananqiexiafan.icu:9999/updatePassword?username=11111111&password=2222222
```

返回示例1：登录状态下

```json
{
    "code": 200,
    "data": {},
    "msg": "修改密码成功"
}
```

返回示例2：未登录或者注销登录或登陆超时

```json
{
    "code": 502,
    "data": {},
    "msg": "未登录不能执行此操作"
}
```

### 退出登录(get)

```url
http://ananqiexiafan.icu:9999/logout
```

注：用户在登录之后才能退出。

#### 请求体(参数：无)

请求示例：

```纯文本
http://ananqiexiafan.icu:9999/logout
```

返回示例1：登录状态下

```json
{
    "code": 200,
    "data": {},
    "msg": "退出登录成功!"
}
```

返回示例2：未登录或者注销登录或登陆超时

```json
{
    "code": 502,
    "data": {},
    "msg": "未登录不能执行此操作"
}
```

### 获取用户详情信息(get)

```url
http://ananqiexiafan.icu:9999/getDetail
```

注：用户在登录之后才能获取。

#### 请求体(参数：无)

请求示例：

```纯文本
http://ananqiexiafan.icu:9999/getDetail
```

返回示例1：登录状态下

```json
{
    "code": 200,
    "data": {
        "data": {
            "Uid": 3,
            "Username": "888d8d8",
            "Avatar": "http://ananqiexiafan.icu/d/tou.png",
            "Birthday": "2022-02-21",
            "Introduction": "n",
            "Qq": "n",
            "Gender": 2,
            "Email": "n",
            "CreateTime": "2023-01-30",
            "Love": 0,
            "Loved": 0,
            "Stare": 0
        }
    },
    "msg": "ok"
}
```

返回示例2：未登录或者注销登录或登陆超时

```json
{
    "code": 502,
    "data": {},
    "msg": "未登录不能执行此操作"
}
```

### 更改用户详情信息(get)

```url
http://ananqiexiafan.icu:9999/updateDetail
```

注：用户在登录之后才能更改。

#### 请求体(参数：qq, gender, introduction, email, birthday)

请求示例：

```纯文本
http://ananqiexiafan.icu:9999/updateDetail?qq=1233333&gender=1&introduction=我是猪gg&email=2259@qq.com&birthday=2020-01-02
```

返回示例1：登录状态下

```json
{
    "code": 200,
    "data": {},
    "msg": "修改成功"
}
```

返回示例2：未登录或者注销登录或登陆超时

```json
{
    "code": 502,
    "data": {},
    "msg": "未登录不能执行此操作"
}
```

### 关注用户(get)

```url
http://ananqiexiafan.icu:9999/love
```

注：用户在登录之后才能更改。

#### 请求体(参数：lovedUser)

请求示例：

```纯文本
http://ananqiexiafan.icu:9999/love?lovedUser=12344s
```

返回示例0：登录状态下已经关注了

```json
{
    "code": 400,
    "data": {},
    "msg": "已经关注了该用户12344s"
}
```

返回示例1：登录状态下关注成功

```json
{
    "code": 200,
    "data": {},
    "msg": "关注12344s成功"
}
```

返回示例2:用户不存在

```clojure
{
    "code": 400,
    "data": {},
    "msg": "该用户不存在"
}
```

返回示例3：未登录或者注销登录或登陆超时

```json
{
    "code": 502,
    "data": {},
    "msg": "未登录不能执行此操作"
}
```



### 获取关注列表(get)

```url
http://ananqiexiafan.icu:9999/getLoveList
```

注：用户在登录可执行。

#### 请求体(参数：无)

请求示例：

```纯文本
http://ananqiexiafan.icu:9999/getLoveList
```

返回示例0：登录状态下获取成功

```json
{
    "code": 200,
    "data": {
        "data": [
            {
                "UsernameLove": "777",
                "UserNameLoved": "1234444"
            },
            {
                "UsernameLove": "777",
                "UserNameLoved": "1234444"
            },
            {
                "UsernameLove": "777",
                "UserNameLoved": "123444"
            },
            {
                "UsernameLove": "777",
                "UserNameLoved": "1234"
            },
            {
                "UsernameLove": "777",
                "UserNameLoved": "12344s"
            }
        ]
    },
    "msg": "ok"
}
```

返回示例3：未登录或者注销登录或登陆超时

```json
{
    "code": 502,
    "data": {},
    "msg": "未登录不能执行此操作"
}
```

### 获取被关注列表(get)

```url
http://ananqiexiafan.icu:9999/getLovedList
```

注：用户在登录可执行。

#### 请求体(参数：无)

请求示例：

```纯文本
http://ananqiexiafan.icu:9999/getLoveList
```

返回示例0：登录状态下获取成功

```json
{
    "code": 200,
    "data": {
        "data": null
    },
    "msg": "ok"
}
```

返回示例3：未登录或者注销登录或登陆超时

```json
{
    "code": 502,
    "data": {},
    "msg": "未登录不能执行此操作"
}
```



### 点赞用户(get)

```url
http://ananqiexiafan.icu:9999/userStar
```

注：用户在登录可执行。

#### 请求体(参数：staredUser)

请求示例：

```纯文本
http://ananqiexiafan.icu:9999/userStar?staredUser=12344s
```

返回示例0：登录状态下已经今日已经点赞了

```json
{
    "code": 400,
    "data": {},
    "msg": "今天该用户已经点过赞了，不能重复点赞"
}
```

返回示例1：登录状态下点赞成功

```json
{
    "code": 200,
    "data": {},
    "msg": "您成功给用户aaaaaa点赞"
}

```

返回示例2:用户不存在

```clojure
{
    "code": 400,
    "data": {},
    "msg": "该用户不存在"
}
```

返回示例3：未登录或者注销登录或登陆超时

```json
{
    "code": 502,
    "data": {},
    "msg": "未登录不能执行此操作"
}
```

### 发布文章(get)

```url
http://ananqiexiafan.icu:9999/insertArticle
```

注：用户在登录可执行。

#### 请求体(参数：title content tags)

请求示例：

```纯文本
http://ananqiexiafan.icu:9999/insertArticle?title=这是文章标题&content=这是文章的内容，字数不大于3000&tags=前端 后端 技巧 学习
```

返回示例0：登录状态下已经今日已经点赞了

```json
{
    "code": 200,
    "data": {
        "articleId": 1
    },
    "msg": "发表文章成功！"
}
```

返回示例1：未登录或者注销登录或登陆超时

```json
{
    "code": 502,
    "data": {},
    "msg": "未登录不能执行此操作"
}
```

### 发布文章(get)

```url
http://ananqiexiafan.icu:9999/insertArticle
```

注：用户在登录可执行。

#### 请求体(参数：title content tags)（文章标题 文章内容 文章标签\[多个标签空格隔开]）

请求示例：

```纯文本
http://ananqiexiafan.icu:9999/insertArticle?title=这是文章标题&content=这是文章的内容，字数不大于3000&tags=前端 后端 技巧 学习
```

返回示例0：ok

```json
{
    "code": 200,
    "data": {
        "articleId": 1
    },
    "msg": "发表文章成功！"
}
```

返回示例1：未登录或者注销登录或登陆超时

```json
{
    "code": 502,
    "data": {},
    "msg": "未登录不能执行此操作"
}
```

### 发布文章评论(get)

```url
http://ananqiexiafan.icu:9999/insertArticleComment
```

注：用户在登录可执行。

#### 请求体(参数：title content)（文章标题 发表的评论的内容）

请求示例：

```纯文本
http://ananqiexiafan.icu:9999/insertArticleComment?title=这是文章标题&content=这是这篇文章的第一条评论
```

返回示例0：ok

```json
{
    "code": 200,
    "data": {},
    "msg": "评论成功!"
}
```

返回示例1：未登录或者注销登录或登陆超时

```json
{
    "code": 502,
    "data": {},
    "msg": "未登录不能执行此操作"
}
```

### 点赞文章(get)

```url
http://ananqiexiafan.icu:9999/insertArticleStar
```

注：用户在登录可执行。

#### 请求体(参数：title)（文章标题）

请求示例：

```纯文本
http://ananqiexiafan.icu:9999/insertArticleStar?title=这是文章标题
```

返回示例0：ok

```json
{
    "code": 200,
    "data": {},
    "msg": "点赞成功"
}
```

返回示例1：未登录或者注销登录或登陆超时

```json
{
    "code": 502,
    "data": {},
    "msg": "未登录不能执行此操作"
}
```

### 收藏文章(get)

```url
http://ananqiexiafan.icu:9999/insertArticleFavorite
```

注：用户在登录可执行。

#### 请求体(参数：title )（文章标题）

请求示例：

```纯文本
http://ananqiexiafan.icu:9999/insertArticleFavorite?title=这是文章标题
```

返回示例0：登录状态下已经今日已经点赞了

```json
{
    "code": 200,
    "data": {},
    "msg": "收藏成功"
}
```

返回示例1：未登录或者注销登录或登陆超时

```json
{
    "code": 502,
    "data": {},
    "msg": "未登录不能执行此操作"
}
```

### 浏览文章(get)

```url
http://ananqiexiafan.icu:9999/insertArticleView
```

注：用户在登录可执行。

#### 请求体(参数：title)（文章标题）

请求示例：

```纯文本
http://ananqiexiafan.icu:9999/insertArticleView?title=这是文章标题
```

返回示例0：登录状态下已经今日已经点赞了

```json
{
    "code": 200,
    "data": {},
    "msg": "浏览成功"
}
```

返回示例1：未登录或者注销登录或登陆超时

```json
{
    "code": 502,
    "data": {},
    "msg": "未登录不能执行此操作"
}
```

### 获取文章信息(get)

```url
http://ananqiexiafan.icu:9999/GetArticle
```

注：用户在登录可执行。

#### 请求体(参数：articleId)(文章id)

请求示例：

```纯文本
http://ananqiexiafan.icu:9999/GetArticle?articleId=1
```

返回示例0：登录状态下已经今日已经点赞了

```json
{
    "code": 200,
    "data": {
        "commentList": [
            {
                "Acid": 1,
                "Aid": 1,
                "User": "11111111",
                "Comment": "这是这篇文章的第一条评论",
                "Time": "2023-02-03 22:23:13"
            }
        ],
        "detail": {
            "Aid": 1,
            "Title": "这是文章标题",
            "Content": "这是文章的内容，字数不大于3000",
            "Author": "11111111",
            "CreateTime": "2023-02-03 22:18:58",
            "StarCount": 1,
            "FavoriteCount": 1,
            "ViewCount": 1,
            "CommentCount": 1
        },
        "tags": [
            "前端",
            "后端",
            "技巧",
            "学习"
        ]
    },
    "msg": "获取成功!"
}
```

返回示例1：未登录或者注销登录或登陆超时

```json
{
    "code": 502,
    "data": {},
    "msg": "未登录不能执行此操作"
}
```

### 获取用户收藏列表（简略，别人）(get)

```url
http://ananqiexiafan.icu:9999/userFavoriteList
```

注：用户在登录可执行。

#### 请求体(参数：username)（用户名）

请求示例：

```纯文本
http://ananqiexiafan.icu:9999/userFavoriteList?username=11111111
```

返回示例0：ok

```json
{
    "code": 200,
    "data": {
        "data": [
            {
                "Axid": 1,
                "Aid": 1,
                "User": "11111111"
            }
        ]
    },
    "msg": "获取用户收藏列表成功"
}
```

返回示例1：未登录或者注销登录或登陆超时

```json
{
    "code": 502,
    "data": {},
    "msg": "未登录不能执行此操作"
}
```

### 获取用户浏览列表(get)

```url
http://ananqiexiafan.icu:9999/userViewList
```

注：用户在登录可执行。

#### 请求体(参数：username)(用户名)

请求示例：

```纯文本
http://ananqiexiafan.icu:9999/userViewList?username=11111111
```

返回示例0：ok

```json
{
    "code": 200,
    "data": {
        "data": [
            {
                "Avid": 1,
                "Aid": 1,
                "User": "11111111",
                "Time": "2023-02-03 22:29:14"
            }
        ]
    },
    "msg": "获取用户浏览列表成功"
}
```

返回示例1：未登录或者注销登录或登陆超时

```json
{
    "code": 502,
    "data": {},
    "msg": "未登录不能执行此操作"
}
```

### 获取用户点赞列表(get)

```url
http://ananqiexiafan.icu:9999/userStarList
```

注：用户在登录可执行。

#### 请求体(参数：username)(用户名)

请求示例：

```纯文本
http://ananqiexiafan.icu:9999/userStarList?username=11111111
```

返回示例0：登录状态下已经今日已经点赞了

```json
{
    "code": 200,
    "data": {
        "data": [
            {
                "Axid": 1,
                "Aid": 1,
                "User": "11111111"
            }
        ]
    },
    "msg": "获取用户点赞列表成功"
}
```

返回示例1：未登录或者注销登录或登陆超时

```json
{
    "code": 502,
    "data": {},
    "msg": "未登录不能执行此操作"
}
```

### 获取热榜(get)

```url
http://ananqiexiafan.icu:9999/getHot
```

注：用户在登录可执行。

#### 请求体(参数：无)

请求示例：

```纯文本
http://ananqiexiafan.icu:9999/getHot
```

返回示例0：ok

```json
{
    "code": 200,
    "data": {
        "data": [
            {
                "Art": {
                    "Aid": 1,
                    "Title": "这是文章标题",
                    "Content": "这是文章的内容，字数不大于3000",
                    "Author": "11111111",
                    "CreateTime": "2023-02-03 22:18:58",
                    "StarCount": 1,
                    "FavoriteCount": 1,
                    "ViewCount": 1,
                    "CommentCount": 1
                },
                "Tags": [
                    "前端",
                    "后端",
                    "技巧",
                    "学习"
                ]
            }
        ]
    },
    "msg": "获取热榜成功"
}
```

返回示例1：未登录或者注销登录或登陆超时

```json
{
    "code": 502,
    "data": {},
    "msg": "未登录不能执行此操作"
}
```

### 获取最新榜(get)

```url
http://ananqiexiafan.icu:9999/GetArticleTimeA
```

注：用户在登录可执行。

#### 请求体(参数：无)

请求示例：

```纯文本
http://ananqiexiafan.icu:9999/GetArticleTimeA
```

返回示例0：登录状态下已经今日已经点赞了

```json
{
    "code": 200,
    "data": {
        "data": [
            {
                "Art": {
                    "Aid": 1,
                    "Title": "这是文章标题",
                    "Content": "这是文章的内容，字数不大于3000",
                    "Author": "11111111",
                    "CreateTime": "2023-02-03 22:18:58",
                    "StarCount": 1,
                    "FavoriteCount": 1,
                    "ViewCount": 1,
                    "CommentCount": 1
                },
                "Tags": [
                    "前端",
                    "后端",
                    "技巧",
                    "学习"
                ]
            }
        ]
    },
    "msg": "获取最新榜成功"
}
```

返回示例1：未登录或者注销登录或登陆超时

```json
{
    "code": 502,
    "data": {},
    "msg": "未登录不能执行此操作"
}
```

### 获取推荐(get)

```url
http://ananqiexiafan.icu:9999/GetCommand
```

注：用户在登录可执行。

#### 请求体(参数：无)

请求示例：

```纯文本
http://ananqiexiafan.icu:9999/GetCommand
```

返回示例0：登录状态下已经今日已经点赞了

```json
{
    "code": 200,
    "data": {
        "data": [
            {
                "Art": {
                    "Aid": 1,
                    "Title": "这是文章标题",
                    "Content": "这是文章的内容，字数不大于3000",
                    "Author": "11111111",
                    "CreateTime": "2023-02-03 22:18:58",
                    "StarCount": 1,
                    "FavoriteCount": 1,
                    "ViewCount": 1,
                    "CommentCount": 1
                },
                "Tags": [
                    "前端",
                    "后端",
                    "技巧",
                    "学习"
                ]
            }
        ]
    },
    "msg": "获取推荐榜成功"
}
```

返回示例1：未登录或者注销登录或登陆超时

```json
{
    "code": 502,
    "data": {},
    "msg": "未登录不能执行此操作"
}
```

### 用户取消收藏文章(get)

```url
http://ananqiexiafan.icu:9999/CancerArticleFavorite
```

注：用户在登录可执行。

#### 请求体(参数：aid)（文章id）

请求示例：

```纯文本
http://ananqiexiafan.icu:9999/CancerArticleFavorite?aid=1
```

返回示例0：登录状态下已经今日已经点赞了

```json
{
    "code": 200,
    "data": {},
    "msg": "取消收藏成功!"
}
```

返回示例1：未登录或者注销登录或登陆超时

```json
{
    "code": 502,
    "data": {},
    "msg": "未登录不能执行此操作"
}
```

### 删除文章评论(get)

```url
http://ananqiexiafan.icu:9999/RemoveArticleComment
```

注：用户在登录可执行。

#### 请求体(参数：acid)(评论id)

请求示例：

```纯文本
http://ananqiexiafan.icu:9999/RemoveArticleComment?acid=1
```

返回示例0：登录状态下已经今日已经点赞了

```json
{
    "code": 200,
    "data": {},
    "msg": "删除评论成功!"
}
```

返回示例1：未登录或者注销登录或登陆超时

```json
{
    "code": 502,
    "data": {},
    "msg": "未登录不能执行此操作"
}
```

### 用户取消对文章点赞(get)

```url
http://ananqiexiafan.icu:9999/CancerArticleStar
```

注：用户在登录可执行。

#### 请求体(参数：aid)(文章id)

请求示例：

```纯文本
http://ananqiexiafan.icu:9999/CancerArticleStar?aid=1
```

返回示例0：登录状态下已经今日已经点赞了

```json
{
    "code": 200,
    "data": {},
    "msg": "取消点赞成功!"
}
```

返回示例1：未登录或者注销登录或登陆超时

```json
{
    "code": 502,
    "data": {},
    "msg": "未登录不能执行此操作"
}
```

### 取消关注用户(get)

```url
http://ananqiexiafan.icu:9999/CancerStarUser
```

注：用户在登录可执行。

#### 请求体(参数：staredUser)(被取消关注的用户)

请求示例：

```纯文本
http://ananqiexiafan.icu:9999/CancerStarUser?staredUser=
```

返回示例0：登录状态下已经今日已经点赞了

```json
{
    "code": 200,
    "data": {},
    "msg": "你成功取消关注111"
}
```

返回示例1：未登录或者注销登录或登陆超时

```json
{
    "code": 502,
    "data": {},
    "msg": "未登录不能执行此操作"
}
```

### 获取自己的收藏列表(get)

```url
http://ananqiexiafan.icu:9999/GetUserFavoriteRecord
```

注：用户在登录可执行。

#### 请求体(参数：无)

请求示例：

```纯文本
http://ananqiexiafan.icu:9999/GetUserFavoriteRecord
```

返回示例0：

```json
{
    "code": 400,
    "data": {},
    "msg": "用户收藏列表为空!"
}
```

返回示例1：未登录或者注销登录或登陆超时

```json
{
    "code": 502,
    "data": {},
    "msg": "未登录不能执行此操作"
}
```

### 获取自己的点赞列表(get)

```url
http://ananqiexiafan.icu:9999/GetUserStarRecord
```

注：用户在登录可执行。

#### 请求体(参数：无)

请求示例：

```纯文本
http://ananqiexiafan.icu:9999/GetUserStarRecord
```

返回示例0：

```json
{
    "code": 400,
    "data": {},
    "msg": "用户点赞列表为空!"
}
```

返回示例1：未登录或者注销登录或登陆超时

```json
{
    "code": 502,
    "data": {},
    "msg": "未登录不能执行此操作"
}
```

### 获取自己的浏览列表(get)

```url
http://ananqiexiafan.icu:9999/GetUserViewRecord
```

注：用户在登录可执行。

#### 请求体(参数：无)

请求示例：

```纯文本
http://ananqiexiafan.icu:9999/GetUserViewRecord
```

返回示例0：

```json
{
    "code": 200,
    "data": {
        "data": [
            {
                "Time": "2023-02-03 22:29:14",
                "Art": {
                    "Aid": 1,
                    "Title": "这是文章标题",
                    "Content": "这是文章的内容，字数不大于3000",
                    "Author": "11111111",
                    "CreateTime": "2023-02-03 22:18:58",
                    "StarCount": 0,
                    "FavoriteCount": 0,
                    "ViewCount": 1,
                    "CommentCount": 0
                },
                "Tags": [
                    "前端",
                    "后端",
                    "技巧",
                    "学习"
                ]
            }
        ]
    },
    "msg": "获取用户浏览列表成功!"
}
```

返回示例1：未登录或者注销登录或登陆超时

```json
{
    "code": 502,
    "data": {},
    "msg": "未登录不能执行此操作"
}
```

### 获取标签列表(get)

```url
http://ananqiexiafan.icu:9999/GetTagList
```

注：用户在登录可执行。

#### 请求体(参数：tag)(标签名)

请求示例：

```纯文本
http://ananqiexiafan.icu:9999/GetTagList?tag=前端
```

返回示例0：

```json
{
    "code": 200,
    "data": {
        "data": [
            {
                "Art": {
                    "Aid": 1,
                    "Title": "这是文章标题",
                    "Content": "这是文章的内容，字数不大于3000",
                    "Author": "11111111",
                    "CreateTime": "2023-02-03 22:18:58",
                    "StarCount": 0,
                    "FavoriteCount": 0,
                    "ViewCount": 1,
                    "CommentCount": 0
                },
                "Tags": [
                    "前端",
                    "后端",
                    "技巧",
                    "学习"
                ]
            }
        ]
    },
    "msg": "获取成功标签页"
}
```

返回示例1：未登录或者注销登录或登陆超时

```json
{
    "code": 502,
    "data": {},
    "msg": "未登录不能执行此操作"
}
```

### 获取关注用户列表(get)

```url
http://ananqiexiafan.icu:9999/GetUserArticle
```

注：用户在登录可执行。

#### 请求体(参数：无)

请求示例：

```纯文本
http://ananqiexiafan.icu:9999/GetUserArticle
```

返回示例0：

```json
{
    "code": 400,
    "data": {},
    "msg": "获取关注为空!"
}
```

返回示例1：未登录或者注销登录或登陆超时

```json
{
    "code": 502,
    "data": {},
    "msg": "未登录不能执行此操作"
}
```

### 获取用户文章列表(get)

```url
http://ananqiexiafan.icu:9999/GetUserArticleList
```

注：用户在登录可执行。

#### 请求体(参数：username)(用户名)

请求示例：

```纯文本
http://ananqiexiafan.icu:9999/GetUserArticleList?username=777
```

返回示例0：

```json
{
    "code": 200,
    "data": {
        "data": [
            {
                "Art": {
                    "Aid": 1,
                    "Title": "aaa但是",
                    "Content": "dasd",
                    "Author": "777",
                    "CreateTime": "2023-02-01 10:09:40",
                    "StarCount": 5,
                    "FavoriteCount": 0,
                    "ViewCount": 1,
                    "CommentCount": 2
                },
                "Tags": [
                    "后端",
                    "运维",
                    "安全",
                    "开发"
                ]
            }
        ]
    },
    "msg": "获取用户文章列表成功"
}
```

返回示例1：未登录或者注销登录或登陆超时

```json
{
    "code": 502,
    "data": {},
    "msg": "未登录不能执行此操作"
}
```
