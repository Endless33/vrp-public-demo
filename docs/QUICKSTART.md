Quick Start

This guide describes the fastest way to evaluate the public VRP demonstration.

No knowledge of the protected runtime implementation is required.

---

Goal

The purpose of the demo is to observe publicly verifiable protocol behavior.

The demonstration focuses on runtime behavior rather than implementation details.

---

Expected Workflow

1. Download the latest VRP Public Demo release.
2. Launch the demonstration application.
3. Run the default verification scenario.
4. Observe the runtime events.
5. Review the final verification verdict.
6. Export the public evidence report if desired.

---

What You Should Observe

Depending on the selected scenario, the demonstration may include:

- Session establishment
- Transport migration
- Session continuity
- Replay rejection
- Stale-state rejection
- Authority validation
- Deterministic runtime behavior
- Final verification verdict

The exact runtime sequence may vary between demonstration scenarios while preserving the same architectural guarantees.

---

Public Verification

The demonstration is intended to validate observable behavior only.

It does not expose:

- protected runtime algorithms
- proprietary implementation details
- internal protocol mechanisms
- protected security logic

---

Additional Documentation

VRP Specification

https://github.com/Endless33/VRP-specification

Runtime Boundary Preview

https://github.com/Endless33/vrp-runtime-boundary-preview

---

Current Status

The public demonstration is under active engineering development.

This document will be updated as new demonstration scenarios become available.