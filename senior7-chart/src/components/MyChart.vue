<template>
    <div>
        <v-widget title="位置信息" content-bg="white">
        <div slot="widget-content">
            <e-chart 
            :path-option="[
                ['dataset.source', locationData],
                ['legend.bottom', '0'],
                ['color', [color.lightBlue.base, color.indigo.base, color.pink.base, color.green.base, color.cyan.base, color.teal.base]],
                ['xAxis.show', false],
                ['yAxis.show', false],
                ['series[0].type', 'pie'],
                ['series[0].avoidLabelOverlap', true],         
                ['series[0].radius', ['50%', '70%']],                      
            ]"
            height="400px"
            width="100%"
            >
            </e-chart>     
        </div>
        </v-widget>  
        <br><br><h4>以下也是图表，不过不是以card形式存在的. {{msg}}</h4>
        <e-chart 
            :path-option="[
                ['dataset.source', locationData],
                ['legend.bottom', '0'],
                ['color', [color.lightBlue.base, color.indigo.base, color.pink.base, color.green.base, color.cyan.base, color.teal.base]],
                ['xAxis.show', false],
                ['yAxis.show', false],
                ['series[0].type', 'pie'],
                ['series[0].avoidLabelOverlap', true],         
                ['series[0].radius', ['50%', '70%']],                      
            ]"
            height="400px"
            width="100%"
            >
        </e-chart> 
    </div>
</template>

<script>
import Material from "vuetify/es5/util/colors";
import EChart from "./chart/echart.js";
import VWidget from "./VWidget.vue";

export default {
  name: "mychart",
  components: {
    EChart,
    VWidget
  },
  data() {
    return {
      msg: "abc?",
      color: Material,
      tempLocation: [
        {
          value: 50,
          name: "China"
        }
      ]
    };
  },
  computed: {
    locationData() {
      const locationData = [
        {
          value: 50,
          name: "China"
        },
        {
          value: 35,
          name: "USA"
        },
        {
          value: 25,
          name: "EU"
        },
        {
          value: 10,
          name: "Russia"
        },
        {
          value: 10,
          name: "Other"
        }
      ];
      return locationData;
    }
  },
  created() {
    console.log("created!");
  },
  mounted() {
    console.log("mounted!");
    var ptr = this;
    window.setTimeout(() => {
      console.log("oo");
      ptr._location = [
        {
          value: 100,
          name: "中国"
        }
      ];
    }, 3000);
  }
};
</script>