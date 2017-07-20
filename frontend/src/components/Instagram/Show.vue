<template>
  <div class="insta-show-container">

    <div class="row no-wrap gt-sm">
      <router-link to="/instagram">
        <button class="deep-purple insta-button-nav"><i>keyboard_arrow_left</i>{{ $t("instagram.buttons.list") }}</button>
      </router-link>
      <router-link :to="{ name: 'instagram-edit', params: { instagram_id: instagram_id }}">
        <button class="deep-purple insta-button-nav">{{ $t("instagram.buttons.edit") }}</button>
      </router-link>
      <button v-if="instagram(instagram_id).state == 'start'" class="insta-button-nav text-white bg-pink-13"
        @click="stopInstagram">{{ $t("instagram.buttons.stop") }}</button>
      <button v-else class="insta-button-nav text-white bg-green"
        @click="startInstagram">{{ $t("instagram.buttons.start") }}</button>
      <button class="pink insta-button-nav" @click="deleteAccount">{{ $t("instagram.buttons.delete") }}</button>
    </div>

    <div class="row wrap sm">
      <router-link to="/instagram">
        <button class="deep-purple insta-button-nav full-width"><i>keyboard_arrow_left</i>{{ $t("instagram.buttons.list") }}</button>
      </router-link>
      <router-link :to="{ name: 'instagram-edit', params: { instagram_id: instagram_id }}">
        <button class="deep-purple insta-button-nav full-width">{{ $t("instagram.buttons.edit") }}</button>
      </router-link>
      <button v-if="instagram(instagram_id).state == 'start'" class="insta-button-nav text-white bg-pink-13 full-width"
        @click="stopInstagram">{{ $t("instagram.buttons.stop") }}</button>
      <button v-else class="insta-button-nav text-white bg-green full-width"
        @click="startInstagram">{{ $t("instagram.buttons.start") }}</button>
      <button class="pink insta-button-nav full-width" @click="deleteAccount">{{ $t("instagram.buttons.delete") }}</button>
    </div>

    <div  v-if="instagram(instagram_id).state != 'start'" class="insta-alert shadow-1 row inline items-center text-black bg-pink-1">
      {{ $t("instagram.toasts.inactive") }}
    </div>

    <div class="row wrap">
      <div class="column insta-column">
        <img class="insta-profile-pic" :src="instagram(instagram_id).info.profile_pic_url">
        <div class="text-pink insta-fullname">{{ instagram(instagram_id).info.full_name }}</div>
      </div>

      <div class="column insta-column">
        <table class="q-table">
          <tbody>
            <tr>
              <td class="text-right insta-property"><span class="text-deep-purple">{{ $t("instagram.fields.media_count") }}:</span></td>
              <td class="text-left insta-property">{{ instagram(instagram_id).info.media_count }}</td>
            </tr>
            <tr>
              <td class="text-right insta-property"><span class="text-deep-purple">{{ $t("instagram.fields.followers_count") }}:</span></td>
              <td class="text-left insta-property">{{ instagram(instagram_id).info.follower_count }}</td>
            </tr>
            <tr>
              <td class="text-right insta-property"><span class="text-deep-purple">{{ $t("instagram.fields.following_count") }}:</span></td>
              <td class="text-left insta-property">{{ instagram(instagram_id).info.following_count }}</td>
            </tr>
            <tr>
              <td class="text-right insta-property"><span class="text-deep-purple">{{ $t("instagram.fields.hours") }}:</span></td>
              <td class="text-left insta-property">{{ instagram(instagram_id).hours.min }} - {{ instagram(instagram_id).hours.max }}</td>
            </tr>
            <tr>
              <td class="text-right insta-property"><span class="text-deep-purple">{{ $t("instagram.fields.state") }}:</span></td>
              <td class="text-left insta-property">
                <span v-if="instagram(instagram_id).state == 'start'" class="label text-white bg-green">{{ instagram(instagram_id).state }}</span>
                <span v-else class="label text-white bg-pink-13">{{ instagram(instagram_id).state }}</span>
              </td>
            </tr>
          </tbody>
        </table>
      </div>

      <div class="column insta-column">
        <table class="q-table">
          <tbody>
            <tr>
              <td class="text-right insta-property"><span class="text-deep-purple">{{ $t("instagram.fields.trusted") }}:</span></td>
              <td class="text-left insta-property">{{ instagram(instagram_id).trusted ? $t("yes") : $t("no") }}</td>
            </tr>
            <tr>
              <td class="text-right insta-property"><span class="text-deep-purple">{{ $t("instagram.fields.like_speed") }}:</span></td>
              <td class="text-left insta-property">{{ instagram(instagram_id).speed.like }}</td>
            </tr>
            <tr>
              <td class="text-right insta-property"><span class="text-deep-purple">{{ $t("instagram.fields.comment_speed") }}:</span></td>
              <td class="text-left insta-property">{{ instagram(instagram_id).speed.comment }}</td>
            </tr>
            <tr>
              <td class="text-right insta-property"><span class="text-deep-purple">{{ $t("instagram.fields.follow_speed") }}:</span></td>
              <td class="text-left insta-property">{{ instagram(instagram_id).speed.follow }}</td>
            </tr>
            <tr>
              <td class="text-right insta-property"><span class="text-deep-purple">{{ $t("instagram.fields.unfollow_speed") }}:</span></td>
              <td class="text-left insta-property">{{ instagram(instagram_id).speed.unfollow }}</td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <channel-list :instagram="instagram(instagram_id)"></channel-list>

  </div>
</template>

<script>
import Vue from 'vue'
import { mapGetters } from 'vuex'
import { Toast } from 'quasar'
import InstagramAPI from '../../api/instagram'
import ChannelList from '../Channel/List.vue'

const REFRESH_COMPONENT_TIMEOUT = 60 * 1000

export default {
  data () {
    return {
      instagram_id: this.$route.params.instagram_id,
      destroyed: false
    }
  },
  components: {
    'channel-list': ChannelList
  },
  computed: {
    ...mapGetters([
      'instagram'
    ])
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
  methods: {
    refreshComponent () {
      var self = this
      if (self.destroyed === true) { return }
      self.$forceUpdate()
      setTimeout(function () { self.refreshComponent() }, REFRESH_COMPONENT_TIMEOUT)
    },
    deleteAccount () {
      InstagramAPI.deleteInstagram(this, this.instagram(this.instagram_id), '/instagram')
      Toast.create({
        html: Vue.t('instagram.toasts.deleting'),
        color: 'white',
        bgColor: '#E91E63', // '#673AB7'
        timeout: 5000
      })
    },
    startInstagram () {
      InstagramAPI.startInstagram(this, this.instagram(this.instagram_id))
      Toast.create({
        html: Vue.t('instagram.toasts.start'),
        color: 'white',
        bgColor: '#4caf50', // '#673AB7'
        timeout: 5000
      })
    },
    stopInstagram () {
      InstagramAPI.stopInstagram(this, this.instagram(this.instagram_id))
      Toast.create({
        html: Vue.t('instagram.toasts.stop'),
        color: 'white',
        bgColor: '#E91E63', // '#673AB7'
        timeout: 5000
      })
    }
  }
}
</script>

<style>
.insta-show-container {
  margin: 20px;
}
.insta-profile-pic {
  width: 250px;
  height: 250px;
}
.insta-thumbnail {
  width: 250px;
  height: 250px;
  margin: 40px 30px 0 0;
}
.insta-fullname {
  margin-top: 20px;
  align-items: flex-end;
  font-weight: 300;
  font-size: xx-large;
}
.insta-column {
  margin-right: 20px;
}
.insta-property {
  margin-bottom: 20px;
  align-items: flex-end;
  font-weight: 300;
  font-size: large;
}
.insta-button-nav {
  margin: 0 20px 20px 0;
}
.insta-alert {
  font-weight: 400;
  font-size: large;
  margin-bottom: 20px;
  padding: 10px;
}
</style>
