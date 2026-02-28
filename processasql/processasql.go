package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5"
)

func main() {

	start := time.Now()

	conn, err := pgx.Connect(context.Background(),
		"postgres://postgres:postgres@localhost:5432/data_benchmark")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(context.Background())

	// =========================
	// MÉTRICAS GLOBAIS
	// =========================
	var faturamentoTotal float64
	var mediaDiaria float64

	err = conn.QueryRow(context.Background(), `
		WITH faturamento_dia AS (
			SELECT data_venda,
				   SUM(quantidade * valor_unitario) AS total_dia
			FROM vendas
			GROUP BY data_venda
		)
		SELECT
			(SELECT SUM(quantidade * valor_unitario) FROM vendas) AS total,
			(SELECT AVG(total_dia) FROM faturamento_dia) AS media
	`).Scan(&faturamentoTotal, &mediaDiaria)

	if err != nil {
		log.Fatal(err)
	}

	// =========================
	// FATURAMENTO POR PRODUTO
	// =========================
	rowsProduto, err := conn.Query(context.Background(), `
		SELECT produto,
		       SUM(quantidade * valor_unitario) AS total_produto
		FROM vendas
		GROUP BY produto
		ORDER BY total_produto DESC
	`)
	if err != nil {
		log.Fatal(err)
	}
	defer rowsProduto.Close()

	fmt.Printf("\nFaturamento Total: %.2f\n", faturamentoTotal)
	fmt.Printf("Média Diária: %.2f\n", mediaDiaria)

	fmt.Println("\nFaturamento por Produto:")
	for rowsProduto.Next() {
		var produto string
		var total float64
		rowsProduto.Scan(&produto, &total)
		fmt.Printf("%-10s : %.2f\n", produto, total)
	}

	// =========================
	// TOP 5 DIAS
	// =========================
	rowsDia, err := conn.Query(context.Background(), `
		SELECT data_venda,
		       SUM(quantidade * valor_unitario) AS total_dia
		FROM vendas
		GROUP BY data_venda
		ORDER BY total_dia DESC
		LIMIT 5
	`)
	if err != nil {
		log.Fatal(err)
	}
	defer rowsDia.Close()

	fmt.Println("\nTop 5 Dias:")
	for rowsDia.Next() {
		var data time.Time
		var total float64
		rowsDia.Scan(&data, &total)
		fmt.Printf("%s : %.2f\n",
			data.Format("2006-01-02"),
			total)
	}

	end := time.Now()

	fmt.Println("\nTempo total:", end.Sub(start))
}