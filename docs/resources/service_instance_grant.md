---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "tsuru_service_instance_grant Resource - terraform-provider-tsuru"
subcategory: ""
description: |-
  Tsuru Service Instance Grant
---

# tsuru_service_instance_grant (Resource)

Tsuru Service Instance Grant

## Example Usage

```terraform
resource "tsuru_service_instance_grant" "instance_grant" {
  service_name     = "service01"
  service_instance = "my-instance"
  team             = "mysupport-team"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **service_instance** (String) Name of service instance
- **service_name** (String) Name of service kind
- **team** (String) Team name

### Optional

- **id** (String) The ID of this resource.


