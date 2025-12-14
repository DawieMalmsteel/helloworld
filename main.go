package main

import (
	"fmt"
	"sync"
	"time"
)

// Đây là thằng đầu bếp
// id: Tên thằng bếp (1, 2, 3)
// jobs: Kênh nhận việc (chỉ được lấy ra <-)
// results: Kênh trả kết quả (chỉ được gửi vào <-)
func dauBep(id int, jobs <-chan string, results chan<- string, wg *sync.WaitGroup) {
	// Khi hàm này kết thúc (đầu bếp đi về), báo cho WaitGroup biết
	defer wg.Done()

	// Vòng lặp vô tận để lấy việc từ channel jobs
	// Cú pháp 'range' này rất hay: Nó tự động dừng khi channel jobs bị close và hết dữ liệu
	for monAn := range jobs {
		fmt.Printf("Bếp %d đang nấu %s...\n", id, monAn)

		// Giả vờ nấu mất 1 giây
		time.Sleep(1 * time.Second)

		// Nấu xong, gửi thông báo vào channel results
		results <- fmt.Sprintf("Món %s do bếp %d nấu xong!", monAn, id)
	}
}

func main() {
	jobs := make(chan string, 100)    // Giá treo order (Buffer to tí cho thoải mái)
	results := make(chan string, 100) // Cửa sổ trả món

	var wg sync.WaitGroup

	// 1. Tuyển 3 thằng đầu bếp vào làm việc
	for i := 1; i <= 3; i++ {
		wg.Add(1) // Điểm danh 1 thằng
		go dauBep(i, jobs, results, &wg)
	}

	// 2. Thằng Main (Phục vụ) ném 10 món ăn vào jobs channel
	danhSachMon := []string{"Cơm gà", "Phở", "Bún", "Miến", "Cháo", "Mì", "Nui", "Hủ tiếu", "Bánh canh", "Súp"}
	for _, mon := range danhSachMon {
		jobs <- mon
	}

	// QUAN TRỌNG: Ném xong phải ĐÓNG channel jobs.
	// Để bọn đầu bếp biết là "Hết order rồi, nấu nốt đống còn lại rồi về đi"
	close(jobs)

	// 3. Gom kết quả
	// Ở đây có 1 cái bẫy. Nếu mày range over results ngay đây, mày sẽ không biết bao giờ dừng.
	// Cách đơn giản nhất cho đêm nay:
	// Tạo 1 goroutine riêng để đợi bọn đầu bếp về hết rồi đóng channel results
	go func() {
		wg.Wait()      // Chờ 3 thằng bếp về hết
		close(results) // Đóng cửa sổ trả món
	}()

	// 4. In kết quả ra (Main đứng chờ ở cửa sổ trả món)
	for ketQua := range results {
		fmt.Println(">>>", ketQua)
	}

	fmt.Println("HẾT KHÁCH, ĐÓNG CỬA!")
}
