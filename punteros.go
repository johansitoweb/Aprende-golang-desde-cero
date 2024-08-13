var x int = 10
var p *int = &x // p es un puntero a x
fmt.Println(*p) // Imprime el valor de x a través del puntero
