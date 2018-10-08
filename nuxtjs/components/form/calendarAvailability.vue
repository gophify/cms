<template>
    <v-container fluid style="max-width: 1200px">
        <div id="calendar"></div>
        <v-btn id="counterChange" flat v-on:click="emitData"></v-btn>
    </v-container>
</template>

<script>
export default {
  props: { model: String, data: Array },
  data () {
    return {
      init: true
    }
  },
  watch: {
    data: function (val, oldVal) {
      if (this.init) {
        this.init = false
        this.initDataMethod()
      }
    }
  },
  methods: {
    emitData () {
      this.$emit('updateFormVal', {model: this.model, data: this.collectDateMethod()})
    },
    initDataMethod () {
      var dateBooked = this.data
      for (var i = 0; i < dateBooked.length; i++) {
        let sdate = dateBooked[i]['s']
        let edate = dateBooked[i]['e']

        if (sdate !== undefined & edate !== undefined) {
          var partsS = sdate.split('-')
          var partsE = edate.split('-')

          let dateS = new Date(partsS[0], parseInt(partsS[1]) - 1, parseInt(partsS[2]))
          let dateE = new Date(partsE[0], parseInt(partsE[1]) - 1, parseInt(partsE[2]))

          let nextDay = dateS
          for (var j = 0; j < 1000; j++) {
            let nextDayD = nextDay.getDate()
            let nextDayMI = nextDay.getMonth()
            let nextDayY = nextDay.getFullYear()
            let nextDayId = 'd_' + nextDayY + '_' + nextDayMI + '_' + nextDayD
            //   console.log(nextDayId)

            let dEl = document.getElementById(nextDayId)
            //   alert(dEl)
            if (typeof (dEl) !== 'undefined' && dEl != null) {
              dEl.classList.add('selected')
            }
            if (nextDay >= dateE) {
              break
            }
            nextDay.setDate(nextDay.getDate() + 1)
          }
        }
      }
    },
    collectDateMethod () {
      var val = ''
      var prevVal = ''
      var dateSelected = []
      var dayclass = document.getElementsByClassName('day')
      var next = true
      var n = 0
      for (let i = 0; i < dayclass.length; i++) {
        val = dayclass[i].getAttribute('data-date')
        if (dayclass[i].classList.contains('selected')) {
          if (next) {
            dateSelected[n] = {}
            dateSelected[n]['s'] = val
            next = false
          }
        } else {
          if (!next) {
            dateSelected[n]['e'] = prevVal
            n++
            next = true
          }
        }
        prevVal = dayclass[i].getAttribute('data-date')
      }
      //   console.log(dateSelected)
      return dateSelected
    },
    displayCalendar (y, m) {
      var htmlContent = ''
      var FebNumberOfDays = ''
      var counter = 1

      var dateNow = new Date()
      var dayNow = dateNow.getDate()
      var monthNow = dateNow.getMonth()

      var date = new Date(y, m)
      var month = date.getMonth()

      var nextMonth = month + 1 // +1; //Used to match up the current month with the correct start date.
      //   var prevMonth = month - 1
      var year = date.getFullYear()

      // Determing if February (28,or 29)
      if (month === 1) {
        if (((year % 100 !== 0) && (year % 4 === 0)) || (year % 400 === 0)) {
          FebNumberOfDays = 29
        } else {
          FebNumberOfDays = 28
        }
      }

      // names of months and week days.
      var monthNames = ['January', 'February', 'March', 'April', 'May', 'June', 'July', 'August', 'September', 'October', 'November', 'December']
      //   var dayNames = ['Sunday', 'Monday', 'Tuesday', 'Wednesday', 'Thrusday', 'Friday', 'Saturday']
      var dayPerMonth = ['31', '' + FebNumberOfDays + '', '31', '30', '31', '30', '31', '31', '30', '31', '30', '31']

      // days in previous month and next one , and day of week.
      var nextDate = new Date(nextMonth + ' 1 ,' + year)
      var weekdays = nextDate.getDay()
      var weekdays2 = weekdays
      var numOfDays = dayPerMonth[month]

      // this leave a white space for days of pervious month.
      while (weekdays > 0) {
        htmlContent += "<td class='monthPre'></td>"

        // used in next loop.
        weekdays--
      }

      // loop to build the calander body.
      while (counter <= numOfDays) {
        // When to start new line.
        if (weekdays2 > 6) {
          weekdays2 = 0
          htmlContent += '</tr><tr>'
        }

        // if counter is current day.
        // highlight current day using the CSS defined in header.

        var m2d = 0
        var d2d = 0
        if (counter <= 9) {
          d2d = '0' + counter
        } else {
          d2d = counter
        }

        if ((m + 1) <= 9) {
          m2d = '0' + (m + 1)
        } else {
          m2d = (m + 1)
        }

        if (counter === dayNow && m === monthNow) {
          htmlContent += "<td class='dayNow day' id='d_" + y + '_' + m + '_' + counter + "' data-date='" + (y + '-' + m2d + '-' + d2d) + "'>" + counter + '</td>'
        } else {
          htmlContent += "<td class='monthNowBAK day' id='d_" + y + '_' + m + '_' + counter + "' data-date='" + (y + '-' + m2d + '-' + d2d) + "'>" + counter + '</td>'
        }

        weekdays2++
        counter++
      }

      // building the calendar html body.
      var calendarBody = "<div class='flex xs12 sm4 md3 px-2 py-2'><div class='flex py-3 calendar'><table width='100%'> <tr class='monthNow'><th colspan='7'>" + monthNames[month] + ' ' + year + '</th></tr>'
      calendarBody += "<tr class='dayNames'>  <td>Sun</td>  <td>Mon</td> <td>Tues</td>" +
                '<td>Wed</td> <td>Thurs</td> <td>Fri</td> <td>Sat</td> </tr>'
      calendarBody += '<tr>'
      calendarBody += htmlContent
      calendarBody += "</tr><tr><td colspan='7'><label>select all</label><input type='checkbox' class='selectalldaysthismonth' /></td></tr></table></div></div>"
      // set the content of div .

      // console.log(calendarBody)
      return calendarBody
    }
  },
  mounted () {
    var calendars = ''
    var dateNow = new Date()
    var y = dateNow.getFullYear()
    var m = dateNow.getMonth()

    for (let i = m; i <= (11 + m); i++) {
      if (i > 11) {
        // console.log(y+1 + ', ' + (i-12));
        calendars += this.displayCalendar((y + 1), (i - 12))
      } else {
        // console.log(y + ', ' + i);
        calendars += this.displayCalendar(y, i)
      }
    }

    document.getElementById('calendar').innerHTML = '<div class="layout row wrap">' + calendars + '</div>'

    /// ///////////////////////////////////////////////
    var classname = document.getElementsByClassName('selectalldaysthismonth')
    var myFunction = function () {
      var es = this.closest('tbody').getElementsByClassName('day')
      for (var i = 0; i < es.length; i++) {
        if (this.checked) {
          es[i].classList.add('selected')
        } else {
          es[i].classList.remove('selected')
        }
      }
      document.getElementById('counterChange').click()
    }
    for (var i = 0; i < classname.length; i++) {
      classname[i].addEventListener('click', myFunction, false)
    }

    /// ///////////////////////////////////////////////
    var selectday = function () {
      if (this.classList.contains('selected')) {
        this.classList.remove('selected')
      } else {
        this.classList.add('selected')
      }
      document.getElementById('counterChange').click()
    }
    var dayclass = document.getElementsByClassName('day')
    for (i = 0; i < dayclass.length; i++) {
      dayclass[i].addEventListener('click', selectday, false)
    }
  }
}
</script>


<style>
#calendar .calendar{
  border: 1px solid #999;
}
#calendar .calendar td,
#calendar .calendar th{
  text-align: right;
  width: 9.286%;
  padding: 2%;
  height: 25px;
  font-size: 11px;
}
#calendar .calendar td.day{
  cursor: pointer;
  font-size: 10px;
}
#calendar .calendar tr.monthNow th{
  padding-bottom: 8px;
  letter-spacing: 1px;
}
#calendar .calendar tr.dayNames{
  background-color: #eee;
  font-weight: bold;
}
#calendar .calendar td.selected{
    background-color: #ffcdd2;
}
#calendar input[type=checkbox]
{
  -ms-transform: scale(2); /* IE */
  -moz-transform: scale(2); /* FF */
  -webkit-transform: scale(2); /* Safari and Chrome */
  -o-transform: scale(2); /* Opera */
  transform: scale(1.5);
  padding: 10px;
  margin-left: 15px;
}
</style>
