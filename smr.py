# programa.py
import sys
import creator
import delete

# Mapeamento de chaves para as funções de criação e exclusão
strategy_map = {
    "create": creator.execute,
    "delete": delete.execute
}

# Verifica se há pelo menos dois argumentos (o primeiro é o nome do script e o segundo é a chave)
if len(sys.argv) > 2:
    print(sys.argv)
    # Dividindo a string de parâmetros em uma lista usando espaço como separador
    parametros = sys.argv[2:]
    print("Parâmetros passados:")
    for parametro in parametros:
        print("-", parametro)
    # Passando a chave e os parâmetros para a função execute da estratégia correspondente
    strategy_key = sys.argv[1]
    if strategy_key in strategy_map:
        strategy_map[strategy_key](parametros[0], parametros[1:])
    else:
        print("Chave não encontrada.")
else:
    print("help")
