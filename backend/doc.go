// Package backend is the root of the contactus-ext backend Go module.
//
// This module holds the public contract surface of the contactus extension —
// model/const shapes, facade interfaces, and caller-satisfied callback
// signatures — and depends only on foundational/core packages, never on
// another extension. Contract packages (e.g. contactusmodels) are migrated in
// by the contactus-ext reference-extraction plan; this root file exists so the
// module has a buildable package from the moment the repo is stood up.
package backend
