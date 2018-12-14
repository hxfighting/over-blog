<template>
    <div class="user-avator-dropdown">
        <Dropdown @on-click="handleClick">
            <Badge :dot="!!messageUnreadCount">
                <Avatar :src="userAvator"/>
            </Badge>
            <Icon :size="18" type="md-arrow-dropdown"></Icon>
            <DropdownMenu slot="list">
                <DropdownItem name="password">修改密码</DropdownItem>
                <DropdownItem name="message">
                    个人信息
                    <!--<Badge style="margin-left: 10px" :count="messageUnreadCount"></Badge>-->
                </DropdownItem>
                <DropdownItem name="logout">退出登录</DropdownItem>
            </DropdownMenu>
        </Dropdown>
    </div>
</template>

<script>
    import './user.less'
    import {mapActions} from 'vuex'

    export default {
        name: 'User',
        props: {
            userAvator: {
                type: String,
                default: ''
            },
            messageUnreadCount: {
                type: Number,
                default: 0
            }
        },
        methods: {
            ...mapActions([
                'handleLogOut'
            ]),
            logout() {
                this.handleLogOut().then(() => {
                    this.$router.push({
                        name: 'login'
                    })
                })
            },
            message() {
                this.$router.push({
                    path: '/info/info_page'
                })
            },
            password() {
                this.$router.push({
                    path: '/password/password_page'
                })
            },
            handleClick(name) {
                switch (name) {
                    case 'logout':
                        this.logout()
                        break
                    case 'message':
                        this.message()
                        break
                    case 'password':
                        this.password()
                        break
                }
            }
        }
    }
</script>
