
# Projeto - Sistema McRonald's AP2 - Atualização do código


*******

Sumário Executivo 
 1. [Partes Envolvidas](#partesenvolvidas)
 2. [Contexto do Problema](#contextoproblema)
 3. [Requisitos](#requisitos)
 4. [Tecnologias e Frameworks Utilizadas](#tecnologia)
 5. [Considerações Finais](#final)
 6. [Referências e Documentação Suporte de Apoio](#ref)

*******

<div id='partesenvolvidas'/> 

## Partes Envolvidas

### Cliente: IBMEC

Partes envolvidas: 

- Victor - Professor da Disciplina de Estrutura de Dados (IBMEC)

### Squad: Alunos IBMEC

Partes envolvidas: 

- Daniel Gripa;

- Leonardo Costa;

- Thalles Diepes;


<div id='contextoproblema'/> 

## Contexto do Problema

O McRonald’s é um food truck em franco crescimento no último ano. Foi fundado por dois amigos apaixonados por boa comida e interessados em trazer uma nova opção de gastronomia no bairro em que vivem.

Atualmente, a empresa já possui um sistema que processa os pedidos, mas esse sistema contém 2 bugs e precisa de alguns outros aprimoramentos.

<div id='requisitos'/>

## Requisitos

Requisitos funcionais
  O programa deve incluir a opção de atualizar um produto, modificando o seu preço. O id, o nome e a descrição devem ser mantidos;
  O programa deve incluir uma nova métrica, chamada ticket médio, que calcula o valor médio de cada pedido (total faturado até o momento, dividido pelo número de pedidos já encerrados);
  O programa deve incluir uma segunda forma de exibir os produtos, ordenados por nome, e não por id;

Requisitos não funcionais
  A estrutura de dados que forma a lista de produtos deve ser uma transformada em uma lista simplesmente encadeada. Todas as operações sobre essa lista (adicionar, buscar, excluir, atualizar, exibir etc.) devem ser refatorados para considerar uma lista encadeada;
  Não há restrição do algoritmo de ordenação utilizado para a exibição dos produtos ordenados pelo nome (pode ser o bubblesort, por exemplo).
  
Bugs encontrados
  *Bug 1*
      1.Selecione a opção 6 (adicionar pedido);
      2.Insira valores quaisquer para um pedido (p.ex., s, depois 1 1 e depois 0 0);
      3.Selecione a opção 7 (expedir pedido);
      4.Selecione a opção 7 (expedir pedido).
      O programa cai num panic e não consegue resolver essa operação. O comportamento desse cenário deve ser idêntico ao caso em que se insere a opção 7 (expedir pedido) assim que o programa é aberto.
  *Bug 2*
      1.Abra o programa, utilizando o arquivo de dados (dados.csv) fornecido;
      2.Selecione a opção 5 (exibir os produtos);
      3.Você verá três produtos criados;
      4.Selecione a opção 1 (cadastrar novo produto);
      5.Insira quaisquer valores para nome, descrição e preço;
      6.Selecione a opção 5 (exibir os produtos).
O programa lista todos os produtos, porém o id do produto criado manualmente é 1, e não 4. Ou seja, o programa não está considerando os ids dos produtos pré-carregados.

<div id='tecnologia'/>

## Tecnologias e Frameworks Utilizadas:

Colocamos abaixo, as especificações de cada tecnologia utilizada durante a fase de desenvolvimento do projeto:

Especificação:

 * **Linguagem de Programação** : Golang.

 * **IDE** : GoLand da JetBrains.

 * **Sistema** : Git & GitHub.

<div id='ref'/>
