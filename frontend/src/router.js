import Vue from 'vue'
import VueRouter from 'vue-router'
import auth from './api/auth'

Vue.use(VueRouter)

const requireAuth = (to, _from, next) => {
  if (!auth.checkAuth()) {
    next({
      path: '/login',
      query: { redirect: to.fullPath }
    })
  }
  else {
    next()
  }
}

const afterAuth = (_to, from, next) => {
  if (auth.checkAuth()) {
    next(from.path)
  }
  else {
    next()
  }
}

const authToDashboard = (_to, from, next) => {
  if (auth.checkAuth()) {
    next('/instagram')
  }
  else {
    next()
  }
}

function load (component) {
  return () => System.import(`components/${component}.vue`)
}

function loadFrom (from, component) {
  return () => System.import(`components/${from}/${component}.vue`)
}

export default new VueRouter({
  mode: 'history',
  routes: [
    { path: '/',
      component: loadFrom('Layout', 'Landing'),
      children: [
        { path: '', component: load('Landing'), beforeEnter: authToDashboard },
        { path: 'login', component: load('Login'), beforeEnter: afterAuth },
        { path: 'signup', component: load('Signup'), beforeEnter: afterAuth }
      ]
    },
    { path: '/instagram',
      component: loadFrom('Layout', 'Dashboard'),
      beforeEnter: requireAuth,
      children: [
        { path: '', name: 'instagram-list', component: loadFrom('Instagram', 'List') },
        { path: '/instagram/new', name: 'instagram-new', component: loadFrom('Instagram', 'Form') },
        { path: '/instagram/:instagram_id/edit', name: 'instagram-edit', component: loadFrom('Instagram', 'Edit') },
        { path: '/instagram/:instagram_id', name: 'instagram-show', component: loadFrom('Instagram', 'Show') },
        { path: '/instagram/:instagram_id/channel/new', name: 'channel-new', component: loadFrom('Channel', 'Form') },
        { path: '/instagram/:instagram_id/channel/:id/edit', name: 'channel-edit', component: loadFrom('Channel', 'Edit') },
        { path: '/instagram/:instagram_id/channel/:id', name: 'channel-show', component: loadFrom('Channel', 'Show') }
      ]
    },
    { path: '*', component: load('Error404') }
  ]
})
