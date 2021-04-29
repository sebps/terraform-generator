./terraform-generator generate module --dir=test/modules --name=instance-configuration
./terraform-generator generate configuration --name=network --dir=test/modules/instance-configuration
./terraform-generator generate variable --name=instance_name --type=string --dir=test/modules/instance-configuration
./terraform-generator generate output --dir=test/modules/instance-configuration --name=instance_ip_address --value=module.instance_configuration.ip
./terraform-generator generate resource --type=aws_ec2 --name=server_instance --dir=test/modules/instance-configuration --configuration=instance
./terraform-generator generate data --type=aws_region --name=current_region --dir=test/modules/instance-configuration --configuration=data