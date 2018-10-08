export const state = () => ({
  sidebar: false,
  activeNav: '',
  breadcrumbs: [],
  mediaPath: 'http://localhost:8080', // localhost/ 172.104.185.200
  apiPath: 'http://172.104.185.200:8080',
  // session: {role: 'root', firstname: 'Nawi', lastname: 'Kartini', uname: 'nawika', email: 'nawi@gmail.com'}
  session: {}
})

export const mutations = {
  toggleSidebar (state) {
    state.sidebar = !state.sidebar
  },

  changeActiveNav (state, activeNav) {
    state.activeNav = activeNav
  },

  changeBreadcrumbs (state, breadcrumbs) {
    state.breadcrumbs = breadcrumbs
  },

  updateSession (state, session) {
    state.session = session
  }
}
