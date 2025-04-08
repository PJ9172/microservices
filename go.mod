module Microservices

go 1.23.0

toolchain go1.23.8

require (
	github.com/gorilla/mux v1.8.1
	github.com/joho/godotenv v1.5.1
	github.com/jordan-wright/email v4.0.1-0.20210109023952-943e75fe5223+incompatible
	gorm.io/driver/mysql v1.5.7
	gorm.io/gorm v1.25.12
)

require github.com/gorilla/securecookie v1.1.2 // indirect

require (
	github.com/go-sql-driver/mysql v1.7.0 // indirect
	github.com/gorilla/sessions v1.4.0
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	golang.org/x/text v0.24.0 // indirect
)
