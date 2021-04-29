# terraform-generator
CLI tool to generate terraform components

## Usage
```tf
terraform-generator generate [component] [flags]
```

## Components
The following components are currently managed :
- module
- variable 
- output
- resource
- data
- configuration

### Module
```tf
terraform-generator generate module --dir=modules --name=instance-configuration
```
This command will generate a modules/instance-configuration directory including the following files README.md, main.tf, outputs.tf, variables.tf and terraform.tf

### Configuration
```tf
terraform-generator generate configuration --name=network --module=modules/instance-configuration
```
This command will generate a network.tf configuration file in the modules/instance-configuration directory.`,

### Variable
```tf
terraform-generator generate variable --name=instance_name --type=string --module=modules/instance-configuration
```
This command will append a variable block with name instance_name and type string at the end of the modules/instance-configuration/variables.tf

### Output
```tf
terraform-generator generate output --module=modules/instance-configuration --name=instance_ip_address --value=module.instance_configuration.ip
```
This command will append an output block with name instance_ip_address and value module.instance_configuration.ip at the end of the modules/instance-configuration/outputs.tf

### Resource
```tf
terraform-generator generate resource --type=aws_s3_bucket --name=static_website_bucket --module=modules/instance-configuration --configuration=resources
```
This command will append a resource block of type "aws_s3_bucket" and name "static_website_bucket" at the end of the modules/instance-configuration/resources.tf

### Data
```tf
terraform-generator generate resource --type=aws_s3_bucket --name=static_website_bucket --module=modules/instance-configuration --configuration=data
```
This command will append a data block of type "aws_s3_bucket" and name "static_website_bucket" at the end of the modules/instance-configuration/data.tf