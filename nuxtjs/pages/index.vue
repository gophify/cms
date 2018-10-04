<template>
  <div class="loginForm">
      <h1>Sign in to Gophify Admin v2</h1>
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
import userlog from '../utils/userlog.js'
export default {
  layout: 'blank',
  data () {
    return {
      loadingpage: true,
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
          axithis.login_loading = false
          if (response.data.Status === 'success') {
            userlog.log(axithis, 'login')
            localStorage.setItem('session', JSON.stringify({role: response.data.Role, firstname: response.data.Firstname, lastname: response.data.Lastname, uname: response.data.Uname, email: response.data.Email}))
            axithis.$router.push({path: '/admiin'})
          } else {
            alert('Invalid Login')
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