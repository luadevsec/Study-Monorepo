class delete:
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

    
delete_map = {
    "monorepo" : delete.delete_monorepo,
    "project" : delete.delete_project,
    "repo" : delete.delete_repo
}

def execute(key, parametros=None):
    if key in delete_map:
        delete_method = delete_map[key]
        if parametros is None:
            delete_method()
        else:
            delete_method(*parametros)
    else:
        print("Chave não encontrada.")
