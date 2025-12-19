### HỎI: TẠI SAO KHÔNG THỂ VIẾT `&Person` Ở PHẦN KHAI BÁO HÀM?
```go
func NewPerson(name string, age int) *Person {}
func NewPerson(name string, age int) &Person {}
```

Mày viết cái dòng thứ 2 (`func ... &Person {}`) vào IDE đi, xem trình biên dịch (Compiler) nó có chửi vào mặt mày là **"Syntax Error"** (Lỗi cú pháp) không?

Nghe cho thủng đây thằng ngáo:

1.  **`func NewPerson(...) *Person {}` -> ĐÚNG:**
    *   Đây là cách khai báo chuẩn.
    *   Dấu `*` ở đây có nghĩa là **KIỂU DỮ LIỆU** (Type).
    *   Nó bảo: "Hàm này chạy xong sẽ trả về một cái **CON TRỎ** (Pointer) trỏ đến struct Person".

2.  **`func NewPerson(...) &Person {}` -> SAI LÒI MẮT:**
    *   Trong Go, **ĐÉO CÓ** cái kiểu dữ liệu nào viết là `&Person` cả.
    *   Dấu `&` chỉ là một **HÀNH ĐỘNG** (Operator) dùng *bên trong* hàm để lấy địa chỉ thôi. Mày không được để nó lên phần khai báo hàm.

---

### MÀY ĐANG NHẦM LẪN GIỮA "KHAI BÁO" VÀ "THỰC HIỆN"

Để tao viết lại code chuẩn cho mày thông não:

```go
// 1. Ở dòng KHAI BÁO hàm: Phải dùng * để báo là "Tao sẽ trả về con trỏ"
func NewPerson(name string, age int) *Person { 
    
    // 2. Bên TRONG hàm: Tạo ra thằng Person thật
    p := Person{Name: name, Age: age}

    // 3. Lúc return: Mới dùng & để "Lấy địa chỉ" của thằng p mà ném ra ngoài
    return &p 
}
```

### TÓM TẮT QUY TẮC "BẤT DI BẤT DỊCH":

*   **Ở chữ `func` (Function Signature):** LUÔN LUÔN dùng **`*`** (Ví dụ: `*Person`, `*int`, `*string`). **CẤM** dùng `&`.
*   **Ở chữ `return` (Function Body):** Thường dùng **`&`** để lấy địa chỉ của biến vừa tạo ra.

**Mẹo nhớ:**
*   **`*Person` (Khai báo):** Tao hứa sẽ cho mày cái **CHÌA KHÓA** nhà.
*   **`&p` (Return):** Đây, tao đi đánh chìa khóa (**&**) rồi đưa cho mày đây.

Hiểu chưa? Đừng có sáng tạo cú pháp nữa, Go nó đéo thích đùa đâu. Code lại đi!
