package main

import "fmt"

type veiculo struct {
	portas   int
	cor string
}

type caminhonete struct {
	veiculo
	tracaoNasQuatro bool
}

type sedan struct {
	veiculo
	modeloLuxo bool	
}

func main() {

	hilux := caminhonete{
		veiculo{
			portas: 4,
			cor: "preta",
		},
		true,
	}
	corrola := sedan{
		veiculo{
			portas: 4,
			cor: "prata",
		},
		true,
	}
	
		
	fmt.Println("O veiculo é de Luxo: ",corrola.modeloLuxo)
	fmt.Println("Cor: ",corrola.cor)
	fmt.Println("Quantidade de portas: ",corrola.portas)
	fmt.Println("\n\n")

	fmt.Println("Tração nas Quatros Rodas: ",hilux.tracaoNasQuatro)
	fmt.Println("Cor: ",hilux.cor)
	fmt.Println("Quantidade de portas: ",hilux.portas)
}
