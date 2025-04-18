# TurboQueue

A scalable, production-ready messaging and scheduling service built in Go

## Overview
This project replicates QStash with advanced features like replica queues, topic partitioning, and real-time subscriptions, demonstrating Go concurrency, distributed systems, and production-grade practices.

## Features
- **Replica Queues:** Messages enqueued to multiple Redis instances per topic for fault tolerance.
- **Topic Partitioning:** 3 Redis queues per topic (e.g., `qstash:queue:topic1:p0`) for load distribution.
- **Dead-Letter Queue:** Failed messages stored in PostgreSQL and Redis, retrievable via `/dead-letters`.
- **Retry Mechanism:** Failed deliveries retry 3 times with exponential backoff (1s, 2s, 4s).
- **Scheduling:** Messages can be delayed (e.g., 60 seconds).
- **WebSocket Subscriptions:** Real-time updates via Redis pub/sub.
- **Authentication:** API key required for all endpoints (`X-API-Key`).
- **Metrics:** Prometheus tracks delivery latency and DLQ size.
- **Tests:** Unit tests for HTTP delivery and message processing.

## Architecture
![Architecture Diagram](https://via.placeholder.com/600x300.png?text=Architecture+Diagram)  
*(Replace with Draw.io/Excalidraw diagram)*  
- **API Layer:** Gin handles REST and WebSocket endpoints.
- **Storage:** PostgreSQL (partitioned `messages`, `topics`, `dead_letters`).
- **Queuing:** Redis with replica queues per topic partition.
- **Workers:** Go goroutines process messages concurrently.

## Tech Stack
- **Go:** Core language with goroutines/channels.
- **PostgreSQL:** Durable storage with partitioning.
- **Redis:** Replica queues and pub/sub.
- **Gin:** HTTP framework.
- **Prometheus:** Metrics.
- **Zerolog:** Logging.
- **Docker:** Containerization.

## Setup
1. Clone the repository:
   ```bash
   git clone <repo-url>
   cd TurboQueue