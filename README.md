# 📊 Benchmark de Processamento de Dados: Python vs. Go#

Este repositório contém os testes de performance que realizei processando 5 milhões de registros de uma base PostgreSQL 18.3 rodando no Fedora 43. 

O objetivo foi comparar a eficiência do "padrão de mercado" (Python/Pandas) com a robustez e velocidade do Go.

## 🚀 Resultados do Benchmark

Os testes consistiram em extrair dados, calcular faturamento total, média diária e o ranking dos Top 5 produtos/dias.


| Tecnologia | Estratégia | Tempo Total | Performance vs. Python |
| :--- | :--- | :--- | :--- |
| **🐍 Python (Pandas)** | Carregamento em Memória | **12,07s** | 1x (Referência) |
| **🐹 Go (Nativo)** | Processamento em Slices | **2,68s** | 4,5x mais rápido |
| **⚡ Go + SQL Otimizado** | Agregação no Banco (pgx) | **1,16s** | **10,4x mais rápido** |


## 🛠️ Stack Tecnológica

| Componente | Especificação |
| :--- | :--- |
| **Sistema Operacional** | Fedora 43 |
| **Linguagens** | Go (Nativo) \| Python (Pandas) |
| **Banco de Dados** | PostgreSQL 18.3 🚀 |
| **Monitoramento** | Mission Center (Gráficos de linha em tempo real) |

##💡 Principais Conclusões

### 🧠 Análise Técnica

| Pilar | Observação do Benchmark |
| :--- | :--- |
| **Eficiência de CPU** | Enquanto o Python mantém um platô alto de consumo, o Go entrega um pulso rápido e libera o hardware. Isso é economia real de infraestrutura. |
| **Engenharia de Dados** | A versão "Go + SQL" prova que delegar a agregação ao banco de dados reduz drasticamente o I/O e o uso de memória. |
| **Escalabilidade** | O Go se consolida como a ferramenta superior para sistemas de missão crítica onde performance e custo são prioridades. |

## 💰 Cloud Cost Analysis (GCP, AWS & Azure)

A performance bruta do Go reflete diretamente na economia de infraestrutura. Em ambientes de nuvem, onde o faturamento é baseado em **recursos alocados x tempo**, a eficiência do Go + SQL gera uma redução drástica no TCO (Total Cost of Ownership).

| Métrica de Custo | 🐍 Python (Pandas) | 🐹 Go (Nativo) | 🚀 Go + SQL |
| :--- | :--- | :--- | :--- |
| **Perfil de Instância** | High Memory ($$$) | Standard ($$) | Micro/Small ($) |
| **Tempo de Bilhetagem** | 12,07s | 2,68s | **1,16s** |
| **Custo de CPU** | 100% (Referência) | ↓ 78% de economia | **↓ 90% de economia** |
| **Tráfego de Rede** | Alto (Extração total) | Médio (Slices) | **Mínimo (Agregação no DB)** |

### O impacto financeiro:
1. **Serverless (Lambda/Functions):** O custo de execução é quase 10x menor, pois o tempo de faturamento cai de 12s para ~1s.
2. **Densidade de Containers (K8s):** No seu **Fedora**, vimos que o Go consome uma fração da RAM do Python. No cluster (GKE/EKS), isso permite rodar muito mais instâncias no mesmo node, reduzindo a necessidade de instâncias caras.
3. **Data Transfer Out:** Ao processar no PostgreSQL 18.3, eliminamos o custo de trafegar 5 milhões de linhas entre o banco e a aplicação.


### 🚀 Como Reproduzir o Benchmark

| Passo | Ação | Comando / Local |
| :--- | :--- | :--- |
| **1** | Preparar o Banco | Use o script de geração de massa na pasta `/scripts` |
| **2** | Executar Python | `python processa.py` |
| **3** | Executar Go (Memória) | `go run processa.go` |
| **4** | Executar Go (SQL) | `go run processa_sql.go` |


