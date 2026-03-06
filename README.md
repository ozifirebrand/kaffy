# Fruit Share

A small Go demo that publishes a list of fruit names to Apache Kafka. It acts as a **Kafka producer** and distributes messages across multiple partitions.

## What it does

- Connects to a Kafka broker and sends **~180 fruit names** (from Açaí to Youngberry) to the topic `multi-lane-broadcast`.
- Spreads messages across **3 partitions** in round-robin fashion (partition = `index % 3`).
- Uses the [segmentio/kafka-go](https://github.com/segmentio/kafka-go) client.

## Prerequisites

- **Go** 1.24+ (see `go.mod`)
- A running **Kafka** broker at `localhost:9092` with the topic `multi-lane-broadcast` (or create the topic with 3 partitions)

## Setting up Kafka on Windows

After installing Kafka’s binary archive, move the extracted folder into `C:\` and rename it to `kafka`. Open a terminal and go to the Windows scripts:

```powershell
cd C:\kafka\bin\windows
```

### Start the broker

Generate a cluster ID and format storage (use the UUID from the first command in the second):

```powershell
.\kafka-storage.bat random-uuid
.\kafka-storage.bat format --standalone -t <UUID> -c ..\..\config\server.properties
.\kafka-server-start.bat ..\..\config\server.properties
```

Replace `<UUID>` with the value printed by `random-uuid` (e.g. `W9OEFRq3Rb-8GbpP5lR7rQ`).

### Create the topic (for this project)

In a new terminal, still under `C:\kafka\bin\windows`:

```powershell
.\kafka-topics.bat --create --topic multi-lane-broadcast --partitions 3 --replication-factor 1 --bootstrap-server localhost:9092
```

### Optional: single-topic producer/consumer (broadcast-events)

Create topic:

```powershell
.\kafka-topics.bat --create --topic broadcast-events --bootstrap-server localhost:9092
```

**Producer** (one terminal):

```powershell
.\kafka-console-producer.bat --topic broadcast-events --bootstrap-server localhost:9092
```

**Consumer** (another terminal):

```powershell
.\kafka-console-consumer.bat --topic broadcast-events --from-beginning --bootstrap-server localhost:9092
```

Type in the producer and watch messages appear in the consumer.

### Multi-lane topic (multi-lane-broadcast)

Create the topic (if not already created above):

```powershell
.\kafka-topics.bat --create --topic multi-lane-broadcast --partitions 3 --replication-factor 1 --bootstrap-server localhost:9092
```

**Consumer 1** (terminal 1):

```powershell
.\kafka-console-consumer.bat --topic multi-lane-broadcast --bootstrap-server localhost:9092 --group my-team
```

**Consumer 2** (terminal 2, same consumer group):

```powershell
.\kafka-console-consumer.bat --topic multi-lane-broadcast --bootstrap-server localhost:9092 --group my-team
```

**Producer** (terminal 3; stop the `broadcast-events` producer with Ctrl+C first if it’s running):

```powershell
.\kafka-console-producer.bat --topic multi-lane-broadcast --bootstrap-server localhost:9092
```

Messages will be split across the consumers in the group. You can also run this project (`go run main.go`) to send the fruit list to `multi-lane-broadcast` and see it in these consumers.

## Run

```bash
go run main.go
```

You should see output like:

```
--- PRODUCER: Sending fruits ---
Sent Açaí to Partition 0
Sent Ackee to Partition 1
...
--- All fruits sent successfully ---
```

## Configuration

| Setting          | Value                 |
|------------------|-----------------------|
| Broker           | `localhost:9092`      |
| Topic            | `multi-lane-broadcast`|
| Partition count  | 3                     |

Edit `main.go` to change the broker, topic, or partition count.

## Dependencies

- `github.com/segmentio/kafka-go` — Kafka client for Go

Install with:

```bash
go mod tidy
```
