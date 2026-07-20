package facade4contactus

import (
	"reflect"
	"testing"

	"github.com/sneat-co/ext-contactus/backend/contactusmodels/const4contactus"
	"github.com/sneat-co/sneat-core-modules/linkage/dbo4linkage"
)

func TestContactReferencesExposeOnlyLinkageIdentity(t *testing.T) {
	t.Parallel()
	if got, want := NewContactFullRef("space_1", "contact_1"), dbo4linkage.NewFullItemRef(const4contactus.ExtensionID, const4contactus.ContactsCollection, "space_1", "contact_1"); !reflect.DeepEqual(got, want) {
		t.Fatalf("NewContactFullRef() = %#v, want %#v", got, want)
	}
	if got, want := NewContactRefSameSpace("contact_1"), dbo4linkage.NewItemRefSameSpace(const4contactus.ExtensionID, const4contactus.ContactsCollection, "contact_1"); !reflect.DeepEqual(got, want) {
		t.Fatalf("NewContactRefSameSpace() = %#v, want %#v", got, want)
	}
}
