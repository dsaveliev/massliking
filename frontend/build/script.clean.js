var
  shell = require('shelljs'),
  path = require('path')
  env = require('./env-utils'),
  config = require('../config')
  config_env = env.prod ? config.build.env : config.dev.env

shell.rm('-rf', path.resolve(__dirname, '../../' + config_env.BUILD_DIR + '/static'))
console.log(' Cleaned build artifacts.\n')
