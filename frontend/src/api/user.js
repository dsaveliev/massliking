import auth from './auth'

const API_URL = process.env.API_URL
const USER_GET_URL = API_URL + 'user'

export default {

  isEmpty (context) {
    return Object.keys(context.$store.state.user).length === 0
  },

  updateUser (context) {
    context.$http.get(USER_GET_URL, {headers: auth.getAuthHeader()})
      .then(response => {
        context.$store.commit('updateUser', { user: response.body })
      }, error => {
        if (error.status === 401) {
          auth.logout(context, '/')
        }
      })
  }

}
