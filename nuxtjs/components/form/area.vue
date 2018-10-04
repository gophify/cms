<template>
    <div class="field_wrapper">
        <v-combobox
        class="area"
        v-model="area"
        :items="items"
        item-text="label"
        item-value="key"></v-combobox>
    </div>
</template>

<script>
  export default {
    props: { data: String, model: String },
    data () {
      return {
        init: true,
        area: {key: '', label: ''},
        items: [
          {key: 'ubud', label: 'Ubud, Bali'},
          {key: 'seminyak', label: 'Seminyak, Bali'},
          {key: 'kuta', label: 'Kuta, Bali'},
          {key: 'canggu', label: 'Canggu, Bali'}
        ]
      }
    },
    methods: {
      emitFormVal () {
        this.$emit('updateFormVal', {model: this.model, data: this.area.key})
      }
    },
    watch: {
      data: function (val, oldVal) {
        if (this.init) {
          this.init = false
          for (let a = 0; a < this.items.length; a++) {
            if (this.items[a].key === this.data) {
              this.area = this.items[a]
            }
          }
        }
      },
      area: function (val, oldVal) {
        this.emitFormVal()
      }
    }
  }
</script>