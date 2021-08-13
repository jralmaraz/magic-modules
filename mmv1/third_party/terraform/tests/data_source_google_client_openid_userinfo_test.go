package google

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccDataSourceGoogleClientOpenIDUserinfo_basicNew(t *testing.T) {
	t.Parallel()

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckGoogleClientOpenIDUserinfo_basic,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.google_client_openid_userinfo.me", "email"),
				),
			},
		},
	})
}

const testAccCheckGoogleClientOpenIDUserinfo_basic = `
data "google_client_openid_userinfo" "me" {}
`
