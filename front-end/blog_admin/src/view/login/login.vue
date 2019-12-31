<style lang="less">
    @import './login.less';
</style>

<template>
    <div class="login">
        <div class="login-con">
            <Card icon="log-in" title="欢迎登录" :bordered="false">
                <div class="form-con">
                    <Form ref="loginForm" :model="form" :rules="rules" @keydown.enter.native="handleSubmit">
                        <FormItem prop="userName">
                            <Input v-model="form.userName" placeholder="请输入用户名">
                                <span slot="prepend">
                                    <Icon :size="16" type="ios-person"></Icon>
                                </span>
                            </Input>
                        </FormItem>
                        <FormItem prop="password">
                            <Input type="password" v-model="form.password" placeholder="请输入密码">
                                <span slot="prepend">
                                    <Icon :size="14" type="md-lock"></Icon>
                                </span>
                            </Input>
                        </FormItem>
                        <FormItem prop="captcha">
                            <Input v-model="form.captcha" placeholder="请输入验证码">
                                <span slot="prepend">
                                    <img :size="14"
                                         :src="captcha_url"
                                         style="width: 100px;height: 21px;cursor: pointer" @click="getCaptcha">
                                </span>
                            </Input>
                        </FormItem>
                        <FormItem>
                            <Button @click="handleSubmit" type="primary" long :disabled="login_disabled">登录</Button>
                        </FormItem>
                    </Form>
                    <p class="login-tip">拖油瓶博客</p>
                </div>
            </Card>
        </div>
    </div>
</template>

<script>
    import {mapActions} from 'vuex'
    import {getCaptcha} from "../../api/data";
    export default {
        props: {
            userNameRules: {
                type: Array,
                default: () => {
                    return [
                        {required: true, message: '账号不能为空', trigger: 'blur'}
                    ]
                }
            },
            passwordRules: {
                type: Array,
                default: () => {
                    return [
                        {required: true, message: '密码不能为空', trigger: 'blur'}
                    ]
                }
            },
            captchaRules: {
                type: Array,
                default: () => {
                    return [
                        {required: true, message: '验证码不能为空', trigger: 'blur'}
                    ]
                }
            }
        },
        data() {
            return {
                form: {
                    userName: '',
                    password: '',
                    captcha: '',
                    key: ''
                },
                login_disabled: false,
                captcha_url: ''
            }
        },
        computed: {
            rules() {
                return {
                    userName: this.userNameRules,
                    password: this.passwordRules,
                    captcha: this.captchaRules
                }
            }
        },
        methods: {
            ...mapActions([
                'handleLogin',
                'getUserInfo'
            ]),
            handleSubmit() {
                this.$refs.loginForm.validate((valid) => {
                    if (valid) {
                        this.login_disabled = true
                        let userName = this.form.userName,
                            password = this.form.password,
                            captcha = this.form.captcha,
                            key = this.form.key;
                        this.handleLogin({userName, password,captcha,key})
                            .then(res => {
                                this.getUserInfo().then(res => {
                                    this.$Message.success('登录成功')
                                    this.$router.push({
                                        name: this.$config.homeName
                                    })
                                    this.login_disabled = false
                                })
                            }).catch(err => {
                            this.login_disabled = false;
                            this.$Message.error(err);
                            this.getCaptcha()
                        })
                    }
                })
            },

            //获取验证码
            getCaptcha(){
                getCaptcha().then(res=>{
                    let data = res.data.data;
                    this.captcha_url = data.img;
                    this.form.key = data.key;
                })
            }
        },
        created() {
            this.getCaptcha();
        }
    }
</script>

<style>

</style>
