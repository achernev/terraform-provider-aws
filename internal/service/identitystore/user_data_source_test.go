package identitystore_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/aws/aws-sdk-go/service/identitystore"
	sdkacctest "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
)

func TestAccIdentityStoreUserDataSource_basic(t *testing.T) {
	dataSourceName := "data.aws_identitystore_user.test"
	resourceName := "aws_identitystore_user.test"
	name := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheck(t)
			testAccPreCheckSSOAdminInstances(t)
		},
		ErrorCheck:               acctest.ErrorCheck(t, identitystore.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckUserDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccUserDataSourceConfig_basic(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "user_name", resourceName, "user_name"),
					resource.TestCheckResourceAttr(dataSourceName, "user_name", name),
					resource.TestCheckResourceAttrPair(dataSourceName, "addresses.0", resourceName, "addresses.0"),
					resource.TestCheckResourceAttrPair(dataSourceName, "user_id", resourceName, "user_id"),
				),
			},
		},
	})
}

func TestAccIdentityStoreUserDataSource_filterUserName(t *testing.T) {
	dataSourceName := "data.aws_identitystore_user.test"
	resourceName := "aws_identitystore_user.test"
	name := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	email := acctest.RandomEmailAddress(acctest.RandomDomainName())

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheck(t)
			testAccPreCheckSSOAdminInstances(t)
		},
		ErrorCheck:               acctest.ErrorCheck(t, identitystore.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckUserDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccUserDataSourceConfig_filterUserName(name, email),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "user_id", resourceName, "user_id"),
					resource.TestCheckResourceAttr(dataSourceName, "user_name", name),
				),
			},
		},
	})
}

func TestAccIdentityStoreUserDataSource_uniqueAttributeUserName(t *testing.T) {
	dataSourceName := "data.aws_identitystore_user.test"
	resourceName := "aws_identitystore_user.test"
	name := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	email := acctest.RandomEmailAddress(acctest.RandomDomainName())

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheck(t)
			testAccPreCheckSSOAdminInstances(t)
		},
		ErrorCheck:               acctest.ErrorCheck(t, identitystore.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckUserDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccUserDataSourceConfig_uniqueAttributeUserName(name, email),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "user_id", resourceName, "user_id"),
					resource.TestCheckResourceAttr(dataSourceName, "user_name", name),
				),
			},
		},
	})
}

func TestAccIdentityStoreUserDataSource_email(t *testing.T) {
	dataSourceName := "data.aws_identitystore_user.test"
	resourceName := "aws_identitystore_user.test"
	name := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	email := acctest.RandomEmailAddress(acctest.RandomDomainName())

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheck(t)
			testAccPreCheckSSOAdminInstances(t)
		},
		ErrorCheck:               acctest.ErrorCheck(t, identitystore.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckUserDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccUserDataSourceConfig_email(name, email),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "user_id", resourceName, "user_id"),
					resource.TestCheckResourceAttr(dataSourceName, "user_name", name),
				),
			},
		},
	})
}

func TestAccIdentityStoreUserDataSource_userID(t *testing.T) {
	dataSourceName := "data.aws_identitystore_user.test"
	resourceName := "aws_identitystore_user.test"
	name := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	email := acctest.RandomEmailAddress(acctest.RandomDomainName())

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheck(t)
			testAccPreCheckSSOAdminInstances(t)
		},
		ErrorCheck:               acctest.ErrorCheck(t, identitystore.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckUserDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccUserDataSourceConfig_id(name, email),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "user_id", resourceName, "user_id"),
					resource.TestCheckResourceAttrPair(dataSourceName, "user_name", resourceName, "user_name"),
				),
			},
		},
	})
}

func TestAccIdentityStoreUserDataSource_nonExistent(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t); testAccPreCheckSSOAdminInstances(t) },
		ErrorCheck:               acctest.ErrorCheck(t, identitystore.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckUserDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccUserDataSourceConfig_nonExistent,
				ExpectError: regexp.MustCompile(`no Identity Store User found matching criteria`),
			},
		},
	})
}

func TestAccIdentityStoreUserDataSource_userIdFilterMismatch(t *testing.T) {
	name1 := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	email1 := acctest.RandomEmailAddress(acctest.RandomDomainName())
	name2 := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	email2 := acctest.RandomEmailAddress(acctest.RandomDomainName())

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t); testAccPreCheckSSOAdminInstances(t) },
		ErrorCheck:               acctest.ErrorCheck(t, identitystore.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckUserDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccUserDataSourceConfig_userIdFilterMismatch(name1, email1, name2, email2),
				ExpectError: regexp.MustCompile(`no Identity Store User found matching criteria`),
			},
		},
	})
}

func TestAccIdentityStoreUserDataSource_externalIdConflictsWithUniqueAttribute(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t); testAccPreCheckSSOAdminInstances(t) },
		ErrorCheck:               acctest.ErrorCheck(t, identitystore.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckUserDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccUserDataSourceConfig_externalIdConflictsWithUniqueAttribute,
				ExpectError: regexp.MustCompile(`Invalid combination of arguments`),
			},
		},
	})
}

func TestAccIdentityStoreUserDataSource_filterConflictsWithExternalId(t *testing.T) {
	name := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	email := acctest.RandomEmailAddress(acctest.RandomDomainName())

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t); testAccPreCheckSSOAdminInstances(t) },
		ErrorCheck:               acctest.ErrorCheck(t, identitystore.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckUserDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccUserDataSourceConfig_filterConflictsWithUniqueAttribute(name, email),
				ExpectError: regexp.MustCompile(`Conflicting configuration arguments`),
			},
		},
	})
}

func TestAccIdentityStoreUserDataSource_userIdConflictsWithExternalId(t *testing.T) {
	name := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	email := acctest.RandomEmailAddress(acctest.RandomDomainName())

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t); testAccPreCheckSSOAdminInstances(t) },
		ErrorCheck:               acctest.ErrorCheck(t, identitystore.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckUserDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccUserDataSourceConfig_userIdConflictsWithUniqueAttribute(name, email),
				ExpectError: regexp.MustCompile(`Conflicting configuration arguments`),
			},
		},
	})
}

func testAccUserDataSourceConfig_base(name, email string) string {
	return fmt.Sprintf(`
data "aws_ssoadmin_instances" "test" {}

resource "aws_identitystore_user" "test" {
  identity_store_id = tolist(data.aws_ssoadmin_instances.test.identity_store_ids)[0]
  display_name      = "Acceptance Test"
  user_name         = %[1]q

  name {
    family_name = "Acceptance"
    given_name  = "Test"
  }

  emails {
    value = %[2]q
  }
}
`, name, email)
}

func testAccUserDataSourceConfig_basic(name string) string {
	return fmt.Sprintf(`
data "aws_ssoadmin_instances" "test" {}

resource "aws_identitystore_user" "test" {
  identity_store_id = tolist(data.aws_ssoadmin_instances.test.identity_store_ids)[0]
  display_name      = %[1]q
  user_name         = %[1]q

  addresses {
    country        = "US"
    formatted      = "Formatted Address 1"
    locality       = "The Locality 1"
    postal_code    = "AAA BBB 1"
    primary        = true
    region         = "The Region 1"
    street_address = "The Street Address 1"
    type           = "The Type 1"
  }

  name {
    family_name = "Acceptance"
    given_name  = "Test"
  }
}

data "aws_identitystore_user" "test" {
  identity_store_id = tolist(data.aws_ssoadmin_instances.test.identity_store_ids)[0]
  user_id           = aws_identitystore_user.test.user_id
}
`, name)
}

func testAccUserDataSourceConfig_filterUserName(name, email string) string {
	return acctest.ConfigCompose(
		testAccUserDataSourceConfig_base(name, email),
		`
data "aws_identitystore_user" "test" {
  identity_store_id = tolist(data.aws_ssoadmin_instances.test.identity_store_ids)[0]

  filter {
    attribute_path  = "UserName"
    attribute_value = aws_identitystore_user.test.user_name
  }
}
`,
	)
}

func testAccUserDataSourceConfig_uniqueAttributeUserName(name, email string) string {
	return acctest.ConfigCompose(
		testAccUserDataSourceConfig_base(name, email),
		`
data "aws_identitystore_user" "test" {
  identity_store_id = tolist(data.aws_ssoadmin_instances.test.identity_store_ids)[0]

  alternate_identifier {
    unique_attribute {
      attribute_path  = "UserName"
      attribute_value = aws_identitystore_user.test.user_name
    }
  }
}
`,
	)
}

func testAccUserDataSourceConfig_email(name, email string) string {
	return acctest.ConfigCompose(
		testAccUserDataSourceConfig_base(name, email),
		`
data "aws_identitystore_user" "test" {
  identity_store_id = tolist(data.aws_ssoadmin_instances.test.identity_store_ids)[0]

  alternate_identifier {
    unique_attribute {
      attribute_path  = "Emails.Value"
      attribute_value = aws_identitystore_user.test.emails[0].value
    }
  }
}
`,
	)
}

func testAccUserDataSourceConfig_id(name, email string) string {
	return acctest.ConfigCompose(
		testAccUserDataSourceConfig_base(name, email),
		`
data "aws_identitystore_user" "test" {
  identity_store_id = tolist(data.aws_ssoadmin_instances.test.identity_store_ids)[0]

  filter {
    attribute_path  = "UserName"
    attribute_value = aws_identitystore_user.test.user_name
  }

  user_id = aws_identitystore_user.test.user_id
}
`)
}

func testAccUserDataSourceConfig_userIdFilterMismatch(name1, email1, name2, email2 string) string {
	return acctest.ConfigCompose(
		testAccUserDataSourceConfig_base(name1, email1),
		fmt.Sprintf(`
resource "aws_identitystore_user" "test2" {
  identity_store_id = tolist(data.aws_ssoadmin_instances.test.identity_store_ids)[0]
  display_name      = "Acceptance Test"
  user_name         = %[1]q

  name {
    family_name = "Acceptance"
    given_name  = "Test"
  }

  emails {
    value = %[2]q
  }
}

data "aws_identitystore_user" "test" {
  identity_store_id = tolist(data.aws_ssoadmin_instances.test.identity_store_ids)[0]

  filter {
    attribute_path  = "UserName"
    attribute_value = aws_identitystore_user.test.user_name
  }

  user_id = aws_identitystore_user.test2.user_id
}
`, name2, email2))
}

const testAccUserDataSourceConfig_nonExistent = `
data "aws_ssoadmin_instances" "test" {}

data "aws_identitystore_user" "test" {
  alternate_identifier {
    unique_attribute {
      attribute_path  = "UserName"
      attribute_value = "does-not-exist"
    }
  }

  identity_store_id = tolist(data.aws_ssoadmin_instances.test.identity_store_ids)[0]
}
`

const testAccUserDataSourceConfig_externalIdConflictsWithUniqueAttribute = `
data "aws_ssoadmin_instances" "test" {}

data "aws_identitystore_user" "test" {
  alternate_identifier {
    external_id {
      id     = "test"
      issuer = "test"
    }

    unique_attribute {
      attribute_path  = "UserName"
      attribute_value = "does-not-exist"
    }
  }

  identity_store_id = tolist(data.aws_ssoadmin_instances.test.identity_store_ids)[0]
}
`

func testAccUserDataSourceConfig_filterConflictsWithUniqueAttribute(name, email string) string {
	return acctest.ConfigCompose(
		testAccUserDataSourceConfig_base(name, email),
		`
data "aws_identitystore_user" "test" {
  identity_store_id = tolist(data.aws_ssoadmin_instances.test.identity_store_ids)[0]

  alternate_identifier {
    unique_attribute {
      attribute_path  = "UserName"
      attribute_value = aws_identitystore_user.test.user_name
    }
  }

  filter {
    attribute_path  = "UserName"
    attribute_value = aws_identitystore_user.test.user_name
  }
}
`)
}

func testAccUserDataSourceConfig_userIdConflictsWithUniqueAttribute(name, email string) string {
	return acctest.ConfigCompose(
		testAccUserDataSourceConfig_base(name, email),
		`
data "aws_identitystore_user" "test" {
  identity_store_id = tolist(data.aws_ssoadmin_instances.test.identity_store_ids)[0]

  alternate_identifier {
    unique_attribute {
      attribute_path  = "UserName"
      attribute_value = aws_identitystore_user.test.user_name
    }
  }

  user_id = aws_identitystore_user.test.user_id
}
`)
}
