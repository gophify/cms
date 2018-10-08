export default {
  auth (e) {
    if (localStorage.getItem('session')) {
      e.$store.commit('updateSession', JSON.parse(localStorage.getItem('session')))
    } else {
      e.$router.push({path: '/'})
      return false
    }
  }
}
