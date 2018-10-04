<template>
    <v-layout row>
        <v-flex style="height:70px;max-width:180px;background-position:left!important" v-on:click="src=$store.state.mediaPath + selectedImage.dir + selectedImage.mfn, imgpopup=true" :class="['xs6', 'pr-3', 'imgwrapper', {'ori': !selectedImage.IsThumb & selectedImage.Lp == ''}, {'l': !selectedImage.IsThumb & selectedImage.Lp == 'l'}, {'p': !selectedImage.IsThumb & selectedImage.Lp == 'p'}]" :style="'background-image: url('+ $store.state.mediaPath + selectedImage.dir + ((selectedImage.IsThumb) ? 'thumbnails-128/' : '') + selectedImage.mfn +');'">
            <v-icon v-if="selectedImage.mfn" color="black">zoom_in</v-icon>
        </v-flex>
        <v-flex xs6 :class="{'pt-4': selectedImage.mfn}">
            <v-btn color="blue" outline small v-on:click="openMedia=true">
                <span v-if="selectedImage.mfn">change image</span>
                <span v-if="!selectedImage.mfn">choose image</span>
            </v-btn>
        </v-flex>

        <formMediaModal :isMultiple="false" @openMedia="openMediaMethod" @selectedImage="selectedImageMethod" v-if="openMedia" />
        <imgpopup :imgpopup="imgpopup" :src="src" @imgpopup="imgpopupMethod"></imgpopup>
    </v-layout>
</template>

<script>
export default {
  props: { model: String, index: Number, selectedImageStr: Object },
  data () {
    return {
      openMedia: false,
      selectedImage: {},
      imgpopup: false,
      src: ''
    }
  },
  methods: {
    openMediaMethod (value) {
      this.openMedia = value
    },
    selectedImageMethod (si) {
      this.selectedImage = si
      this.$emit('media', {model: this.model, index: this.index, media: si})
    },
    imgpopupMethod (val) {
      this.imgpopup = val
    }
  },
  watch: {
    selectedImageStr: function (val, oldVal) {
      if (!val) {
        val = {}
      }
      this.selectedImage = val
      this.$emit('media', {model: this.model, index: this.index, media: val})
    }
  },
  beforeMount () {
    if (this.selectedImageStr) {
      this.selectedImage = this.selectedImageStr
    } else {
      this.selectedImage = {}
    }
  }
}
</script>