package main

import (
	"context"
	"fmt"
	"log"
	"sort"
	"time"

	"github.com/jackc/pgx/v5"
)

type Venda struct {
	Produto       string
	Quantidade    int
	ValorUnitario float64
	DataVenda     time.Time
}

type ProdutoValor struct {
	Produto string
	Valor   float64
}

type DiaValor struct {
	Data  time.Time
	Valor float64
}

func main() {

	start := time.Now()

	conn, err := pgx.Connect(context.Background(),
		"postgres://postgres:postgres@localhost:5432/data_benchmark")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(context.Background())

	rows, err := conn.Query(context.Background(),
		"SELECT produto, quantidade, valor_unitario, data_venda FROM vendas")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	total := 0.0
	porProduto := make(map[string]float64)
	porDia := make(map[time.Time]float64)

	count := 0

	for rows.Next() {
		var v Venda
		err := rows.Scan(&v.Produto, &v.Quantidade, &v.ValorUnitario, &v.DataVenda)
		if err != nil {
			log.Fatal(err)
		}

		faturamento := float64(v.Quantidade) * v.ValorUnitario

		total += faturamento
		porProduto[v.Produto] += faturamento
		porDia[v.DataVenda] += faturamento
		count++
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	// Média diária
	somaDias := 0.0
	for _, v := range porDia {
		somaDias += v
	}
	mediaDiaria := somaDias / float64(len(porDia))

	// Ordenar produtos
	var listaProdutos []ProdutoValor
	for produto, valor := range porProduto {
		listaProdutos = append(listaProdutos, ProdutoValor{produto, valor})
	}

	sort.Slice(listaProdutos, func(i, j int) bool {
		return listaProdutos[i].Valor > listaProdutos[j].Valor
	})

	// Ordenar dias
	var listaDias []DiaValor
	for data, valor := range porDia {
		listaDias = append(listaDias, DiaValor{data, valor})
	}

	sort.Slice(listaDias, func(i, j int) bool {
		return listaDias[i].Valor > listaDias[j].Valor
	})

	end := time.Now()

	// ===== SAÍDA =====

	fmt.Println("Linhas processadas:", count)
	fmt.Printf("\nFaturamento Total: %.2f\n", total)

	fmt.Println("\nFaturamento por Produto:")
	for _, item := range listaProdutos {
		fmt.Printf("%-10s : %.2f\n", item.Produto, item.Valor)
	}

	fmt.Printf("\nMédia Diária: %.2f\n", mediaDiaria)

	fmt.Println("\nTop 5 Dias:")
	for i := 0; i < 5 && i < len(listaDias); i++ {
		fmt.Printf("%s : %.2f\n",
			listaDias[i].Data.Format("2006-01-02"),
			listaDias[i].Valor)
	}

	fmt.Println("\nTempo total:", end.Sub(start))
}
