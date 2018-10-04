<template>   
    <formBase :loading="loading" :cparent="cparent" :formType="formType" :formProps="formProps" :formVal="formVal" :formId="formId" :contentId="contentId" mod=""></formBase>
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
      breadcrumbs: [],
      cparent: ''
    }),
    beforeMount () {
      if (this.$route.params.child) {
        this.cparent = this.$route.params.id
        this.formId = this.$route.params.child
      } else {
        this.formId = this.$route.params.fid
      }
      // INIT GENERATE AND FILL FROM
      const axithis = this

      if (this.formId !== '') {
        axios.post(this.$store.state.apiPath + '/form-generator/api',
          {
            key: 'getformjson',
            formId: this.formId,
            formType: 'regular'
          }
        ).then(function (response) {
          axithis.formProps = JSON.parse(response.data.Data.Json)
          axithis.$store.commit('changeActiveNav', axithis.formProps.form_name)
          axithis.breadcrumbs.push({label: response.data.Data.Title, to: '/admiin/' + axithis.formId + '/list/'})
          axithis.breadcrumbs.push({label: 'add new form', to: ''})
          axithis.$store.commit('changeBreadcrumbs', axithis.breadcrumbs)

          axithis.loading = false
          userlog.log(axithis, {type: 'open_add_form', table_name: axithis.formProps.table_name})
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