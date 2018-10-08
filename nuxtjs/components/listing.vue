<template>
    <div>
        <v-flex class="header light-green darken-1 white--text pt-4 mb-2" style="position: relative;">
            <h3 :class="['mb-4 pt-3 ml-4', {'mb-5': !headerProps.tabs}]">
                <v-icon dark large class="mr-2">{{ headerProps.icon }}</v-icon>
                {{ headerProps.title }} 
                <small class="ml-1" style="color: #C5E1A5;" v-if="headerProps.subtitle">({{ headerProps.title }})</small>
            </h3>
            <div style="float: right; position: absolute; top: 80px; right: 5px; text-align: right;">
                <div v-for="(info, i) in headerProps.infos" :key="i">
                    {{ info.label }}
                    <v-chip label small :class="{'blue-grey lighten-1 white--text': i == 1}">
                        <v-icon left>{{ info.icon }}</v-icon> {{ info.info }}
                    </v-chip>
                </div>
            </div>
            <v-tabs class="mainTabs"
                dark
                next-icon="arrow_forward"
                prev-icon="arrow_back"
                show-arrows
                slider-color="yellow"
                >
                <v-tab
                  v-for="(item, i) in headerProps.tabs"
                  :key="i" v-on:click="activeTab = i"
                >
                  <v-icon small class="mr-1">{{ item.icon }}</v-icon> {{ item.label }}
                </v-tab>
            </v-tabs>       
        </v-flex>  

        <div v-if="loading" :style="'background: url('+ $store.state.mediaPath +'/assets/images/main_loading.gif) no-repeat center center white; display: table; height: calc(100vh - 320px); width: 100%;'">
        </div>
        
        <div v-if="!loading" class="content">
          <div v-show="activeTab == ti" v-for="(tab, ti) in headerProps.tabs" :key="ti">
            <div v-if="ti==0">
              <div class="pt-2">
                <v-layout listing py-2 row th>
                  <v-flex xs1 text-xs-center pl-2>no.</v-flex>

                  <v-flex :class="'pl-1 '+ c.Class" v-for="(c, ci) in column" :key="ci">{{ c.L }}</v-flex>

                  <v-flex xs2 text-xs-center hidden-xs-only>actions</v-flex>
                </v-layout>

                <v-layout listing py-1 row v-for="(data, i) in list" :key="i">
                  <v-flex xs1 text-xs-center pl-2 middle>{{ i+1 }}</v-flex>                  
                  
                  <v-flex :class="'pl-2 middle '+ c.Class" v-for="(c, ci) in column" :key="ci">
                    <!-- <span v-if="c.Type === 'image'">{{ data[c.M] }}</span> -->
                    <span v-if="c.Type === 'image'" :style="'float: left; width: 100%; max-width: 55px; height: 38px; display: table; background-size: cover; background-position: center; background-image: url('+ $store.state.mediaPath + ((data[c.M] != null & data[c.M].length > 0) ? JSON.parse(data[c.M])[0].dir + ((JSON.parse(data[c.M])[0].IsThumb == 'true') ? 'thumbnails-128/' : '') + JSON.parse(data[c.M])[0].mfn : '') +');'" v-on:click="src=$store.state.mediaPath + ((data[c.M] != null) ? JSON.parse(data[c.M])[0].dir + JSON.parse(data[c.M])[0].mfn : ''), imgpopup=true"></span>
                    <!-- <v-icon v-if="c.Type === 'image'" class="hidden-xs-only" small style="float: left;" v-on:click="src=$store.state.mediaPath + ((data[c.M] != null & data[c.M].length > 0) ? JSON.parse(data[c.M])[0].dir + JSON.parse(data[c.M])[0].mfn : ''), imgpopup=true" 
                    color="green">
                      zoom_in</v-icon> -->
                    
                    
                    <nuxt-link v-if="c.M=='title'" :to="($route.params.child ? '/admiin/'+ $route.params.fid +'/' + $route.params.id + '/' + tableName + '/' + data.id + '/edit/' : '/admiin/'+ tableName +'/' + data.id + '/edit/')">
                      {{ (data['title'] != null) ? ((data['title'].length > 48) ? (data['title'].substring(0, 48) + '...') : data['title']) : '' }}
                    </nuxt-link>
                    <span v-if="c.M !='title' & c.Type !='image'">{{ data[c.M] }}</span>
                  </v-flex>

                  <v-flex xs2 text-xs-center mt-2>
                    <nuxt-link class="hidden-sm-and-down mr-3" :to="($route.params.child ? '/admiin/'+ $route.params.fid +'/' + $route.params.id + '/' + tableName + '/' + data.id + '/edit/' : '/admiin/'+ tableName +'/' + data.id + '/edit/')">
                      <v-icon small color="orange">
                          edit
                      </v-icon>
                    </nuxt-link>

                    <nuxt-link class="hidden-sm-and-down" v-if="tbl_child.name" :to="'/admiin/'+ tableName +'/' + data.id + '/' + tbl_child.name + '/list/'">
                      <v-btn round small dark outline color="primary">
                          {{ tbl_child.viewChildLbl }}
                      </v-btn>
                    </nuxt-link>

                    <v-menu :nudge-width="100" class="hidden-md-and-up">
                      <v-toolbar-title slot="activator">
                        <v-icon color="blue" dark class="more_vert">more_vert</v-icon>
                      </v-toolbar-title>

                      <v-list>
                        <v-list-tile>
                          <nuxt-link :to="($route.params.child ? '/admiin/'+ $route.params.fid +'/' + $route.params.id + '/' + tableName + '/' + data.id + '/edit/' : '/admiin/'+ tableName +'/' + data.id + '/edit/')">
                            <v-btn round small dark outline color="primary">
                              <v-icon small color="orange" left>edit</v-icon>edit
                            </v-btn>
                          </nuxt-link>
                        </v-list-tile>
                        <v-list-tile v-if="tbl_child.name">
                          <nuxt-link v-if="tbl_child.name" :to="'/admiin/'+ tableName +'/' + data.id + '/' + tbl_child.name + '/list/'">
                            <v-btn round small dark outline color="primary">
                                {{ tbl_child.viewChildLbl }}
                            </v-btn>
                          </nuxt-link>
                        </v-list-tile>
                      </v-list>
                    </v-menu>

                  </v-flex>                    
                </v-layout>
              </div>
            </div>
          </div>    
        </div>

      <imgpopup :imgpopup="imgpopup" :src="src" @imgpopup="imgpopupMethod"></imgpopup>
      <div v-if="!loading" class="floating-btn">
        <nuxt-link v-if="$route.params.child" :to="(parentTbl.name == '') ? '/admiin/'+ $route.params.fid +'/list/' : '/admiin/'+ parentTbl.name +'/'+ parentTbl.id +'/'+ $route.params.fid +'/list/'">
          <v-btn color="warning" round large dark><v-icon left>arrow_back</v-icon>Back</v-btn>
        </nuxt-link>            
        <nuxt-link :to="addnew_link">
          <v-btn round large dark color="primary">
              <v-icon>add</v-icon><span class="hidden-xs-only"> add new</span>
          </v-btn>
        </nuxt-link>
      </div>               
    </div>
</template>

<script>
import axios from 'axios'
import userlog from '../utils/userlog.js'
import auth from '../utils/auth.js'
export default {
  props: { tableName: String, parent: String },
  data () {
    return {
      column: [{}],
      activeTab: 0,
      loading: true,
      list: [],
      headerProps: {},
      tbl_child: {},
      addnew_link: '',
      cparent: '',
      parentTbl: {},
      src: '',
      imgpopup: false
    }
  },
  methods: {
    imgpopupMethod (val) {
      this.imgpopup = val
    }
  },
  mounted () {
    auth.auth(this)

    if (this.$route.params.child) {
      this.cparent = this.$route.params.id
      this.addnew_link = '/admiin/' + this.$route.params.fid + '/' + this.$route.params.id + '/' + this.$route.params.child + '/add/'
    } else {
      this.addnew_link = '/admiin/' + this.tableName + '/add/'
    }

    const axithis = this
    const formId = this.tableName
    axios.post(this.$store.state.apiPath + '/crud/api',
      {
        cparent_tblname: this.$route.params.fid,
        cparent: this.cparent,
        formId: formId,
        key: 'list'
      }
    ).then(function (response) {
      userlog.log(axithis, {type: 'listing', page: response.data.Title})
      axithis.parentTbl = {id: response.data.Parent_id, name: response.data.Parent_tblname}
      axithis.list = response.data.Datas
      axithis.column = response.data.Cols
      axithis.headerProps.icon = response.data.Icon
      axithis.headerProps.title = response.data.Title
      axithis.$store.commit('changeActiveNav', response.data.Title)
      axithis.tbl_child = {name: response.data.Child, viewChildLbl: response.data.ViewChildLbl}

      var breadcrumbs = [
        {label: axithis.headerProps.title, to: ''},
        {label: 'listing', to: ''}
      ]
      axithis.$store.commit('changeBreadcrumbs', breadcrumbs)

      axithis.loading = false
    })
      .catch(function (err) {
        console.log(err)
        alert('error')
        axithis.loading = false
      })
  },
  beforeMount () {
    this.headerProps = {
      tabs: [
        {icon: 'description', label: 'Listing'}
      ],
      title: 'Listing'
    }
  }
}
</script>