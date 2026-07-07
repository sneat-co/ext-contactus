# ext-contactus / frontend

Home of the `@sneat/extension-contactus-contract` nx library — contactus's runtime-light public
frontend contract (interfaces, DTO/model types, enums, and Angular `InjectionToken`s).

The library and its nx workspace configuration are migrated here from the `contactus` main repo
by the `contactus-ext` reference-extraction plan (frontend relocation task). Until then this
placeholder reserves the `frontend/` layout slot.

The package depends only on foundational/core peers (`@sneat/core`, `@sneat/dto`,
`@sneat/space-models`, `@sneat/auth-models`, `@angular/core`, `rxjs`) — never on another extension.
