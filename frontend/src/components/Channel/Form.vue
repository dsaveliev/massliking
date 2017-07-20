<template>
  <div class="channel-form-container">
    <div class="row no-wrap gt-sm">
      <router-link :to="{ name: 'instagram-show', params: { instagram_id: instagram_id }}">
        <button class="deep-purple channel-button-nav">{{ $t("channel.buttons.back_to_account") }}</button>
      </router-link>
    </div>

    <div class="row wrap sm">
      <router-link :to="{ name: 'instagram-show', params: { instagram_id: instagram_id }}">
        <button class="deep-purple channel-button-nav full-width">{{ $t("channel.buttons.back_to_account") }}</button>
      </router-link>
    </div>

    <wrapper>
      <div class="card shadow-3 channel-form-card">
        <div class="card-title bg-wite text-pink">
          <big class="channel-form-title">{{ $t("channel.titles.create") }}</big>
        </div>
        <div class="card-content card-force-top-padding column">

          <div class="floating-label channel-input">
            <input required class="full-width"
              :class="{'has-error': $v.form.value.$error}"
              v-model.trim="form.value"
              @input="$v.form.value.$touch()">
            <div v-if="$v.form.$error && !$v.form.value.required" class="text-red">{{ $t("channel.errors.value.required") }}</div>
            <label>{{ $t("channel.fields.value") }}</label>
          </div>

          <div v-if="form.action === 'comment'" class="floating-label channel-input">
            <q-chips required class="full-width"
              v-model.trim="form.comments"
              @input="$v.form.value.$touch()"></q-chips>
            <label>{{ $t("channel.fields.comments") }}</label>
          </div>

          <div class="floating-label channel-input">
            <q-select required class="full-width"
              type="radio"
              :options="actionOptions"
              :class="{'has-error': $v.form.action.$error}"
              v-model.trim="form.action"
              @input="$v.form.action.$touch()"></q-select>
            <div v-if="$v.form.$error && !$v.form.action.required" class="text-red">{{ $t("channel.errors.action.required") }}</div>
            <label>{{ $t("channel.fields.action") }}</label>
          </div>

          <div class="floating-label channel-input">
            <q-select required class="full-width"
              type="radio"
              :options="targetOptions"
              :class="{'has-error': $v.form.target.$error}"
              v-model.trim="form.target"
              @input="$v.form.target.$touch()"></q-select>
            <div v-if="$v.form.$error && !$v.form.target.required" class="text-red">{{ $t("channel.errors.target.required") }}</div>
            <label>{{ $t("channel.fields.target") }}</label>
          </div>

          <div class="channel-input">
            <button class="deep-purple big full-width" @click="submit">{{ $t("channel.buttons.create") }}</button>
          </div>

        </div>
      </div>
    </wrapper>

  </div>
</template>

<script>
import Vue from 'vue'
import { mapGetters } from 'vuex'
import Wrapper from '../Wrapper.vue'
import { required } from 'vuelidate/lib/validators'
import { Toast } from 'quasar'
import ChannelAPI from '../../api/channel'

export default {
  data () {
    return {
      instagram_id: this.$route.params.instagram_id,
      actionOptions: [
        { label: 'Like', value: 'like' },
        { label: 'Comment', value: 'comment' },
        { label: 'Follow', value: 'follow' },
        { label: 'Unfollow', value: 'unfollow' }
      ],
      targetOptions: [
        { label: 'Followers', value: 'followers' },
        { label: 'Subscriptions', value: 'subscriptions' },
        { label: 'Hashtag', value: 'hashtag' },
        { label: 'Likes', value: 'likes' },
        { label: 'Comments', value: 'comments' }
      ],
      form: {
        action: '',
        target: '',
        value: '',
        comments: []
      }
    }
  },
  validations: {
    form: {
      action: { required },
      target: { required },
      value: { required }
    }
  },
  components: {
    'wrapper': Wrapper
  },
  computed: {
    ...mapGetters([
      'instagram'
    ])
  },
  methods: {
    submit () {
      this.$v.form.$touch()
      if (this.$v.form.$error) {
        Toast.create.warning({
          html: Vue.t('common.errors.fields.invalid'),
          color: 'saddlebrown'
        })
        return
      }
      ChannelAPI.createChannel(this, this.instagram(this.instagram_id), this.form, '/instagram/' + this.instagram_id)
    }
  }
}
</script>

<style>
.channel-form-container {
  margin: 20px;
}
.channel-button-nav {
  margin: 0 20px 20px 0;
}
.channel-input {
  padding: 1rem 0;
}
.channel-form-card {
  max-width: 600px;
}
.channel-form-title {
  font-weight: 300;
}
</style>
