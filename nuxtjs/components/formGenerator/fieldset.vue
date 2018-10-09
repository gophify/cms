<template>
    <v-flex xs12>
        <v-layout my-1 row v-for="(field, fi) in fields" :key="fi">
            <v-flex xs2 pr-3 pt-1>
                <v-icon small color="blue" v-if="fi > 0" v-on:click="arrow_upward_clicked(fi)" style="float: right;">arrow_upward</v-icon> &nbsp; 
                <v-icon small color="blue" v-if="fi < (fields.length-1)" v-on:click="arrow_downward_clicked(fi)" style="float: right;">arrow_downward</v-icon> &nbsp; 
                <v-icon small color="red" v-on:click="delete_clicked(fi)" style="float: right;">delete</v-icon>
                <v-icon class="mt-3" color="primary" v-if="fi == (fields.length-1)" dark @click="addField" style="float: right;">add</v-icon>
            </v-flex>
            <v-flex xs1>
                <select v-model="fields[fi]['type']" v-on:change="emitData">
                    <option value="input">input</option>
                    <option value="image">image</option>
                    <option value="textarea">textarea</option>
                </select>
            </v-flex>
            <v-flex xs2>
                <input type="text" v-model="fields[fi]['model']" v-on:change="emitData" />
            </v-flex>
            <v-flex xs2>
                <input type="text" v-model="fields[fi]['label']" v-on:change="emitData" />
            </v-flex>
            <v-flex xs2>
                <select v-model="fields[fi]['col']" v-on:change="emitData" multiple style="height: 36px;">
                    <option :value="'xs'+i" v-for="i in 12" :key="'xs'+i">{{ 'xs'+i }}</option>
                    <option :value="'sm'+i" v-for="i in 12" :key="'sm'+i">{{ 'sm'+i }}</option>
                    <option :value="'md'+i" v-for="i in 12" :key="'md'+i">{{ 'md'+i }}</option>
                </select>
            </v-flex>
            <v-flex xs1>
              <input v-on:change="emitData" type="number" v-model="fields[fi]['size']" v-if="fields[fi]['type'] == 'input' || fields[fi]['type'] == 'textarea'" />                                                                                                                                                             
            </v-flex>
            <v-flex xs2></v-flex>            
        </v-layout>                       
    </v-flex>     
</template>

<script>
  export default {
    props: { index: Number, dataInit: Array },
    data () {
      return {
        fields: [{}]
      }
    },
    methods: {
      emitData () {
        this.$emit('updataData', {index: this.index, data: this.fields})
      },
      arrow_upward_clicked (i) {
        var len = this.fields.length
        var newData = []
        for (let j = 0; j < len; j++) {
          if (j === i) {
            newData.push(this.fields[i - 1])
          } else if (j === (i - 1)) {
            newData.push(this.fields[i])
          } else {
            newData.push(this.fields[j])
          }
        }
        this.fields = newData
      },
      arrow_downward_clicked (i) {
        var len = this.fields.length
        var newData = []
        for (let j = 0; j < len; j++) {
          if (j === i) {
            newData.push(this.fields[i + 1])
          } else if (j === (i + 1)) {
            newData.push(this.fields[i])
          } else {
            newData.push(this.fields[j])
          }
        }
        this.fields = newData
      },
      delete_clicked (i) {
        var len = this.fields.length
        var newData = []
        for (let j = 0; j < len; j++) {
          if (i !== j) {
            newData.push(this.fields[j])
          }
        }
        this.fields = newData
      },
      addField () {
        this.fields.push({})
      }
    },
    beforeMount () {
      if (this.dataInit) {
        this.fields = this.dataInit
      }
    }
  }
</script>