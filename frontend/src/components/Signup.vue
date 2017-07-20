<template>
  <wrapper>
    <div class="row signup-container">

      <div class="card shadow-3">
        <div class="card-title bg-wite text-pink flex">
          <big>{{ $t("common.titles.signup") }}</big>
        </div>
        <div class="card-content card-force-top-padding column">
          <div class="floating-label signup-input">
            <input required class="full-width"
              :class="{'has-error': $v.form.username.$error}"
              v-model.trim="form.username"
              @input="$v.form.username.$touch()">
            <div v-if="$v.form.$error && !$v.form.username.required" class="text-red">{{ $t("common.errors.email.required") }}</div>
            <div v-if="$v.form.$error && !$v.form.username.email" class="text-red">{{ $t("common.errors.email.invalid") }}</div>
            <label>{{ $t("common.fields.email") }}</label>
          </div>

          <div class="floating-label signup-input">
            <input required type="password" class="full-width"
              :class="{'has-error': $v.form.password.$error}"
              v-model.trim="form.password"
              @input="$v.form.password.$touch()">
            <div v-if="$v.form.$error && !$v.form.password.required" class="text-red">{{ $t("common.errors.password.required") }}</div>
            <div v-if="$v.form.$error && !$v.form.password.minLength" class="text-red">{{ $t("common.errors.password.length") }}</div>
            <label>{{ $t("common.fields.password") }}</label>
          </div>

          <div class="floating-label signup-input">
            <input required type="password" class="full-width"
              :class="{'has-error': $v.form.repeatPassword.$error}"
              v-model.trim="form.repeatPassword"
              @input="$v.form.repeatPassword.$touch()">
            <div v-if="$v.form.$error && !$v.form.repeatPassword.sameAsPassword" class="text-red">{{ $t("common.errors.password_confirmation.mismatch") }}</div>
            <label>{{ $t("common.fields.password_confirmation") }}</label>
          </div>

          <div class="signup-input">
            <button class="deep-purple big full-width" @click="submit">{{ $t("common.buttons.signup") }}</button>
          </div>
        </div>
      </div>

    </div>
  </wrapper>
</template>

<script>
import Vue from 'vue'
import { required, email, sameAs, minLength } from 'vuelidate/lib/validators'
import { Toast } from 'quasar'
import Wrapper from './Wrapper.vue'
import auth from '../api/auth'

export default {
  data () {
    return {
      form: {
        username: '',
        password: '',
        repeatPassword: ''
      },
      errorAlreadyExists: false
    }
  },
  validations: {
    form: {
      username: { required, email },
      password: { required, minLength: minLength(7) },
      repeatPassword: { required, sameAsPassword: sameAs('password') }
    }
  },
  components: {
    'wrapper': Wrapper
  },
  watch: {
    errorAlreadyExists: (val) => {
      if (val) {
        Toast.create.negative({
          html: Vue.t('common.errors.credentials.exists'),
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
      this.errorAlreadyExists = false
      auth.signup(this, this.form, '/instagram')
    }
  }
}
</script>

<style>
.signup-container {
  justify-content: center;
  min-width: 50%;
}
.signup-input {
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
