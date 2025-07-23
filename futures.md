## futures table


| Feature                 | How to Build                              |
| ----------------------- | ----------------------------------------- |
| **Multiple Partitions** | Hash key → partition index                |
| **Replication**         | Write to multiple logs, simulate brokers  |
| **Consumer Offsets**    | Track last read message in file or memory |
| **Durability & Flush**  | `flush()` after writes                    |
| **Batched Reads**       | Read N messages at a time                 |
| **Acknowledgments**     | Broker sends `ACK` to producer            |
| **Compression**         | Optional gzip/zlib before writing         |

