# listas descritivas de importações para certificar que o programa não é malicioso
from sys import argv # usado para carregar parametros na chamada do programa
from os import getcwd # usado para carregar o caminho do smr visualmente
from os import path # usado para filtrar o caminho vindo do getcwd

class Creator:
    @staticmethod
    def create_monorepo(name: str = "Studys", local: str = "./"):
        print("create_monorepo - name:", name, "local:", local)
        # Implementação para criar um monorepo
        pass

    @staticmethod
    def create_project(name: str = "new_project"):
        print("create_project - name:", name)
        # Implementação para criar um projeto
        pass

    @staticmethod
    def create_repo(name: str = "new_repo"):
        print("create_repo - name:", name)
        # Implementação para criar um repositório
        pass
    
    @staticmethod
    def help():
        print("metodo não implementado")
        
    def execute(key, parametros=None):
        if key in creator_map:
            creator_method = creator_map[key]
            if parametros is None:
                creator_method()
            else:
                creator_method(*parametros)
        else:
            error([key, parametros])
# Mapeamento de chaves para os métodos
creator_map = {
    "monorepo" : Creator.create_monorepo,
    "project" : Creator.create_project,
    "repo" : Creator.create_repo
}


class Delete:
    @staticmethod
    def delete_monorepo(name: str = "Studys", local: str = "./"):
        print("delete_monorepo - name:", name, "local:", local)
        # Implementação para criar um monorepo
        pass

    @staticmethod
    def delete_project(name: str = "new_project"):
        print("delete_project - name:", name)
        # Implementação para criar um projeto
        pass

    @staticmethod
    def delete_repo(name: str = "new_repo"):
        print("delete_repo - name:", name)
        # Implementação para criar um repositório
        pass

    @staticmethod
    def help():
        print("metodo não implementado")

    def execute(key, parametros=None):
        if key in delete_map:
            delete_method = delete_map[key]
            if parametros is None:
                delete_method()
            else:
                delete_method(*parametros)
        else:
            error([key, parametros])

delete_map = {
    "monorepo" : Delete.delete_monorepo,
    "project" : Delete.delete_project,
    "repo" : Delete.delete_repo
}

class Helper:
    summary_map = [
        'delete : \tdeleta um arquivo e limpa seus registros\n\t\tdelete <type> [group] <target> [clear]',
        'create : \tcria um arquivo e registra\n\t\tcreate <type> [group] <name> [register]',
        'help : \tmostra essa mensagem ou a ajuda de um metodo especifico\n\t\thelp [method]'
    ]
    help_map = {
        'delete' : Delete.help,
        'create' : Creator.help,
        'help' : help,
    }

    @staticmethod
    def help():
        help_messenger = [
            "serio ? kkkkk você ta chamando o help do help ? kkkk mds"
        ]
        for line in help_messenger:
            print(line)
    
    @staticmethod
    def execute(metodo : str = "all"):
        if metodo == "all":
            for help in Helper.summary_map:
                print(help)
            return
        if metodo in Helper.help_map:
            Helper.help_map[metodo]()
            return
        error(metodo)
        




# Mapeamento de chaves para as funções de criação e exclusão
strategy_map = {
    "create": Creator.execute,
    "delete": Delete.execute,
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


