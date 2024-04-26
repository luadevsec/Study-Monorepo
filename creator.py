# creator.py
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

# Mapeamento de chaves para os métodos
creator_map = {
    "monorepo" : Creator.create_monorepo,
    "project" : Creator.create_project,
    "repo" : Creator.create_repo
}

def execute(key, parametros=None):
    if key in creator_map:
        creator_method = creator_map[key]
        if parametros is None:
            creator_method()
        else:
            creator_method(*parametros)
    else:
        print("Chave não encontrada.")
