# Frontend contract lib — `@sneat/extension-contactus-contract`

This repo (`ext-contactus`) is the sanctioned home of the contactus **contract**
surface (per the `extension-contract-repo` convention). The **backend** contract
is already here; the **frontend** contract lib still needs to be finished.

## Status

- ✅ Backend contract (Go): `backend/` (`contactusmodels`, `facade4contactus`).
- 🚧 Frontend contract (Angular): `frontend/libs/extensions/contactus/contract/` — **started
  in this PR** with the two cross-extension surfaces that were missing and were
  blocking a clean consumer integration (requoter picking drivers):
  - **`contacts-selector.contract.ts`** — `CONTACTS_SELECTOR` `InjectionToken` +
    `IContactsSelectorService` + `IContactsSelectorProps`. The service-capability
    pattern: a sibling extension injects the token; contactus `-shared`'s
    `ContactsSelectorService` is bound to it by the host app.
  - **`contacts-list.contract.ts`** — `CONTACTUS_CONTACTS_LIST_TAG` +
    `IContactsListProps`. The **component** pattern: the contacts-list component
    (authored in `-shared`) is registered as a custom element; consumers use the
    tag with contract-typed props — no impl import.
  - **`contact-ref.ts`** — a minimal, foundational `IContactRef` used by both.

## Why these two patterns

Cross-extension reuse must stay **contract-only** (an extension never imports
another extension's `-shared`/`-internal`). Two shapes cover everything:

| Need | Contract exposes | Impl lives in | Wired by |
|---|---|---|---|
| A **service/capability** (e.g. open a picker) | `InjectionToken` + interface | `-shared`/`-internal` | host app binds token → impl |
| A **UI component** (e.g. contacts-list) | custom-element **tag** const + **props** interface | `-shared` (registered as a custom element) | host app registers the element |

Generic, non-contactus components belong in a **foundational** lib
(`@sneat/ui` / `@sneat/ng`), which any extension may import directly.

## ✅ Full API recovered

The complete `0.12.1` contract source (42 files — DTOs, contexts, apidto, and the
existing tokens `CONTACT_SERVICE`, `CONTACTUS_SPACE_SERVICE`, `CONTACT_NAV_SERVICE`,
`CONTACTUS_NAV_SERVICE`, `CONTACT_GROUP_SERVICE`, `CONTACT_ROLE_SERVICE`,
`INVITE_SERVICE`, enums like `ContactRoleDriver`, helpers like `addSpace` /
`filterContactsByTextAndRole` / `validateContactRequest`) was **recovered from
`contactus` git history** — it was removed in commit `875cc5b` (the extraction),
so `875cc5b^` still had it — and moved here. So this lib now holds the full API
**plus** the new picker/list cross-extension surfaces. No `.d.ts` reconstruction
needed.

## TODO to finish

1. **Bootstrap the nx workspace** here: done. `frontend/` now has its own
   `package.json`, `nx.json`, `tsconfig.base.json`, and `eslint.config.mjs`, and
   the contract library lives under `libs/extensions/contactus/contract` like the
   other extension workspaces.
2. **Runtime provider binds ALL tokens.** The implementation lib's
   `provideContactus()` MUST register the concrete impl for **every**
   contract token in one place — including the new **`CONTACTS_SELECTOR`**
   (`{ provide: CONTACTS_SELECTOR, useClass: ContactsSelectorAdapter }`, where the
   adapter maps `IContactsSelectorProps` → the existing shared
   `IContactSelectorOptions.componentProps` and `IContactWithBriefAndSpace` →
   `IContactRef`). Registering all tokens through the single `provide…Internal()`
   is the platform convention — the host app calls it once and every capability is
   wired. The contacts-list component (`-shared`) is registered as the
   `sneat-contacts-list` custom element in the same wiring.
3. **Publish** the new contract + shared versions via the repo release workflow,
   then bump consumers (requoter injects `CONTACTS_SELECTOR` / uses the tag).
