class Command:
    commandName = None
    
    @classmethod
    def default(cls):
        print("metodo default não implementado")
        cls.help()

    def help():
        print("metodo help não implementado")
        
    @classmethod
    def error(cls, parametros):
        print(f"erro ao executar {parametros[0]} de {cls.commandName}")
        cls.help()

    @classmethod
    def execute(cls, key=None, *parametros):
        if key not in cls.key_map:
            cls.error([key, parametros])
            return
        command = cls.key_map[key]
        if parametros is None:
            command()
            return
        command(*parametros)

    key_map = {
        None: default
    }