# listas descritivas de importações para certificar que o programa não é malicioso
from sys import argv # usado para carregar parametros na chamada do programa
from os import getcwd # usado para carregar o caminho do smr visualmente
from os import path # usado para filtrar o caminho vindo do getcwd

# Importações de facades
from Creator import Creator
from Deleter import Deleter
from Helper import Helper




# Mapeamento de chaves para as funções de criação e exclusão
strategy_map = {
    "create": Creator.execute,
    "delete": Deleter.execute,
    "help": Helper.execute
}

def Loader(entrada : str):
    strategy_key = entrada[0]
    key = parametros = None
    if len(entrada) > 1: 
        key = entrada[1]
    if len(entrada) > 2:
        parametros = entrada[2:]
    
    if strategy_key == 'help':
        if key is None:  # Verifica se não há chave (key)
            Helper.execute()  # Chama a função sem argumentos
        else:
            Helper.execute(key)  # Chama a função com a chave (key)
    elif strategy_key in strategy_map:
        if key is None:  # Verifica se não há chave (key)
            strategy_map[strategy_key]()  # Chama a função sem argumentos
        else:
            strategy_map[strategy_key](key, parametros)  # Chama a função com a chave (key) e parâmetros
    else:
        print(f"comando '{strategy_key}' não encontrado")


# definição de facades
def sysLoader():
    Loader(argv[1:])

def inputLoader():

    
    run = True
    while(run):
        local = path.basename(getcwd())
        entrada = input(f"{local} smr> ")
        entrada = entrada.split(" ")
        if entrada[0] == "exit":
            run = False
            continue
        if len(entrada) >= 1:
            Loader(entrada)
            continue
        error(entrada)
    
    
def error(entrada):
    print(f"comando '{' '.join(map(str, entrada))}' invalido ou faltante, digite 'help' para obter ajuda")

# Verifica se há pelo menos dois argumentos (o primeiro é o nome do script e o segundo é a chave)
if len(argv) > 2: sysLoader()
else: inputLoader()


