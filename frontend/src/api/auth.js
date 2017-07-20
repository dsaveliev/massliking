import router from '../router'

const API_URL = process.env.API_URL
const LOGIN_URL = API_URL + 'login'
const SIGNUP_URL = API_URL + 'signup'
const REFRESH_URL = API_URL + 'user/refresh_token'

const TOKEN_ID = 'massliking_token_id'

export default {

  user: {
    authenticated: localStorage.getItem(TOKEN_ID) || false
  },

  checkAuth () {
    const jwt = localStorage.getItem(TOKEN_ID)
    return this.user.authenticated === jwt
  },

  login (context, creds, redirect) {
    context.$http.post(LOGIN_URL, creds)
      .then(response => {
        localStorage.setItem(TOKEN_ID, response.body.token)
        this.user.authenticated = response.body.token

        if (redirect) {
          router.push({path: redirect})
        }
      }, error => {
        if (error.status === 401) {
          context.errorInvalidCredentials = true
        }
      })
  },

  signup (context, creds, redirect) {
    context.$http.post(SIGNUP_URL, creds)
      .then(response => {
        if (this.user.authenticated) {
          localStorage.removeItem(TOKEN_ID)
          this.user.authenticated = false
        }

        this.login(context, creds, redirect)
      }, error => {
        if (error.status === 401) {
          context.errorAlreadyExists = true
        }
      })
  },

  logout (context, redirect) {
    localStorage.removeItem(TOKEN_ID)
    this.user.authenticated = false
    context.$store.commit('deleteUser')
    context.$store.commit('deleteInstagramList')

    if (redirect) {
      router.push({path: redirect})
    }
  },

  refreshToken (context, redirect) {
    context.$http.get(REFRESH_URL, {headers: this.getAuthHeader()})
      .then(response => {
        localStorage.setItem(TOKEN_ID, response.body.token)
        this.user.authenticated = response.body.token
      }, error => {
        if (error.status !== 200) {
          this.logout(context, redirect)
        }
      })
  },

  getAuthHeader () {
    return {
      'Authorization': 'Bearer ' + localStorage.getItem(TOKEN_ID)
    }
  }
}
