# ext-contactus / frontend

Home of the `@sneat/extension-contactus-contract` Nx library: contactus's
runtime-light public frontend contract surface (DTOs, contexts, enums, and
Angular `InjectionToken`s).

The workspace is now self-contained in this repo and follows the standard
extension layout:

- `frontend/package.json`
- `frontend/nx.json`
- `frontend/libs/extensions/contactus/contract`

The package depends only on foundational peers (`@sneat/core`, `@sneat/dto`,
`@sneat/space-models`, `@sneat/auth-models`, `@angular/core`, `rxjs`) and does
not import another extension's implementation.
