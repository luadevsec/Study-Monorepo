from Command import Command

class User(Command):
    commandName = "User"
    
    

    def help():
        print("metodo help implementado")
        pass

    key_map = {
        
        "help" : help,
    }

User.execute()