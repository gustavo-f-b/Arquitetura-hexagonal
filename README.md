
# Arquitetura Hexagonal (Ports and Adapters)

  

![Ports & Adapters](https://codesoapbox.dev/wp-content/uploads/2022/08/hexagonal_simple_2.svg)

  

## Resumo

-  **Domínio da Aplicação:** É o núcleo central do sistema, contendo toda a lógica de negócios. Não deve ter dependências externas.

-  **Ports (Portas):** São interfaces que definem os contratos para a comunicação entre o domínio e o mundo exterior. Existem portas primárias (usadas para expor a funcionalidade do domínio) e portas secundárias (usadas para interagir com serviços externos).

-  **Adapters (Adaptadores):** São implementações concretas dessas interfaces. Eles "adaptam" a comunicação entre o núcleo da aplicação e o mundo exterior. Por exemplo, você pode ter um adaptador que implementa a lógica para salvar dados em um banco de dados específico ou para expor uma API REST.

-  **Direção das Dependências:** Um aspecto crucial dessa arquitetura é que as dependências sempre apontam para dentro. Ou seja, o domínio não depende de nada externo, enquanto os adaptadores dependem das portas (interfaces) definidas pelo domínio.

  

A grande vantagem do uso da arquitetura *Ports and Adapters* é que ela permite uma separação clara entre a lógica de negócios e as tecnologias externas. Isso facilita a troca de componentes externos sem afetar o núcleo da aplicação. Além disso, torna a aplicação mais testável, pois o domínio pode ser testado isoladamente, sem necessidade de integrações externas.

  

>  **Referências**
>
> O projeto foi desenvolvido durante as aulas da [Full Cycle](https://fullcycle.com.br), abaixo estão outras referências que li para poder entender um pouco mais sobre.
>
>  [Ports & Adapters (aka hexagonal) architecture explained](https://codesoapbox.dev/ports-adapters-aka-hexagonal-architecture-explained/)
> [Ports & Adapters Architecture](https://herbertograca.com/2017/09/14/ports-adapters-architecture/)
