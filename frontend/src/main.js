// === DEFAULT / CUSTOM STYLE ===
// WARNING! always comment out ONE of the two require() calls below.
// 1. use next line to activate CUSTOM STYLE (./src/themes)
// require(`./themes/app.${__THEME}.styl`)
// 2. or, use next line to activate DEFAULT QUASAR STYLE
require(`quasar/dist/quasar.${__THEME}.css`)
// ==============================

import Vue from 'vue'
import { sync } from 'vuex-router-sync'

import './localization'

import VueResource from 'vue-resource'
Vue.use(VueResource)

import Vuelidate from 'vuelidate'
Vue.use(Vuelidate)

import router from './router'

import Quasar from 'quasar'
Vue.use(Quasar)

import App from './App'

import {store} from './store/index'

sync(store, router)

Quasar.start(() => {
  /* eslint-disable no-new */
  new Vue({
    el: '#q-app',
    store,
    router,
    render: h => h(App)
  })
})
