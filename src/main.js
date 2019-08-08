import Vue from 'vue'
import App from './App'
import router from './router'

// これなに？
Vue.config.productionTip = false

new Vue({
	el: '#app',
	// router: routerじゃなくてもいい？
	router,
	// このカギカッコの書き方何...?
	components: { App },
	// これは何...?
	// おそらく、id="app"にApp.vueのコンポーネントAppを差し込んでいる？
	template: '<App/>'
	// componentsとtemplateの違いは？
})