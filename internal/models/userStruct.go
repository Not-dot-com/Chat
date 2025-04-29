package models

import "time"

// Для базы данных автоматически задаются имена элементов каждого элемента,
// для json файла всё, кроме пароля
type UserStruct struct {
	ID         int       `json:"ID" db:"ID"`
	Login      string    `json:"Login" db:"Login"`
	Password   string    `json:"-" db:"Password"`
	FirstName  string    `json:"FirstName" db:"FirstName"`
	SecondName string    `json:"SecondName" db:"SecondName"`
	RegData    time.Time `json:"RegData" db:"RegData"`
}

// Есть некие сомнения по поводу такой структуры
// Нужно добавить ещё дату регистрации и дату обновления данных пользователем
// Дату обновления хочу сделать отдельной структурой, чтоб можно было внутри
// UserRegStruct.UpdateDate хранить несколько дат обновления.
// Id сделать автоматически увеличивающимся для каждого пользователя
//
