package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Generar un contexto vacío
	ctx := context.Background()

	// Agregar un valor al contexto
	ctx = context.WithValue(ctx, "my-key", "my value")
	ctx = context.WithValue(ctx, "my-key-int", 5)

	viewContext(ctx)

	// Generar un contexto con un timeout
	ctx2, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel() // Se ejecuta sí o sí

	// Ejecutar un proceso en un goroutine
	myProcess(ctx2)
	// Esperar a que el contexto se cancele a través de channels
	select {
	case <-ctx2.Done():
		fmt.Println("context canceled")
		fmt.Println(ctx2.Err())
	}

	fmt.Println("--------------------")

}

// viewContext imprime el valor asociado con la clave "my-key" del contexto
func viewContext(ctx context.Context) {
	fmt.Printf("my value is '%s'\n", ctx.Value("my-key"))
	fmt.Printf("my other value in the context is %d\n", ctx.Value("my-key-int"))
}

func myProcess(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("context canceled")
			return
		default:
			fmt.Println("working...")
		}
		time.Sleep(1 * time.Second)
	}
}
