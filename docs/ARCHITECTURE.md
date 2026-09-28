# VRP Runtime Boundary Architecture

**Status:** Public

## Purpose

VRP Runtime Boundary provides a stable public interface between the protected VRP runtime and external applications.

It exposes observable behavior only.

The protected runtime implementation remains private.

---

## Design Goals

- Stable public API
- Observable runtime behavior
- No protected runtime disclosure
- Reproducible engineering demonstrations
- Independent integration surface

---

## Repository Structure

```
Protected Runtime
        │
        ▼
VRP Runtime Boundary
        │
        ▼
Public Applications
```

The Runtime Boundary is the only supported public integration point.

Applications should depend on the Runtime Boundary instead of protected runtime code.

---

## Current Public API

Current public capabilities include:

- Session creation
- Transport creation
- Transport switching
- Evidence generation

These operations intentionally demonstrate observable protocol behavior without exposing protected implementation details.

---

## Design Principle

```
SESSION ≠ TRANSPORT
```

Logical session identity is independent from the underlying transport.

Transport replacement preserves logical session continuity.

---

## Current Demonstration

The public demo demonstrates:

1. Session creation
2. Transport attachment
3. Transport migration
4. Evidence generation
5. Continuity preserved

---

## Future Public API

Planned additions include:

- Recovery lifecycle
- Replay rejection
- Evidence export
- Scenario execution
- Public verification helpers

These features will remain implementation-independent and expose only observable behavior.