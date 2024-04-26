# programa.py
import sys
import creator
import delete

# Mapeamento de chaves para as funções de criação e exclusão
strategy_map = {
    "create": creator.execute,
    "delete": delete.execute
}

def sysLoader():
    strategy_key = sys.argv[1]
    key = sys.argv[2]
    parametros = sys.argv[3:]
    if strategy_key in strategy_map:
        strategy_map[strategy_key](key, parametros)
    else:
        print(f"comando '{strategy_key}' não encontrado") 

def inputLoader():
    print("help file")
# Verifica se há pelo menos dois argumentos (o primeiro é o nome do script e o segundo é a chave)
if len(sys.argv) > 2: sysLoader()
else: inputLoader()
