<template>
  <wrapper>
    <div class="row login-container">

      <div class="card shadow-3">
        <div class="card-title bg-wite text-pink flex">
          <big>{{ $t("common.titles.login") }}</big>
        </div>
        <div class="card-content card-force-top-padding column">
          <div class="floating-label login-input">
            <input required class="full-width"
              :class="{'has-error': $v.form.username.$error}"
              v-model.trim="form.username"
              @input="$v.form.username.$touch()">
            <div v-if="$v.form.$error && !$v.form.username.required" class="text-red">{{ $t("common.errors.email.required") }}</div>
            <div v-if="$v.form.$error && !$v.form.username.email" class="text-red">{{ $t("common.errors.email.invalid") }}</div>
            <label>{{ $t("common.fields.email") }}</label>
          </div>

          <div class="floating-label login-input">
            <input required type="password" class="full-width"
              :class="{'has-error': $v.form.password.$error}"
              v-model.trim="form.password"
              @input="$v.form.password.$touch()">
            <div v-if="$v.form.$error && !$v.form.password.required" class="text-red">{{ $t("common.errors.password.required") }}</div>
            <label>{{ $t("common.fields.password") }}</label>
          </div>

          <div class="login-input">
            <button class="deep-purple big full-width" @click="submit">{{ $t("common.buttons.login") }}</button>
          </div>
        </div>
      </div>

    </div>
  </wrapper>
</template>

<script>
import Vue from 'vue'
import { required, email } from 'vuelidate/lib/validators'
import { Toast } from 'quasar'
import Wrapper from './Wrapper.vue'
import auth from '../api/auth'

export default {
  data () {
    return {
      form: {
        username: '',
        password: ''
      },
      errorInvalidCredentials: false
    }
  },
  validations: {
    form: {
      username: { required, email },
      password: { required }
    }
  },
  components: {
    'wrapper': Wrapper
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
      auth.login(this, this.form, '/instagram')
    }
  }
}
</script>

<style>
.login-container {
  justify-content: center;
  min-width: 50%;
}
.login-input {
  padding: 1rem 0;
}
.card {
  overflow: auto;
  min-width: 300px;
  justify-content: center;
  flex-wrap: nowrap;
}
.card-title {
  font-weight: 300;
  justify-content: center;
}
</style>
