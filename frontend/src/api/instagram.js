import auth from './auth'
import channel from './channel'
import router from '../router'

const API_URL = process.env.API_URL

export default {

  isEmptyList (context) {
    return context.$store.state.instagramList.length === 0
  },

  updateInstagramList (context) {
    var url = API_URL + 'instagrams'
    context.$http.get(url, {headers: auth.getAuthHeader()})
      .then(response => {
        context.$store.commit('updateInstagramList', { instagramList: response.body })
        channel.updateAllChannels(context)
      }, error => {
        console.log(error)
      })
  },

  deleteInstagram (context, instagram, redirect) {
    var url = API_URL + 'instagrams/' + instagram.id
    context.$http.delete(url, {headers: auth.getAuthHeader()})
      .then(response => {
        context.$store.commit('deleteInstagram', { instagram: instagram })
        if (redirect) { router.push({path: redirect}) }
      }, error => {
        console.log(error)
      })
  },

  createInstagram (context, form, redirect) {
    var url = API_URL + 'instagrams'
    context.$http.post(url, form, {headers: auth.getAuthHeader()})
      .then(response => {
        var instagram = response.body
        context.$store.commit('addInstagram', { instagram: instagram })
        channel.updateAccountChannels(context, instagram, redirect)
      }, error => {
        console.log(error)
        if (error.status === 400) { context.errorInvalidCredentials = true }
      })
  },

  editInstagram (context, instagram, form, redirect) {
    var url = API_URL + 'instagrams/' + instagram.id
    context.$http.put(url, form, {headers: auth.getAuthHeader()})
      .then(response => {
        var instagram = response.body
        context.$store.commit('updateInstagram', { instagram: instagram })
        channel.updateAccountChannels(context, instagram, redirect)
      }, error => {
        console.log(error)
        if (error.status === 400) { context.errorInvalidCredentials = true }
      })
  },

  startInstagram (context, instagram) {
    var url = API_URL + 'instagrams/' + instagram.id + '/start'
    context.$http.get(url, {headers: auth.getAuthHeader()})
      .then(response => {
        var instagram = response.body
        context.$store.commit('updateInstagram', { instagram: instagram })
        channel.updateAccountChannels(context, instagram, '')
      }, error => {
        console.log(error)
      })
  },

  stopInstagram (context, instagram) {
    var url = API_URL + 'instagrams/' + instagram.id + '/stop'
    context.$http.get(url, {headers: auth.getAuthHeader()})
      .then(response => {
        var instagram = response.body
        context.$store.commit('updateInstagram', { instagram: instagram })
        channel.updateAccountChannels(context, instagram, '')
      }, error => {
        console.log(error)
      })
  }

}
