<template>
  <div class="datepicker">
    <v-flex xs12 v-if="type==1">
      <v-menu
        ref="menu"
        :close-on-content-click="false"
        v-model="menu"
        :nudge-right="40"
        :return-value.sync="data"
        lazy
        transition="scale-transition"
        offset-y
        full-width
        min-width="290px"
      >
        <v-text-field
          slot="activator"
          v-model="data"
          prepend-icon="event"
          readonly
        ></v-text-field>
        <v-date-picker v-model="data" no-title scrollable>
          <v-spacer></v-spacer>
          <v-btn flat color="primary" @click="menu = false">Cancel</v-btn>
          <v-btn flat color="primary" @click="$refs.menu.save(data)">OK</v-btn>
        </v-date-picker>
      </v-menu>
    </v-flex>
    <v-flex xs12 v-if="type==2">
      <v-dialog
        ref="dialog"
        v-model="modal"
        :return-value.sync="data"
        full-width
        width="290px"
      >
        <v-text-field
          flat
          slot="activator"
          v-model="data"
          prepend-inner-icon="event"
          readonly
        ></v-text-field>
        <v-date-picker v-model="data" scrollable>
          <v-spacer></v-spacer>
          <v-btn flat color="primary" @click="modal = false">Cancel</v-btn>
          <v-btn flat color="primary" @click="$refs.dialog.save(data)">OK</v-btn>
        </v-date-picker>
      </v-dialog>
    </v-flex>
    <v-flex xs12 v-if="type==3">
      <v-menu
        ref="menu2"
        :close-on-content-click="false"
        v-model="menu2"
        :nudge-right="40"
        :return-value.sync="data"
        lazy
        transition="scale-transition"
        offset-y
        full-width
        min-width="290px"
      >
        <v-text-field
          slot="activator"
          v-model="data"
          prepend-icon="event"
          readonly
        ></v-text-field>
        <v-date-picker v-model="data" @input="$refs.menu2.save(data)"></v-date-picker>

      </v-menu>
    </v-flex>
  </div>
</template>

<script>
  export default {
    props: { data: String, model: String },
    data () {
      return {
        type: 2,
        menu: false,
        modal: false,
        menu2: false
      }
    },
    methods: {
      emitFormVal () {
        this.$emit('updateFormVal', {model: this.model, data: this.data})
      }
    },
    watch: {
      data: function (val, oldVal) {
        this.emitFormVal()
      }
    }
  }
</script>

<style>
.datepicker .v-input__prepend-inner{
  margin-top: 18px!important;
  margin-left: 15px!important;
}
</style>
