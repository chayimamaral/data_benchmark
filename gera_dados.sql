CREATE TABLE vendas (
    id SERIAL PRIMARY KEY,
    produto VARCHAR(50),
    quantidade INT,
    valor_unitario DECIMAL(10,2),
    data_venda DATE
);

-- Gerando 5 milhões de linhas aleatórias
INSERT INTO vendas (produto, quantidade, valor_unitario, data_venda)
SELECT 
    (ARRAY['Teclado', 'Mouse', 'Monitor', 'Gabinete', 'Headset'])[floor(random() * 5 + 1)],
    floor(random() * 10 + 1)::int,
    (random() * 500 + 50)::decimal(10,2),
    CURRENT_DATE - (random() * 365)::int
FROM generate_series(1, 5000000);