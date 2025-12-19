vậy cái này là sao thằng lồn code 1 thì tao đặt pointer tới 1 cái struct Person thì chạy được, còn code 2 tao đặt value cho cái pointer chưa init tức là nil pointer á thì chạy đéo được giải ngố cho tao cái man:

```go
	// code 1: chạy được
	var person *Person
	person = &Person{Name: "Alice", Age: 30}

	// code 2: đéo chạy được
	var person *Person
	*person = Person{Name: "Alice", Age: 30}

```

**Đụ má, mày đang đâm đầu vào cái lỗi kinh điển nhất của mấy thằng mới chơi Pointer: "NIL POINTER DEREFERENCE" (Truy cập vào vùng nhớ rỗng).**

Nghe tao giải thích bằng cái ví dụ "Cái Nhà" tao vừa dạy mày lúc nãy là mày thông não ngay.

---

### CODE 1: "XÂY NHÀ TRƯỚC, LẤY CHÌA KHÓA SAU" (CHẠY NGON)

```go
var person *Person              // 1. Mày mua cái móc chìa khóa (Biến con trỏ), hiện tại chưa có chìa (nil).
person = &Person{Name: "Alice"} // 2. Mày XÂY CÁI NHÀ (Struct) -> Lấy địa chỉ (&) -> Móc vào chìa khóa.
```
*   **Giải thích:**
    *   `Person{...}`: Go nó cấp phát bộ nhớ (Xây nhà thật) cho con Alice.
    *   `&`: Lấy địa chỉ cái nhà đó.
    *   `person =`: Gán cái địa chỉ đó cho biến con trỏ.
    *   -> **Kết quả:** `person` đang giữ địa chỉ của một vùng nhớ có thật. Hợp lệ vãi lồn.

---

### CODE 2: "CỐ ĐẤM ĂN XÔI VÀO HƯ KHÔNG" (ĂN CỨT)

```go
var person *Person              // 1. Mày mua cái móc chìa khóa. Hiện tại nó là NIL (Rỗng, đéo trỏ đi đâu cả).
*person = Person{Name: "Alice"} // 2. Mày dùng dấu * (Phá cửa) vào cái hư không???
```
*   **Tại sao chết?**
    *   Dòng 1: `person` đang là `nil`. Tức là nó đéo có địa chỉ nhà nào cả. Nó đang trỏ vào "hư vô".
    *   Dòng 2: `*person` nghĩa là: *"Tao muốn PHÁ CỬA vào cái nhà mà biến person đang trỏ tới để nhét con Alice vào"*.
    *   **VẤN ĐỀ:** Làm đéo gì có cái nhà nào? `person` đang là `nil` mà? Mày đang cố phá cửa một căn nhà tàng hình?
    *   **Hậu quả:** Runtime Panic: `invalid memory address or nil pointer dereference`. Chương trình sập ngay lập tức.

---

### MỐI QUAN HỆ NHÂN QUẢ:
Mày muốn dùng `*person` (dereference) để ghi dữ liệu, thì bắt buộc cái `person` đó **PHẢI ĐANG TRỎ VÀO MỘT VÙNG NHỚ CÓ THẬT** trước đã.

Nếu mày muốn viết kiểu Code 2 mà vẫn chạy được, mày phải "cấp đất" cho nó trước bằng hàm `new()`:

```go
var person *Person
person = new(Person) // Cấp phát vùng nhớ rỗng cho person (Giờ nó đéo phải nil nữa, nó có nhà rồi nhưng nhà trống)

*person = Person{Name: "Alice", Age: 30} // Giờ thì phá cửa vào được rồi, vì nhà đã có.
```

**CHỐT LẠI:**
*   **Code 1:** Tạo nhà -> Lấy địa chỉ -> Gán cho con trỏ. (**Chuẩn**)
*   **Code 2:** Có con trỏ rỗng -> Đòi chui vào bên trong con trỏ rỗng để sửa. (**Ngu**)

Hiểu chưa thằng lồn? Đừng bao giờ chọc (`*`) vào một thằng đang là `nil`. Nó đấm cho sập chương trình đấy!
