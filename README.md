## franz

making a distributed event streaming platform from scratch in Python.

### Main components that I'm targeting to build

- Broker
  Server that stores messages by topic and partition.

- Producer
  Sends messages to a topic/partition on the broker.

- Consumer
  Pulls messages from topic/partition with offset tracking.

- Topic
  A named stream of messages.

- Partition
  Subset of topic; gurantees ordering.

- Offset
  The position of a message in a partition.

- Commit Log
  Persistent append-only log per partition.
  
### Flow and main architecture

```
franz/
├── broker.py        # Manages topics, partitions, and persistence
├── producer.py      # Client that sends messages to the broker
├── consumer.py      # Client that reads messages with offset tracking
├── message.py       # Defines message format
├── storage.py       # Append-only file-based log
├── server.py        # Starts broker and handles socket clients
└── client_utils.py  # Shared socket protocol helpers
```
