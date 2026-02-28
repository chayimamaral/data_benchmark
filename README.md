📊 #Benchmark de Processamento de Dados: Python vs. Go#

Este repositório contém os testes de performance que realizei processando 5 milhões de registros de uma base PostgreSQL 18.3 rodando no Fedora 43. 

O objetivo foi comparar a eficiência do "padrão de mercado" (Python/Pandas) com a robustez e velocidade do Go.

🚀 ##Resultados do Benchmark

Os testes consistiram em extrair dados, calcular faturamento total, média diária e o ranking dos Top 5 produtos/dias.


| Tecnologia | Estratégia | Tempo Total | Performance vs. Python |
| :--- | :--- | :--- | :--- |
| **🐍 Python (Pandas)** | Carregamento em Memória | **12,07s** | 1x (Referência) |
| **🐹 Go (Nativo)** | Processamento em Slices | **2,68s** | 4,5x mais rápido |
| **⚡ Go + SQL Otimizado** | Agregação no Banco (pgx) | **1,16s** | **10,4x mais rápido** |


🛠️ Stack Tecnológica

### 🛠️ Stack Tecnológica

| Componente | Especificação |
| :--- | :--- |
| **Sistema Operacional** | Fedora 43 |
| **Linguagens** | Go (Nativo) \| Python (Pandas) |
| **Banco de Dados** | PostgreSQL 18.3 🚀 |
| **Monitoramento** | Mission Center (Gráficos de linha em tempo real) |

Sistema Operacional: Fedora 43

Linguagens: Go (Nativo) | Python (Pandas)

Banco de Dados: PostgreSQL 18.3 🚀

Monitoramento: Mission Center (Gráficos de linha em tempo real)



💡 Principais Conclusões

Eficiência de CPU: Enquanto o Python mantém um platô alto de consumo, o Go entrega um pulso rápido e libera o hardware. Isso é economia real de infraestrutura.

Engenharia de Dados: A versão "Go + SQL" prova que delegar a agregação ao banco de dados reduz drasticamente o I/O e o uso de memória.

Escalabilidade: O Go se consolida como a ferramenta superior para sistemas de missão crítica onde performance e custo são prioridades.

###Como reproduzir
###Preparar o Banco:### Use o script de geração de massa de dados na pasta /scripts.
###Executar Python:### python processa.py
###Executar Go (Memória):### go run processa.go
###Executar Go (SQL Otimizado):### go run processa_sql.go
