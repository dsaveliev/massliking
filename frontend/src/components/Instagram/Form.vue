<template>
  <div class="insta-form-container">
    <div class="row no-wrap gt-sm">
      <router-link to="/instagram">
        <button class="deep-purple insta-button-nav">{{ $t("instagram.buttons.list") }}</button>
      </router-link>
    </div>

    <div class="row wrap sm">
      <router-link to="/instagram">
        <button class="deep-purple insta-button-nav full-width">{{ $t("instagram.buttons.list") }}</button>
      </router-link>
    </div>

    <wrapper>
      <div class="card shadow-3 insta-form-card">
        <div class="card-title bg-wite text-pink">
          <big class="insta-form-title">{{ $t("instagram.titles.create") }}</big>
        </div>
        <div class="card-content card-force-top-padding column">

          <div class="floating-label insta-input">
            <input required class="full-width"
              :class="{'has-error': $v.form.username.$error}"
              v-model.trim="form.username"
              @input="$v.form.username.$touch()">
            <div v-if="$v.form.$error && !$v.form.username.required" class="text-red">{{ $t("instagram.errors.username.required") }}</div>
            <label>{{ $t("instagram.fields.username") }}</label>
          </div>

          <div class="floating-label insta-input">
            <input required type="password" class="full-width"
              :class="{'has-error': $v.form.password.$error}"
              v-model.trim="form.password"
              @input="$v.form.password.$touch()">
            <div v-if="$v.form.$error && !$v.form.password.required" class="text-red">{{ $t("instagram.errors.password.required") }}</div>
            <label>{{ $t("instagram.fields.password") }}</label>
          </div>

          <div class="insta-input full-width">
            <label>{{ $t("instagram.fields.hours") }}: {{ form.hours.min }} - {{ form.hours.max }}</label>
            <p></p>
            <q-double-range
              class="deep-purple"
              v-model="form.hours"
              :label=true :labelAlways=true :snap=true :markers=true
              :min="0" :max="23"
            ></q-double-range>
          </div>

          <div class="insta-input full-width">
            <label>{{ $t("instagram.fields.like_speed") }}: {{ form.speed.like }}</label>
            <p></p>
            <q-range
              class="deep-purple"
              v-model="form.speed.like"
              :label=true :labelAlways=true :step=5 :snap=true :markers=true
              :min="0" :max="130"
            ></q-range>
          </div>

          <div class="insta-input full-width">
            <label>{{ $t("instagram.fields.comment_speed") }}: {{ form.speed.comment }}</label>
            <p></p>
            <q-range
              class="deep-purple"
              v-model="form.speed.comment"
              :label=true :labelAlways=true :snap=true :markers=true
              :min="0" :max="14"
            ></q-range>
          </div>

           <div class="insta-input full-width">
             <label>{{ $t("instagram.fields.follow_speed") }}: {{ form.speed.follow }}</label>
            <p></p>
            <q-range
              class="deep-purple"
              v-model="form.speed.follow"
              :label=true :labelAlways=true :step=5 :snap=true :markers=true
              :min="0" :max="130"
            ></q-range>
          </div>

          <div class="insta-input full-width">
            <label>{{ $t("instagram.fields.unfollow_speed") }}: {{ form.speed.unfollow }}</label>
            <p></p>
            <q-range
              class="deep-purple"
              v-model="form.speed.unfollow"
              :label=true :labelAlways=true :step=10 :snap=true :markers=true
              :min="0" :max="300"
            ></q-range>
          </div>

          <div class="insta-input full-width">
            <q-checkbox class="deep-purple" v-model.trim="form.trusted"></q-checkbox>
            <label>
              &nbsp;{{ $t("instagram.fields.trusted") }}
            </label>
          </div>

          <div class="insta-input">
            <button class="deep-purple big full-width" @click="submit">{{ $t("instagram.buttons.create") }}</button>
          </div>

        </div>
      </div>
    </wrapper>

  </div>
</template>

<script>
import Vue from 'vue'
import Wrapper from '../Wrapper.vue'
import { required } from 'vuelidate/lib/validators'
import { Toast } from 'quasar'
import instagram from '../../api/instagram'

export default {
  data () {
    return {
      form: {
        username: '',
        password: '',
        trusted: false,
        hours: {
          min: 0,
          max: 23
        },
        speed: {
          like: 65,
          comment: 7,
          follow: 65,
          unfollow: 150
        }
      },
      errorInvalidCredentials: false
    }
  },
  validations: {
    form: {
      username: { required },
      password: { required }
    }
  },
  watch: {
    errorInvalidCredentials: (val) => {
      if (val) {
        Toast.create.negative({
          html: Vue.t('common.errors.credentials.invalid'),
          timeout: 5000
        })
      }
    }
  },
  components: {
    'wrapper': Wrapper
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
      this.errorInvalidCredentials = false
      instagram.createInstagram(this, this.form, '/instagram')
    }
  }
}
</script>

<style>
.q-range-label {
  font-size: large;
}
.insta-form-container {
  margin: 20px;
}
.insta-button-nav {
  margin: 0 20px 20px 0;
}
.insta-input {
  padding: 1rem 0;
}
.insta-form-card {
  max-width: 600px;
}
.insta-form-title {
  font-weight: 300;
}
</style>
