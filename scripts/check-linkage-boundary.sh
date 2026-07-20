#!/usr/bin/env bash
# The cross-extension Contactus helpers may expose record keys and linkage
# references, but never transactional, query, or mutation capabilities.
set -euo pipefail

forbidden='Run(Readwrite|Readonly)Transaction|\.Get\(|\.Set\(|\.Delete\(|NewWorker|NewQuery|GetMulti|SetMulti'
if rg -n --glob '*.go' "$forbidden" backend/dal4contactus backend/facade4contactus; then
	echo 'Contactus linkage helpers must expose record/linkage data only, never persistence operations.' >&2
	exit 1
fi

imports="$(rg -n --glob '*.go' 'github\.com/sneat-co/(sneat-go/|sneat-bots/|[^/]+/backend)' backend/dal4contactus backend/facade4contactus || true)"
unexpected="$(printf '%s\n' "$imports" | grep -v 'github.com/sneat-co/ext-contactus/backend/' || true)"
if [ -n "$unexpected" ]; then
	printf '%s\n' "$unexpected" >&2
	echo 'Contactus linkage helpers may import only their own contract vocabulary and core packages.' >&2
	exit 1
fi

echo 'Contactus linkage boundary holds.'
