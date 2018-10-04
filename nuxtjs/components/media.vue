<template>
        <v-card>
            <v-card-title>
                <v-btn color="green darken-1" dark :outline="!tab_media_list" large v-on:click="tab_media_list=true, tab_media_upload=false">Media</v-btn>
                <v-btn color="green darken-1" dark :outline="!tab_media_upload" large v-on:click="tab_media_list=false, tab_media_upload=true">Upload</v-btn>
            </v-card-title><hr>

            <div v-if="photos_loading" :style="'background: url('+ $store.state.mediaPath +'/assets/images/main_loading.gif) no-repeat center center white; display: table; height: calc(100vh - 320px); width: 100%;'">
            </div>
            
            <v-card-text v-if="tab_media_upload">
                <div class="container">
                    <v-alert v-model="alertErrImgUpload" v-if="errImgUpload.length > 0"
                      transition="scale-transition"
                      dismissible
                      type="error"
                      outline
                    >
                        <div v-for="(item, i) in errImgUpload" :key="i">
                          Failed to upload <b>{{ item.Fn }}</b>, error: {{ item.Err }}
                        </div>
                    </v-alert>                  
                    <div style="display: table; margin: 0 auto 15px;">
                        <v-btn color="green darken-1" outline large v-on:click="pickFile()">Choose image to upload</v-btn>
                        <v-btn large v-on:click="submitFiles()" :loading="media_upload_loading" :disabled="media_upload_loading" color="blue-grey" class="white--text" v-on:click.native="loader = 'loading3'">
                            Upload
                        <v-icon right dark>cloud_upload</v-icon>
                    </v-btn>                
                        <input style="display: none" type="file" accept="image/*" id="files" ref="files" multiple v-on:change="handleFilesUpload()"/>
                    </div>

                    <div style="display: table; max-width: 800px; margin: auto;">
                        <div style="margin: 2px; width: 190px; height: 150px; float: left;" v-for="(file, key) in temp_images" :key="key">
                            <div :style="'position: relative; background-image: url('+ file.url +'); background-size: cover; height: 150px; width: 100%;'">
                                <span style="position: absolute; bottom:0; right:0; background-color: rgba(0,0,0,0.8); text-align:center; color: white; width: 100%; padding: 5px;">{{ file.name }}</span>
                                <v-btn style="position: absolute; top: -5px; right: -5px;" fab dark small color="red" v-on:click="removeFile( key )"><v-icon>close</v-icon></v-btn>    
                            </div>
                        </div>
                    </div>
                </div>
            </v-card-text>
            
            <v-card-text v-if="tab_media_list">
                <v-layout row wrap>
                    <v-flex xs6 sm4 md3 lg2 px-1 py-1 v-for="(img, i) in form_images" :key="i">
                        <div v-on:click="src=$store.state.mediaPath + img.Dir + img.Fn, imgpopup=true" :style="'background: url('+ $store.state.mediaPath + img.Dir + ((img.IsThumb) ? 'thumbnails-128/' : '') + img.Fn +');'" :class="['grid', 'imgwrapper', {'ori': !img.IsThumb & img.Lp == ''}, {'l': !img.IsThumb & img.Lp == 'l'}, {'p': !img.IsThumb & img.Lp == 'p'}]">
                          <v-icon color="black">zoom_in</v-icon>
                        </div>
                    </v-flex>
                </v-layout>
            </v-card-text>
            <imgpopup :imgpopup="imgpopup" :src="src" @imgpopup="imgpopupMethod"></imgpopup>
        </v-card>
</template>

<script>
import axios from 'axios'
import userlog from '../utils/userlog.js'
import auth from '../utils/auth.js'
export default {
  props: { isMultiple: Boolean },
  data () {
    return {
      openMedia: true,
      selectedImages: [],
      selectedImage: {},
      files: [],
      form_images: [],
      temp_images: [],
      tab_media_list: true,
      photos_loading: false,
      tab_media_upload: false,
      media_upload_loading: false,
      imgpopup: false,
      src: '',
      errImgUpload: [],
      alertErrImgUpload: false
    }
  },
  mounted () {
    auth.auth(this)

    this.photos_loading = true
    const axithis = this
    axios
      .get(this.$store.state.apiPath + '/media-api')
      .then(function (response) {
        axithis.photos_loading = false
        axithis.form_images = response.data
      })
  },
  methods: {
    addFiles () {
      this.$refs.files.click()
    },
    submitFiles () {
      if (this.$store.state.session.role === 'demo') {
        userlog.log(this, {type: 'media_upload_denied_on_media_manager'})
        alert('You are logged in as a guest\nSo sorry you cannot add/edit/delete the data here.\n\nThankyou :)')
        return false
      }

      let formData = new FormData()
      formData.append('key', 'upload')
      for (var i = 0; i < this.files.length; i++) {
        let file = this.files[i]
        formData.append('multiplefiles', file)
      }
      const axithis = this
      this.media_upload_loading = true
      axios.post(this.$store.state.apiPath + '/media-upload',
        formData,
        {
          headers: {
            'Content-Type': 'multipart/form-data'
          }
        }
      ).then(function (response) {
        var newImages = []
        axithis.errImgUpload = []
        for (var r = 0; r < response.data.length; r++) {
          if (response.data[r].Status === 'failed') {
            axithis.errImgUpload.push({Err: JSON.stringify(response.data[r].Err), Fn: response.data[r].Fn})
          } else {
            newImages.push({Id: response.data[r].Id, Dir: response.data[r].Dir, Fn: response.data[r].Fn, Lp: response.data[r].Lp, IsThumb: response.data[r].IsThumb})
          }
        }
        axithis.form_images = newImages.concat(axithis.form_images)

        if (axithis.errImgUpload.length === 0) {
          axithis.tab_media_list = true
          axithis.tab_media_upload = false
        } else {
          axithis.alertErrImgUpload = true
        }

        axithis.media_upload_loading = false

        axithis.temp_images = []
        axithis.files = []
      })
        .catch(function (err) {
          console.log(err)
        })
    },
    handleFilesUpload () {
      let uploadedFiles = this.$refs.files.files
      for (var i = 0; i < uploadedFiles.length; i++) {
        this.files.push(uploadedFiles[i])
        const fn = uploadedFiles[i].name
        const fr = new FileReader()
        fr.readAsDataURL(uploadedFiles[i])
        fr.addEventListener('load', () => {
          this.temp_images.push({url: fr.result, name: fn})
        })
      }
    },
    getFileUrl (file) {
      const fr = new FileReader()
      fr.readAsDataURL(file)
      fr.addEventListener('load', () => {
        return fr.result
      })
    },

    removeFile (key) {
      this.files.splice(key, 1)
      this.temp_images.splice(key, 1)
    },

    pickFile () {
      this.$refs.files.click()
    },

    imageSelect (mid) {
      if (!this.isMultiple) {
        this.unselectAll()
      }
      var e = document.getElementById(mid)
      if (e.classList.contains('active__img')) {
        e.classList.remove('active__img')
      } else {
        e.classList.add('active__img')
      }
    },
    unselectAll () {
      Array.from(document.querySelectorAll('.active__img')).forEach((element, index) => {
        element.classList.remove('active__img')
      })
    },
    collectImageSelect () {
      this.$emit('openMedia', false)
      Array.from(document.querySelectorAll('.active__img')).forEach((element, index) => {
        this.selectedImages.push({'mid': element.getAttribute('data-mid'), 'mfn': element.getAttribute('data-mfn'), 'Lp': element.getAttribute('data-Lp'), 'IsThumb': element.getAttribute('data-IsThumb')})
      })
      if (this.isMultiple) {
        this.$emit('selectedImages', this.selectedImages)
      } else {
        this.$emit('selectedImage', this.selectedImages[0])
      }
    },
    imgpopupMethod (val) {
      this.imgpopup = val
    }
  },

  beforeMount () {
    this.$store.commit('changeActiveNav', 'Media manager')
    let breadcrumbs = [
      {label: 'Media manager', to: '/admiin/media'}
    ]
    this.$store.commit('changeBreadcrumbs', breadcrumbs)
  }
}
</script>

<style type="text/css">
.imgwrapper.active__img{
  border: 3px solid #f00;
  display: table;
}
</style>