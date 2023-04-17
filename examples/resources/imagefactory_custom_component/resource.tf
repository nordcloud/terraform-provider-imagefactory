// Copyright 2021-2023 Nordcloud Oy or its affiliates. All Rights Reserved.

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
  name            = "Install Apache"
  description     = "Install Apache HTTP Server on Microsoft Windows"
  stage           = "BUILD"
  cloud_providers = ["AWS", "AZURE"]
  os_types        = ["WINDOWS"]
  content {
    script      = <<-EOT
      ---
      - name: Installing Apache HTTP Server
        hosts: all

        tasks:
          - name: Create directory structure
            ansible.windows.win_file:
              path: C:\ansible_examples
              state: directory

          - name: Download the Apache installer
            win_get_url:
              url: https://archive.apache.org/dist/httpd/binaries/win32/httpd-2.2.25-win32-x86-no_ssl.msi
              dest: C:\ansible_examples\httpd-2.2.25-win32-x86-no_ssl.msi

          - name: Install MSI
            win_package:
              path: C:\ansible_examples\httpd-2.2.25-win32-x86-no_ssl.msi
              state: present
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
