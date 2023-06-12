package q5

import "strings"

//Uma frase é um palíndromo se, após converter todas as letras maiúsculas em letras minúsculas e remover todos os
//caracteres não alfanuméricos, ela for lida da mesma forma da esquerda para a direita e vice-versa. Caracteres
//alfanuméricos incluem letras e números.
//
//Dada uma string "s", retorne verdadeiro se for um palíndromo e falso caso contrário.

func IsPalindrome(s string) bool {

	s1 := ""
	for _, letras := range strings.ToLower(s) {
		if strings.ContainsAny(string(letras), "bcdfghjklmnpqrstvwxyzaeiou0123456789") {
			s1 += strings.ToLower(string(letras))
		}
	}
	transformadorrune := []rune(s1)
	var reverso []rune
	for i := len(transformadorrune) - 1; i >= 0; i-- {
		reverso = append(reverso, transformadorrune[i])
	}
	if strings.ToLower(s1) == strings.ToLower(string(reverso)) {
		return true
	} else {
		return false
	}
}
