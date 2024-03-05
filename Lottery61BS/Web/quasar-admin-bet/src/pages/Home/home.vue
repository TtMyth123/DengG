<template>
  <q-page class="q-pa-sm">

  </q-page>
</template>

<script>
import ECharts from "vue-echarts";
import * as echarts from 'echarts';
export default {
  name: "home",
  components:{
    ECharts
  },
  data() {
    return{
      Day7BetYieldChartData:{
        option:{
          "xAxis": {
            "show": false,
            "type": "category",
            "axisLine": {"lineStyle": {"color": "rgba(0, 0, 0 , .54)", "type": "dashed"}},
            "axisTick": {
              "show": true,
              "alignWithLabel": true,
              "lineStyle": {"show": true, "color": "rgba(0, 0, 0 , .54)", "type": "dashed"}
            },
            "axisLabel": {"show": false}
          },
          "yAxis": {
            "show": false,
            "type": "value",
            "axisLine": {"lineStyle": {"color": "rgba(0, 0, 0 , .54)", "type": "dashed"}},
            "axisLabel": {"show": false},
            "splitLine": {"lineStyle": {"type": "dashed"}},
            "axisTick": {
              "show": true,
              "lineStyle": {"show": true, "color": "rgba(0, 0, 0 , .54)", "type": "dashed"}
            }
          },
          "series": [{
            "type": "bar",
            "barGap": "-100%",
            "itemStyle": {"normal": {"color": "rgba(0,0,0,0.1)"}},
            "barWidth": "50%"
          }, {"barWidth": "50%", "type": "bar", "itemStyle": {"normal": {"color": "#ffffff"}}}],
          "dataset": {
            "source": [{"label": "0D", "max": 500, "sales": 220}, {
              "label": "1D",
              "max": 500,
              "sales": 182
            }, {"label": "2D", "max": 500, "sales": 191}, {
              "label": "3D",
              "max": 500,
              "sales": 234
            }, {"label": "4D", "max": 500, "sales": 290}, {
              "label": "5D",
              "max": 500,
              "sales": 330
            }, {"label": "6D", "max": 500, "sales": 310}, {
              "label": "7D",
              "max": 500,
              "sales": 123
            }, {"label": "8D", "max": 500, "sales": 442}, {
              "label": "9D",
              "max": 500,
              "sales": 321
            }, {"label": "10D", "max": 500, "sales": 90}, {
              "label": "11D",
              "max": 500,
              "sales": 149
            }, {"label": "12D", "max": 500, "sales": 210}, {
              "label": "13D",
              "max": 500,
              "sales": 122
            }, {"label": "14D", "max": 500, "sales": 133}, {
              "label": "15D",
              "max": 500,
              "sales": 334
            }, {"label": "16D", "max": 500, "sales": 198}, {
              "label": "17D",
              "max": 500,
              "sales": 123
            }, {"label": "18D", "max": 500, "sales": 125}, {"label": "19D", "max": 500, "sales": 220}]
          }
        },
      }
    }
  },

  methods:{
    reloadGet7DayBetYieldChart(){

    },
  },
  mounted() {
    this.reloadGet7DayBetYieldChart()
  }
}
</script>

<style scoped>
</style>
