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
- Runtime Boundary API
- Public evidence generation
- Deterministic observable runtime behavior

Future public demonstrations will include:

- Session recovery
- Replay rejection
- Stale-state rejection
- Authority validation
- Multi-scenario execution
- Public evidence reports

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

# Current Demonstration

The current demo demonstrates:

1. Create logical session
2. Attach transport
3. Migrate transport
4. Preserve logical session
5. Generate public evidence
6. Export JSON evidence

Example output:

```
======================================
VRP Runtime Boundary
======================================

Version: v0.1.0
Principle: SESSION ≠ TRANSPORT

Session: demo-session
State: ACTIVE
Transport: udp:A

----- TRANSPORT MIGRATION -----

Transport switched
Active transport: udp:B

Evidence
Scenario: Transport Migration
Verdict : PASS

FINAL VERDICT
CONTINUITY PRESERVED
```

---

# Building

Clone the repository:

```bash
git clone https://github.com/Endless33/vrp-public-demo.git

cd vrp-public-demo
```

Build:

```bash
go build ./cmd/vrp-public-demo
```

Run:

```bash
go run ./cmd/vrp-public-demo
```

or

```bash
go run ./cmd/vrp-public-demo migration
```

---

# Generated Evidence

Running the demo automatically generates a public engineering evidence report.

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

# Roadmap

Planned public demonstrations include:

- Session Recovery
- Replay Rejection
- Authority Validation
- Stale-State Rejection
- Scenario CLI
- Multiple Evidence Reports
- Automated Demonstration Suite

---

# Status

**Status:** Active Engineering Development

Current public release:

- Runtime Boundary API
- Public Demo
- Transport Migration Demonstration
- JSON Evidence Export

Additional public demonstrations are under active development.

---

# Design Philosophy

The architecture is public.

Observable behavior is public.

Engineering evidence is public.

Protected runtime implementation remains private.