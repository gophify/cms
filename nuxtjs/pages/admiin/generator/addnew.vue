<template>
    <div>
        <formBase :formProps="formProps" :formVal="formVal" mod="generator"></formBase>
        <formGenerator :formId="formId" :formProps="formProps" @formProps="formPropsMethod" @openGenerator="openGeneratorMethod" v-if="openFormGenerator" />
        <v-btn small dark color="red" v-on:click="openFormGenerator=true">open form generator</v-btn>
    </div>
</template>

<script>
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
    }
  }
</script>