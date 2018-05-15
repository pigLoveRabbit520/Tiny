import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

/* Layout */
import Layout from '../views/layout/Layout'

/**
 * hidden: true                   if `hidden:true` will not show in the sidebar(default is false)
 * alwaysShow: true               if set true, will always show the root menu, whatever its child routes length
 *                                if not set alwaysShow, only more than one route under the children
 *                                it will becomes nested mode, otherwise not show the root menu
 * redirect: noredirect           if `redirect:noredirect` will no redirct in the breadcrumb
 * name:'router-name'             the name is used by <keep-alive> (must set!!!)
 * meta : {
    title: 'title'               the name show in submenu and breadcrumb (recommend set)
    icon: 'svg-name'             the icon show in the sidebar,
  }
 **/
export const constantRouterMap = [
    { path: '/login', component: () => import('@/views/login/index'), hidden: true },
    { path: '/404', component: () => import('@/views/404'), hidden: true },

    {
        path: '/',
        component: Layout,
        redirect: '/dashboard',
        name: 'Dashboard',
        hidden: true,
        children: [{
            path: 'dashboard',
            component: () => import('@/views/dashboard/index')
        }]
    },
    {
        path: '/manage',
            component: Layout,
            name: 'Manage',
            meta: { title: '管理', icon: 'form' },
            children: [
                {
                  path: 'posts',
                  name: 'Posts',
                  component: () => import('@/views/dashboard/index'),
                  meta: { title: '文章', icon: 'table' }
                },
                {
                  path: 'pages',
                  name: 'Pages',
                  component: () => import('@/views/dashboard/index'),
                  meta: { title: '独立页面', icon: 'tree' }
                },
                {
                  path: 'comments',
                  name: 'Comments',
                  component: () => import('@/views/dashboard/index'),
                  meta: { title: '评论', icon: 'tree' }
                },
                {
                  path: 'categories',
                  name: 'Categories',
                  component: () => import('@/views/dashboard/index'),
                  meta: { title: '分类', icon: 'tree' }
                }
            ]
    },

    {
        path: '/settings',
        component: Layout,
        meta: { title: '设置', icon: 'form' },
        children: [
            {
                path: 'general',
                name: 'General',
                component: () => import('@/views/settings/index'),
                meta: { title: '常规', icon: 'form' }
            },
            {
              path: 'tree',
              name: 'Tree',
              component: () => import('@/views/dashboard/index'),
              meta: { title: 'Tree', icon: 'tree' }
            }
        ]
    },

    { path: '*', redirect: '/404', hidden: true }
]

export default new Router({
    // mode: 'history', //后端支持可开
    scrollBehavior: () => ({ y: 0 }),
    routes: constantRouterMap
})
