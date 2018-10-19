<template>
  <div id="app">
    <img alt="Vue logo" src="./assets/logo.png">
    <Icons/>
    <Icons2/>
    <HelloWorld msg="Welcome to Your Vue.js App"/>
  </div>
</template>

<script>
import HelloWorld from "./components/HelloWorld.vue";
import Icons from "./components/Icons.vue";
import Icons2 from "./components/Icons2.vue";

export default {
  name: "app",
  components: {
    HelloWorld,
    Icons,
    Icons2
  }
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
