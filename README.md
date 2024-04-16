# Prova Detran CLI

Simulado de provas do Detran.

## Como instalar e utilizar

```
go install github.com/jefersonf/prova-detran
```

Em seguida inicie um novo simulado.

```
prova-detran simulado
```

Para executar a partir do código fonte basta clonar este repositório.

```
cd prova-detran
go build .
```
e em seguida inicie um novo simulado.

```
./prova-detran simulado
```

## Exemplo

![exemplo de simulado](/sample-output.png)

## Banco de questões

As questões presentes nessa ferramenta foram extraídas de livros e materiais de apoio que obtive acesso por meio do CFC (Centro de Formação de Condutores) a qual iniciei meu processo para a primeira habilitação. Além claro, de banco de questões públicas alinhadas ao formato do exame do Detran.

As questões possuem quatro alternativas cada e podem ou não trazer imagens¹ de apoio tais como placas de sinalização. 

¹Dada a limitação de apresetação de imagens diretamente no terminal foi utilizado a representação em [Braille Unicode](https://unicode.org/charts/nameslist/c_2800.html) quando possível.

## Opções (`help`)

```
Estude para o exame teórico do Detran por meio do Prova Detran,
        uma ferramenta de linha de comando que gera simulados da prova 
        a partir de questões reais e para todos os conteúdos abordados.
        Documentação completa está disponível em github.com/jefersonf/prova-detran

Usage:
  provadetran [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  resultados  Mostra o resultado do simulado
  simulado    Inicia um novo simulado
  version     Imprime a versão do Prova Detran

Flags:
  -h, --help   help for provadetran
```

## Mini-FAQ

>As questões seguem o formato do exame oficial?

>**Sim**
---
>As questões abordam todos os temas cobrados?

>**Sim**
---
>É possível visualizar as placas de sinalização que os enunciados fazem referência?

>**Depende! Nesses casos uma representação gráfica aproximada é utilizada no próprio terminal, devido as limitações comuns a uma CLI**

## Como contruir

É possível contruir de duas formas:

1. Adicionando questões ao banco de questões da CLI. (Checar o formato do [banco de questões](#banco-de-questoes))
2. Melhorando e/ou evoluindo o código! Clone e crie uma issue, PRs são bem-vindos (_There's always room for improvements_).
3. Divulgue! Se é brasileiro sabe bem como é "peso" tirar a primeira habilitação, a tão cara CNH. Criei essa ferramenta para estudar a legislação de Trânsito enquanto praticava outra coisa, o desenvolvimento de CLIs em Go. Espero que seja útil se te serve para ao menos um dos casos.  