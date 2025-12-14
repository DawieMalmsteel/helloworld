câu 1 đéo được vì nó nhận vào là value chứ đéo phải pointer muốn sửa thì phải sửa thành
```go
func (s *Student) UpdateName(newName string) {
    s.Name = newName
}
```


câu 2 mày đã thêm `json ` tags vào struct lồn đâu mà đòi nó Marshal ra chuỗi json? với lại code mày viết cũng chưa export ra nữa nếu nó ở package khác thì mày có thêm tags json vào nó cũng đéo biết nữa 
sửa thành:
```go
type User struct {
    Username string `json:"username"`
    Password string `json:"password"`
}
```

câu 3 đơn giản là mày chỉ cần xem nó có implement đủ các methods của interface hay không thôi nếu chưa đủ thì lsp nó báo lỗi cho mày rồi bố mày code bằng nvim nên bớt coi thường lại 
- còn về phần interface{} thì nó là any type dùng nó khác lồn gì js đâu bởi vì bất kì kiểu dữ liệu nào cũng có thể implement nó được


vừa lòng mày chưa cho tao đi tắm cái
