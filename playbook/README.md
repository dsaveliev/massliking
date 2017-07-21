# Deploy Massliking on hosting
This methods tested with $5 droplet on DigitalOcean and works pretty well

## Simple one
1. Make a droplet on DO or your favorite hosting service
2. Add your IP address to `./playbook/default.inventory` file
3. Open `./playbook/roles/web/templates/prod.env.js` and add your IP address into `API_URL` value
4. Go to `./playbook` directory and run `./rollout.sh`
5. For subsequent deploys use `./deploys.sh`
## Custom deploy
1. Make a droplet on DO or your favorite hosting service
2. Add your IP address to `default.inventory` file
3. Fork massliking repo
4. I used [Ansible Vault](http://docs.ansible.com/ansible/latest/playbooks_vault.html) for sensetive files:
    1. Make local file `~/.vault_pass.txt` and save arbitrary password into it
    2. Encrypt your ssh key into file named `deploy_key`
    3. Save it into `./playbook/roles/deploy_keys/files/`
5. Edit `./playbook/roles/web/vars/main.yml` and add your repo
6. Edit `./playbook/roles/postgresql/vars/main.yml` if you want to change database settings
7. Edit `./playbook/roles/web/templates/production.yml` and respectively change `db_string` value
8. Open `./playbook/roles/web/templates/prod.env.js` and add your IP address into `API_URL` value
9. Go to `./playbook` directory and run `./rollout.sh`
10. For subsequent deploys use `./deploys.sh`
