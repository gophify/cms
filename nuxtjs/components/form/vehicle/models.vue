<template>
  <div>
      <v-combobox
      v-model="vhcmodel"
      :items="items"
      class="vhcmodel"
      ></v-combobox>
  </div>
</template>

<script>
  export default {
    props: { data: String, model: String },
    data () {
      return {
        init: true,
        vhcmodel: '',
        items: [
          'Small Cars',
          'Midsize Cars',
          'Large Cars',
          'SUVs',
          'Crossovers',
          'Hybrids',
          'Trucks',
          'Luxury',
          'Sports Cars',
          'Convertibles',
          'Vans',
          'Certfied Pre-Owned'
        ]
      }
    },
    methods: {
      emitFormVal () {
        this.$emit('updateFormVal', {model: this.model, data: this.vhcmodel})
      }
    },
    watch: {
      data: function (val, oldVal) {
        if (this.init) {
          this.init = false
          this.vhcmodel = val
        }
      },
      vhcmodel: function (val, oldVal) {
        this.emitFormVal()
      }
    }
  }
</script>

<style>
.vhcmodel input{
    margin-left: 30px;
    font-size: 22px;
}
</style>