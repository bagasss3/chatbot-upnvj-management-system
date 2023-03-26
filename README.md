# chatbot-upnvj-management-system

Chatbot Management System for UPNVJ

How to Run Server:
go run main.go server

How to create sql migration file:
goose -dir ./database/migration create [name of file(without bracket)] sql

How to run migration:
go run main.go migrate --direction=down
go run main.go migrate --direction=up
