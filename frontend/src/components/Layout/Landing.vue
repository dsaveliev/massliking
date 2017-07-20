<template>
  <q-layout>
    <div slot="header" class="toolbar deep-purple">
      <wrapper>
        <q-toolbar-title :padding="0">
          <div class="logo-container">
          <router-link to="/">
              <span class="logo-text text-white">{{ $t("common.titles.header") }}</span>
          </router-link>
          </div>
        </q-toolbar-title>

        <router-link to="/login">
          <button class="deep-purple">{{ $t("common.buttons.login") }}</button>
        </router-link>

        <router-link to="/signup">
          <button class="pink">{{ $t("common.buttons.signup") }}</button>
        </router-link>
      </wrapper>
    </div>

    <router-view class="layout-view"></router-view>

    <div slot="footer" class="toolbar deep-purple">
      <wrapper>
        {{ $t("common.titles.footer") }}
      </wrapper>
    </div>
  </q-layout>
</template>

<script>
import Wrapper from './../Wrapper.vue'

export default {
  data () {
    return {}
  },
  components: {
    'wrapper': Wrapper
  }
}
</script>

<style>
.logo-text {
  font-size: x-large;
}
.logo-container {
  display: flex;
  align-items: center;
  margin: 5px 0 0 -12px;
}
</style>
