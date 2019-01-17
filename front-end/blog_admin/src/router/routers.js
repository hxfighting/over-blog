import Main from '@/components/main'
import parentView from '@/components/parent-view'

/**
 * iview-admin中meta除了原生参数外可配置的参数:
 * meta: {
 *  title: { String|Number|Function }
 *         显示在侧边栏、面包屑和标签栏的文字
 *         使用'{{ 多语言字段 }}'形式结合多语言使用，例子看多语言的路由配置;
 *         可以传入一个回调函数，参数是当前路由对象，例子看动态路由和带参路由
 *  hideInBread: (false) 设为true后此级路由将不会出现在面包屑中，示例看QQ群路由配置
 *  hideInMenu: (false) 设为true后在左侧菜单不会显示该页面选项
 *  notCache: (false) 设为true后页面在切换标签后不会缓存，如果需要缓存，无需设置这个字段，而且需要设置页面组件name属性和路由配置的name一致
 *  access: (null) 可访问该页面的权限数组，当前路由设置的权限会影响子路由
 *  icon: (-) 该页面在左侧菜单、面包屑和标签导航处显示的图标，如果是自定义图标，需要在图标名称前加下划线'_'
 *  beforeCloseName: (-) 设置该字段，则在关闭当前tab页时会去'@/router/before-close.js'里寻找该字段名对应的方法，作为关闭前的钩子函数
 * }
 */

export default [
    {
        path: '/login',
        name: 'login',
        meta: {
            title: '登录 - 登录',
            hideInMenu: true
        },
        component: () => import('@/view/login/login.vue')
    },
    {
        path: '/',
        name: '_home',
        redirect: '/home',
        component: Main,
        meta: {
            hideInMenu: true,
            notCache: true
        },
        children: [
            {
                path: '/home',
                name: 'home',
                meta: {
                    hideInMenu: true,
                    title: '首页',
                    notCache: true,
                    icon: 'md-home'
                },
                component: () => import('@/view/single-page/home')
            }
        ]
    },
    {
        path: '/article',
        name: '文章管理',
        component: Main,
        meta: {
            hideInBread: true
        },
        children: [
            {
                path: 'article_page',
                name: '文章管理',
                meta: {
                    icon: 'ios-book',
                    title: '文章管理'
                },
                component: () => import('@/view/article.vue')
            }
        ]
    },
    {
        path: '/category',
        name: '分类管理',
        component: Main,
        meta: {
            hideInBread: true
        },
        children: [
            {
                path: 'category_page',
                name: '分类管理',
                meta: {
                    icon: 'md-list',
                    title: '分类管理'
                },
                component: () => import('@/view/category.vue')
            }
        ]
    },
    {
        path: '/comment',
        name: '评论管理',
        component: Main,
        meta: {
            hideInBread: true
        },
        children: [
            {
                path: 'comment_page',
                name: '评论管理',
                meta: {
                    icon: 'ios-chatboxes',
                    title: '评论管理'
                },
                component: () => import('@/view/comment.vue')
            }
        ]
    },
    {
        path: '/user',
        name: '会员管理',
        component: Main,
        meta: {
            hideInBread: true
        },
        children: [
            {
                path: 'user_page',
                name: '会员管理',
                meta: {
                    icon: 'ios-people',
                    title: '会员管理'
                },
                component: () => import('@/view/user.vue')
            }
        ]
    },
    {
        path: '/tag',
        name: '标签管理',
        component: Main,
        meta: {
            hideInBread: true
        },
        children: [
            {
                path: 'tag_page',
                name: '标签管理',
                meta: {
                    icon: 'md-pricetags',
                    title: '标签管理'
                },
                component: () => import('@/view/tag.vue')
            }
        ]
    },
    {
        path: '/config',
        name: '配置管理',
        component: Main,
        meta: {
            hideInBread: true
        },
        children: [
            {
                path: 'config_page',
                name: '配置管理',
                meta: {
                    icon: 'ios-construct',
                    title: '配置管理'
                },
                component: () => import('@/view/webConfig.vue')
            }
        ]
    },
    {
        path: '/chat',
        name: '说说管理',
        component: Main,
        meta: {
            hideInBread: true
        },
        children: [
            {
                path: 'chat_page',
                name: '说说管理',
                meta: {
                    icon: 'ios-chatbubbles',
                    title: '说说管理'
                },
                component: () => import('@/view/chat.vue')
            }
        ]
    },
    {
        path: '/link',
        name: '友联管理',
        component: Main,
        meta: {
            hideInBread: true
        },
        children: [
            {
                path: 'link_page',
                name: '友联管理',
                meta: {
                    icon: 'ios-link',
                    title: '友联管理',
                    notCache:true
                },
                component: () => import('@/view/link.vue')
            }
        ]
    },
    {
        path: '/contact',
        name: '留言管理',
        component: Main,
        meta: {
            hideInBread: true
        },
        children: [
            {
                path: 'contact_page',
                name: '留言管理',
                meta: {
                    icon: 'md-mail',
                    title: '留言管理'
                },
                component: () => import('@/view/contact.vue')
            }
        ]
    },
    {
        path: '/rotation',
        name: '轮播图管理',
        component: Main,
        meta: {
            hideInBread: true
        },
        children: [
            {
                path: 'rotation_page',
                name: '轮播图管理',
                meta: {
                    icon: 'ios-image',
                    title: '轮播图管理'
                },
                component: () => import('@/view/rotation.vue')
            }
        ]
    },
    {
        path: '/photo',
        name: '个人照片管理',
        component: Main,
        meta: {
            hideInBread: true
        },
        children: [
            {
                path: 'photo_page',
                name: '个人照片管理',
                meta: {
                    icon: 'ios-image',
                    title: '个人照片管理'
                },
                component: () => import('@/view/photo.vue')
            }
        ]
    },
    {
        path: '/info',
        name: '个人中心',
        component: Main,
        meta: {
            hideInBread: true,
            hideInMenu: true
        },
        children: [
            {
                path: 'info_page',
                name: '个人中心',
                meta: {
                    icon: 'md-notifications',
                    title: '个人中心'
                },
                component: () => import('@/view/adminInfo.vue')
            }
        ]
    },
    {
        path: '/password',
        name: '密码修改',
        component: Main,
        meta: {
            hideInBread: true,
            hideInMenu: true
        },
        children: [
            {
                path: 'password_page',
                name: '密码修改',
                meta: {
                    icon: 'md-notifications',
                    title: '密码修改'
                },
                component: () => import('@/view/resetPassword.vue')
            }
        ]
    },
    {
        path: '/error_logger',
        name: 'error_logger',
        meta: {
            hideInBread: true,
            hideInMenu: true
        },
        component: Main,
        children: [
            {
                path: 'error_logger_page',
                name: 'error_logger_page',
                meta: {
                    icon: 'ios-bug',
                    title: '错误收集'
                },
                component: () => import('@/view/single-page/error-logger.vue')
            }
        ]
    },
    {
        path: '/401',
        name: 'error_401',
        meta: {
            hideInMenu: true
        },
        component: () => import('@/view/error-page/401.vue')
    },
    {
        path: '/500',
        name: 'error_500',
        meta: {
            hideInMenu: true
        },
        component: () => import('@/view/error-page/500.vue')
    },
    {
        path: '*',
        name: 'error_404',
        meta: {
            hideInMenu: true
        },
        component: () => import('@/view/error-page/404.vue')
    }
]
