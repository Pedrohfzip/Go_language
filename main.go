package main
import "fmt"
func main(){
	//variaveis
	var numero int64
	var username string = "Pedro"
	
	fmt.Println(username)
	fmt.Println(numero)
	funcaoNome(username)
}

//funcao
func funcaoNome(nome string) string{	
	fmt.Println(nome)
	return nome
}

//arrays



