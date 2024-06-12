package campaign

import "testing"

func TestNewCampaign(t *testing.T) {
	name := "Campaign x"
	content := "Body"
	contacts := []string{"email@teste.com", "email2@teste.com"}
	lenContacts := len(contacts)

	campaign := NewCampaign(name, content, contacts)

	if campaign.Name != name {
		t.Errorf("expected " + name)
	}

	if campaign.Content != content {
		t.Errorf("expected " + content)
	}

	if len(campaign.Contact) != lenContacts {
		t.Errorf("expected that the number of contacts equal")
	}
}
