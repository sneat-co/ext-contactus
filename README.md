# contactus-ext

Public **contract surface** for the `contactus` extension — the first repo to follow the
[`extension-contract-repo`](https://github.com/sneat-co/sneat-libs/blob/main/spec/features/extension-contract-repo/README.md)
convention.

It holds only what other extensions, the apps, and the `contactus` main repo need in order to
*talk to* contactus — facade interfaces, DTOs, model/const shapes, and the callback signatures
contactus asks callers to satisfy. It contains **no** contactus implementation.

## Layout

```
contactus-ext/
├── backend/    # Go module github.com/sneat-co/contactus-ext/backend
└── frontend/   # nx lib published as @sneat/extension-contactus-contract
```

## The load-bearing invariant

`contactus-ext` depends **only on foundational/core code — never on another extension.** Because
it has no edge back to any sibling, `sibling → contactus-ext` can never form a dependency cycle,
and (frontend) importing it never triggers the prebuilt-bundle peer-resolution wall.

An interface or type belongs here **only if its entire signature is expressible in contactus-own +
foundational/core types**. If a signature references a *consumer's* types, that interface is the
consumer's contract and stays consumer-owned (e.g. `ContactusAccess`, which speaks invitus/spaceus
types, deliberately is **not** here). The CI check in `.github/workflows/ci.yml` enforces the
invariant.

## Status

Reference extraction in progress — tracked by `sneat-co/sneat-libs` Plan `contactus-ext`. The
backend module, repo metadata, and the invariant CI check are scaffolded here; the frontend
contract lib and the backend contract packages are migrated in by the later plan tasks.
