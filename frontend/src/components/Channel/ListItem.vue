<template>
  <router-link
    class="item item-link item-inset-delimiter multiple-lines"
    :to="{ name: 'channel-show', params: { instagram_id: instagram.id, id: channel.id }}">
      <div class="item-content">
        <span class="label bg-pink text-white"> {{ channel.action }} </span>
        <span class="label bg-deep-purple text-white"> {{ channel.value }} </span>
        <span class="label bg-pink text-white"> {{ channel.target }} </span>
        <span class="label bg-deep-purple-1 text-black">
          <span class="left-detail">{{ $t("channel.fields.conversion").toLowerCase() }}</span>
          <i>trending_up</i>&nbsp;{{ channel.followers_count }}&nbsp;[&nbsp;{{ (channel.followers_count / (channel.leads_count + 1.0)).toFixed(2) }}&nbsp;]
        </span>
        <span class="label bg-deep-purple-1 text-black">
          <span class="left-detail">{{ $t("channel.fields.actions_count").toLowerCase() }}</span>
          <i>thumb_up</i>&nbsp;{{ channel.leads_count }}
        </span>
        <span class="label bg-deep-purple-1 text-black">
          <span class="left-detail">{{ $t("channel.fields.queue").toLowerCase() }}</span>
          <i>person</i>&nbsp;{{ channel.targets_count }}
        </span>
        <span :class="['label', 'text-bold', 'text-white', stateColor]">
          <span class="left-detail">state&nbsp;</span>
          {{ channel.state }}
        </span>
      </div>
  </router-link>
</template>

<script>
export default {
  props: ['instagram', 'channel'],
  data () {
    return {}
  },
  computed: {
    stateColor () {
      if (this.channel.state === 'start') {
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
