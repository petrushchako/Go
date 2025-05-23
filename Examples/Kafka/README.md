# Go Kafka producer and consumer example

## **Step 1: Set up Kafka using Docker**
The easiest way to run Kafka locally for development is using Docker Compose.
1.  **Create a `docker-compose.yml` file:** In a directory of your choice, create a file named `docker-compose.yml` with the following content:

    ```yaml
    version: '3.8'
    services:
      zookeeper:
        image: 'bitnami/zookeeper:latest'
        ports:
          - '2181:2181'
        environment:
          - ALLOW_ANONYMOUS_LOGIN=yes
      kafka:
        image: 'bitnami/kafka:latest'
        ports:
          - '9092:9092'
          - '9094:9094'
        environment:
          - KAFKA_BROKER_ID=1
          - KAFKA_LISTENERS=PLAINTEXT://0.0.0.0:9092,CONTROLLER://0.0.0.0:9094
          - KAFKA_ADVERTISED_LISTENERS=PLAINTEXT://localhost:9092
          - KAFKA_ZOOKEEPER_CONNECT=zookeeper:2181
          - KAFKA_CONTROLLER_QUORUM_VOTERS=1@kafka:9094
          - ALLOW_PLAINTEXT_LISTENER=yes
        depends_on:
          - zookeeper
    ```

2.  **Start Kafka:** Open your terminal in the same directory where you saved `docker-compose.yml` and run the command:

    ```bash
    docker-compose up -d
    ```

    This command will download the Zookeeper and Kafka images (if you don't have them already) and start them in detached mode (`-d`).

3.  **Verify Kafka is running:** You can check if the containers are running using:

    ```bash
    docker ps
    ```

    You should see containers named `your_directory_name_zookeeper_1` and `your_directory_name_kafka_1` (where `your_directory_name` is the name of the directory you created).

<br><br><br>

## **Step 2: Write a Go script to produce messages to Kafka**
1.  **Create a Go project directory:** Create a new directory for your Go project (e.g., `kafka-go-producer`).

2.  **Initialize Go modules (optional but recommended):** Navigate into your project directory in the terminal and run:

    ```bash
    go mod init github.com/yourusername/kafka-go-producer
    ```

    Replace `github.com/yourusername/kafka-go-producer` with your actual repository path or a descriptive name.

3.  **Install the Go Kafka client library:** We'll use the `confluent-kafka-go` library, which is a popular and well-maintained client. In your project directory, run:

    ```bash
    go get github.com/confluentinc/confluent-kafka-go/kafka
    ```

4.  **Create the producer Go file (`producer.go`):** Create a file named `producer.go` in your project directory with the following code:

    ```go
    package main

    import (
        "fmt"
        "log"
        "os"

        "github.com/confluentinc/confluent-kafka-go/kafka"
    )

    func main() {
        broker := "localhost:9092"
        topic := "my-topic"

        p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": broker})
        if err != nil {
            log.Fatalf("Failed to create producer: %s\n", err)
            os.Exit(1)
        }

        defer p.Close()

        // Delivery report handler for produced messages
        go func() {
            for e := range p.Events() {
                switch ev := e.(type) {
                case *kafka.Message:
                    if ev.TopicPartition.Error != nil {
                        fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
                    } else {
                        fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
                    }
                }
            }
        }()

        for i := 0; i < 10; i++ {
            message := fmt.Sprintf("Hello Kafka message #%d from Go!", i)
            err = p.Produce(&kafka.Message{
                TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
                Value:          []byte(message),
            }, nil)
            if err != nil {
                log.Printf("Failed to produce message: %v\n", err)
            }

            // To ensure messages are delivered, you might want to flush the producer
            p.Flush(100) // Wait for outstanding messages to be delivered (in milliseconds)
        }

        fmt.Println("Producer finished sending messages.")
    }
    ```

5.  **Run the producer script:** In your terminal, within the `kafka-go-producer` directory, run:

    ```bash
    go run producer.go
    ```

    You should see output indicating that messages are being delivered to the `my-topic`.

<br><br><br>

## **Step 3: Write a Go script to consume messages from Kafka**
Now, let's create a Go script to consume the messages we just produced.

1.  **Create the consumer Go file (`consumer.go`):** In your project directory, create a file named `consumer.go` with the following code:

    ```go
    package main

    import (
        "fmt"
        "log"
        "os"
        "os/signal"
        "syscall"

        "github.com/confluentinc/confluent-kafka-go/kafka"
    )

    func main() {
        broker := "localhost:9092"
        topic := "my-topic"
        group := "my-consumer-group"

        c, err := kafka.NewConsumer(&kafka.ConfigMap{
            "bootstrap.servers": broker,
            "group.id":          group,
            "auto.offset.reset": "earliest", // Start consuming from the beginning of the topic
        })

        if err != nil {
            log.Fatalf("Failed to create consumer: %s\n", err)
            os.Exit(1)
        }

        defer c.Close()

        err = c.SubscribeTopics([]string{topic}, nil)
        if err != nil {
            log.Fatalf("Failed to subscribe to topic %s: %v\n", topic, err)
            os.Exit(1)
        }

        sigchan := make(chan os.Signal, 1)
        signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

        fmt.Println("Consumer started. Press Ctrl+C to exit.")

        for {
            select {
            case sig := <-sigchan:
                fmt.Printf("Caught signal %v: terminating\n", sig)
                return
            default:
                msg, err := c.ReadMessage(100 * time.Millisecond) // Timeout after 100ms
                if err == nil {
                    fmt.Printf("Received message on %s [%d] at offset %v: %s\n",
                        *msg.TopicPartition.Topic, msg.TopicPartition.Partition, msg.TopicPartition.Offset, string(msg.Value))
                } else if !err.(kafka.Error).IsTimeout() {
                    // The client will automatically try to recover from all errors.
                    fmt.Printf("Consumer error: %v (%v)\n", err, msg)
                }
            }
        }
    }
    ```

6.  **Run the consumer script:** In your terminal, within the `kafka-go-producer` directory, run:

    ```bash
    go run consumer.go
    ```

    You should see the consumer start receiving the messages that the producer sent. You can run multiple instances of the consumer to see how Kafka handles consumer groups.

<br><br><br>

## **Step 4: Clean up (optional)**
When you're finished practicing, you can stop and remove the Docker containers:

```bash
docker-compose down
```

**Important Considerations:**
* **Error Handling:** The provided scripts have basic error handling. In a production environment, you'll need more robust error handling and logging.
* **Configuration:** You might need to adjust the Kafka broker address and topic name in the Go scripts if your Docker setup is different.
* **Asynchronous Production:** The producer example uses a delivery report handler, which is crucial for understanding if your messages were successfully delivered. Production in real-world applications is often asynchronous.
* **Consumer Groups:** The consumer example uses a consumer group (`my-consumer-group`). Kafka uses consumer groups to allow multiple consumers to read from a topic without duplicating messages (each message is delivered to one consumer within a group).
* **Serialization:** For more complex data structures, you'll likely want to use a serialization format like JSON, Protocol Buffers, or Avro and serialize/deserialize your Go structs before sending/after receiving messages.