<template>
  <div class="insta-list-container">
    <router-link :to="{ name: 'instagram-new' }">
      <button class="deep-purple insta-button-add">{{ $t("instagram.buttons.add") }}</button>
    </router-link>
    <div class="list">
      <insta-list-item
        v-for="instagram in instagramList"
        :instagram="instagram"
        :key="instagram.id"
      ></insta-list-item>
    </div>
  </div>
</template>

<script>
import { mapState } from 'vuex'
import InstaListItem from './ListItem.vue'
import InstagramAPI from '../../api/instagram'

const REFRESH_COMPONENT_TIMEOUT = 60 * 1000

export default {
  data () {
    return {
      destroyed: false
    }
  },
  beforeCreate () {
    var self = this
    if (InstagramAPI.isEmptyList(self)) {
      InstagramAPI.updateInstagramList(self)
    }
  },
  created () {
    this.refreshComponent()
  },
  beforeDestroy () {
    this.destroyed = true
  },
  components: {
    'insta-list-item': InstaListItem
  },
  computed: mapState([
    'instagramList'
  ]),
  methods: {
    refreshComponent () {
      var self = this
      if (self.destroyed === true) { return }
      self.$forceUpdate()
      setTimeout(function () { self.refreshComponent() }, REFRESH_COMPONENT_TIMEOUT)
    }
  }
}
</script>

<style>
.insta-list-container {
  margin: 20px;
}
.insta-button-add {
  margin-bottom: 20px;
}
</style>
