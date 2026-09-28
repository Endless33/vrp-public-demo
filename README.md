VRP Public Demo

VRP Public Demo is a standalone demonstration application for the public architecture of VRP (Veil Routing Protocol).

Its purpose is to allow engineers, researchers, and organizations to observe publicly verifiable runtime behavior without requiring access to the protected runtime implementation.

---

What this demonstrates

The demo focuses on observable protocol behavior, including:

- Session establishment
- Session continuity
- Transport migration
- Replay rejection
- Stale-state rejection
- Authority validation
- Deterministic runtime behavior
- Public evidence generation

The demonstration is intentionally limited to publicly documented architectural behavior.

Protected runtime implementation details are not included.

---

Engineering Principles

VRP is built around one architectural principle:

«SESSION ≠ TRANSPORT»

The logical session remains canonical.

Transport becomes replaceable.

Observable behavior should remain correct while network conditions change.

---

Project Scope

This repository is intended for:

- engineering demonstrations
- public evaluation
- architectural understanding
- independent verification
- reproducible runtime behavior

It is not the protected VRP runtime.

---

Public Documentation

VRP Specification

https://github.com/Endless33/VRP-specification

Runtime Boundary Preview

https://github.com/Endless33/vrp-runtime-boundary-preview

---

Status

Active engineering development.

Public demonstration project.

---

Architecture is public.

Protected runtime implementation remains private.