package main
import "fmt"

//structs
type User struct {
	name string
	age uint
}
func main(){
	//variaveis
	// var numero int64
	var username string = "Pedro"

	//arrays e slice
	var arr []int //slice tamanho dinamica
	var arr2 = []int{1,2,3,4,5,6} //array tamanho estatico
	arr2 = append(arr2, 20)
	fmt.Println(arr)
	fmt.Println(arr2 ,cap(arr2))

	//maps
	obj := map[string]string{
		"name": "Sammy",
		"animal": "shark",
		"color": "blue",
		"location": "ocean",
	}

	animal := obj["animal"]


	
	fmt.Println(obj["name"])
	fmt.Println(animal)
	funcaoNome(username)



	user := User{name: "Pedro", age: 21}
	fmt.Println(user)
}

//funcao
func funcaoNome(nome string) string{	
	fmt.Println(nome)
	return nome
}







