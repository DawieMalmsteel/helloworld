**Lú vì syntax quá khó đọc**

---
**Lú cái lồn gì? Bình tĩnh! Vấn đề chỉ quay quanh SEND và RECEIVE**

Mày lú là do mày đang cố nhớ vẹt. Cái cú pháp này nhìn thì như ma trận nhưng quy tắc của nó **SIÊU ĐƠN GIẢN**.

Tao chỉ cho mày cái **MẸO NHÌN HÌNH** này, bảo đảm mày đéo bao giờ quên được nữa.

Nguyên tắc vàng: **MŨI TÊN (`<-`) NÓ CHỈ HƯỚNG ĐI CỦA DỮ LIỆU.**

Coi chữ `chan` là cái **ỐNG NƯỚC**.

---

### 1. `chan<-` (Mũi tên chọc VÀO đít chữ chan)
*   **Hình ảnh:** `Data` -> `chan`
*   **Ý nghĩa:** Dữ liệu ĐI VÀO ống.
*   **Quyền hạn:** **CHỈ ĐƯỢC GỬI (SEND ONLY)**. Mày chỉ được ném đồ vào, cấm thò tay lấy ra.
*   **Ví dụ:** Thằng Phục Vụ. Nó chỉ được ném order vào bếp.
    ```go
    // jobs chan<- string
    jobs <- "Cơm gà" // OK: Ném cơm gà VÀO channel
    a := <-jobs      // LỖI NGAY: Đéo được lấy ra!
    ```

### 2. `<-chan` (Mũi tên chui RA KHỎI đầu chữ chan)
*   **Hình ảnh:** `chan` -> `Data`
*   **Ý nghĩa:** Dữ liệu ĐI RA từ ống.
*   **Quyền hạn:** **CHỈ ĐƯỢC NHẬN (RECEIVE ONLY)**. Mày chỉ được hứng đồ rơi ra, cấm nhét ngược vào.
*   **Ví dụ:** Thằng Đầu Bếp. Nó chỉ được lấy order ra để nấu.
    ```go
    // jobs <-chan string
    monAn := <-jobs  // OK: Hứng món ăn TỪ channel ra
    jobs <- "Phở"    // LỖI NGAY: Đéo được nhét ngược vào!
    ```

### 3. `chan` (Đéo có mũi tên nào)
*   **Ý nghĩa:** **LƯỠNG TÍNH (BIDIRECTIONAL)**. Thích ra thì ra, thích vào thì vào.
*   Thường dùng khi mày `make` channel ở hàm `main`. Sau đó khi truyền vào hàm con, Go sẽ tự ép kiểu về 2 loại trên để bảo mật.

---

**TẠI SAO PHẢI PHÂN RA LẰNG NHẰNG THẾ?**
Để **CHỐNG NGU**.

Ví dụ hàm `dauBep`:
```go
func dauBep(jobs <-chan string) { ... }
```
Nếu trong hàm này mày lỡ tay code `jobs <- "Cơm"`, trình biên dịch (Compiler) sẽ vả vào mồm mày ngay: *"Ê thằng ngu, mày là đầu bếp, mày chỉ được lấy order thôi, ai cho mày tạo order?"*.

Nó giúp code mày an toàn, thằng nào làm việc nấy, không dẫm chân nhau.

**TÓM LẠI:**
*   **`chan<-`**: Dữ liệu **VÀO** (Gửi).
*   **`<-chan`**: Dữ liệu **RA** (Nhận).
*   **Mẹo nhớ:** Nhìn cái mũi tên so với chữ `chan`.

Đã thông cái não chưa? Giờ quay lại code tiếp đi, đừng để cái mũi tên dọa mày!
