#!/bin/bash
ansible-playbook -i default.inventory -D server.yml --vault-password-file ~/.vault_pass.txt
