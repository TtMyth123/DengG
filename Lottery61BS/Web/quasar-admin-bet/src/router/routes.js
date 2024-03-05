const routes = [
  {
    path: '/',
    redirect: '/Bet',
    component: () => import('layouts/MainLayout.vue'),
    children: [
      {path: '',
        redirect: '/Bet',
        component: () => import('pages/Home/home.vue')},
      {
        path: '/Home',
        component: () => import('pages/Home/home.vue')
      },
      {
        path: '/UserProfile',
        component: () => import('pages/Home/UserProfile.vue'),
        meta: { title: '我的信息' }
      },

      {
        path: '/SysUserList',
        component: () => import('pages/SysUser/SysUserList.vue')
      },

      {
        path: '/SixSiteType',
        component: () => import('pages/SixLottery/SixSiteType/SixSiteType.vue')
      },
      {
        path: '/SixSiteUser',
        component: () => import('pages/SixLottery/SixSiteUser/SixSiteUser.vue')
      },
      {
        path: '/Bet',
        component: () => import('pages/SixLottery/Bet/Bet.vue')
      },
    ]
  },

  {
    path: '/Login',
    component: () => import('pages/RegLogin/Login.vue')
  },

  {
    path: '/:catchAll(.*)*',
    component: () => import('pages/Error404.vue')
  },
]

export default routes
