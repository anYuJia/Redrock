// User 表示用户信息
type User struct {
ID       int    // 用户 ID
Username string // 用户名
Password string // 密码
}

// BorrowRecord 表示借阅记录
type BorrowRecord struct {
ID        int       // 借阅记录 ID
BookID    int       // 书籍 ID
UserID    int       // 用户 ID
BorrowedAt time.Time // 借阅时间
DueAt      time.Time // 应还时间
ReturnedAt time.Time // 归还时间
IsOverdue  bool      // 是否逾期
}

// Book 表示书籍信息
type Book struct {
ID         int       // 书籍 ID
Title      string    // 书名
Author     string    // 作者
ISBN       string    // ISBN 编号
Status     string    // 状态（可借、已借、预约等）
BorrowedBy int       // 借阅者 ID（如果有的话）
CreatedAt  time.Time // 入库时间
}
