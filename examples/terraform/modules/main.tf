module "alpha" {
  source = "./alpha"
  var1 = module.alpha.out1
}

output "result" {
  value = module.alpha.
}
