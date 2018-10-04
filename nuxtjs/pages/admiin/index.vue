<template>
    <div>
        <v-flex class="header light-green darken-1 white--text pt-4 mb-2" style="position: relative;">
            <h3 :class="['mb-4 pt-3 ml-4', {'mb-5': !headerProps.tabs}]">
                <v-icon dark large class="mr-2">dashboard</v-icon>
                {{ headerProps.title }} 
                <small class="nowrap ml-1 hidden-sm-and-down" style="color: #C5E1A5;">({{subtitle}})</small>
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
                    :key="i" v-on:click="subtitle=item.label, activeTab = i, updateBreadcrumbs(item.label)"
                >
                    <v-icon small class="mr-1">{{ item.icon }}</v-icon> {{ item.label }}
                </v-tab>
            </v-tabs>       
        </v-flex>
        <div v-show="activeTab == 0">
            <hotelCalendar></hotelCalendar>
        </div>
    </div>
</template>

<script>
// import axios from 'axios'
export default {
  data () {
    return {
      activeTab: 0,
      headerProps: {},
      subtitle: 'Room availability'
    }
  },
  methods: {
    updateBreadcrumbs (title) {
      let breadcrumbs = [
        {label: title, to: '/admiin'}
      ]
      this.$store.commit('changeBreadcrumbs', breadcrumbs)
    }
  },
  beforeMount () {
    this.headerProps = {
      tabs: [
        {icon: 'hotel', label: 'Room availability'},
        {icon: 'help', label: 'User Guide'},
        {icon: 'directions_car', label: 'Driver availability'},
        {icon: 'attach_money', label: 'Latest payments'}
        // {icon: 'hotel', label: 'Latest hotel reservations'},
        // {icon: 'directions_car', label: 'Latest car rental'},
        // {icon: 'add_shopping_cart', label: 'Latest transaction'},
        // {icon: 'multiline_chart', label: 'Log latest activities'}
      ],
      title: 'Dashboard'
    }
    let breadcrumbs = [
      {label: 'Room Availability', to: '/admiin'}
    ]
    this.$store.commit('changeBreadcrumbs', breadcrumbs)
  }
}
</script>

<style>

</style>