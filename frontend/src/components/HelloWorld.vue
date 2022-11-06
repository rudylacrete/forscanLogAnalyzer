<script lang="ts" setup>
import {reactive} from 'vue'
import {Greet, LoadLogFile} from '../../wailsjs/go/main/App'

const data = reactive({
  name: "",
  resultText: "Please enter your name below 👇",
  chartOptions: {
    chart: {
      id: 'vuechart-example'
    },
    xaxis: {
      type: "numeric"
    },
    tooltip: {
      x: {
        show: false
      },
      y: {
        formatter: (val: any) => {
          return `${val.toFixed(2)} bar`
        }
      }
    }
  },
  series: [{
    name: "Turbo",
    data: Array<number>()
  }],
})

function greet() {
  Greet(data.name).then(result => {
    data.resultText = result;
    return LoadLogFile();
  }).then(logs => {
    console.log(logs);
    const turboIndex = logs.fields.indexOf("turbo");
    for(let i = 0; i < logs.values.length; i++) {
      data.series[0].data.push(logs.values[i][turboIndex]);
    }
  }).catch(err => {
    console.error(err);
  })
}

</script>

<template>
  <main>
    <div id="result" class="result">{{ data.resultText }}</div>
    <div id="input" class="input-box">
      <input id="name" v-model="data.name" autocomplete="off" class="input" type="text"/>
      <button class="btn" @click="greet">Greet</button>
    </div>
    <apexchart width="500" type="line" :options="data.chartOptions" :series="data.series"></apexchart>
  </main>
</template>

<style scoped>
.result {
  height: 20px;
  line-height: 20px;
  margin: 1.5rem auto;
}

.input-box .btn {
  width: 60px;
  height: 30px;
  line-height: 30px;
  border-radius: 3px;
  border: none;
  margin: 0 0 0 20px;
  padding: 0 8px;
  cursor: pointer;
}

.input-box .btn:hover {
  background-image: linear-gradient(to top, #cfd9df 0%, #e2ebf0 100%);
  color: #333333;
}

.input-box .input {
  border: none;
  border-radius: 3px;
  outline: none;
  height: 30px;
  line-height: 30px;
  padding: 0 10px;
  background-color: rgba(240, 240, 240, 1);
  -webkit-font-smoothing: antialiased;
}

.input-box .input:hover {
  border: none;
  background-color: rgba(255, 255, 255, 1);
}

.input-box .input:focus {
  border: none;
  background-color: rgba(255, 255, 255, 1);
}
</style>
