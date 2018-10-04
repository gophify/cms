<template>
    <formBase :formType="formType" :formProps="formProps" :formVal="formVal" :formId="formId" mod=""></formBase>
</template>

<script>
  import axios from 'axios'
  import userlog from '../../../../utils/userlog.js'
  export default {
    data: () => ({
      formType: 'independent',
      formProps: {fields: [], tabs: []},
      formVal: {},
      loading: true,
      formId: '0',
      contentId: '0',
      breadcrumbs: []
    }),
    beforeMount () {
      // INIT GENERATE AND FILL FROM
      this.formId = this.$route.params.fid
      this.contentId = this.$route.params.id
      const axithis = this

      axithis.breadcrumbs.push({label: 'Independent pages', to: '/admiin/independents'})

      if (this.formId > 0) {
        axios.post(this.$store.state.apiPath + '/form-generator/api',
          {
            key: 'getformjson',
            formId: this.formId,
            formType: this.formType
          }
        ).then(function (response) {
          axithis.formProps = JSON.parse(response.data.Data.Json)
          axithis.breadcrumbs.push({label: response.data.Data.Title, to: ''})

          if (axithis.formType === 'independent') {
            axios.post(axithis.$store.state.apiPath + '/crud/api',
              {
                formId: axithis.formId,
                key: 'readIndependent'
              }
            ).then(function (response) {
              var fs = axithis.formProps.fields
              var formVal = {}
              for (let i = 0; i < fs.length; i++) {
                if (response.data[fs[i].model]) {
                  formVal[fs[i].model] = response.data[fs[i].model].V
                } else {
                  if (fs[i].type === 'input' || fs[i].type === 'textarea') {
                    formVal[fs[i].model] = ''
                  } else if (fs[i].type === 'images' || fs[i].type === 'fieldset' || fs[i].type === 'calendar') {
                    formVal[fs[i].model] = []
                  }
                }
              }
              axithis.formVal = formVal

              axithis.breadcrumbs.push({label: 'form edit', to: ''})
              axithis.$store.commit('changeBreadcrumbs', axithis.breadcrumbs)
  
              axithis.loading = false
              userlog.log(axithis, {type: 'open_edit_independent_form', content_title: axithis.formVal.title})
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