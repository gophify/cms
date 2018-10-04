<template>
    <v-combobox
    v-model="brand"
    :items="items"
    class="brand"
></v-combobox>
</template>

<script>
  export default {
    props: { data: String, model: String },
    data () {
      return {
        init: true,
        brand: '',
        items: [
          'Toyota',
          'Ford',
          'Suzuki',
          'Nissan',
          'Mazda',
          'BMW',
          'Mitsubishi',
          'Daihatsu',
          'Isuzu',
          'Volvo'
        ]
      }
    },
    methods: {
      emitFormVal () {
        this.$emit('updateFormVal', {model: this.model, data: this.brand})
      }
    },
    watch: {
      data: function (val, oldVal) {
        if (this.init) {
          this.init = false
          this.brand = val
        }
      },
      brand: function (val, oldVal) {
        this.emitFormVal()
      }
    }
  }
</script>

<style>
.brand input{
    margin-left: 30px;
    font-size: 22px;
}
</style>
