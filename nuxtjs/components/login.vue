<template>
  <div>
      <h1>Sign in to Gophify Admin</h1>
      <v-form ref="form" v-model="valid" lazy-validation>
          <v-text-field
              prepend-icon="account_circle"
              v-model="email"
              :rules="emailRules"
              label="E-mail"
              required outline
            ></v-text-field>
                        
        <v-text-field type="password"
            v-model="password"
            prepend-icon="vpn_key"
            :rules="passwRules"
            label="Password"
            required outline
          ></v-text-field>

          <v-btn :loading="login_loading" v-on:click="onSubmit" class="teal lighten-1 white--text" dark large>Sign In</v-btn>
      </v-form>    
  </div>
</template>

<script>
import axios from 'axios'
export default {
  layout: 'login',
  data () {
    return {
      login_loading: false,
      valid: false,
      password: '',
      passwRules: [
        v => !!v || 'Password is required'
      ],
      email: '',
      emailRules: [
        v => !!v || 'E-mail is required',
        v => /.+@.+/.test(v) || 'E-mail must be valid'
      ]
    }
  },
  methods: {
    onSubmit () {
      if (this.$refs.form.validate()) {
        const axithis = this
        this.login_loading = true
        axios.post(this.$store.state.apiPath + '/auth/api',
          {
            email: this.email,
            password: this.password
          },
          {
            headers: {
              'Content-Type': 'multipart/form-data'
            }
          }
        ).then(function (response) {
          // console.log(response.data)
          axithis.login_loading = false
          if (response.data.Status === 'success') {
            window.location.href = '/admiin/'
          } else {
            alert(JSON.stringify(response.data.Err))
          }
        })
          .catch(function (err) {
            console.log(err)
            axithis.login_loading = false
          })
      }
    }
  }
}
</script>

<style>
  body{
    background-color: rgb(83, 110, 121);
    font-family: 'Open Sans Condensed', sans-serif!important;
  }
  .error--text{
    color: red;
  }
  h1{
    margin: 15% 5% 15px;
    letter-spacing: 2px;
    font-size: 42px;
    color: white;
  }
  button{
    height: 60px!important;
    width: 250px!important;
    margin: 20px 5%!important;
  }
  .v-btn__content{
    font-size: 26px!important;
  }
  input{
    font-family: 'Open Sans Condensed', sans-serif!important;
    padding: 50px 30px 20px!important;
    color: #888;
    font-size: 32px;
  }
  .v-messages{
    font-size: 18px;
    font-weight: bold;
    min-height: 26px;
    margin-right: 20px;
  }
  .v-input input{
    max-height: 90px;
  }
  .v-input{
    padding: 0 5%;
    background-color: rgb(243, 252, 252);
    border-bottom: 1px dashed rgba(83, 110, 121, .3);
  }
  .v-input__slot{
    border: none!important;
  }
</style>