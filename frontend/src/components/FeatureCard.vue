<template>
  <div class="card-container flex">
    <div class="card">
      <div class="card-title bg-deep-purple text-white">
        {{ title }}
      </div>
      <div class="card-content card-force-top-padding">
        {{ content }}
      </div>
    </div>
  </div>
</template>

<script>
export default {
  props: ['title', 'content'],
  data () {
    return {}
  }
}
</script>

<style>
.card-container {
  max-width: 320px;
}
.card-title {
  font-weight: 300;
  display: flex;
  justify-content: center;
}
.card-content {
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
  align-items: space-start;
}
</style>
