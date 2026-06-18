# geo-dispatch-stream 🚀 🛵

Um laboratório prático (Proof of Concept) desenvolvido para simular a ingestão e o processamento de fluxos contínuos de dados (Stream Processing) de geolocalização em tempo real.

O projeto foi inspirado nos desafios de engenharia de grandes plataformas de logística e _delivery_ (como o iFood), simulando o rastreio de múltiplos estafetas em simultâneo e a computação de eventos de proximidade diretamente em memória.

Este projeto também serviu como um desafio pessoal para explorar o ecossistema da linguagem **Go (Golang)** e do **Apache Kafka**, utilizando Inteligência Artificial como ferramenta de aceleração de aprendizagem de novas tecnologias.

---

## 🏗️ Como Funciona a Arquitetura?

O sistema baseia-se numa **Arquitetura Orientada a Eventos (EDA)** e divide-se em três componentes principais localizadas no mesmo binário graças à concorrência nativa do Go:

1. **Simulador de Telemetria (Producers):** Utiliza _Goroutines_ em background para simular múltiplos estafetas (`ENTREGADOR-01`, `ENTREGADOR-02`, `ENTREGADOR-03`). Cada um gera e envia coordenadas de GPS dinâmicas (JSON) para o Apache Kafka a cada segundo.
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

**2. Iniciar o Apache Kafka (via Docker):**
Execute o comando abaixo para subir o Kafka e o Zookeeper em background:

```bash
docker-compose up -d
```

_(Aguarde cerca de 10 a 15 segundos para que o broker do Kafka inicialize completamente)._

**3. Descarregar as dependências do Go:**

```bash
go mod tidy
```

**4. Executar a aplicação:**

```bash
go run main.go
```

## 📺 Entender os Logs do Terminal

Ao executar o projeto, verá um fluxo contínuo de dados no seu terminal com a seguinte convenção de cores:

- 🟢 Logs em Verde (📤 [ENTREGADOR-XX]): Representam os dados de GPS simulados a serem enviados pelas Goroutines para o tópico do Kafka.

- 🔵 Logs em Azul (📥 [PROCESSADOR]): Representam o motor de processamento a ler os eventos do Kafka em tempo real.

- 🟡 Logs em Amarelo (🔔 [ALERTA]): Representam a regra de negócio a ser executada com sucesso quando um estafeta cruza a barreira geográfica definida.

## 📝 Licença

Este projeto está sob a licença MIT. Sinta-se à vontade para utilizar, partilhar e evoluir o código!

Desenvolvido com 🧠, ⚡ e a ajuda de Inteligência Artificial para acelerar o processo de engenharia.
