<template>
    <v-content>
          <div class="loadingpage" v-if="loadingpage" :style="'background-image: url('+ $store.state.mediaPath +'/assets/images/main_loading.gif);'"></div>
          <v-app>
              <v-navigation-drawer v-model="mainmenu_drawer" dark :class="['blue-grey', 'darken-1', {'drawerDesktop': ww_1280}]" :mini-variant="true" mini-variant-width="170" app>
              
              <a href="/" style="
                  color: #fff;
                  font-weight: bold;
                  letter-spacing: 2px;
                  display: table;
                  width: 100%;
                  margin-top: 30px;
                  font-size: 26px;
                  margin-left: 14px;
              ">Gophify</a>              
              
              <!-- <a href="/"><img :src="$store.state.apiPath + '/assets/images/outpost_logo_header.png'" width="115" class="mt-5 pt-3 ml-3" /></a> -->
              <ul class="subheading menuleft mt-2">
                  <li v-for="(menu, i) in mainMenu" :key="i" v-on:click="$store.commit('changeActiveNav', menu.title)">
                      <nuxt-link :to="menu.to" :class="{ 'active': $store.state.activeNav == menu.title }">
                      <v-icon small>{{ menu.icon }}</v-icon>{{ menu.title }}</nuxt-link>
                  </li>
              </ul>
              </v-navigation-drawer>                  
              
              <v-flex class="hidden-xs-only breadcrumbs light-green darken-3 pl-3 py-2">
                  <span v-if="!mainmenu_drawer" v-on:click="mainmenu_drawer=true" style="cursor: pointer;">
                      <v-icon style="color: white!important; margin-right: 20px;">menu</v-icon> | &nbsp; &nbsp; 
                  </span>        
                  <nuxt-link to="/admiin"><v-icon small dark style="margin-bottom: 2px;">language</v-icon> Dashboard</nuxt-link>
                  <nuxt-link :to="b.to" v-for="(b, i) in $store.state.breadcrumbs" :key="i">
                      <v-icon small dark>keyboard_arrow_right</v-icon>
                      {{ b.label }}</nuxt-link>
                  <div style="float: right;" class="pr-3"><span>Welcome, {{ $store.state.session.firstname }} {{ $store.state.session.lastname }}</span> &nbsp; <a v-on:click="logout()" style="cursor: pointer;"><u>(logout)</u></a></div>
              </v-flex>

              <v-toolbar dense dark color="light-green darken-3" class="hidden-sm-and-up">
                <v-toolbar-side-icon v-if="!mainmenu_drawer" v-on:click="mainmenu_drawer=true" style="cursor: pointer;"></v-toolbar-side-icon>

                <v-toolbar-title>Gophify</v-toolbar-title>

                <v-spacer></v-spacer>

                <v-menu>
                  <v-toolbar-title slot="activator">
                    <v-icon dark>account_circle</v-icon>
                  </v-toolbar-title>
                  <div class="px-2 py-1" style="color: #333; background-color:white">Welcome, {{ $store.state.session.firstname }} {{ $store.state.session.lastname }}<br>
                  <a v-on:click="logout()" style="cursor: pointer;"><u>(logout)</u></a></div>
                </v-menu>
                <!-- <v-btn icon>
                  <v-icon>more_vert</v-icon>
                </v-btn> -->
              </v-toolbar>              
              <nuxt />              
          </v-app>
    </v-content>
</template>

<script>
  import axios from 'axios'
  import auth from '../utils/auth.js'
  export default {
    data () {
      return {
        mainmenu_drawer: true,
        ww_600: false,
        ww_1280: false,
        mainMenu: [
          { icon: 'dashboard', title: 'Dashboard', to: '/admiin' },
          { icon: 'pages', title: 'Custom Pages', to: '/admiin/independents' }
        ],
        loadingpage: true
      }
    },
    methods: {
      logout () {
        this.loadingpage = true
        const axithis = this
        axios
          .get(this.$store.state.apiPath + '/logout-api')
          .then(response => {
            localStorage.removeItem('session')
            axithis.loadingpage = false
            axithis.$router.push({path: '/'})
            return false
          })
          .catch(function (err) {
            axithis.loadingpage = false
            console.log(err)
          })
      }
    },
    mounted () {
      auth.auth(this)
  
      this.loadingpage = false
      const axithis = this
      axios
        .post(this.$store.state.apiPath + '/form-generator/api', {key: 'listallforms', formType: 'r'})
        .then(response => {
          // console.log(response.data)
          for (var i = 0; i < response.data.length; i++) {
            if (response.data[i].Parent === '') {
              axithis.mainMenu.push({icon: response.data[i].Icon, title: response.data[i].Title, to: '/admiin/' + response.data[i].Name + '/list/'})
            }
          }
          axithis.mainMenu.push({ icon: 'image', title: 'Media manager', to: '/admiin/media' })
          axithis.mainMenu.push({ icon: 'settings', title: 'Settings', to: '/admiin/settings' })
        })
        .catch(function (err) {
          console.log(err)
        })
    },
    beforeMount () {
      if (window.innerWidth >= 1280) {
        this.ww_1280 = true
        this.mainmenu_drawer = true
      }
      if (window.innerWidth <= 600) {
        this.ww_600 = true
      }
    }
  }
</script>