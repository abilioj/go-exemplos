package entity

type User struct {
	id    int
	nome  string
	idade int
}

func (u *User) New(id int, nome string, idade int) {
	u.id = id
	u.nome = nome
	u.idade = idade
}

func (u *User) GetId() int {
	return u.id
}
func (u *User) GetNome() string {
	return u.nome
}
func (u *User) GetIdade() int {
	return u.idade
}
