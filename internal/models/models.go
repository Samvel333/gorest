package models

type Person struct {
	ID          string `json:"id"` // I Prefer store it as UUID (string)
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Patronymic  string `json:"patronymic,omitempty"` // omitempty указывает JSON-кодировщику исключить поле JSON, если значение поля считается пустым.
	Age         int    `json:"age"`
	Gender      string `json:"gender"`
	Nationality string `json:"nationality"`
}