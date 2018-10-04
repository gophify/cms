import axios from 'axios'
export default {
  log (e, action) {
    if (e.$store.state.session.role !== 'root') {
      const axithis = e
      if (action !== null && typeof action === 'object') {
        action = JSON.stringify(action)
      }
      const action2 = action
      axios
        .post(e.$store.state.apiPath + '/userlog-api',
          {
            role: axithis.$store.state.session.role,
            fullname: axithis.$store.state.session.firstname + ' ' + axithis.$store.state.session.lastname,
            email: axithis.$store.state.session.email,
            action: action2
          })
    }
  }
}
