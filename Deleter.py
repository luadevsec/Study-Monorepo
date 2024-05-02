from Command import Command

class Deleter(Command):
    commandName = "Deleter"
    
    

    def help():
        print("metodo help implementado")
        pass

    
    key_map = {
        
        "help" : help,
    }
