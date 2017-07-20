<template>
    <wrapper padding=true>
      <div>
        <div class="landing-container column">
          <div class="card shadow-3 gt-sm">
            <div class="card-content">
              <h2 class="landing-title text-pink">{{ $t("landing.title") }}</h2>
              <h5 class="landing-motto">{{ $t("landing.motto") }}</h5>
            </div>
          </div>
          <div class="card shadow-3 sm">
            <div class="card-content">
              <h5 class="landing-title text-pink">{{ $t("landing.title") }}</h5>
              <span class="landing-motto">{{ $t("landing.motto") }}</span>
            </div>
          </div>
          <!--<div class="landing-submit">-->
            <!--<router-link to="/signup">-->
              <!--<button class="pink big">{{ $t("landing.submit") }}</button>-->
            <!--</router-link>-->
          <!--</div>-->
        </div>

        <div class="row landing-features wrap">
          <feature-card :title="$t('landing.features[0].title')" :content="$t('landing.features[0].content')" />
          <feature-card :title="$t('landing.features[1].title')" :content="$t('landing.features[1].content')" />
          <feature-card :title="$t('landing.features[2].title')" :content="$t('landing.features[2].content')" />
        </div>

        <div class="row landing-features wrap">
          <feature-card :title="$t('landing.features[3].title')" :content="$t('landing.features[3].content')" />
          <feature-card :title="$t('landing.features[4].title')" :content="$t('landing.features[4].content')" />
          <feature-card :title="$t('landing.features[5].title')" :content="$t('landing.features[5].content')" />
        </div>
      </div>
    </wrapper>
</template>

<script>
import FeatureCard from './FeatureCard.vue'
import Wrapper from './Wrapper.vue'

export default {
  data () {
    return {}
  },
  components: {
    'feature-card': FeatureCard,
    'wrapper': Wrapper
  }
}
</script>

<style>
.landing-container {
  align-items: center;
  justify-content: flex-baseline;
  padding: 1rem;
}
.landing-title {
  font-weight: 300;
}
.landing-motto {
  font-weight: 300;
}
.landing-submit {
  padding: 2rem;
}
.landing-features {
  padding-top: 1rem;
  justify-content: space-around;
  align-items: stretch;
}
</style>
