provider "aws" {
  region     = "us-east-1"
}

resource "aws_instance" "rmq" {
  ami           = "ami-04505e74c0741db8d "
  instance_type = "t2.micro"
  key_name = "rabbitmq"
  vpc_security_group_ids = ["sg-0e6f56a8099ee33e4"]
}