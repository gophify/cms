import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

const _519f7f5b = () => import('../pages/admiin/index.vue' /* webpackChunkName: "pages/admiin/index" */).then(m => m.default || m)
const _c21352ac = () => import('../pages/admiin/settings.vue' /* webpackChunkName: "pages/admiin/settings" */).then(m => m.default || m)
const _48506f80 = () => import('../pages/admiin/independents.vue' /* webpackChunkName: "pages/admiin/independents" */).then(m => m.default || m)
const _0155ed6d = () => import('../pages/admiin/media.vue' /* webpackChunkName: "pages/admiin/media" */).then(m => m.default || m)
const _82dd38bc = () => import('../pages/admiin/generator/addnew.vue' /* webpackChunkName: "pages/admiin/generator/addnew" */).then(m => m.default || m)
const _42122ff2 = () => import('../pages/admiin/generator/_id.vue' /* webpackChunkName: "pages/admiin/generator/_id" */).then(m => m.default || m)
const _625f162a = () => import('../pages/admiin/independent/_fid/edit.vue' /* webpackChunkName: "pages/admiin/independent/_fid/edit" */).then(m => m.default || m)
const _d0378b4a = () => import('../pages/admiin/_fid/add.vue' /* webpackChunkName: "pages/admiin/_fid/add" */).then(m => m.default || m)
const _012190d4 = () => import('../pages/admiin/_fid/list.vue' /* webpackChunkName: "pages/admiin/_fid/list" */).then(m => m.default || m)
const _5c6079f5 = () => import('../pages/admiin/_fid/_id/edit.vue' /* webpackChunkName: "pages/admiin/_fid/_id/edit" */).then(m => m.default || m)
const _dcef824c = () => import('../pages/admiin/_fid/_id/_child/add.vue' /* webpackChunkName: "pages/admiin/_fid/_id/_child/add" */).then(m => m.default || m)
const _3bfe1c35 = () => import('../pages/admiin/_fid/_id/_child/list.vue' /* webpackChunkName: "pages/admiin/_fid/_id/_child/list" */).then(m => m.default || m)
const _66903476 = () => import('../pages/admiin/_fid/_id/_child/_ccid/edit.vue' /* webpackChunkName: "pages/admiin/_fid/_id/_child/_ccid/edit" */).then(m => m.default || m)
const _94f7f9e4 = () => import('../pages/index.vue' /* webpackChunkName: "pages/index" */).then(m => m.default || m)



if (process.client) {
  window.history.scrollRestoration = 'manual'
}
const scrollBehavior = function (to, from, savedPosition) {
  // if the returned position is falsy or an empty object,
  // will retain current scroll position.
  let position = false

  // if no children detected
  if (to.matched.length < 2) {
    // scroll to the top of the page
    position = { x: 0, y: 0 }
  } else if (to.matched.some((r) => r.components.default.options.scrollToTop)) {
    // if one of the children has scrollToTop option set to true
    position = { x: 0, y: 0 }
  }

  // savedPosition is only available for popstate navigations (back button)
  if (savedPosition) {
    position = savedPosition
  }

  return new Promise(resolve => {
    // wait for the out transition to complete (if necessary)
    window.$nuxt.$once('triggerScroll', () => {
      // coords will be used if no selector is provided,
      // or if the selector didn't match any element.
      if (to.hash) {
        let hash = to.hash
        // CSS.escape() is not supported with IE and Edge.
        if (typeof window.CSS !== 'undefined' && typeof window.CSS.escape !== 'undefined') {
          hash = '#' + window.CSS.escape(hash.substr(1))
        }
        try {
          if (document.querySelector(hash)) {
            // scroll to anchor by returning the selector
            position = { selector: hash }
          }
        } catch (e) {
          console.warn('Failed to save scroll position. Please add CSS.escape() polyfill (https://github.com/mathiasbynens/CSS.escape).')
        }
      }
      resolve(position)
    })
  })
}


export function createRouter () {
  return new Router({
    mode: 'history',
    base: '/',
    linkActiveClass: 'nuxt-link-active',
    linkExactActiveClass: 'nuxt-link-exact-active',
    scrollBehavior,
    routes: [
		{
			path: "/admiin",
			component: _519f7f5b,
			name: "admiin"
		},
		{
			path: "/admiin/settings",
			component: _c21352ac,
			name: "admiin-settings"
		},
		{
			path: "/admiin/independents",
			component: _48506f80,
			name: "admiin-independents"
		},
		{
			path: "/admiin/media",
			component: _0155ed6d,
			name: "admiin-media"
		},
		{
			path: "/admiin/generator/addnew",
			component: _82dd38bc,
			name: "admiin-generator-addnew"
		},
		{
			path: "/admiin/generator/:id?",
			component: _42122ff2,
			name: "admiin-generator-id"
		},
		{
			path: "/admiin/independent/:fid?/edit",
			component: _625f162a,
			name: "admiin-independent-fid-edit"
		},
		{
			path: "/admiin/:fid/add",
			component: _d0378b4a,
			name: "admiin-fid-add"
		},
		{
			path: "/admiin/:fid/list",
			component: _012190d4,
			name: "admiin-fid-list"
		},
		{
			path: "/admiin/:fid/:id?/edit",
			component: _5c6079f5,
			name: "admiin-fid-id-edit"
		},
		{
			path: "/admiin/:fid/:id?/:child?/add",
			component: _dcef824c,
			name: "admiin-fid-id-child-add"
		},
		{
			path: "/admiin/:fid/:id?/:child?/list",
			component: _3bfe1c35,
			name: "admiin-fid-id-child-list"
		},
		{
			path: "/admiin/:fid/:id?/:child?/:ccid?/edit",
			component: _66903476,
			name: "admiin-fid-id-child-ccid-edit"
		},
		{
			path: "/",
			component: _94f7f9e4,
			name: "index"
		}
    ],
    
    
    fallback: false
  })
}
