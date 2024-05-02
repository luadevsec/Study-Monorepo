# Study Monorepo
 um gerenciador de monorepos para estudo


atualmente eu estou com um bug bizonho no comando help, e gostaria de falar um pouco sobre ele

eu estou usando (ou tentando) o design partner strategy que consiste em criar uma estrutura de map
com as funções de forma padronizada para ser chamado facilmente, o problema é que meu comando help
não recebe parametros e toda a logica do meu programa quebrou por conta disso, então eu cheguei a uma conclusão.
realmente quem fala que devemos fazer classes abstratas, interfaces e padrões está certo kkkkk eu nunca tive esse tipo de problema e não teria se eu tivesse um grande bloco de if e else mas aplicando um padrão de projeto eu acabo sendo obrigado a utilizar boas praticas, não só por charme mas principalmente por necessidade e isso só se aprende se lascando na pele kkkkk então esse erro que eu to pelando para resolver é horrivel mas importante. estarei agora refatorando todo o projeto para exceções como essa não estragarem meu fluxo. cada vez que eu evoluo como programadora eu percebo que tudo se baseia em trade-of usar if tem suas vantagens e não usar também tem, não tem muito que um certo ou errado depende do uso e da implementação. strategy é muito bom e charmoso (quando funciona) mas colocar strategy pra substituir 4 ifs é um over-engenier e isso só sofrendo que se aprende kkkkk


 Roadmap plain
 ---------------------------------------------------------

 - [x] implementar o padrão strategy para escalabilidade de funções
 - [x] aceitar parametros na chamada do sistema
    - [x] criação do sysloader
    - [x] criação do inputloader
 - [x] sysloader usando a arquitetura monolitica
 - [x] inputloader em loop com saida do programa
    - [ ] criar o metodo menu help
        - [x] sair do monolito, criação de classes externas e independentes
        - [x] implementar um padrão de comandos
        - [ ] refatorar o fluxo de execução
    - [ ] criar o metodo exit
 - [ ] criar metodo de usar/transformar projetos
 - [ ] criar a implementação de create
