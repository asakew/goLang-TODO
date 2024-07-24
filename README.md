# goLang-TODO
This is a TODO list for goLang projects.
- goFiber + Postgres + GORM
  Todo app + goFiber + Postgres + GORM + index.html get: add, delete, update
- 
## Demo Screenshot
![github](/public/img/screenshot.png)

### Start Commands
```bash
git clone https://github.com/asakew/goLang-TODO.git
go get github.com/gofiber/fiber/v2
go get gorm.io/driver/postgres
go get gorm.io/gorm
go run main.go
```
running server:  http://127.0.0.1:3000/

#### DB myConnection:
- dsn := "host=localhost user=postgres password=superUser7 dbname=marina port=5432 sslmode=disable TimeZone=Asia/Shanghai"

## Source Code
- web framework: goFiber http://goFiber.oi/
- css framework: mdbootstrap: https://mdbootstrap.com/
