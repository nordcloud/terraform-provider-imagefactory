// Copyright 2021-2024 Nordcloud Oy or its affiliates. All Rights Reserved.

# An example of a SHELL component

resource "imagefactory_custom_component" "shell_component" {
  name            = "Install nginx"
  description     = "Install nginx on Ubuntu"
  stage           = "BUILD"
  cloud_providers = ["AWS", "AZURE"]
  os_types        = ["LINUX"]
  content {
    script      = <<-EOT
      apt-get update && apt-get install nginx -y
    EOT
    provisioner = "SHELL"
  }
}

output "shell_component" {
  value = imagefactory_custom_component.shell_component
}

# An example of a Powershell component

resource "imagefactory_custom_component" "powershell_component" {
  name            = "Install nginx"
  description     = "Install nginx Server on Microsoft Windows"
  stage           = "BUILD"
  cloud_providers = ["AWS", "AZURE"]
  os_types        = ["WINDOWS"]
  content {
    script      = <<-EOT
      $ErrorActionPreference = 'Stop';
      $ProgressPreference = 'SilentlyContinue';
      Invoke-WebRequest -Method Get -Uri https://nginx.org/download/nginx-1.25.4.zip -OutFile c:\nginx-1.25.4.zip ;
      Expand-Archive -Path c:\nginx-1.25.4.zip -DestinationPath c:\ ;
      Remove-Item c:\nginx-1.25.4.zip -Force
    EOT
    provisioner = "POWERSHELL"
  }
}

output "powershell_component" {
  value = imagefactory_custom_component.powershell_component
}

# An example of a Ansible playbook component

resource "imagefactory_custom_component" "ansible_component" {
  name            = "Install nginx"
  description     = "Install nginx using ansible playbook"
  stage           = "BUILD"
  cloud_providers = ["AWS", "AZURE"]
  os_types        = ["LINUX"]
  content {
    script      = <<-EOT
      ---
      # playbook.yml
      - name: set up webserver
        hosts: all

        tasks:
          - name: ensure nginx is at the latest version
            package:
              name: nginx
              state: present

          - name: start nginx
            service:
              name: nginx
              state: started
              enabled: yes
    EOT
    provisioner = "ANSIBLE"
  }
}

output "ansible_component" {
  value = imagefactory_custom_component.ansible_component
}
