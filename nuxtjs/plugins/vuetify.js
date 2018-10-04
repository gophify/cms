import Vue from 'vue'
import {
  Vuetify,
  VApp,
  VTabs,
  VBadge,
  VAlert,
  VMenu,
  VList,
  VDialog,
  VDatePicker,
  VAutocomplete,
  VAvatar,
  VProgressLinear,
  VForm,
  VParallax,
  VTextField,
  VCombobox,
  VSelect,
  VNavigationDrawer,
  VChip,
  VCard,
  VBtn,
  VIcon,
  VGrid,
  VSpeedDial,
  VToolbar,
  VTextarea
} from 'vuetify'
import login from '@/components/login'
import hotelCalendar from '@/components/hotelCalendar'
import media from '@/components/media'
import imgpopup from '@/components/imgpopup'
import listing from '@/components/listing'
import edit from '@/components/edit'
import add from '@/components/add'

import formGenerator from '@/components/formGenerator/main'
import formGeneratorFs from '@/components/formGenerator/fieldset'
import formBase from '@/components/form/base'

import formVehicleBrands from '@/components/form/vehicle/brands'
import formVehicleModels from '@/components/form/vehicle/models'
import formAmenities from '@/components/form/amenities'
import formArea from '@/components/form/area'

import formMediaModal from '@/components/form/mediaModal'
import calendarAvailability from '@/components/form/calendarAvailability'
import formMultiFs from '@/components/form/multiFs'
import formDatepicker from '@/components/form/datepicker'
import formImages from '@/components/form/images'
import formImage from '@/components/form/image'
import formLabel from '@/components/form/label'

Vue.use(Vuetify, {
  components: {
    VApp,
    VTabs,
    VBadge,
    VAlert,
    VMenu,
    VList,
    VDialog,
    VDatePicker,
    VAutocomplete,
    VAvatar,
    VProgressLinear,
    VForm,
    VParallax,
    VTextField,
    VCombobox,
    VSelect,
    VNavigationDrawer,
    VChip,
    VCard,
    VBtn,
    VIcon,
    VGrid,
    VSpeedDial,
    VToolbar,
    VTextarea
  }
})
Vue.component('login', login)
Vue.component('hotelCalendar', hotelCalendar)
Vue.component('media', media)
Vue.component('imgpopup', imgpopup)
Vue.component('listing', listing)
Vue.component('edit', edit)
Vue.component('add', add)

Vue.component('formGenerator', formGenerator)
Vue.component('formGeneratorFs', formGeneratorFs)

Vue.component('formMediaModal', formMediaModal)
Vue.component('formBase', formBase)

Vue.component('formVehicleBrands', formVehicleBrands)
Vue.component('formVehicleModels', formVehicleModels)
Vue.component('formAmenities', formAmenities)
Vue.component('formArea', formArea)

Vue.component('calendarAvailability', calendarAvailability)
Vue.component('formMultiFs', formMultiFs)
Vue.component('formDatepicker', formDatepicker)
Vue.component('formImages', formImages)
Vue.component('formImage', formImage)
Vue.component('formLabel', formLabel)
