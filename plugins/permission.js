import { boot } from 'quasar/wrappers'
import { LoadingBar, Loading, QSpinnerGears } from 'quasar'
// import { Allowlist } from 'src/settings'
// import { GetToken } from 'src/utils/cookies'
import { useUserStore } from 'src/stores/user'
import { usePermissionStore } from 'src/stores/permission'
import { postAction } from 'src/api/manage'
import { i18n } from './i18n'

const userStore = useUserStore()
const permissionStore = usePermissionStore()

LoadingBar.setDefaults({
    color: 'red',
    size: '4px',
    position: 'top'
})

function startLoading() {
    Loading.show({
        // spinner: QSpinnerGears,
        message: i18n.global.t('System') + i18n.global.t('Loading')
    })
    LoadingBar.start()
}

function stopLoading() {
    Loading.hide()
    LoadingBar.stop()
}

export default boot(({ router, store }) => {
    router.beforeEach((to, from, next) => {
        startLoading()
        const token = userStore.GetToken()
        if (token) {
            if (!permissionStore.userMenu.length) {
                permissionStore.GetUserMenu().then(res => {
                    // 在vue-router4中，addRoutes被废弃，改为了addRoute，循环调用
                    // 动态添加鉴权路由表
                    if (res) {
                        res.forEach(item => {
                            router.addRoute(item)
                        })
                        next({ ...to, replace: true })
                    } else {
                        store.dispatch('user/HandleLogout')
                        next({ path: '/', replace: true })
                    }
                })
            } else {
                next()
            }
            stopLoading()
        } else {
            const st = location.search.replace('?ticket=', '')
            const svc = "http://" + window.location.host + "/";
            const serviceUrl = encodeURIComponent(svc)
            const params = {
                app_id: 'gqa',
                ticket: st,
                service: serviceUrl
            }
            if (st) {
                validateTicket(params)
            } else {
                window.location.href = "http://192.168.44.121/sso/login?appid=gqa&service=" + serviceUrl
            }
        }
    })
    router.afterEach(() => {
        stopLoading()
    })
})

const validateTicket = async (params) => {
    const vt = await postAction('public/plugin-weaverCas/validate-ticket', params)
    if (vt.code === 1) {
        const login = await userStore.CasLogin(vt.data.records)
        if (login) {
            location.replace(location.href.replace(location.search, ''))
        } else {
            const svc = "http://" + window.location.host + "/";
            const serviceUrl = encodeURIComponent(svc)
            window.location.href = "http://192.168.44.121/sso/login?appid=gqa&service=" + serviceUrl
        }
    } else {
        const svc = "http://" + window.location.host + "/";
        const serviceUrl = encodeURIComponent(svc)
        window.location.href = "http://192.168.44.121/sso/login?appid=gqa&service=" + serviceUrl
    }
}
