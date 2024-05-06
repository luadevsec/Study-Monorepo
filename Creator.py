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

    def help():
        help_messenger = [
            "serio ? kkkkk você ta chamando o help do help ? kkkk mds"
        ]
        for line in help_messenger:
            print(line)
    
    key_map = {
        "monorepo" : create_monorepo,
        "project" : create_project,
        "repo" : create_repo,
        "help" : help,
        None: default
    }

