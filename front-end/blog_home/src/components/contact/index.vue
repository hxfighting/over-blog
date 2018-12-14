<template>
    <div>
        <div class="box">
            <center>
                <div class="box-header">
                    <h1 class="center">Contact</h1>
                </div>
            </center>
            <div class="box-content">
                <div id="contact_form">
                    <div id="ff">
                        <label>
                            <span>请输入你的名字:</span>
                            <input type="text" name="name" id="name" required placeholder="任意名字都可以哦！">
                        </label>
                        <label>
                            <span>请输入你的邮箱:</span>
                            <input type="email" name="email" id="email" required placeholder="请输入正确邮箱,错误邮箱收不到回复哦！">
                        </label>
                        <label>
                            <span>在这里写下你想说的话:</span>
                            <textarea name="content" id="message"
                                      placeholder="说点你对这个家伙想说的话！,如果是交换友链，请写下你的友链地址和友链名字。"></textarea>
                        </label>
                        <center><input class="sendButton" type="submit" name="Submit" value="确定"></center>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
    export default {
        name: "index"
    }
</script>

<style scoped>

</style>
