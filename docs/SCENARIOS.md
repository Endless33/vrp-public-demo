Demonstration Scenarios

This document describes the public demonstration scenarios available in the VRP Public Demo.

The purpose of these scenarios is to demonstrate observable protocol behavior.

Protected runtime implementation details are intentionally excluded.

---

Scenario 1

Session Establishment

Goal

Create a new session and verify successful initialization.

Expected Result

- Session established
- Runtime initialized
- Canonical state created

Final Verdict

PASS

---

Scenario 2

Transport Migration

Goal

Replace the active transport without replacing the logical session.

Expected Result

- Original transport detached
- Replacement transport attached
- Session preserved

Final Verdict

CONTINUITY PRESERVED

---

Scenario 3

Replay Rejection

Goal

Inject a replayed packet.

Expected Result

- Replay detected
- Packet rejected
- Canonical state unchanged

Final Verdict

REPLAY REJECTED

---

Scenario 4

Stale-State Rejection

Goal

Inject stale protocol state.

Expected Result

- Stale state detected
- State rejected
- Canonical authority preserved

Final Verdict

STALE STATE REJECTED

---

Scenario 5

Authority Validation

Goal

Validate canonical authority during runtime.

Expected Result

- Authority verified
- Session remains canonical
- Runtime continues

Final Verdict

AUTHORITY PRESERVED

---

Scenario 6

Complete Demonstration

Goal

Execute all public demonstration scenarios.

Expected Result

Every scenario completes successfully.

Final Verdict

PUBLIC VERIFICATION PASSED

---

Additional scenarios will be added as the public demonstration evolves.