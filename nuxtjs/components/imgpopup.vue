<template>
  <v-layout row justify-center>
    <v-dialog
      v-model="imgpopup"
      max-width="760"
    >
      <v-card>
        <v-card-text>
          <img :src="src" style="max-width: 100%;"/>
        </v-card-text>

        <v-card-actions>
          <v-spacer></v-spacer>

          <v-btn
            color="green darken-1"
            flat="flat"
            @click="imgpopup = false">
            Close
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </v-layout>
</template>

<script>
  export default {
    props: { imgpopup: Boolean, src: String },
    data () {
      return {
      }
    },
    watch: {
      imgpopup: function (val) {
        this.$emit('imgpopup', val)
      }
    }
  }
</script>