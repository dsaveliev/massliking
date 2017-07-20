#!/bin/bash
ansible-playbook -i default.inventory -D web.yml --vault-password-file ~/.vault_pass.txt
