package facade4contactus

import (
	"github.com/sneat-co/ext-contactus/backend/contactusmodels/const4contactus"
	"github.com/sneat-co/sneat-core-modules/linkage/dbo4linkage"
	"github.com/sneat-co/sneat-go-core/coretypes"
)

// NewContactFullRef returns a Contactus item reference in another space. It is
// extension-owned glue for consumers that link to a contact without importing
// the Contactus implementation.
func NewContactFullRef(spaceID coretypes.SpaceID, contactID string) dbo4linkage.ItemRef {
	return dbo4linkage.NewFullItemRef(const4contactus.ExtensionID, const4contactus.ContactsCollection, spaceID, contactID)
}

// NewContactRefSameSpace returns a Contactus item reference in the current
// space. It is intentionally a linkage reference, not a Contactus DTO.
func NewContactRefSameSpace(contactID string) dbo4linkage.ItemRef {
	return dbo4linkage.NewItemRefSameSpace(const4contactus.ExtensionID, const4contactus.ContactsCollection, contactID)
}
