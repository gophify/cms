<template>
    <formBase :loading="loading" :formType="formType" :formProps="formProps" :formVal="formVal" :formId="formId" :contentId="contentId" mod=""></formBase>
</template>

<script>
  import axios from 'axios'
  import userlog from '../utils/userlog.js'
  export default {
    data: () => ({
      formType: 'regular',
      formProps: {fields: [], tabs: []},
      formVal: {},
      loading: true,
      formId: '',
      contentId: '0',
      breadcrumbs: []
    }),
    beforeMount () {
      // INIT GENERATE AND FILL FROM
      if (this.$route.params.child) {
        this.formId = this.$route.params.child
        this.contentId = this.$route.params.ccid
      } else {
        this.formId = this.$route.params.fid
        this.contentId = this.$route.params.id
      }
      const axithis = this

      if (this.formId !== '') {
        axios.post(this.$store.state.apiPath + '/form-generator/api',
          {
            key: 'getformjson',
            formId: this.formId,
            formType: 'regular'
          }
        ).then(function (response) {
          // axithis.loading = true
          axithis.formProps = JSON.parse(response.data.Data.Json)
          axithis.breadcrumbs.push({label: response.data.Data.Title, to: '/form/' + axithis.formId + '/list/'})

          if (axithis.contentId > 0) {
            axios.post(axithis.$store.state.apiPath + '/crud/api',
              {
                contentId: axithis.contentId,
                formId: axithis.formId,
                key: 'read'
              }
            ).then(function (response) {
              // console.log(response.data)
              var fs = axithis.formProps.fields
              var formVal = {}
              for (let i = 0; i < fs.length; i++) {
                if (fs[i].type === 'input' || fs[i].type === 'textarea' || fs[i].type === 'datepicker' || fs[i].type === 'area' || fs[i].type === 'vehicleBrands' || fs[i].type === 'vehicleModels') {
                  formVal[fs[i].model] = response.data[fs[i].model]
                } else if (fs[i].type === 'images' || fs[i].type === 'fieldset' || fs[i].type === 'calendar') {
                  if (!response.data[fs[i].model]) {
                    formVal[fs[i].model] = []
                  } else {
                    if ((response.data[fs[i].model]).toString() !== '') {
                      formVal[fs[i].model] = JSON.parse(response.data[fs[i].model])
                    }
                  }
                } else if (fs[i].type === 'amenities') {
                  formVal[fs[i].model] = response.data[fs[i].model]
                }
              }
  
              axithis.breadcrumbs.push({label: (formVal['title'].length > 32) ? (formVal['title'].substring(0, 32) + '...') : formVal['title'], to: ''})
              axithis.breadcrumbs.push({label: 'edit form', to: ''})
              axithis.$store.commit('changeBreadcrumbs', axithis.breadcrumbs)
              axithis.$store.commit('changeActiveNav', axithis.formProps.form_name)
  
              axithis.formVal = formVal
              axithis.loading = false
              if (!axithis.$route.query.success) {
                userlog.log(axithis, {type: 'open_edit_form', table_name: axithis.formProps.table_name, content_title: axithis.formVal.title})
              }
            })
              .catch(function (err) {
                console.log(err)
                alert('error')
                axithis.loading = false
              })
          }
        })
          .catch(function (err) {
            console.log(err)
            alert('error')
            axithis.loading = false
          })
      }
    }
  }
</script>