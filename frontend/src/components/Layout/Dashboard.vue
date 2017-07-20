<template>
  <q-layout>

    <div slot="header" class="toolbar deep-purple">
      <button class="hide-on-drawer-visible" @click="$refs.drawer.open()">
        <i>menu</i>
      </button>
      <q-toolbar-title :padding="0">
        <router-link to="/instagram">
          <span class="logo-text text-white">{{ $t("common.titles.header") }}</span>
        </router-link>
      </q-toolbar-title>
      <button class="deep-purple" @click="logout">{{ $t("common.buttons.logout") }}</button>
    </div>

    <q-drawer ref="drawer">
      <div class="list highlight platform-delimiter">
        <div class="list-label">
          <div class="toolbar-greeting text-deep-purple">
            {{ $t("dashboard.titles.greeting") }}, {{ user.username }}!
          </div>
        </div>
        <q-drawer-link icon="account_box" to="/instagram">
          {{ $t("dashboard.titles.accounts") }}&nbsp;&nbsp;
          <span class="sidemenu-label label bg-deep-purple-4 text-white">{{ instagramList.length }}</span>
        </q-drawer-link>
        <q-drawer-link icon="security" to="/security">{{ $t("dashboard.titles.security") }}</q-drawer-link>
        <q-drawer-link icon="mail" to="/contacts">{{ $t("dashboard.titles.contacts") }}</q-drawer-link>
        <!--<q-drawer-link icon="credit_card" to="/billing">{{ $t("dashboard.titles.billing") }}</q-drawer-link>-->
      </div>
    </q-drawer>

    <router-view class="layout-view"></router-view>

    <div slot="footer" class="toolbar deep-purple">
      <wrapper>
        {{ $t("common.titles.footer") }}
      </wrapper>
    </div>
  </q-layout>
</template>

<script>
import { mapState } from 'vuex'
import Wrapper from './../Wrapper.vue'
import auth from '../../api/auth'
import user from '../../api/user'
import instagram from '../../api/instagram'

const REFRESH_TIMEOUT = 15 * 60 * 1000
// TODO: Увеличить таймаут
const REFRESH_ACCOUNTS_TIMEOUT = 60 * 1000

export default {
  data () {
    return {}
  },
  beforeCreate () {
    var self = this
    if (user.isEmpty(self)) {
      user.updateUser(self)
    }
    if (instagram.isEmptyList(self)) {
      instagram.updateInstagramList(self)
    }
  },
  created () {
    this.refreshToken()
    this.refreshAccounts()
  },
  components: {
    'wrapper': Wrapper
  },
  computed: mapState([
    'user',
    'instagramList'
  ]),
  methods: {
    logout () {
      auth.logout(this, '/')
    },
    refreshToken () {
      var self = this
      auth.refreshToken(self, '/')
      setTimeout(function () { self.refreshToken() }, REFRESH_TIMEOUT)
    },
    refreshAccounts () {
      var self = this
      instagram.updateInstagramList(self)
      setTimeout(function () { self.refreshAccounts() }, REFRESH_ACCOUNTS_TIMEOUT)
    }
  }
}
</script>

<style>
.logo-text {
  font-size: x-large;
}
.toolbar-content {
  display: flex;
  align-items: center;
}
.sidemenu-label {
  font-size: small;
}
.toolbar-greeting {
  font-weight: 400;
  font-size: large;
}
</style>
