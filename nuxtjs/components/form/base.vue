<template>
  <div><div id="top"></div>
    <v-flex class="header light-green darken-1 white--text pt-4 mb-3" style="position: relative;">
        <h3 v-if="formType == 'regular'" :class="['mb-4 pt-3 ml-4', {'mb-5': !formProps.tabs}]">
            <v-icon dark large class="mr-2">edit</v-icon>
            {{ formProps.form_name }}
            <small class="nowrap ml-1" style="color: #C5E1A5;" v-if="formVal['title']">({{ (formVal['title'].length > 32) ? (formVal['title'].substring(0, 32) + '...') : formVal['title'] }})</small>
            <small class="nowrap ml-1" style="color: #C5E1A5;" v-if="!formVal['title']">(Add new)</small>
        </h3>

        <h3 v-if="formType != 'regular'" :class="['mb-4 pt-3 ml-4', {'mb-5': !formProps.tabs}]">
            <v-icon dark large class="mr-2">edit</v-icon>
            Custom page
            <small class="nowrap ml-1" style="color: #C5E1A5;">({{ (formProps.form_name != undefined) ? ((formProps.form_name.length > 32) ? (formProps.form_name.substring(0, 32) + '...') : formProps.form_name) : '...' }})</small>
        </h3>

        <div style="float: right; position: absolute; top: 80px; right: 5px; text-align: right;">
            <div v-for="(info, i) in formProps.infos" :key="i">
                {{ info.label }}
                <v-chip label small :class="{'blue-grey lighten-1 white--text': i == 1}">
                    <v-icon left>{{ info.icon }}</v-icon> {{ info.info }}
                </v-chip>
            </div>
        </div>

          <v-tabs class="mainTabs" v-if="mainTabs"
              dark
              next-icon="arrow_forward"
              prev-icon="arrow_back"
              show-arrows
              slider-color="yellow"
              >
              <v-tab
                v-for="(item, i) in formProps.tabs"
                :key="i" v-on:click="activeTab = i"
              >
                <v-icon small class="mr-1">{{ item.icon }}</v-icon> {{ item.label }}
              </v-tab>
          </v-tabs>       
    </v-flex>

    <div v-if="loading" :style="'background: url('+ $store.state.mediaPath +'/assets/images/main_loading.gif) no-repeat center center white; display: table; position: absolute; z-index:99; height: calc(100vh - 210px); width: 100%;'">
    </div>
    <div class="content">
        <v-alert v-model="message"
          transition="scale-transition"
          dismissible
          type="success"
          style="font-size: 14px; font-weight: bold; text-align: center;"
          class="green lighten-5 green--text"
          outline
        >Data changes have been saved successfully
        </v-alert>

        <v-form ref="form" class="mt-2 mb-5" style="display: table; width: 100%;" v-model="valid" lazy-validation> 
          <div v-show="activeTab == ti" v-for="(tab, ti) in formProps.tabs" :key="ti">
              <v-layout row wrap>
                  <v-flex :class="(f.col != undefined) ? ((f.col.length > 1) ? f.col : (f.col).join(' ')) : ''" v-for="(f, i) in formProps.fields" :key="i">
                      <formLabel :label="f.label" v-if="f.tab == ti" />
                      
                      <v-text-field flat :class="f.class" :rules="[v => !!v || 'Item is required']"
                        v-model="formVal[f.model]" :maxlength="f.size"
                        required v-if="f.type=='input' && f.tab == ti"
                      ></v-text-field>

                      <v-textarea class="input_textarea" :rows="f.size" :rules="[v => !!v || 'Item is required']"
                        v-model="formVal[f.model]"
                        hint="Hint text"
                        v-if="f.type=='textarea' && f.tab == ti"
                      ></v-textarea>

                      <formDatepicker :class="f.class" :data="formVal[f.model]" :model="f.model" @updateFormVal="updateFormValMethod" v-if="f.type=='datepicker' && f.tab == ti"></formDatepicker>

                      <formArea :class="f.class" :data="formVal[f.model]" :model="f.model" @updateFormVal="updateFormValMethod" v-if="f.type=='area' && f.tab == ti"></formArea>

                      <formAmenities :data="formVal[f.model]" :model="f.model" @updateFormVal="updateFormValMethod" v-if="f.type=='amenities' && f.tab == ti"></formAmenities>

                      <formMultiFs :fields="f.fields" :data="formVal[f.model]" :model="f.model" @updateFormVal="updateFormValMethod" v-if="f.type=='fieldset' && f.tab == ti"></formMultiFs>

                      <formImages :data="formVal[f.model]" :model="f.model" @updateFormVal="updateFormValMethod" :colWrapper="(typeof f.col==='object') ? f.col : []" v-if="f.type=='images' && f.tab == ti" />

                      <calendarAvailability :data="formVal[f.model]" :model="f.model" @updateFormVal="updateFormValMethod" v-if="f.type=='calendar' && f.tab == ti"></calendarAvailability>

                      <formVehicleBrands :class="f.class" :data="formVal[f.model]" :model="f.model" @updateFormVal="updateFormValMethod" v-if="f.type=='vehicleBrands' && f.tab == ti"></formVehicleBrands>

                      <formVehicleModels :class="f.class" :data="formVal[f.model]" :model="f.model" @updateFormVal="updateFormValMethod" v-if="f.type=='vehicleModels' && f.tab == ti"></formVehicleModels>
                  </v-flex>
              </v-layout>
          </div>
        </v-form>
        
        <p>&nbsp;</p><p>&nbsp;</p><p>&nbsp;</p>
        <p>&nbsp;</p><p>&nbsp;</p><p>&nbsp;</p>    
    </div>

    <div class="floating-btn" v-if="mod == ''">        
      <div v-if="formType != 'independent'">
          <nuxt-link :to="($route.params.child ? '/admiin/'+ $route.params.fid +'/' + $route.params.id + '/' + $route.params.child + '/list/' : '/admiin/'+ $route.params.fid +'/list/')">
            <v-btn color="warning" round large dark><v-icon left>arrow_back</v-icon>Back</v-btn>
          </nuxt-link>            
      </div>
      <div v-if="formType == 'independent'">
          <nuxt-link :to="'/admiin/independents'">
            <v-btn color="warning" round large dark><v-icon left>arrow_back</v-icon>Back</v-btn>
          </nuxt-link>            
      </div>      
      <div class="hidden-xs-only">
          <v-btn type="submit" :loading="saving_loading" round large color="success" v-on:click="onSubmit">
            <v-icon left>save</v-icon>Save & Publish</v-btn>
      </div>
      <div>
          <v-speed-dial direction = "top" transition="slide-y-reverse-transition">
              <v-btn slot="activator" color="indigo" dark fab>
                  <v-icon>more_vert</v-icon>
              </v-btn>
              <v-btn large dark color="red" @click="dialog_delete_confirm = true">
                  <v-icon left>delete</v-icon> move to trash
              </v-btn> 
              <v-btn class="hidden-sm-and-up" type="submit" :loading="saving_loading" large color="success" v-on:click="onSubmit">
                <v-icon left>save</v-icon>Save & Publish
              </v-btn>               
              <!-- <v-btn large dark color="info">
                  <v-icon left>save</v-icon> Save & Unpublish
              </v-btn> -->
          </v-speed-dial>
      </div> 
    </div>

    <v-dialog v-model="dialog_delete_confirm" max-width="290">
      <v-card>
        <v-card-title class="headline">Are you sure want to delete this?</v-card-title>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn color="green darken-1" flat v-on:click="onDelete()">Yes, delete now!</v-btn>
          <v-btn color="green darken-1" flat @click.native="dialog_delete_confirm = false">Cancel</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>    
  </div>
</template>

<script>
  import axios from 'axios'
  import userlog from '../../utils/userlog.js'
  import auth from '../../utils/auth.js'
  export default {
    props: { loading: Boolean, formType: String, formProps: Object, formVal: Object, formId: String, contentId: String, mod: String },
    data: () => ({
      title: '',
      mainTabs: false,
      activeTab: 0,
      saving_loading: false,
      message: false,
      valid: true,
      dialog_delete_confirm: false,
      name: '',
      nameRules: [
        v => !!v || 'Name is required',
        v => (v && v.length <= 10) || 'Name must be less than 10 characters'
      ],
      email: '',
      emailRules: [
        v => !!v || 'E-mail is required',
        v => /.+@.+/.test(v) || 'E-mail must be valid'
      ],
      select: null,
      checkbox: false
    }),
    computed: {
      cparent: function () {
        if (this.$route.params.child) {
          return this.$route.params.id
        }
        return ''
      }
    },

    watch: {
      formProps: function (val, oldVal) {
        if (val !== []) {
          this.mainTabs = true
        }
      }
    },

    methods: {
      updateFormValMethod (val) {
        this.formVal[val.model] = val.data
      },
      onDelete () {
        if (this.$store.state.session.role === 'demo') {
          alert('You are logged in as a guest\nSo sorry you cannot add/edit/delete the data here.\n\nThankyou :)')
          return false
        }
      },
      onSubmit () {
        if (this.$store.state.session.role === 'demo') {
          userlog.log(this, {type: 'submit_denied', table_name: this.formProps.table_name, content_title: this.formVal.title})
          alert('You are logged in as a guest\nSo sorry you cannot add/edit/delete the data here.\n\nThankyou :)')
          return false
        }

        if (this.$refs.form.validate()) {
          var fs = this.formProps.fields
          var fv

          if (this.formType === 'independent') {
            fv = {}
            for (let i = 0; i < fs.length; i++) {
              fv[fs[i].model] = {T: fs[i].type, V: this.formVal[fs[i].model]}
            }
          } else {
            fv = []
            for (let i = 0; i < fs.length; i++) {
              if (fs[i].type === 'images' || fs[i].type === 'fieldset' || fs[i].type === 'calendar' || fs[i].type === 'amenities') {
                fv.push({T: fs[i].type, M: fs[i].model, V: JSON.stringify(this.formVal[fs[i].model])})
              } else {
                fv.push({T: fs[i].type, M: fs[i].model, V: this.formVal[fs[i].model]})
              }
            }
          }

          const axithis = this
          this.saving_loading = true
          axios.post(this.$store.state.apiPath + '/crud/api',
            {
              content: JSON.stringify(fv),
              contentId: this.contentId,
              cparent: this.cparent,
              formId: this.formId,
              formType: this.formType,
              key: 'save'
            }
          ).then(function (response) {
            if (response.data.Status === 'success') {
              var redirect = ''
              if (axithis.contentId === '0') {
                userlog.log(axithis, {type: 'success_add_content', table_name: axithis.formProps.table_name, content_title: axithis.formVal.title})
                if (axithis.$route.params.child) {
                  redirect = '/admiin/' + axithis.$route.params.fid + '/' + axithis.$route.params.id + '/' + axithis.$route.params.child + '/' + response.data.Id + '/edit/?success=true'
                } else {
                  redirect = '/admiin/' + axithis.formId + '/' + response.data.Id + '/edit/?success=true'
                }
                axithis.$router.push({path: redirect})
              } else {
                userlog.log(axithis, {type: 'success_edit_content', table_name: axithis.formProps.table_name, content_title: axithis.formVal.title})
                location.href = '#top'
                axithis.message = true
              }
            } else {
              userlog.log(axithis, {type: 'submit_error', url: window.location.href, error_msg: JSON.stringify(response.data.Err)})
              alert(JSON.stringify(response.data.Err))
            }
            axithis.saving_loading = false
          })
            .catch(function (err) {
              userlog.log(axithis, {type: 'submit_catch_error', url: window.location.href})
              console.log(err)
              alert('error')
              axithis.saving_loading = false
            })
        } else {
          alert('Please fill all the required fields, check them on all tabs if necessary.')
        }
      }
    },

    mounted () {
      auth.auth(this)
  
      if (this.$route.query.success) {
        this.message = true
      }
    },
    beforeMount () {
  
    }
  }
</script>