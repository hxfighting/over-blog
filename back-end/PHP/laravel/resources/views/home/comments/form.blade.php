<div  class="comment_form hx_comment" style="display: none">
    @if(isset($parentId))
        <input type="hidden" class="co_pid" name="co_pid" value="{{$parentId}}">
    @endif
    <div class="form-group hx_group">
        <input class="pull-right co_email" type="email" name="co_email" placeholder="请输入接收回复的邮箱" required="required"/>
        <textarea id="co_content" name="co_content" class="form-control co_content" required="required" placeholder="回复内容"></textarea>
    </div>
    <button type="submit" class="btn btn-success co_reply">回复</button>
    <button type="submit" class="btn btn-success co_cancel">取消</button>
</div>