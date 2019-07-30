Vue.component("hdr", {
  template: "#hdr"
});

Vue.component("index", {
  template: "#index"
});
Vue.component("skill", {
  template: "#skill"
});
Vue.component("sns", {
  template: "#sns"
});
Vue.component("tech", {
  template: "#tech"
});

var router = new VueRouter({
  routes: [
    {
      path: "/",
      component: {
        template: index
      }
    },
    {
      path: "/skill",
      component: {
        template: skill
      }
    },
    {
      path: "/sns",
      component: {
        template: sns
      }
    }
  ]
});

var app = new Vue({
  el: "#app",
  router: router
});
