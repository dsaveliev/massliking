<template>
  <router-link
    class="item item-link item-inset-delimiter text-deep-purple multiple-lines"
    :to="{ name: 'instagram-show', params: { instagram_id: instagram.id }}">
      <img class="item-primary thumbnail" :src="instagram.info.profile_pic_url">
      <div class="item-content">
        <span class="label bg-deep-purple text-white">{{ instagram.info.full_name }}</span>
        <span class="label bg-deep-purple-1 text-black">
          <span class="left-detail">{{ $t("instagram.fields.channels_count").toLowerCase() }}</span>
          <i>compare_arrows</i>&nbsp;{{ instagram.channels.length }}
        </span>
        <span class="label bg-deep-purple-1 text-black">
          <span class="left-detail">{{ $t("instagram.fields.media_count").toLowerCase() }}</span>
          <i>photo_library</i>&nbsp;{{ instagram.info.media_count }}
        </span>
        <span class="label bg-deep-purple-1 text-black">
          <span class="left-detail">{{ $t("instagram.fields.followers_count").toLowerCase() }}</span>
          <i>people</i>&nbsp;{{ instagram.info.follower_count }}
        </span>
        <span class="label bg-deep-purple-1 text-black">
          <span class="left-detail">{{ $t("instagram.fields.following_count").toLowerCase() }}</span>
          <i>people_outline</i>&nbsp;{{ instagram.info.following_count }}
        </span>
        <span :class="['label', 'text-bold', 'text-white', stateColor]">
          <span class="left-detail">state&nbsp;</span>
          {{ instagram.state }}
        </span>
      </div>
  </router-link>
</template>

<script>
export default {
  props: ['instagram'],
  data () {
    return {}
  },
  computed: {
    stateColor () {
      if (this.instagram.state === 'start') {
        return 'bg-green'
      }
      else {
        return 'bg-pink-13'
      }
    }
  }
}
</script>

<style>
</style>
