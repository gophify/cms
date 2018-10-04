<template>
    <v-container>
    <fieldset v-for="i in ((data != undefined) ? data.length : [])" :key="i">
        <div class="action">
            <v-icon class="mx-1" small color="blue" v-if="i > 1" v-on:click="arrow_upward_clicked(i-1)">arrow_upward</v-icon>
            <v-icon class="mx-1" small color="blue" v-if="i < data.length" v-on:click="arrow_downward_clicked(i-1)">arrow_downward</v-icon>
            <v-icon class="mx-1" small color="red" v-on:click="delete_clicked(i-1)">delete</v-icon>
        </div>
        <v-layout row wrap class="fscontent">
            <v-flex mb-2 :class="(f.col != undefined) ? ((f.col.length > 1) ? f.col : (f.col).join(' ')) : []" v-for="(f, fi) in fields" :key="fi">
                <label>{{ f.label }}</label>
                <formImage v-if="f.type=='image'" :selectedImageStr="data[i-1][f.model]" @media="getMedia" :model="f.model" :index="i-1" class="mb-2" />
                <input v-on:change="populateData" type="text" v-if="f.type=='input'" v-model="data[i-1][f.model]" />
                <textarea v-on:change="populateData" rows="4" v-if="f.type=='textarea'" v-model="data[i-1][f.model]"></textarea>
            </v-flex>
        </v-layout>
    </fieldset>

    <div style="margin: auto; display: table;">
        <v-btn large dark color="info" v-on:click="addMore">
            <v-icon left>add</v-icon> add section
        </v-btn>
    </div>
    </v-container>    
</template>

<script>
  export default {
    props: { fields: Array, data: Array, model: String },
    data: () => ({
    }),
    methods: {
      populateData () {
        var md = {model: this.model, data: this.data}
        this.$emit('updateFormVal', md)
      },
      getMedia (v) {
        this.data[v.index][v.model] = v.media
        this.populateData()
      },
      addMore () {
        if (this.data === undefined) {
          this.data = []
        }
        this.data.push({})
      },
  
      arrow_upward_clicked (i) {
        var len = this.data.length
        var newData = []
        for (let j = 0; j < len; j++) {
          if (j === i) {
            newData.push(this.data[i - 1])
          } else if (j === (i - 1)) {
            newData.push(this.data[i])
          } else {
            newData.push(this.data[j])
          }
        }
        this.data = newData
        console.log(this.data)
      },
      arrow_downward_clicked (i) {
        var len = this.data.length
        var newData = []
        for (let j = 0; j < len; j++) {
          if (j === i) {
            newData.push(this.data[i + 1])
          } else if (j === (i + 1)) {
            newData.push(this.data[i])
          } else {
            newData.push(this.data[j])
          }
        }
        this.data = newData
      },
      delete_clicked (i) {
        var len = this.data.length
        var newData = []
        for (let j = 0; j < len; j++) {
          if (i !== j) {
            newData.push(this.data[j])
          }
        }
        this.data = newData
        this.data.length = this.data.length
      },
      emitFormVal () {
        var data = {model: this.model, data: this.data}
        this.$emit('updateFormVal', data)
      }
    },
    watch: {
      data: function (val, oldVal) {
        this.emitFormVal()
      }
    }
  }
</script>