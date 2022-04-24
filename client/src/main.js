import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'

import ElementPlus from 'element-plus'
import 'element-plus/lib/theme-chalk/index.css'

import Card from '@/components/Base/Card.vue'
import Center from '@/components/Base/Center.vue'
import LinkCard from '@/components/Base/LinkCard.vue'
import Grid from '@/components/Base/Grid.vue'
import ScrollContainer from '@/components/Base/ScrollContainer.vue'

const app = createApp(App)

app.use(router)
app.use(store)

app.component('Card', Card)
app.component('LinkCard', LinkCard)
app.component('Grid', Grid)
app.component('Center', Center)
app.component('ScrollContainer', ScrollContainer)

app.use(ElementPlus)

app.mount('#app')
