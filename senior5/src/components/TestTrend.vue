<template>
  <div id="test">
  <h3>我是内部的趋势图</h3>
    <trend
      :data="[0, 2, 5, 9, 5, 10, 3, 5, 0, 0, 1, 8, 2, 9, 0]"
      :gradient="['#6fa8dc', '#42b983', '#2c3e50']"
      auto-draw
      smooth>
    </trend>
  </div>
</template>

<script>
export default {
  name: "test",
  components: {}
};
</script>
