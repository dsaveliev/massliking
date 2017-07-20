import Vue from 'vue'
import VueI18n from 'vue-i18n'

const locales = {
  en: {
    yes: 'Yes',
    no: 'No',
    dashboard: {
      titles: {
        greeting: 'Hi',
        accounts: 'Accounts',
        security: 'Security',
        billing: 'Billing',
        contacts: 'Contacts'
      }
    },
    instagram: {
      buttons: {
        show: 'Show account info',
        list: 'Accounts list',
        edit: 'Edit account',
        create: 'Create account',
        delete: 'Delete account',
        add: 'Add new account',
        start: 'Start instagram',
        stop: 'Stop instagram'
      },
      toasts: {
        deleting: 'Deleting account',
        start: 'Start instagram',
        stop: 'Stop instagram',
        inactive: 'Warning: There is impossible any account changes in the current state.'
      },
      titles: {
        create: 'New instagram account',
        edit: 'Edit instagram account'
      },
      fields: {
        username: 'Username',
        password: 'Password',
        hours: 'Activity hours (UTC)',
        like_speed: 'Like speed',
        comment_speed: 'Comment speed',
        follow_speed: 'Follow speed',
        unfollow_speed: 'Unfollow speed',
        trusted: 'Trusted account',
        media_count: 'Media count',
        followers_count: 'Followers count',
        following_count: 'Following count',
        state: 'State',
        channels_count: 'Channels count'
      },
      errors: {
        username: {
          required: 'Username is required.'
        },
        password: {
          required: 'Password is required.'
        }
      }
    },
    channel: {
      buttons: {
        add: 'Add new channel',
        edit: 'Edit channel',
        create: 'Create channel',
        delete: 'Delete channel',
        back_to_account: 'Back to account',
        start: 'Start channel',
        stop: 'Stop channel'
      },
      toasts: {
        deleting: 'Deleting channel',
        start: 'Start channel',
        stop: 'Stop channel'
      },
      titles: {
        create: 'New channel',
        edit: 'Edit channel'
      },
      fields: {
        value: 'Channel value',
        target: 'Target',
        action: 'Action',
        state: 'State',
        conversion: 'Conversion',
        actions_count: 'Actions count',
        queue: 'Queue',
        comments: 'Comments'
      },
      errors: {
        value: {
          required: 'Channel value is required.'
        },
        target: {
          required: 'Target value is required.'
        },
        action: {
          required: 'Action value is required.'
        }
      }
    },
    common: {
      buttons: {
        login: 'Log In',
        signup: 'Sign Up',
        logout: 'Log Out'
      },
      titles: {
        login: 'Log In',
        signup: 'Sign Up',
        header: 'MassLiking',
        footer: 'MassLiking © 2017'
      },
      fields: {
        email: 'Email',
        password: 'Password',
        password_confirmation: 'Password confirmation'
      },
      errors: {
        fields: {
          invalid: 'Please review fields again.'
        },
        credentials: {
          invalid: 'Incorrect Username and/or Password.',
          exists: 'User with this credentials already exists.'
        },
        email: {
          required: 'Email is required.',
          invalid: 'Email is invalid.'
        },
        password: {
          required: 'Password is required.',
          length: 'Password must have at least 6 letters.'
        },
        password_confirmation: {
          mismatch: 'Passwords must be identical.'
        }
      }
    },
    landing: {
      title: 'Get social media superpowers',
      motto: 'Accelerate your life on Instagram for more targeted likes, comments and follows',
      submit: 'Get 3 free days',
      features: [
        {
          title: 'No downloads',
          content: 'You can use MassLiking straight from the web on all browsers. You don\'t need to download or install anything to enjoy our service, which is why MassLiking is the safest Instagram bot available.'
        },
        {
          title: 'Automate everything',
          content: 'You can easily automate your liking, commenting and following activities based on specific hashtags and geolocations, as well as unfollow users from different sources.'
        },
        {
          title: 'Start and close',
          content: 'MassLiking conveniently works on our servers, which means you can feel free to logout, change accounts, or even close your browser window after you start your MassLiking activity.'
        },
        {
          title: '3 day free trial',
          content: 'We\'re confident that you\'ll love our service. Try MassLiking free for 3 days and you\'ll see why Instagrammers continue to use our service.'
        },
        {
          title: 'Full control',
          content: 'We provide tons of filters and customization options to help you increase your Instagram followers, likes and comments for the target audiences that make sense for you.'
        },
        {
          title: 'Safe to use',
          content: 'Our one of a kind service will automatically reduce its speed to ensure that your account is safe from hitting Instagram limits and we offer multiple speed settings for more advanced users.'
        }
      ]
    }
  }
}

Vue.use(VueI18n)
Vue.config.lang = 'en'

Object.keys(locales).forEach(function (lang) {
  Vue.locale(lang, locales[lang])
})
