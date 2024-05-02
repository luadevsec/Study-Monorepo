from Command import Command

class Creator(Command):
    commandName = "Creator"
    

    def create_monorepo(name: str = "Studys", local: str = "./"):
        print("create_monorepo - name:", name, "local:", local)
        # Implementação para criar um monorepo
        pass

    def create_project(name: str = "new_project"):
        print("create_project - name:", name)
        # Implementação para criar um projeto
        pass

    def create_repo(name: str = "new_repo"):
        print("create_repo - name:", name)
        # Implementação para criar um repositório
        pass

    
    def default():
        print("metodo default implementado")
        Creator.help()
        pass
    
    key_map = {
        "monorepo" : create_monorepo,
        "project" : create_project,
        "repo" : create_repo,
        "help" : help,
        None: default
    }

