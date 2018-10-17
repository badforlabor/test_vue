import Vue from 'vue'
import {
  Vuetify,
  VApp,
  VNavigationDrawer,
  VFooter,
  VList,
  VBtn,
  VIcon,
  VGrid,
  VToolbar,
  VCard,
  VTextField,
  VForm,
  transitions
} from 'vuetify'
import 'vuetify/src/stylus/app.styl'
import zhHans from 'vuetify/es5/locale/zh-Hans'

Vue.use(Vuetify, {
  components: {
    VApp,
    VNavigationDrawer,
    VFooter,
    VList,
    VBtn,
    VIcon,
    VGrid,
    VToolbar,
    VCard,
    VTextField,
    VForm,
    transitions
  },
  theme: {
  },
  customProperties: true,
  iconfont: 'md',
  lang: {
    locales: { zhHans },
    current: 'zh-Hans'
  },
})
