package main

import "fmt"

func main() {

	estadosBrasileiros := make([]string, 27, 27)

	estadosBrasileiros = []string{"Acre", "Alagoas", "Amapá", "Amazonas", "Bahia", "Ceara", "Distrito Federal", "Espírito Santo", "Goiás", "Maranhão", "Mato Grosso", "Mato Grosso do Sul", "Minas Gerais", "Pará", "Paraíba", "Paraná", "Pernambuco", "Piauí", "Rio de Janeiro", "Rio Grande do Norte", "Rio Grande do Sul", "Rondônia", "Roraima", "Santa Catarina", "São Paulo", "Sergipe", "Tocantins"}

	fmt.Println(len(estadosBrasileiros), cap(estadosBrasileiros))
	fmt.Println(estadosBrasileiros)
	fmt.Println("\n\n")

	estadosBrasileiros = append(estadosBrasileiros, "Juazeiro do Norte")

	fmt.Println(len(estadosBrasileiros), cap(estadosBrasileiros))
	fmt.Println(estadosBrasileiros)
}
