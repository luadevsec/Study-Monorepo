class Command:
    commandName = None
    

    # metodos default para serem sobrescritos
    
    def default():
        print("metodo default não implementado")


# padrão de execução        

    @staticmethod
    @classmethod
    def error(cls, parametros):
        print(f"erro ao executar {parametros[0]} de {cls.commandName}")
        print("digite help para obter ajuda")

    @classmethod
    def execute(cls, key=None, *parametros):
        def run(command, *parametros):
            if parametros == ():
                command()
                return
            command(*parametros)

        if key not in cls.key_map:
            if key not in Command.key_map:
                cls.error([key, *parametros])
                return
            run(Command.key_map[key], *parametros)
            return
        run(cls.key_map[key], *parametros)

    @classmethod
    def help(cls):    
        for line in cls.help_messenger:
            print(line)


# padrão de implementação
    key_map = {
        None: default,
        'help': help,
    }

    help_messenger = [
        "comando de ajuda ainda não implementado",
        "por favor, implemente o comando de ajuda para a classe atual",
        "adicionando um array de help_messenger com as mensagens de ajuda",
        "ou sobrescrevendo o metodo help() da classe atual"
    ]