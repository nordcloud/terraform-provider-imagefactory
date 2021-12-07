# Terraform Provider ImageFactory

Run the following command to build the provider

```shell
go build -o terraform-provider-imagefactory
```

## Test sample configuration

First, build and install the provider.

```shell
make install
```

Then, run the following command to initialize the workspace and apply the sample configuration.

```shell
terraform init && terraform apply
```

## Install the released provider

1. Download the Terraform provider zip package from the Github releases that is right for your OS.
2. Unzip package
3. Run `./install.sh`
