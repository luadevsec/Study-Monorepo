# listas descritivas de importações para certificar que o programa não é malicioso
from sys import argv # usado para carregar parametros na chamada do programa
from os import getcwd # usado para carregar o caminho do smr visualmente
from os import path # usado para filtrar o caminho vindo do getcwd

# Importações de facades
from Creator import Creator
from Deleter import Deleter
from Helper import Helper
from Test import Test

# Mapeamento de chaves para as funções de execução
command_map = {
    "create": Creator.execute,
    "delete": Deleter.execute,
    "help": Helper.execute,
    "test": Test.execute
}

def error(code, *parametros): # função que trata os erros
    error_map = {
        0: f"entrada vazia ou invalida: {parametros}",
        1: f"comando '{parametros[0]}' não encontrado"
    }
    print(error_map[code])
    Helper.execute()

def Loader(entrada): # função basica que trata a entrada e chama a função correta
    command_key = entrada[0]
    if not (command_key in command_map): #fail first
        error(1, command_key)
        return
    command = command_map[command_key]
    if len(entrada) == 1:
        command()
        return
    command(entrada[1], entrada[2:])

def sysLoader (): # função que carrega o programa com os parametros passados na chamada do terminal
    Loader(argv[1:]) # o primeiro parametro é o nome do programa, o 1 em diante é os parametros

def inputLoader (): # função que carrega o programa com os parametros passados no input do terminal
    instance = path.basename(getcwd())
    run = True 
    while(run):
        entrada = input(f"{instance} smr> ")
        entrada = entrada.split(" ")
        if entrada[0] == "exit":
            run = False
            continue
        if len(entrada) >= 1:
            Loader(entrada)
            continue
        error(0, entrada)

def start(): 
    entrada = argv[1:]

    if len(entrada) == 0:
        inputLoader()
        return
    sysLoader() 




start() # chamada da função start