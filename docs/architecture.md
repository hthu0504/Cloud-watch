# Architecture

```text
Server Simulator
       ↓
    Go API
       ↓
     Kafka
       ↓
    Workers
       ↓
 ┌─────┴─────┐
 ↓           ↓
DynamoDB     S3
       ↓
   React UI