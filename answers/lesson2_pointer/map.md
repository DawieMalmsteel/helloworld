```go
// Ví dụ về map nil pointer trong Go
package main

import "fmt"

func main() {
    var m map[string]int // Khai báo map nhưng chưa khởi tạo (nil map)
    // m == nil
    fmt.Println("Giá trị ban đầu của m:", m) // In ra: map[]
    fmt.Println("m có phải nil không?", m == nil) // In ra: true

    // Truy cập phần tử trên map nil KHÔNG gây panic, chỉ trả về giá trị zero của kiểu value
    fmt.Println("Truy cập key 'a' trên map nil:", m["a"]) // In ra: 0

    // Nhưng nếu cố gắng gán giá trị cho map nil sẽ gây panic
    // m["a"] = 1 // panic: assignment to entry in nil map

    // Để sử dụng map, cần khởi tạo với make
    m = make(map[string]int)
    m["a"] = 1
    fmt.Println("Sau khi khởi tạo và gán:", m) // In ra: map[a:1]

    // Hoặc khai báo map rỗng ngay từ đầu
    n := map[string]int{}
    n["b"] = 2
    fmt.Println("Map n sau khi gán:", n) // In ra: map[b:2]
}

/*
Giải thích:
- Map nil là map chưa được khởi tạo (giá trị mặc định là nil).
- Có thể đọc từ map nil (trả về zero value), nhưng không thể ghi vào map nil (sẽ panic).
- Luôn dùng make để khởi tạo map trước khi ghi dữ liệu.
*/
```
