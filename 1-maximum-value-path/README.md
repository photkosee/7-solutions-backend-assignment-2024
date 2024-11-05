## 1. จงหาเส้นทางที่มีค่ามากที่สุด

![Figure 1-1](../files/max-path.png)

- ซึ่งจะอยู่ใน Array ดังต่อไปนี้ `[[59], [73, 41], [52, 40, 53], [26, 53, 6, 34]]`
- โดยเส้นทางที่มีค่ามากที่สุดจะเป็นตามจุดสีแดง
- แต่ละ node ห้ามย้อนกลับ (ต้องขึ้นลงเป็นทางเดียว) และเชื่อมกัน
- คำตอบให้อยู่ในรูปของ จำนวนรวมของเส้นทางที่ผ่าน ซึ่งจากตัวอย่างคือ `237`

ให้เขียนโปรแกรมภาษา GO โดยใช้ input จากไฟล์นี้ <https://github.com/7-solutions/backend-challenge/blob/main/files/hard.json> และแสดงผลเป็นค่าที่ได้จากการคำนวณ

# Test case

- input = `[[59], [73, 41], [52, 40, 53], [26, 53, 6, 34]]` output = `237`
- input = <https://github.com/7-solutions/backend-challenge/blob/main/files/hard.json> output = `7273`

### How to run
```
# ./1-maximum-value-path
go run ./main.go
```