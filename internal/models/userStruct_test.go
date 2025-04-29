package models

import (
	"reflect"
	"testing"
	"time"
)

func TestUserStruct(t *testing.T) {
	// Ожидаемые типы:
	intType := reflect.TypeOf(1)
	stringType := reflect.TypeOf("1")
	dateType := reflect.TypeOf(time.Time{})

	testCases := []UserStruct{
		{
			ID:         1,
			Login:      "UserExmpl1",
			Password:   "PassExmpl1",
			FirstName:  "FrstNameExmpl1",
			SecondName: "ScndNameExmpl1",
			RegData:    time.Date(2025, 04, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			ID:         2,
			Login:      "UserExmpl2",
			Password:   "PassExmpl2",
			FirstName:  "FrstNameExmpl2",
			SecondName: "ScndNameExmpl2",
			RegData:    time.Date(2025, 04, 1, 0, 0, 0, 0, time.UTC),
		},
	}

	if idType := reflect.TypeOf(testCases[0].ID); idType != intType {
		t.Errorf("В ID ожидается тип: %v , а на самом деле тип: %v", intType, idType)
	}
	if loginType := reflect.TypeOf(testCases[0].Login); loginType != stringType {
		t.Errorf("В Login ожидается тип: %v, а на самом деле тип: %v", stringType, loginType)
	}
	if passwordType := reflect.TypeOf(testCases[0].Password); passwordType != stringType {
		t.Errorf("В Password ожидается тип: %v, а на самом деле тип: %v", stringType, passwordType)
	}
	if firstNameType := reflect.TypeOf(testCases[0].FirstName); firstNameType != stringType {
		t.Errorf("В FirstName ожидается тип: %v, а на самом деле тип: %v", stringType, firstNameType)
	}
	if secondNameType := reflect.TypeOf(testCases[0].SecondName); secondNameType != stringType {
		t.Errorf("В SecondName ожидается тип: %v, а на самом деле тип: %v", stringType, secondNameType)
	}
	if dataStructType := reflect.TypeOf(testCases[0].RegData); dataStructType != dateType {
		t.Errorf("В RegData ожидается тип: %v, а на самом деле тип: %v", dateType, dataStructType)
	}

}

// Библия структуры
// 1. Структура должна создавать по своему подобию переменные
// 2. Нужно проверять появление новых элементов структуры
// 3. Нужно проверять тип данных элемента структуры
// 4. Сразу проверять, что айдишник итерируется сам
// 5. Проверить валидность значений даты
