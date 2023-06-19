package q2

//Escreva uma função para encontrar o prefixo comum mais longo entre um array de strings.
//
//Se não houver um prefixo comum, retorne uma string vazia "".

func LongestCommonPrefix(strs []string) string {
	prefixo := strs[0]
	for i := 1; i < len(strs); i++ {
		for !strings.HasPrefix(strs[i], prefixo) {
			prefixo = prefixo[:len(prefixo)-1]
		}
	}
	if len(strs) == 0 || len(prefixo) == 0 {
		return ""
	}
	return ""
}
