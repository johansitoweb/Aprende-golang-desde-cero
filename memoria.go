// Asignación de memoria
var x *int = new(int) // Asigna memoria para un entero
*x = 10              // Asigna valor a la dirección de memoria
fmt.Println(*x)      // Imprime el valor
