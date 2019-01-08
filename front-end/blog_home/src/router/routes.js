const mainPage = () => import("../pages/Index");
const category = () => import("../components/category/Index");
const tag = () => import("../components/tag/Index");
const indexLeftContent = () => import("../components/content/leftSide");
const indexHeaderBottom = () => import('../components/header/bottom');
const chat = () => import('../components/chat/index');
const contact = () => import('../components/contact/index');
const article = () => import('../components/detail/detail');
const articleHeader = () => import('../components/detail/detailHeaderBottom');
const routes = [
  {
    path: '/',
    component: mainPage,
    redirect: '/index',
    children: [
      {
        path: 'index',
        components: {
          default: indexHeaderBottom,
          h_content: indexLeftContent
        }
      },
      {
        path: 'category/:id',
        name: 'category',
        components: {
          default: '',
          h_content: category
        },
        beforeEnter (to, from, next) {
          let id = to.params.id;
          if(isNaN(id)){
            next(error)
          }else {
            next()
          }
        },
        beforeUpdate (to, from, next) {
          let id = to.params.id;
          if(isNaN(id)){
            next(error)
          }else {
            next()
          }
        },
      },
      {
        path: 'tag/:id',
        name: 'tag',
        components: {
          default: '',
          h_content: tag
        },
        beforeEnter (to, from, next) {
          let id = to.params.id;
          if(isNaN(id)){
            next(error)
          }else {
            next()
          }
        },
        beforeUpdate (to, from, next) {
          let id = to.params.id;
          if(isNaN(id)){
            next(error)
          }else {
            next()
          }
        },
      },
      {
        path: 'chat',
        components: {
          default: '',
          h_content: chat
        }
      },
      {
        path: 'contact',
        components: {
          default: '',
          h_content: contact
        }
      },
      {
        path: 'article/:id',
        name: 'article',
        components: {
          default: articleHeader,
          h_content: article
        },
        beforeEnter (to, from, next) {
          let id = to.params.id;
          if(isNaN(id)){
            next(error)
          }else {
            next()
          }
        },
        beforeUpdate (to, from, next) {
          let id = to.params.id;
          if(isNaN(id)){
            next(error)
          }else {
            next()
          }
        },
      }
    ]
  }
]

// Always leave this as last one
if (process.env.MODE !== 'ssr') {
  routes.push({
    path: '*',
    component: () => import('pages/Error404.vue')
  })
}

export default routes
