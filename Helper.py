from Command import Command
from Creator import Creator
from Deleter import Deleter

class Helper(Command):
    commandName = "Helper"
    

    summary_map = [
    'delete : \tdeleta um arquivo e limpa seus registros\n\t\tdelete <type> [group] <target> [clear]',
    'create : \tcria um arquivo e registra\n\t\tcreate <type> [group] <name> [register]',
    'help : \tmostra essa mensagem ou a ajuda de um metodo especifico\n\t\thelp [method]'
    ]   
        
    key_map = {
        'delete' : Deleter.help,
        'create' : Creator.help,
        'help' : help,
    }

    def default():
        for help in Helper.summary_map:
            print(help)

    def help():
        help_messenger = [
            "serio ? kkkkk você ta chamando o help do help ? kkkk mds"
        ]
        for line in help_messenger:
            print(line)
