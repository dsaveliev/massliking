#!/bin/bash
ansible-playbook -i default.inventory -D rollout.yml --vault-password-file ~/.vault_pass.txt
