# ⚛️ Forge OS — The Reactor

> "Where raw events meet high-velocity execution. The pulse of the Forge."

---

## 🔥 What is The Reactor?

If **Forge OS Backend** (NestJS) is the **Architect** — designing the blueprints and managing the sacred structures of thought — then **The Reactor** (Go) is the **Engine**.

The Reactor is a high-performance, event-driven service designed to handle the "heavy lifting" of the Forge. It doesn't sleep, it doesn't hesitate. It listens to the stream of consciousness (**Redis Streams**) and transmutes raw data into real-time reality.

It is responsible for:
- **Instant Reflexes:** Handling high-concurrency WebSockets for the Forge Chamber.
- **Continuous Evolution:** Calculating Gamification logic (XP, Levels) at sub-millisecond speeds.
- **The Chronicler:** Sinking massive amounts of Audit Logs without breaking a sweat.
- **Presence:** Tracking the cosmic movement of the user across the system.

---

## 🏗 The Architecture of Power

The Reactor is built with **Go**, chosen for its brutal efficiency and fearless concurrency. It follows a **Clean Architecture** to ensure the core logic remains untainted by external drivers.

### The Flow (Event-Driven)
1. **NestJS (Producer):** Captures intent and pushes events to Redis Streams.
2. **Redis Streams (The Forge):** Acts as the persistent nervous system.
3. **The Reactor (Consumer):** Processes events via Goroutines and emits impacts.

### Structure Overview
```text
trhgatu-forge-os-reactor/
├── cmd/                # Entry points (The Spark)
├── internal/
│   ├── sensors/        # Stream Consumers (Listening to the Forge)
│   ├── transmuters/    # Core Logic (The Alchemy: XP, Logs, Stats)
│   ├── emitters/       # Outbound (Websocket, Notifications)
│   └── domain/         # Pure Entities & Value Objects
├── pkg/                # Shared Utilities (DB Drivers, Logger)
└── docker-compose.yml
```

---

## 🧪 Core Subsystems (The Transmuters)

### ⚡ Gamification Engine
Processes every action within the Forge. Whether it's a journal entry or a project sync, the Reactor calculates the earned XP and triggers Level-Up events instantly.

### 📡 Presence Radar
Tracks active sessions and "pulse" signals across the Forge OS, providing the real-time heartbeat of your digital extension.

### 📜 High-Velocity Logger
A dedicated pipeline for system-wide audit logs, ensuring every "mutation of reality" is recorded without impacting the performance of the main API.

---

## 🛠 Tech Stack

- **Go (Golang)** - The core material.
- **Redis Streams** - The nervous system.
- **Goroutines** - The microscopic workers.
- **MongoDB** - The long-term memory.
- **Socket.io / NATS** - For real-time emission.

---

## 🌠 Role in the Ecosystem

The Reactor is not a standalone app. It is a **Kinetic Extension**. It exists to ensure that as your mind and data grow, the system remains fast, fluid, and responsive.

**NestJS manages the Truth. The Reactor executes the Impact.**

---

## 📜 License

Private • Personal OS • Built for trhgatu