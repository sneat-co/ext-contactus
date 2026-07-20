package const4contactus

import "testing"

func TestKnownPetKinds(t *testing.T) {
	if !IsKnownPetPetKind(PetKindDog) {
		t.Fatal("dog should be known")
	}
	if IsKnownPetPetKind("dragon") {
		t.Fatal("dragon should not be known")
	}
}

func TestKnownSpaceMemberRole(t *testing.T) {
	if !IsKnownSpaceMemberRole(SpaceMemberRoleMember, nil) {
		t.Fatal("standard role should be known")
	}
	if !IsKnownSpaceMemberRole("custom", []SpaceMemberRole{"custom"}) {
		t.Fatal("configured custom role should be known")
	}
	if IsKnownSpaceMemberRole("custom", []SpaceMemberRole{"other"}) {
		t.Fatal("unconfigured custom role should not be known")
	}
}
