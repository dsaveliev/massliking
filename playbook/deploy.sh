#!/bin/bash
ansible-playbook -i default.inventory -D deploy.yml --vault-password-file ~/.vault_pass.txt
