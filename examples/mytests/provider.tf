terraform {
  required_providers {
    bigipv2 = {
      version = "1.0.0"
      source = "terraform.local/net/bigipv2"
  }
 }
}


provider "bigipv2" {
  address = "172.16.0.129"
  username = "admin"
  password = "123***sss"
}

