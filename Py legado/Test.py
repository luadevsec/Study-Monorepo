# a classe de teste é um teste para verificar se a classe command funciona
# e se o fluxo de execução do programa está correto
# se essa classe funcionar, então qualquer comando pode funcionar tambem
# se for implementado corretamente

from Command import Command

class Test(Command):
    commandName = "Test"
    
    help_messenger = [
        "eu sou um teste, aqui está oq eu posso fazer:",
        "test : \tmostra essa mensagem",
        "help : \tmostra essa mensagem"
    ]

    key_map = {
        
    }