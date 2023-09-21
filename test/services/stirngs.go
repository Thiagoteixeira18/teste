package services_test

const (
	espanhol            = "espanhol"
	frances             = "frances"
	prefixoOlaPortugues = "Olá, "
	prefixoOlaEspanhol  = "Hola, "
	prefixoOlaFrances   = "Bonjour, "
)

func Ola(nome string, idioma string) string {
	if nome == "" {
		nome = "Mundo"
	}

	return prefixodeSaudacao(idioma) + nome
}

func prefixodeSaudacao(idioma string) (prefixo string) {
	switch idioma {
	case frances:
		prefixo = prefixoOlaFrances
	case espanhol:
		prefixo = prefixoOlaEspanhol
	default:
		prefixo = prefixoOlaPortugues
	}
	return
}
