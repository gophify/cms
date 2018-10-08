<template>
  <div>
    <v-flex class="pt-3 pb-2" style="display:table;width:100%;">
        <div style="float: left; margin-right: 20px;"><span style="width: 16px; height: 16px; background-color: red; display: block; float: left; margin-top: 0; margin-right: 6px;"></span> reserved & confirmed</div>
        <div style="float: left;"><span style="width: 16px; height: 16px; background-color: #FFCDD2; display: block; float: left; margin-top: 0; margin-right: 6px;"></span> blocked by admin</div>
        <div class="hidden-xs-only" style="float: right; margin-right: 20px;">Nb: click on the room name to block the dates</div>
    </v-flex>
    <div v-if="loadingpage" :style="'background: url('+ $store.state.mediaPath +'/assets/images/main_loading.gif) no-repeat center center white; display: table; height: calc(100vh - 320px); width: 100%;'">
    </div>
    <div class="cal-avail mb-3" v-for="p in properties" v-if="properties.length > 0">
        <v-flex class="border_section">
            <span class="icon_section"><v-icon dark small>check</v-icon></span>
            <span class="label_section">{{ p.Title }}<!-- <small class="hidden-xs-only">({{ p.area }})</small>--></span>
        </v-flex>
        <div class="cal-hscroll">
            <div>
                <div class="cal-el-nowrap" v-for="(r, ri) in p.Rooms">
                    <div style="cursor: pointer;" :class="[{'first' : ri == 0}, 'cal-room']">
                      <h4 class="w100"><nuxt-link :to="'/admiin/roomtype/'+ r.Rt_id +'/room/'+ r.Id +'/edit/'">{{ r.Title }}</nuxt-link></h4>
                    </div>
                    <div class="cal-per-month" v-for="cal in cals">
                        <div class="bold caption pl-2" style="background-color: #fafafa; color:#536E79;" v-if="ri == 0">
                            {{ cal.moName }} {{ cal.year }}
                        </div>
                        <div class="cal-el-nowrap">
                            <span :id="'d_'+cal.year+'_'+cal.mo+'_'+d+'_'+r.Id" :class="['cal-d']" v-for="d in cal.no">{{ d }}</span>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios'
import userlog from '../utils/userlog.js'
export default {
  data () {
    return {
      cals: [],
      properties: [],
      availability: [],
      dateBooked: [],
      loadingpage: true
    }
  },
  updated () {
    if (this.availability !== undefined) {
      this.markCalendar()
      userlog.log(this, {type: 'success_open_availability_on_dashboard'})
    } else {
      userlog.log(this, {type: 'failed_open_availability_on_dashboard'})
      this.$router.push({path: '/admiin'})
    }
  },
  methods: {
    markCalendar () {
      for (var i = 0; i < this.availability.length; i++) {
        if (this.availability[i]['Availability'] !== undefined) {
          for (var k = 0; k < this.availability[i]['Availability'].length; k++) {
            var sdate = this.availability[i]['Availability'][k]['S']
            var edate = this.availability[i]['Availability'][k]['E']

            var partss = sdate.split('-')
            var partse = edate.split('-')

            var dates = new Date(partss[0], parseInt(partss[1]) - 1, parseInt(partss[2]))
            var datee = new Date(partse[0], parseInt(partse[1]) - 1, parseInt(partse[2]))
            var nextDay = dates
            for (var j = 0; j < 1000; j++) {
              var nextDayd = nextDay.getDate()
              var nextDaymi = nextDay.getMonth()
              var nextDayy = nextDay.getFullYear()
              var nextDayid = 'd_' + nextDayy + '_' + nextDaymi + '_' + nextDayd + '_' + this.availability[i]['Room_id']
              if (document.getElementById(nextDayid) != null) {
                document.getElementById(nextDayid).classList.add('blockByAdmin')
              }
              if (nextDay >= datee) {
                break
              }
              nextDay.setDate(nextDay.getDate() + 1)
            }
          }
        }
      }
    }
  },
  mounted () {
    const axithis = this
    axios
      .get(this.$store.state.apiPath + '/dashboard-hotelcalav')
      .then(response => {
        axithis.availability = response.data.Availability
        axithis.properties = response.data.Properties
        var monthNames = ['January', 'February', 'March', 'April', 'May', 'June', 'July', 'August', 'September', 'October', 'November', 'December']
        var dayPerMonth = [31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31]

        var dateNow = new Date()
        var y = dateNow.getFullYear()
        var m = dateNow.getMonth()
        var moi, year
        for (var i = m; i <= (11 + m); i++) {
          if (i > 11) {
            moi = i - 12
            year = y + 1
          } else {
            moi = i
            year = y
          }

          if (moi === 1) {
            if ((year % 100 !== 0) & (year % 4 === 0) || (year % 400 === 0)) {
              dayPerMonth[moi] = 29
            } else {
              dayPerMonth[moi] = 28
            }
          }
          axithis.cals.push({'moName': monthNames[moi], 'no': dayPerMonth[moi], 'year': year, 'mo': moi})
          axithis.loadingpage = false
        }
      })
      .catch(function (err) {
        console.log(err)
      })
  }
}
</script>