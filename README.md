# Massliking
## Synopsis
Full featured mass-following service with an emphasis on marketing concepts.

Written with [Golang](https://golang.org/) and [Quasar](http://quasar-framework.org/) ([VueJS](https://vuejs.org/) under the hood)

This is a relatively stable application, but still in alpha.

[Demo site](http://massliking.com)
## Core concepts
Each target account is considered as a separate channel for attracting followers.

Statistics per each channel gives you the ability to optimise attracting speed by throwing out slow channels.

This ability along with targets filtering makes your growth much faster (I hope it does, huh).
## Instagram client
An instabot package from the backend is a copy of https://github.com/instabot-py/instabot.py.git

There is a [separate document](LIMITS.md) for instagram limits used in the service.

Also, there is an [API description](backend/README.md) for backend.
## Requirements
* `Go` >= 1.8.3
* `Node.js` >= 4.2.6
* `NPM` >= 3.5.2
* `PostgreSQL` >= 9.5.6
#### For deploy
* `Ansible` >= 2.3.1.0

_Maybe it works with earlier versions, this is my current env._
## How to run local instance
1. Create a new database in PostgreSQL
2. Rename `.sample` files from `./config` directory
3. Open `config/development.yml` and edit `db_string` value
4. Compile project with `./bin/build.sh`
5. Run service with `./bin/run.sh`
6. [Here we go!](http://localhost:8080)
## How to deploy on hosting via Ansible
Check out [these instructions](playbook/README.md) for deployment via [Ansible](https://www.ansible.com/).
## TODO
- Implement per-channel filters for targets
- Implement user-friendly interface for channels (sorting, filters etc.)
- Implement signup confirmation via email
- Cover backend(at least) with tests
- Rewrite channels worker (fill queues in separate goroutines)
- Plug in the state machine library to the models
