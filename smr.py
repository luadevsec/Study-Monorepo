# programa.py
import sys

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
        
    def execute(key, parametros=None):
        if key in creator_map:
            creator_method = creator_map[key]
            if parametros is None:
                creator_method()
            else:
                creator_method(*parametros)
        else:
            print("Chave não encontrada.")
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

    def execute(key, parametros=None):
        if key in delete_map:
            delete_method = delete_map[key]
            if parametros is None:
                delete_method()
            else:
                delete_method(*parametros)
        else:
            print("Chave não encontrada.")

delete_map = {
    "monorepo" : Delete.delete_monorepo,
    "project" : Delete.delete_project,
    "repo" : Delete.delete_repo
}


# Mapeamento de chaves para as funções de criação e exclusão
strategy_map = {
    "create": Creator.execute,
    "delete": Delete.execute
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


