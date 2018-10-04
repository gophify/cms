<template>
    <div>
        <v-flex class="header light-green darken-1 white--text pt-4 mb-4" style="position: relative;">
            <h3 :class="['display-1 mb-4 pt-3 ml-4', {'mb-5': !headerProps.tabs}]">
                <v-icon dark large class="mr-2">settings</v-icon>
                {{ headerProps.title }} 
                <small class="ml-1" style="color: #C5E1A5;" v-if="headerProps.subtitle">({{ headerProps.subtitle }})</small>
            </h3>
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

        <div class="content">
          <div v-show="activeTab == ti" v-for="(tab, ti) in headerProps.tabs" :key="ti">
            <div v-if="ti==0">
              <div class="px-5">
                <v-layout listing py-2 row th>
                  <v-flex xs1 text-xs-center>no.</v-flex>
                  <v-flex xs7>Title</v-flex>
                  <!-- <v-flex xs1 text-xs-center>actions</v-flex> -->
                </v-layout>

                <v-layout listing py-2 row v-for="(data, i) in list" :key="i">
                  <v-flex xs1 text-xs-center>{{ i+1 }}</v-flex>                  
                  
                  <v-flex xs7>
                    <nuxt-link v-if="data.Type == 'i'" :to="'/admiin/generator/' + data.Id + '/'">{{ data['Title'] }}</nuxt-link>
                    <nuxt-link v-if="data.Type != 'i'" :to="'/admiin/generator/' + data.Name + '/'">{{ data['Title'] }}</nuxt-link>
                  </v-flex>

                  <!-- <v-flex xs1 text-xs-center>
                    <nuxt-link :to="'/admiin/generator/' + data.Id + '/edit/'"><v-icon small color="orange">edit</v-icon></nuxt-link>
                    &nbsp; &nbsp; <v-icon small color="blue">arrow_forward</v-icon></v-flex> -->
                </v-layout>
              </div>
            </div>
          </div>    
        </div>

      <div class="floating-btn">
        <nuxt-link :to="'/admiin/generator/addnew/'">
          <v-btn round large dark color="primary">
              <v-icon left>add</v-icon> add new
          </v-btn>
        </nuxt-link>
      </div>               
    </div>
</template>

<script>
import axios from 'axios'
// import userlog from '../../utils/userlog.js'
import auth from '../../utils/auth.js'
export default {
  data () {
    return {
      column: [{}],
      activeTab: 0,
      loading: false,
      list: [],
      headerProps: {}
    }
  },
  mounted () {
    auth.auth(this)

    const axithis = this
    this.loading = true
    axios
      .post(this.$store.state.apiPath + '/form-generator/api', {key: 'listallforms', formType: ''})
      .then(response => {
        // console.log(response.data)
        axithis.list = response.data
        axithis.loading = false
      })
      .catch(function (err) {
        console.log(err)
        axithis.loading = false
      })
  },
  beforeMount () {
    this.headerProps = {
      tabs: [
        {icon: 'settings', label: 'Form Generator'},
        {icon: 'settings', label: 'Settings 1'},
        {icon: 'settings', label: 'Settings 2'}
      ],
      title: 'Settings'
    }
    this.$store.commit('changeActiveNav', 'Settings')

    let breadcrumbs = [
      {label: 'Settings', to: '/admiin/settings'}
    ]
    this.$store.commit('changeBreadcrumbs', breadcrumbs)
  }
}
</script>