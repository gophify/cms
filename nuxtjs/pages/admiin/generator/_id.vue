<template>
    <div>
        <formBase :formProps="formProps" :formVal="formVal" mod="generator"></formBase>
        <formGenerator :formId="formId" :formProps="formProps" @formProps="formPropsMethod" @openGenerator="openGeneratorMethod" v-if="openFormGenerator" />
        <v-btn small dark color="red" v-on:click="openFormGenerator=true">open form generator</v-btn>
    </div>
</template>

<script>
  import axios from 'axios'
  import userlog from '../../../utils/userlog.js'
  export default {
    data: () => ({
      formVal: {},
      formId: '0',
      formType: '',
      openFormGenerator: true,
      formProps: {fields: [], tabs: []}
    }),

    methods: {
      formPropsMethod (val) {
        this.formProps = val
      },
      openGeneratorMethod (val) {
        this.openFormGenerator = val
      }
    },
    mounted () {
      var breadcrumbs = {}
      breadcrumbs = [
        {label: 'Villa Suarti and Resorts', to: '/'},
        {label: 'The Suite Room', to: '/'},
        {label: 'Edit Room', to: '/'}
      ]
      this.$store.commit('changeBreadcrumbs', breadcrumbs)
      this.formId = this.$route.params.id
      // alert(this.formId)
      if (!isNaN(parseInt(this.$route.params.id))) {
        this.formType = 'independent'
      } else {
        this.formType = 'regular'
      }
  
      if (this.formId !== '0') {
        const axithis = this
        this.loading = true
        axios.post(this.$store.state.apiPath + '/form-generator/api',
          {
            key: 'getformjson',
            formId: this.formId,
            formType: this.formType
          }
        ).then(function (response) {
          // console.log(response.data.Data.Title)
          userlog.log(axithis, {type: 'open_generator', form_title: response.data.Data.Title})
          if (response.data.Status === 'success') {
            axithis.formProps = JSON.parse(response.data.Data.Json)
            axithis.formProps['table_parent'] = response.data.Data.Parent
            axithis.formProps['table_child'] = response.data.Data.Child
            axithis.formProps['viewChildLbl'] = response.data.Data.ViewChildLbl
            axithis.loading = false
          } else {
            alert(JSON.stringify(response.data.Err))
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