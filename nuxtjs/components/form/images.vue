<template>
    <span>
        <v-container fluid>
            <v-layout wrap>
                <v-flex xs12 :md6="colWrapper == 'xs12'" m-y-4 v-for="(photo, i) in data" :key="i">
                    <v-layout xs12 mx-1 my-1 px-1 pt-1 class="grey lighten-4" style="position: relative;">
                        <v-flex v-on:click="src=$store.state.mediaPath + photo.dir + photo.mfn, imgpopup=true" :class="['xs4', 'mr-3', 'imgwrapper', {'ori': !photo.IsThumb & photo.Lp == ''}, {'l': !photo.IsThumb & photo.Lp == 'l'}, {'p': !photo.IsThumb & photo.Lp == 'p'}]" :style="'background: url('+ $store.state.mediaPath + photo.dir + ((photo.IsThumb) ? 'thumbnails-128/' : '') + photo.mfn +');'">
                          <v-icon color="black">zoom_in</v-icon>
                        </v-flex>
                        <v-flex class="xs8 pt-4 pb-2">                    
                            <span style="font-size: 9px; display: table; float: left; width: 34px; margin-left: 6px; padding-top: 12px;">title:</span>
                            <input v-on:change="emitFormVal" v-model="data[i]['title']" style="width:calc(100% - 60px);margin-top:9px;margin-bottom:2px;border-radius:4px;padding:3px 8px;background-color: #F1F8E9;border:1px solid #AED581!important;">                                       
                            <span style="font-size: 9px; display: table; float: left; width: 34px; margin-left: 6px; padding-top: 12px;">alt:</span>
                            <input v-model="data[i]['alt']" type="text" style="width: 60%;margin-top:5px;margin-bottom:2px;border-radius:4px;padding:3px 8px;background-color: #F1F8E9;border:1px solid #AED581!important;">
                        </v-flex>

                        <div style="position: absolute; right: 10px; top: 5px;">
                            <v-icon small color="blue" v-if="i > 0" v-on:click="arrow_upward_clicked(i)">arrow_upward</v-icon> &nbsp; 
                            <v-icon small color="blue" v-if="i < (data.length-1)" v-on:click="arrow_downward_clicked(i)">arrow_downward</v-icon> &nbsp; 
                            <v-icon small color="red" v-on:click="delete_clicked(i)">close</v-icon>
                        </div>                        
                    </v-layout>
                </v-flex>
            </v-layout>
          <v-flex xs12 ml-4 mt-3>
              <v-btn style="width:240px;font-size:11px;font-weight:bold;" dark color="success" v-on:click.stop="openMedia = true">select from media <v-icon right>cloud_upload</v-icon></v-btn>
          </v-flex>
        </v-container>

        <formMediaModal :isMultiple="true" :lastNumber="lastNumber" @openMedia="openMediaMethod" @selectedImages="selectedImagesMethod" v-if="openMedia" />
        <imgpopup :imgpopup="imgpopup" :src="src" @imgpopup="imgpopupMethod"></imgpopup>
    </span>
</template>

<script>
export default {
  props: { data: Array, model: String, colWrapper: Array },
  data () {
    return {
      init: true,
      openMedia: false,
      lastNumber: 0,
      imgpopup: false,
      src: ''
    }
  },
  methods: {
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
      this.emitFormVal()
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
      this.emitFormVal()
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
      this.emitFormVal()
    },
    openMediaMethod (value) {
      this.openMedia = value
    },
    selectedImagesMethod (v) {
      if (this.data === undefined) {
        this.data = v
      } else {
        this.data = this.data.concat(v)
      }
      this.emitFormVal()
    },
    emitFormVal () {
      var imgs = {model: this.model, data: this.data}
      this.$emit('updateFormVal', imgs)
    },
    imgpopupMethod (val) {
      this.imgpopup = val
    }
  }
}
</script>