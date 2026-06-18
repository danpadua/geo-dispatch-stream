# geo-dispatch-stream 🚀 🛵

Um laboratório prático (Proof of Concept) desenvolvido para simular a ingestão e o processamento de fluxos contínuos de dados (Stream Processing) de geolocalização em tempo real.

O projeto foi inspirado nos desafios de engenharia de grandes plataformas de logística e _delivery_ (como o iFood), simulando o rastreio de múltiplos estafetas em simultâneo e a computação de eventos de proximidade diretamente em memória.

Este projeto também serviu como um desafio pessoal para explorar o ecossistema da linguagem **Go (Golang)** e do **Apache Kafka**, utilizando Inteligência Artificial como ferramenta de aceleração de aprendizagem de novas tecnologias.

---

## 🏗️ Como Funciona a Arquitetura?

O sistema baseia-se numa **Arquitetura Orientada a Eventos (EDA)** e divide-se em três componentes principais localizadas no mesmo binário graças à concorrência nativa do Go:

1. **Simulador de Telemetria (Producers):** Utiliza _Goroutines_ em background para simular múltiplos estafetas (`MOTO-01`, `MOTO-02`, `MOTO-03`). Cada um gera e envia coordenadas de GPS dinâmicas (JSON) para o Apache Kafka a cada segundo.
2. **Mensajaria (Broker):** O **Apache Kafka** atua como uma esteira distribuída de alta performance, recebendo e ordenando o fluxo massivo de dados sem criar gargalos.
3. **Processador de Stream (Consumer):** Um motor que consome as mensagens do Kafka em tempo real, analisa a latitude/longitude de cada estafeta e dispara um alerta visual sempre que um motorista entra no raio de proximidade configurado (Geofencing).

---

## 🛠️ Tecnologias Utilizadas

- **Go (Golang):** Escolhida pela sua excelente performance e gestão nativa de concorrência (_Goroutines_).
- **Apache Kafka:** Plataforma de streaming de eventos distribuídos para ingestão de dados de alta frequência.
- **Docker & Docker Compose:** Para isolar e subir o ambiente do Kafka localmente de forma rápida.

---

## 🚀 Como Executar o Projeto Localmente

### Pré-requisitos

Antes de começar, certifique-se de que tem instalado na sua máquina:

- [Docker](https://www.docker.com/) e Docker Compose.
- [Go](https://go.dev/) (versão 1.20 ou superior).

### Passo a Passo

**1. Clonar o repositório:**

```bash
git clone [https://github.com/TEU_UTILIZADOR/geo-dispatch-stream.git](https://github.com/TEU_UTILIZADOR/geo-dispatch-stream.git)
cd geo-dispatch-stream
```
