import Vue from "vue";
import Router from "vue-router";
import About from '@/components/about'
import skill from '@/components/skill'
import sns from '@/components/sns'
// import HelloWorld from '@/components/'

// vueとvueRouterにVue.useをしているが、useは他はどのようなものに使う？
Vue.use(Router);

export default new Router({
  routes: [
    {
      path: "/",
      name: 'About',
      component: About
    },
    {
      path: "/skill",
      name: 'skill',
      component: skill
    },
    {
      path: "/sns",
      component: sns
    }
  ]
});
