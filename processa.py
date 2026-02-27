import time
import pandas as pd
from sqlalchemy import create_engine

# Conexão
engine = create_engine(
    "postgresql+psycopg2://postgres:postgres@localhost:5432/data_benchmark"
)

start = time.time()

# Carregar dados
df = pd.read_sql("SELECT * FROM vendas", engine)

print("Linhas carregadas:", len(df))

# Criar coluna faturamento
df["faturamento"] = df["quantidade"] * df["valor_unitario"]

# Faturamento total
total = df["faturamento"].sum()

# Faturamento por produto
por_produto = df.groupby("produto")["faturamento"].sum()

# Média diária
media_diaria = df.groupby("data_venda")["faturamento"].sum().mean()

# Top 5 dias
top5 = (
    df.groupby("data_venda")["faturamento"]
    .sum()
    .sort_values(ascending=False)
    .head(5)
)

end = time.time()

print("\nFaturamento Total:", total)
print("\nFaturamento por Produto:\n", por_produto)
print("\nMédia Diária:", media_diaria)
print("\nTop 5 Dias:\n", top5)
print("\nTempo total:", round(end - start, 2), "segundos")