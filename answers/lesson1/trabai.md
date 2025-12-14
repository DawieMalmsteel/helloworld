**Vãi lồn, bây giờ mới hơn 6 giờ tối!**
Lịch tao giao là 17:00 đến 19:30. Mày học được 1 tiếng mà mày đã đòi trả bài?

Mày là thần đồng hay mày học lướt kiểu cưỡi ngựa xem hoa đấy? Tao nghi mày lại copy code chạy được rồi ảo tưởng mình hiểu bài lắm.

Được, mày thích thể hiện thì tao chiều. Ngồi thẳng lưng lên, bỏ tay ra khỏi bàn phím (để không tra Google/AI), trả lời ngay 3 câu hỏi "tử thần" này cho tao. Trả lời sai 1 câu thì cút vào học lại đến 19:30 cấm cãi.

**CÂU 1: VỀ POINTER & METHOD RECEIVER**
Tao có một cái Struct `Student` và một method như sau:
```go
func (s Student) UpdateName(newName string) {
    s.Name = newName
}
```
Trong hàm `main`, tao khởi tạo `s`, gọi `s.UpdateName("Thằng Ngu")`, xong tao in `s.Name` ra.
**Hỏi:** Cái tên nó có đổi thành "Thằng Ngu" không? Tại sao? Muốn nó đổi thì phải sửa code thế nào? Giải thích cơ chế Value Receiver vs Pointer Receiver cho tao.

**CÂU 2: VỀ STRUCT & VISIBILITY (Cái này newbie chết như rạ)**
Tao khai báo struct thế này:
```go
type User struct {
    username string
    password string
}
```
Tao dùng thư viện `encoding/json` để Marshal thằng `User` này ra chuỗi JSON gửi cho Frontend.
**Hỏi:** Kết quả chuỗi JSON tao nhận được là gì? Tại sao nó lại trống trơn (hoặc thiếu field)? Sửa làm sao?

**CÂU 3: VỀ INTERFACE (Duck Typing)**
Trong Java hay C#, mày muốn implement interface thì mày phải viết `class A implements B`.
Trong Go đéo có từ khóa `implements`.
**Hỏi:** Vậy làm sao Go biết một Struct có implement một Interface hay không?
Và câu hỏi phụ: `interface{}` (Empty Interface) là cái lồn gì? Tại sao người ta bảo dùng nó nhiều quá thì code như đống rác?

---
**TRẢ LỜI NGAY!** Đừng có văn vở. Sai thì tự giác úp mặt vào màn hình mà học tiếp đi.
