package entity

type User struct {
	ID    int
	Nome  string
	Idade int
}

func (u *User) New(id int, nome string, idade int) {
	u.ID = id
	u.Nome = nome
	u.Idade = idade
}
