<template>
  <div class="channel-show-container">

    <div class="row no-wrap gt-sm">
      <router-link :to="{ name: 'instagram-show', params: { instagram_id: instagram_id }}">
        <button class="deep-purple channel-button-nav"><i>keyboard_arrow_left</i>{{ $t("channel.buttons.back_to_account") }}</button>
      </router-link>
      <router-link :to="{ name: 'channel-edit', params: { instagram_id: instagram_id, id: id }}">
        <button class="deep-purple channel-button-nav">{{ $t("channel.buttons.edit") }}</button>
      </router-link>
      <button v-if="channel(instagram_id, id).state == 'start'" class="channel-button-nav text-white bg-pink-13"
        @click="stopChannel">{{ $t("channel.buttons.stop") }}</button>
      <button v-else class="channel-button-nav text-white bg-green"
        @click="startChannel">{{ $t("channel.buttons.start") }}</button>
      <button class="pink channel-button-nav" @click="deleteChannel">{{ $t("channel.buttons.delete") }}</button>
    </div>

    <div class="row wrap sm">
      <router-link :to="{ name: 'instagram-show', params: { instagram_id: instagram_id }}">
        <button class="deep-purple channel-button-nav full-width"><i>keyboard_arrow_left</i>{{ $t("channel.buttons.back_to_account") }}</button>
      </router-link>
      <router-link :to="{ name: 'channel-edit', params: { instagram_id: instagram_id, id: id }}">
        <button class="deep-purple channel-button-nav full-width">{{ $t("channel.buttons.edit") }}</button>
      </router-link>
      <button v-if="channel(instagram_id, id).state == 'start'" class="channel-button-nav text-white bg-pink-13 full-width"
        @click="stopChannel">{{ $t("channel.buttons.stop") }}</button>
      <button v-else class="channel-button-nav text-white bg-green full-width"
        @click="startChannel">{{ $t("channel.buttons.start") }}</button>
      <button class="pink channel-button-nav full-width" @click="deleteChannel">{{ $t("channel.buttons.delete") }}</button>
    </div>

    <div class="text-pink channel-fullname">Channel for {{ instagram(instagram_id).info.full_name }}</div>

    <div class="row wrap">
      <div class="column channel-column">
        <table class="q-table">
          <tbody>
            <tr>
              <td class="text-right channel-property"><span class="text-deep-purple">{{ $t("channel.fields.value") }}:</span></td>
              <td class="text-left channel-property">{{ channel(instagram_id, id).value }}</td>
            </tr>
            <tr>
              <td class="text-right channel-property"><span class="text-deep-purple">{{ $t("channel.fields.action") }}:</span></td>
              <td class="text-left channel-property">{{ channel(instagram_id, id).action }}</td>
            </tr>
            <tr>
              <td class="text-right channel-property"><span class="text-deep-purple">{{ $t("channel.fields.target") }}:</span></td>
              <td class="text-left channel-property">{{ channel(instagram_id, id).target }}</td>
            </tr>
            <tr>
              <td class="text-right channel-property"><span class="text-deep-purple">{{ $t("channel.fields.state") }}:</span></td>
              <td class="text-left channel-property">
                <span v-if="channel(instagram_id, id).state == 'start'" class="label text-white bg-green">{{ channel(instagram_id, id).state }}</span>
                <span v-else class="label text-white bg-pink-13">{{ channel(instagram_id, id).state }}</span>
              </td>
            </tr>
          </tbody>
        </table>
      </div>

      <div class="column channel-column">
        <table class="q-table">
          <tbody>
            <tr>
              <td class="text-right channel-property"><span class="text-deep-purple">{{ $t("channel.fields.conversion") }}:</span></td>
              <td class="text-left channel-property">
                {{ channel(instagram_id, id).followers_count }}&nbsp;[&nbsp;{{ (channel(instagram_id, id).followers_count / (channel(instagram_id, id).leads_count + 1.0)).toFixed(2) }}&nbsp;]
              </td>
            </tr>
            <tr>
              <td class="text-right channel-property"><span class="text-deep-purple">{{ $t("channel.fields.actions_count") }}:</span></td>
              <td class="text-left channel-property">{{ channel(instagram_id, id).leads_count }}</td>
            </tr>
            <tr>
              <td class="text-right channel-property"><span class="text-deep-purple">{{ $t("channel.fields.queue") }}:</span></td>
              <td class="text-left channel-property">{{ channel(instagram_id, id).targets_count }}</td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

  </div>
</template>

<script>
import Vue from 'vue'
import { mapGetters } from 'vuex'
import { Toast } from 'quasar'
import ChannelAPI from '../../api/channel'
import InstagramAPI from '../../api/instagram'

export default {
  data () {
    return {
      id: this.$route.params.id,
      instagram_id: this.$route.params.instagram_id
    }
  },
  beforeCreate () {
    var self = this
    if (InstagramAPI.isEmptyList(self)) {
      InstagramAPI.updateInstagramList(self)
    }
  },
  computed: {
    startedChannel () {
      return this.channel(this.instagram_id, this.id).state === 'start'
    },
    stateColor () {
      if (this.channel(this.instagram_id, this.id).state === 'start') {
        return 'bg-green'
      }
      else {
        return 'bg-pink-13'
      }
    },
    ...mapGetters([
      'instagram',
      'channel'
    ])
  },
  methods: {
    deleteChannel () {
      ChannelAPI.deleteChannel(this,
        this.instagram(this.instagram_id),
        this.channel(this.instagram_id, this.id),
        '/instagram/' + this.instagram_id)
      Toast.create({
        html: Vue.t('channel.toasts.deleting'),
        color: 'white',
        bgColor: '#E91E63', // '#673AB7'
        timeout: 5000
      })
    },
    startChannel () {
      ChannelAPI.startChannel(this,
        this.instagram(this.instagram_id),
        this.channel(this.instagram_id, this.id))
      Toast.create({
        html: Vue.t('channel.toasts.start'),
        color: 'white',
        bgColor: '#4caf50', // '#673AB7'
        timeout: 5000
      })
    },
    stopChannel () {
      ChannelAPI.stopChannel(this,
        this.instagram(this.instagram_id),
        this.channel(this.instagram_id, this.id))
      Toast.create({
        html: Vue.t('channel.toasts.stop'),
        color: 'white',
        bgColor: '#E91E63', // '#673AB7'
        timeout: 5000
      })
    }
  }
}
</script>

<style>
.channel-show-container {
  margin: 20px;
}
.channel-fullname {
  margin: 20px 0;
  align-items: flex-end;
  font-weight: 300;
  font-size: xx-large;
}
.channel-column {
  margin-right: 20px;
}
.channel-property {
  margin-bottom: 20px;
  align-items: flex-end;
  font-weight: 300;
  font-size: large;
}
.channel-button-nav {
  margin: 0 20px 20px 0;
}
</style>
