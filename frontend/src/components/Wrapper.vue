<template>
  <div class="wrapper-container auto" v-bind:class="{'wrapper-padding': padding}">
    <div class="wrapper-element auto">
      <slot></slot>
    </div>
  </div>
</template>

<script>
export default {
  props: ['padding'],
  data () {
    return {}
  }
}
</script>

<style>
.wrapper-padding {
  padding: 2rem 1rem;
}
.wrapper-container {
  display: flex;
  justify-content: center;
}
.wrapper-element {
  display: flex;
  max-width: 1140px;
  justify-content: center;
}
</style>
