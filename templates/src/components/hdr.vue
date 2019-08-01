<template>
  <div class="hdr">
    <!-- ヘッダ部分のコンポーネントで、VueRouterを適用していく -->
    <ul>
      <li>
        <router-link to="/">トップページ</router-link>
      </li>
      <li>
        <router-link to="/skill">スキル</router-link>
      </li>
      <li>
        <router-link to="/sns">sns</router-link>
      </li>
    </ul>
  </div>
</template>

<script>
export default {
  name: "hdr"
};
</script>