<template>
  <div class="channel-list-container">
    <router-link :to="{ name: 'channel-new' }">
      <button class="deep-purple channel-button-add">{{ $t("channel.buttons.add") }}</button>
    </router-link>

    <div class="list">
      <channel-list-item
        v-for="channel in instagram.channels"
        :instagram="instagram"
        :channel="channel"
        :key="channel.id"
      ></channel-list-item>
    </div>
  </div>
</template>

<script>
import ChannelListItem from './ListItem.vue'

export default {
  props: ['instagram'],
  data () {
    return {}
  },
  components: {
    'channel-list-item': ChannelListItem
  }
}
</script>

<style>
.channel-list-container {
}
.channel-button-add {
  margin: 20px 0;
}
</style>
