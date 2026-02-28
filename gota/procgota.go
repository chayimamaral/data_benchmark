package main

import (
	"context"
	"fmt"
	"log"
	"sort"
	"time"

	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
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

	rows, err := conn.Query(context.Background(),
		"SELECT produto, quantidade, valor_unitario, data_venda FROM vendas")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Criar slices para DataFrame
	var produtos []string
	var quantidades []int
	var valores []float64
	var datas []string

	for rows.Next() {
		var produto string
		var quantidade int
		var valor float64
		var data time.Time

		err := rows.Scan(&produto, &quantidade, &valor, &data)
		if err != nil {
			log.Fatal(err)
		}

		produtos = append(produtos, produto)
		quantidades = append(quantidades, quantidade)
		valores = append(valores, valor)
		datas = append(datas, data.Format("2006-01-02"))
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	// Criar DataFrame
	df := dataframe.New(
		series.New(produtos, series.String, "produto"),
		series.New(quantidades, series.Int, "quantidade"),
		series.New(valores, series.Float, "valor_unitario"),
		series.New(datas, series.String, "data_venda"),
	)

	fmt.Println("Linhas carregadas:", df.Nrow())

	// Criar coluna faturamento
	//faturamento := df.Col("quantidade").Float().
	//		Multiply(df.Col("valor_unitario").Float())

	quant := df.Col("quantidade").Float()
	val := df.Col("valor_unitario").Float()

	faturamento := make([]float64, len(quant))

	for i := range quant {
		faturamento[i] = quant[i] * val[i]
	}

	df = df.Mutate(series.New(faturamento, series.Float, "faturamento"))

	// Faturamento total
	total := df.Col("faturamento").Sum()

	// Faturamento por produto
	groupProduto := df.GroupBy("produto")
	resultProduto := groupProduto.Aggregation(
		[]dataframe.AggregationType{dataframe.Aggregation_SUM},
		[]string{"faturamento"},
	)

	// Faturamento por dia
	groupDia := df.GroupBy("data_venda")
	resultDia := groupDia.Aggregation(
		[]dataframe.AggregationType{dataframe.Aggregation_SUM},
		[]string{"faturamento"},
	)

	// Converter resultado dia para slice para ordenar
	type DiaValor struct {
		Data  string
		Valor float64
	}

	var dias []DiaValor
	for i := 0; i < resultDia.Nrow(); i++ {
		dias = append(dias, DiaValor{
			Data:  resultDia.Col("data_venda").Elem(i).String(),
			Valor: resultDia.Col("faturamento_SUM").Elem(i).Float(),
		})
	}

	sort.Slice(dias, func(i, j int) bool {
		return dias[i].Valor > dias[j].Valor
	})

	mediaDiaria := resultDia.Col("faturamento_SUM").Mean()

	end := time.Now()

	// ===== SAÍDA =====

	fmt.Printf("\nFaturamento Total: %.2f\n", total)

	fmt.Println("\nFaturamento por Produto:")
	fmt.Println(resultProduto)

	fmt.Printf("\nMédia Diária: %.2f\n", mediaDiaria)

	fmt.Println("\nTop 5 Dias:")
	for i := 0; i < 5 && i < len(dias); i++ {
		fmt.Printf("%s : %.2f\n", dias[i].Data, dias[i].Valor)
	}

	fmt.Println("\nTempo total:", end.Sub(start))
}
