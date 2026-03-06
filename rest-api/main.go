package main

import (
	"fmt"
	"rest-api/http"
	"rest-api/postgres/simple_connection"
	"rest-api/todo"
)

func main() {
	conn := simple_connection.CheckConnection()

	fmt.Println(conn)

	todoList := todo.NewList()
	httpHandlers := http.NewHTTPHandlers(todoList, conn)
	httpServer := http.NewHTTPServer(httpHandlers)
	if err := httpServer.StartServer(); err != nil {
		fmt.Println("failed to start server")
	}
}

/*
============================SQL==================================
||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||
====Вставка
--------INSERT INTO users (full_name, phone_number)
--------VALUES ('Иванов Иван Иванович', '+7 999 888 66 55');
======Без телефона
--------INSERT INTO users (full_name)
--------VALUES ('Иванов Иван Иванович');
||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||||
====Получение данных
--------SELECT full_name, phone_number
--------FROM users
======Получение где номер телефона не NULL
--------SELECT full_name, phone_number
--------FROM users
--------WHERE phone_number IS NOT NULL
====Удаление из БД по id
--------DELETE FROM users WHERE id=2
*/
