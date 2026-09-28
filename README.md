# VRP Public Demo

**VRP Public Demo** is a standalone engineering demonstration for the public architecture of **VRP (Veil Routing Protocol)**.

Its purpose is to allow engineers, researchers, organizations, and independent evaluators to observe publicly verifiable runtime behavior without requiring access to the protected VRP runtime implementation.

The demonstration focuses exclusively on publicly documented architectural behavior.

The protected runtime implementation is **not included**.

---

# What this demonstrates

The current public demonstration includes:

- Session establishment
- Session continuity
- Transport migration
- Session recovery
- Replay rejection
- Stale-state rejection
- Authority validation
- Runtime Boundary API
- Interactive demonstration menu
- CLI scenario execution
- Public evidence generation
- Deterministic observable runtime behavior

---

# Engineering Principle

VRP is built around one architectural principle:

```
SESSION ≠ TRANSPORT
```

The logical session remains canonical.

Underlying transports may change without changing logical session identity.

Observable behavior remains stable while network conditions change.

---

# Repository Purpose

This repository is intended for:

- Engineering demonstrations
- Public evaluation
- Architectural understanding
- Independent verification
- Runtime Boundary evaluation
- Reproducible public demonstrations

This repository is **not** the protected VRP runtime.

---

# Repository Structure

```
cmd/
    vrp-public-demo/

internal/
    evidence/

docs/

evidence/
```

---

# Requirements

- Go 1.25 or newer
- Linux, macOS or Windows
- Git

Verify your Go installation:

```bash
go version
```

---

# Installation

Clone the repository:

```bash
git clone https://github.com/Endless33/vrp-public-demo.git

cd vrp-public-demo
```

Download dependencies:

```bash
go mod tidy
```

Verify the project:

```bash
go fmt ./...

go test ./...

go vet ./...
```

---

# Running the Demo

## Interactive Mode

Start the interactive demonstration:

```bash
go run ./cmd/vrp-public-demo
```

You'll see:

```
======================================
VRP PUBLIC DEMO
======================================

1. Session Establishment
2. Transport Migration
3. Session Recovery
4. Replay Rejection
5. Stale-State Rejection
6. Authority Validation
7. Full Demonstration
0. Exit
```

Simply enter the number of the scenario you want to execute.

---

## Command Line Mode

Individual scenarios can also be executed directly.

Session Establishment

```bash
go run ./cmd/vrp-public-demo session
```

Transport Migration

```bash
go run ./cmd/vrp-public-demo migration
```

Session Recovery

```bash
go run ./cmd/vrp-public-demo recovery
```

Replay Rejection

```bash
go run ./cmd/vrp-public-demo replay
```

Stale-State Rejection

```bash
go run ./cmd/vrp-public-demo stale
```

Authority Validation

```bash
go run ./cmd/vrp-public-demo authority
```

Run every demonstration sequentially

```bash
go run ./cmd/vrp-public-demo full
```

Display help

```bash
go run ./cmd/vrp-public-demo help
```

---

# Generated Evidence

Some demonstrations generate a public engineering evidence report.

Example:

```
evidence/
└── transport-migration.json
```

Example report:

```json
{
  "version": "v0.1.0",
  "scenario": "Transport Migration",
  "verdict": "PASS",
  "principle": "SESSION ≠ TRANSPORT",
  "session_id": "demo-session",
  "transport": "udp:B",
  "timestamp": "..."
}
```

---

# Runtime Boundary

The demo communicates exclusively through the public Runtime Boundary API.

Protected runtime implementation remains outside the public interface.

---

# Public Documentation

## VRP Specification

https://github.com/Endless33/VRP-specification

## Runtime Boundary Preview

https://github.com/Endless33/vrp-runtime-boundary-preview

---

# Current Public Scenarios

- Session Establishment
- Transport Migration
- Session Recovery
- Replay Rejection
- Stale-State Rejection
- Authority Validation
- Full Demonstration

---

# Status

**Status:** Active Engineering Development

Current public release includes:

- Runtime Boundary API
- Interactive Public Demo
- Scenario CLI
- Transport Migration
- Session Recovery
- Replay Rejection
- Stale-State Rejection
- Authority Validation
- Public JSON Evidence Export

---

# Design Philosophy

The architecture is public.

Observable behavior is public.

Engineering evidence is public.

Protected runtime implementation remains private.