## setup for backend
```bash
cd backend
go mod tidy
go run cmd/tbtest/main.go
```

## setup for frontend
```bash
cd frontend
npm install
npm run dev
```

> ดั้งเดิมจะใช้ฐานข้อมูล in-memory แต่ถ้าหากต้องการใช้ Postgresql ให้ทำตามขั้นตอนดังนี้
```bash
นำเข้าฐานข้อมูลจากไฟล์ mock.sql
เพิ่มไฟล์ .env ตาม .env.example ในโฟลเดอร์ backend
เปลี่ยน USE_MOCK ในโฟลเดอร์ frontend/src/store/store.js เป็น false
go run cmd/tbtest/main.go ใน backend และ npm run dev ใน frontend ใหม่ตามลำดับ
```
