package main

import "fmt"

func main() {
	// YOU NEED TO SET THIS
	// Propriedade da SOFTSOL e CREFOCUS
	prazo := 360
	valorEmprestimo := 200_000.0
	taxaJurosAnual := 8.0
	valorMaximo := 4000.0
	// TODO atualizar valorMaximo sobre inflacao
	// TODO aumentar salario sobre o tempo?
	// TODO calculadora FGTS, com invetimento do saque aniversao, e aumento de saldo em 3% ao ano
	//

	valorRestante := valorEmprestimo
	taxaJurosMensal := (taxaJurosAnual / 12) / 100
	parcela := valorEmprestimo / float64(prazo)

	fmt.Printf("Taxa mensal: %f\n", taxaJurosMensal)
	fmt.Printf("Valor parcela mensal: %f\n", parcela)

	for i := range prazo {
		parcelaAtual := parcela + (taxaJurosMensal * valorRestante)

		if parcelaAtual >= valorRestante {
			pago := parcelaAtual - valorRestante
			fmt.Printf("Ultima parcela: %d Valor: %f\n", i, pago)
			return
		}

		if parcelaAtual > valorMaximo {
			realPago := valorMaximo - (taxaJurosMensal * valorRestante)
			valorRestante -= realPago
			fmt.Println(realPago)
			fmt.Println(valorRestante)
			fmt.Printf("Parcela: %d Valor: %f\n", i, valorMaximo)
		} else {
			amortizado := valorMaximo - parcelaAtual
			valorRestante -= parcela
			valorRestante -= amortizado
			fmt.Printf("Parcela: %d Valor: %f\n", i, parcelaAtual)
			fmt.Printf("Parcela: %d Amortizado: %f\n", i, amortizado)
		}
	}
}
