# Definindo as funções que serão utilizadas como estratégias
def create_r(repo):
    print("create repo %s" % repo)

def create_p(project):
    print("create project %s" % project)

# Criando um dicionário para mapear chaves às funções
strategy_map = {
    'create repo': create_r,
    'create project': create_p,
}

# Função para chamar a estratégia com base na chave fornecida
def execute_strategy(key_with_param):
    parts = key_with_param.split()  # Dividindo a string em partes usando espaço como separador
    if len(parts) >= 2:
        key = ' '.join(parts[:-1])  # Juntando as partes da chave
        param = parts[-1]  # O último elemento é o parâmetro
        if key in strategy_map:
            strategy = strategy_map[key]
            strategy(param)  # Chamando a função da estratégia com o parâmetro
        else:
            print("Chave não encontrada.")
    else:
        print("Formato inválido. Use 'chave parametro'.")

# Exemplo de uso
execute_strategy('create repo ss')
execute_strategy('create project java')
execute_strategy('chave3')
execute_strategy('chave4')  # Testando uma chave que não existe
