<template>
  <div id="app">
  <h1>测试趋势图</h1>
    <trend
      :data="[0, 2, 5, 9, 5, 10, 3, 5, 0, 0, 1, 8, 2, 9, 0]"
      :gradient="['#6fa8dc', '#42b983', '#2c3e50']"
      auto-draw
      smooth>
    </trend>
    <test/>
    <TestTrend/>
  </div>
</template>


<script>
import TestTrend from "./components/TestTrend.vue";
export default {
  name: "app",
  components: { TestTrend }
};
</script>

<style>
#app {
  font-family: "Avenir", Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
