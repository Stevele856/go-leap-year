## Yêu cầu
Nhập ngày tháng năm → Hiển thị thông tin sau :

- Thứ mấy trong tuần (Monday, Tuesday,...)
- Tuần thứ mấy trong năm (week 1-53)
- Ngày thứ mấy trong năm (day 1-366)
- Quý nào (Q1, Q2, Q3, Q4)
- Còn bao nhiêu ngày đến cuối năm


Tính khoảng cách giữa 2 ngày:

- Số ngày chênh lệch
- Số tuần, tháng


Tính ngày sau/trước N ngày:

- Nhập ngày + số ngày muốn cộng/trừ
- Hiển thị ngày kết quả


Kiểm tra năm nhuận
Hiển thị lịch tháng (dạng ASCII calendar)


## Step

1. Logic năm nhuận:
```go
- Chia hết cho 400 → nhuận
- Chia hết cho 100 → không nhuận
- Chia hết cho 4 → nhuận
- Còn lại → không
```

- ví dụ nhập 3/2/2026 thì ra thứ 3


=> Các bước 
- validate year
- validate month
- determine maxDay of that month
- validate day <= maxDay



====== // =========


=> Doomsday

# Doomsday algorithm

| Ngày  | Luôn cùng thứ |
| ----- | ------------- |
| 4/4   | ✅             |
| 6/6   | ✅             |
| 8/8   | ✅             |
| 10/10 | ✅             |
| 12/12 | ✅             |
| 9/5   | ✅             |
| 5/9   | ✅             |
| 7/11  | ✅             |
| 11/7  | ✅             |

👉 Nếu biết Doomsday của năm đó là thứ mấy, bạn có thể suy ra mọi ngày khác bằng cách cộng / trừ.

# Ý tưởng tổng thể

- Tính Doomdays của năm
- Lấy mốc Doomdays của tháng
- Tính chênh lệch ngày
- Suy ra weekday

# Century Anchor day
- 0: Sunday
- 1: Monday
- 2: Tuesday
- 3: Wenesday
- 4: Thursday
- 5: Friday
- 6: Saturday

- Hàm `centuryAnchor()` nhận vào một năm và tính toán "anchor day" (ngày neo) cho thế kỷ đó. Đây là ngày trong tuần tương ứng với một số ngày đặc biệt trong năm (Ví dụ: 4/4, 6/6, 8/8, 10/10, 12/12)

```go
func centuryAnchor(year int) int {
    century := year / 100
    switch century % 4 {
    case 0:
        return 2
    case 1: 
        return 0
    case 2:
        return 5
    default:
        return 3
    }
}
```

1. Cách hoạt động: 
+ Tại sao chia 100: Để biết thể kỷ nào
+ Tại sao % 4: để biết vị trí trong chu kỳ 4 thế kỷ

Từ thuật toán Doomsday của John Conway. Ông đã tính toán và chứng minh rằng:

+ Các thế kỷ có century % 4 = 0 (như 1600, 2000, 2400) → anchor day rơi vào Thứ Ba
+ Các thế kỷ có century % 4 = 1 (như 1700, 2100) → anchor day rơi vào Chủ Nhật
+ Các thế kỷ có century % 4 = 2 (như 1800, 2200) → anchor day rơi vào Thứ Sáu
+ Các thế kỷ có century % 4 = 3 (như 1900, 2300) → anchor day rơi vào Thứ Tư

- Hàm DoomsDay `DoomdayOfYear()` tính DoomsDay của một năm cụ thể tức là ngày trong tuần mà các ngày đặc biệt 4/4,6/6, 8/8, 10/10, 12/12 rơi vào. 

```go
func DoomdaysOfYear(year int) int {
    y := year % 100 // lấy  2 chữ số cuối của năm 
    a := y / 12 // Số lần chia hết cho 12
    b := y % 12 // Phần dư khi chia cho 12
    c := b / 4 // Số năm nhuận trong phần dư
    
    d := a + b + c
    
    return (d + centuryAnchor(year)) % 7 // Cộng với anchor day của thế kỷ, rồi lấy mod 7
} 
```

