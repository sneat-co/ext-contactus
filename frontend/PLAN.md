# Frontend contract lib — `@sneat/extension-contactus-contract`

This repo (`contactus-ext`) is the sanctioned home of the contactus **contract**
surface (per the `extension-contract-repo` convention). The **backend** contract
is already here; the **frontend** contract lib still needs to be finished.

## Status

- ✅ Backend contract (Go): `backend/` (`contactusmodels`, `facade4contactus`).
- 🚧 Frontend contract (Angular): `frontend/libs/contactus-contract/` — **started
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

## TODO to finish the extraction (separate, larger step)

1. **Bootstrap the nx workspace** here (`frontend/` currently has no `nx.json` /
   root `package.json`) — mirror `assetus/frontend` / `contactus/frontend`
   (`ng-package.json`, `project.json`, `tsconfig.*`, publishable target).
2. **Migrate the full published `0.12.1` API** into this lib — the DTOs/types,
   enums (`ContactType`, `ContactRoleDriver`, …) and the existing tokens
   (`CONTACT_SERVICE`, `CONTACTUS_SPACE_SERVICE`, `CONTACT_NAV_SERVICE`, …) and
   helpers (`addSpace`, `filterContactsByTextAndRole`, `validateContactRequest`, …).
   The package currently ships only `.d.ts` + compiled JS (no `.ts` source and no
   `repository` link), so this needs the original source or a `.d.ts`-guided
   reconstruction (helpers re-implemented).
3. **Bind the new tokens in contactus `-shared`**: `{ provide: CONTACTS_SELECTOR,
   useExisting: ContactsSelectorService }` (adapter mapping `IContactsSelectorProps`
   → the existing `IContactSelectorOptions.componentProps`), and register the
   contacts-list component as the `sneat-contacts-list` custom element.
4. **Publish** the new contract + shared versions via the repo release workflow,
   then bump consumers (requoter injects `CONTACTS_SELECTOR` / uses the tag).
