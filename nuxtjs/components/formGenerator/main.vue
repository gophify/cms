<template>
  <div>
    <v-dialog persistent scrollable
      v-model="dialog"
      width="1024"
    >
      <v-card>
        <v-card-title
          class="headline grey lighten-2"
          primary-title
        >
          Form Generator

          <v-spacer></v-spacer>
          
          <v-btn :loading="saving2json"
            color="primary"
            flat v-on:click="migrationMethod"
          >
            Database Migration
          </v-btn>
        </v-card-title>
        
        <v-card-text v-show="migration">
          
          <v-layout row>
            <v-flex xs6>
                <pre>{{ migrationInfo }}</pre>
            </v-flex>
            <v-flex xs6 arial>
                {{ sqlCreateTable }}
                <v-btn :loading="createTableLoading"
                  color="primary"
                  flat v-on:click="createTable"
                >
                  create now
                </v-btn>
                <v-alert v-if="createTableResult.Status"
                  :value="true"
                  :color="createTableResult.Status" class="mr-5"
                  outline
                >
                  <pre v-if="createTableResult.Status == 'error'">{{ createTableResult.Error }}</pre>
                  <pre v-if="createTableResult.Status == 'success'">{{ createTableResult.Msg }}</pre>
                </v-alert>              
            </v-flex>            
          </v-layout>


          <v-btn :loading="saving2json" class="mt-5"
            color="primary"
            flat v-on:click="migration=false"
          >
            back
          </v-btn>          
        </v-card-text>

        <v-card-text class="formGenerator" v-show="!migration">           
            <v-tabs
                dark
                color="blue"
                show-arrows
            >
                <v-tabs-slider color="yellow"></v-tabs-slider>

                <v-tab
                v-for="(section, si) in sections"
                :href="'#tab-' + si"
                :key="si"
                >
                <b>{{ section }}</b>
                </v-tab>

                <v-tabs-items class="mt-3">
                    <v-tab-item id="tab-1">
                        <div class="px-1">
                            <v-layout row v-if="formProps['fields'].length>0">
                                <v-flex xs2 v-if="formProps['tabs'][0] !== {}">tab:</v-flex>
                                <v-flex xs1>type:</v-flex>
                                <v-flex xs2>model:</v-flex>
                                <v-flex xs2>label:</v-flex>
                                <v-flex xs2>col:</v-flex>
                                <v-flex xs1>class:</v-flex>
                                <v-flex xs1>mxlgth/rows:</v-flex>
                                <v-flex xs1 pl-2>actions</v-flex>
                            </v-layout>
                            <v-layout my-1 row wrap v-for="(field, fi) in formProps['fields']" :key="fi">
                                <v-flex xs2 v-if="formProps['tabs'][0] !== {}">
                                    <select v-model="formProps['fields'][fi]['tab']">
                                        <option :value="ti" v-for="(tab, ti) in formProps['tabs']" :key="ti">{{ tab.label }}</option>
                                    </select>
                                </v-flex>
                                <v-flex xs1>
                                    <select v-model="formProps['fields'][fi]['type']">
                                        <option value="input">input</option>
                                        <!-- <option value="image">image</option> -->
                                        <option value="images">images</option>
                                        <option value="textarea">textarea</option>
                                        <option value="fieldset">fieldset</option>
                                        <option value="calendar">Calendar</option>
                                        <option value="datepicker">Datepicker</option>
                                        <option value="area">Area</option>
                                        <option value="amenities">Amenities</option>
                                        <option value="vehicleBrands">Vehicle Brands</option>
                                        <option value="vehicleModels">Vehicle Models</option>
                                    </select>
                                </v-flex>                            
                                <v-flex xs2>
                                    <input type="text" v-model="formProps['fields'][fi]['model']" />
                                </v-flex>
                                <v-flex xs2>
                                    <input type="text" v-model="formProps['fields'][fi]['label']" />
                                </v-flex>
                                <v-flex xs2>
                                    <select v-model="formProps['fields'][fi]['col']" multiple style="height: 36px;">
                                        <option :value="'xs'+i" v-for="i in 12" :key="'xs'+i">{{ 'xs'+i }}</option>
                                        <option :value="'sm'+i" v-for="i in 12" :key="'sm'+i">{{ 'sm'+i }}</option>
                                        <option :value="'md'+i" v-for="i in 12" :key="'md'+i">{{ 'md'+i }}</option>
                                    </select>
                                </v-flex>             
                                <v-flex xs1>
                                    <select v-model="formProps['fields'][fi]['class']">
                                        <option value="input_lg">input_lg</option>
                                        <option value="input_md">input_md</option>
                                    </select>
                                </v-flex>                                
                                <v-flex xs1>
                                    <input type="number" v-model="formProps['fields'][fi]['size']" v-if="formProps['fields'][fi]['type'] == 'input' || formProps['fields'][fi]['type'] == 'vehicleBrands' || formProps['fields'][fi]['type'] == 'vehicleModels' || formProps['fields'][fi]['type'] == 'textarea'" />
                                </v-flex>
                                <v-flex xs1 pl-2>
                                    <v-icon small color="blue" v-if="fi > 0" v-on:click="arrow_upward_clicked(fi)">arrow_upward</v-icon> 
                                    <v-icon small color="blue" v-if="fi < (formProps['fields'].length-1)" v-on:click="arrow_downward_clicked(fi)">arrow_downward</v-icon> 
                                    <v-icon small color="red" v-on:click="delete_clicked(fi)">delete</v-icon>
                                </v-flex>

                                <formGeneratorFs @updataData="updataDataMethod" :dataInit="formProps['fields'][fi]['fields']" :index="fi" v-if="formProps['fields'][fi]['type'] == 'fieldset'" />
                            </v-layout>
                        </div>
                        
                        <v-btn small color="primary" dark @click="addField" style="float: right;">
                            <v-icon left small>add</v-icon> add field
                        </v-btn>                        
                    </v-tab-item>                   
                    <v-tab-item id="tab-0">
                        <v-layout row wrap my-2>
                            <v-flex xs2 mt-1 text-xs-right>Form title: &nbsp; </v-flex>
                            <v-flex xs5><input type="text" v-model="formProps['form_name']" /></v-flex>

                            <v-flex xs1 mt-1 text-xs-right>Icon: &nbsp; </v-flex>
                            <v-flex xs4>
                              <input type="text" v-model="formProps['form_icon']" />
                            </v-flex>                            
                        </v-layout>
                        <v-layout row wrap my-2>
                            <v-flex xs2 mt-1 text-xs-right>Form type: &nbsp; </v-flex>
                            <v-flex :xs2="formProps['form_type'] == 'r'" :xs10="formProps['form_type'] != 'r'">
                              <select v-model="formProps['form_type']">
                                  <option value="r">Regular</option>
                                  <option value="i">Independent, for custom/irregular page purposes</option>
                              </select>
                            </v-flex>

                            <v-flex xs1 mt-1 text-xs-right v-if="formProps['form_type'] == 'r'">Name: &nbsp; </v-flex>
                            <v-flex xs3 v-if="formProps['form_type'] == 'r'">
                              <input type="text" v-model="formProps['table_name']" />
                            </v-flex>

                            <v-flex xs1 mt-1 text-xs-right v-if="formProps['form_type'] == 'r'">Parent: &nbsp; </v-flex>
                            <v-flex xs3 v-if="formProps['form_type'] == 'r'">
                              <select v-model="formProps['table_parent']">
                                  <option :value="tp.Name" v-for="(tp, tpi) in table_parents" :key="tpi">{{ tp.Title }}</option>
                              </select>                              
                            </v-flex>                            
                        </v-layout>
                        <v-layout row wrap my-2 v-if="formProps['table_child'] != ''">
                            <v-flex xs2 mt-1 text-xs-right>Table Child: &nbsp; </v-flex>
                            <v-flex xs4><input readonly type="text" v-model="formProps['table_child']" /></v-flex>

                            <v-flex xs2 mt-1 text-xs-right>View child label: &nbsp; </v-flex>
                            <v-flex xs4>
                              <input type="text" v-model="formProps['viewChildLbl']" />
                            </v-flex>                            
                        </v-layout>                        
                        <v-layout row wrap my-2>
                            <v-flex xs2 mt-1 text-xs-right>Tabs: &nbsp; </v-flex>
                            <v-flex xs10>
                                <v-layout row v-for="(tab, ti) in this.formProps['tabs']" :key="ti" class="mb-1">
                                    <v-flex xs5><input type="text" v-model="formProps['tabs'][ti]['icon']" placeholder="icon" /></v-flex>
                                    <v-flex xs6 pl-1><input type="text" v-model="formProps['tabs'][ti]['label']" placeholder="label" /></v-flex>
                                    <v-flex xs1 pt-2><v-icon color="red" v-if="ti>0" v-on:click="removeTab(ti)">close</v-icon></v-flex>
                                </v-layout>
                                
                                <v-btn small color="primary" dark @click="addTab">
                                    <v-icon left small>add</v-icon> add tab
                                </v-btn>
                            </v-flex>
                        </v-layout>
                    </v-tab-item>
                </v-tabs-items>
            </v-tabs>
        </v-card-text>

        <hr />
        <v-card-actions>
          <v-spacer></v-spacer>
          
          <v-btn
            color="primary"
            flat v-on:click="previewFormMethod"
          >
            Preview
          </v-btn>

          <v-btn
            color="primary"
            flat v-on:click="saveJson"
          >
            Generate & Save
          </v-btn>

          <v-btn
            color="primary"
            flat
            @click="close"
          >Close</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </div>
</template>

<script>
import axios from 'axios'
export default {
  props: { formId: String, formProps: Object },
  data () {
    return {
      loading: false,
      migration: false,
      migrationInfo: '',
      createTableLoading: false,
      sqlCreateTable: '',
      createTableResult: {},
      saving2json: false,
      dialog: true,
      sections: ['common settings', 'fields', 'button actions'],
      table_parents: []
    }
  },
  mounted () {
    this.loading = true
    const axithis = this
    axios
      .post(this.$store.state.apiPath + '/form-generator/api', {key: 'listallforms', formType: 'r'})
      .then(response => {
        // console.log(response.data)
        axithis.loading = false
        axithis.table_parents = response.data
      })
      .catch(function (err) {
        console.log(err)
        alert('error')
        axithis.loading = false
      })
  },
  methods: {
    arrow_upward_clicked (i) {
      var len = this.formProps['fields'].length
      var newData = []
      for (let j = 0; j < len; j++) {
        if (j === i) {
          newData.push(this.formProps['fields'][i - 1])
        } else if (j === (i - 1)) {
          newData.push(this.formProps['fields'][i])
        } else {
          newData.push(this.formProps['fields'][j])
        }
      }
      this.formProps['fields'] = newData
    },
    arrow_downward_clicked (i) {
      var len = this.formProps['fields'].length
      var newData = []
      for (let j = 0; j < len; j++) {
        if (j === i) {
          newData.push(this.formProps['fields'][i + 1])
        } else if (j === (i + 1)) {
          newData.push(this.formProps['fields'][i])
        } else {
          newData.push(this.formProps['fields'][j])
        }
      }
      this.formProps['fields'] = newData
    },
    delete_clicked (i) {
      var len = this.formProps['fields'].length
      var newData = []
      for (let j = 0; j < len; j++) {
        if (i !== j) {
          newData.push(this.formProps['fields'][j])
        }
      }
      this.formProps['fields'] = newData
    },
    createTable () {
      if (this.$store.state.session.role !== 'root') {
        alert('ACCESS DENIED')
        return false
      }
      const axithis = this
      this.createTableLoading = true
      axios.post(this.$store.state.apiPath + '/form-generator/api',
        {
          id: this.formId,
          type: this.formProps.form_type,
          key: 'createtable',
          sql: this.sqlCreateTable
        }
      ).then(function (response) {
        console.log(response.data)
        axithis.createTableResult = response.data
        axithis.createTableLoading = false
      })
        .catch(function (err) {
          axithis.createTableLoading = false
          alert('error')
          console.log(err)
        })
    },
    migrationMethod () {
      this.formId = this.$route.params.id
      const axithis = this
      this.migrationInfoLoading = true
      axios.post(this.$store.state.apiPath + '/form-generator/api',
        {
          key: 'getformjson',
          formId: this.formId,
          formType: this.formProps.formType
        }
      ).then(function (response) {
        // console.log(response.data)
        var fields = JSON.parse(response.data.Data.Json).fields
        axithis.migrationInfo = fields

        var sqlCreateTable = 'CREATE TABLE ' + response.data.Data.Name + ' (id INT AUTO_INCREMENT, parent Char(8), '
        for (let i = 0; i < fields.length; i++) {
          if (fields[i].model) {
            if (fields[i].type === 'input' || fields[i].type === 'vehicleBrands' || fields[i].type === 'vehicleModels') {
              sqlCreateTable += fields[i].model + ' Varchar(' + fields[i].size + '), '
            } else if (fields[i].type === 'textarea') {
              sqlCreateTable += fields[i].model + ' Text, '
            } else {
              sqlCreateTable += fields[i].model + ' Text, ' // Json
            }
          }
        }
        sqlCreateTable += 'PRIMARY KEY (id));'

        axithis.sqlCreateTable = sqlCreateTable

        axithis.migration = true
        axithis.migrationInfoLoading = false
      })
        .catch(function (err) {
          axithis.migrationInfoLoading = false
          alert('error migrationMethod')
          console.log(err)
        })
    },
    close () {
      this.$emit('openGenerator', false)
    },
    previewFormMethod () {
      this.close()
      this.$emit('formProps', this.formProps)
    },
    addField () {
      this.formProps['fields'].push({})
    },
    addTab () {
      this.formProps['tabs'].push({})
    },
    removeTab (i) {
      this.formProps['tabs'].splice(i, 1)
    },
    updataDataMethod (val) {
      this.formProps['fields'][val.index]['fields'] = val.data
    },
    saveJson () {
      if (this.$store.state.session.role !== 'root') {
        alert('ACCESS DENIED')
        return false
      }
      const axithis = this
      this.saving2json = true
      axios.post(this.$store.state.apiPath + '/form-generator/api',
        {
          key: 'savejson',
          formId: this.formId,
          icon: this.formProps.form_icon,
          name: this.formProps.table_name,
          title: this.formProps.form_name,
          type: this.formProps.form_type,

          parent: this.formProps.table_parent,
          child: this.formProps.table_child,
          viewChildLbl: this.formProps.viewChildLbl,

          json: JSON.stringify(this.formProps)
        }
      ).then(function (response) {
        // console.log(response.data)
        if (response.data.Status === 'success') {
          axithis.$emit('generateForm', axithis.formProps)
          axithis.close()
          var redirect = '/admiin/generator/' + response.data.Id + '/'
          axithis.$router.push({path: redirect})
        } else {
          alert(JSON.stringify(response.data.Err))
        }
        axithis.saving2json = false
      })
        .catch(function (err) {
          console.log(err)
          alert('error')
        })
    }
  }
}
</script>