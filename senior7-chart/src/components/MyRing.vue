<template>
    <div>
        <v-widget title="自定义环状图" content-bg="white">
        <div slot="widget-content">
            <ve-ring :data="chartData" :settings="chartSettings"></ve-ring>
        </div>
        </v-widget>
    </div>
</template>

<script>
import VWidget from "./VWidget.vue";

export default {
  components: {
    VWidget
  },

  data() {
    return {
      chartData: {
        columns: ["日期", "访问用户"],
        rows: [
          { 日期: "1/1", 访问用户: 1393 },
          { 日期: "1/2", 访问用户: 3530 },
          { 日期: "1/3", 访问用户: 2923 },
          { 日期: "1/4", 访问用户: 1723 },
          { 日期: "1/5", 访问用户: 3792 },
          { 日期: "1/6", 访问用户: 4593 }
        ]
      },
      chartSettings: {
        hoverAnimation: false,
        legendLimit: 7
      }
    };
  },

  mounted() {
    let ptr = this;
    window.setTimeout(() => {
      ptr.timer();
    }, 2000);
  },
  methods: {
    timer() {
      this.chartData.rows.push({
        日期: "1/" + (this.chartData.rows.length + 1),
        访问用户: 4593
      });
    }
  }
};
</script>