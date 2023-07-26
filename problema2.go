

/* Desafio 2 - Crie um codigo que print "pin"" para todos  os numeros divisiveis por 3 entre 1 e 100 e
"pan" para todos  os numeros divisiveis por 5 entre 1 e 100 */



package main
import "fmt"

func main(){
	for i := 1 ; i <= 100; i++ {
		if i%3 == 0 {fmt.Println("Pin!")} else if  i%5 == 0 {fmt.Println("Pan!")}   else {fmt.Println(i)}
		
}
}